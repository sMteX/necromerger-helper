package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sMteX/necro-prestige-planner/internal/calculator"
	"github.com/sMteX/necro-prestige-planner/internal/data"
	"github.com/sMteX/necro-prestige-planner/internal/db/sqlc"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sMteX/necro-prestige-planner/internal/tui/shared"
	"github.com/sqlc-dev/pqtype"
)

type Handler struct {
	Queries sqlc.Querier
}

// HealthHandler godoc
//
//	@Summary	Health check
//	@Tags		system
//	@Produce	plain
//	@Success	200	{string}	string	"OK"
//	@Router		/health [get]
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// ListPlansHandler godoc
//
//	@Summary	List all saved plans
//	@Tags		plans
//	@Produce	json
//	@Success	200	{array}		PlanSummary
//	@Failure	500	{string}	string	"internal server error"
//	@Router		/api/plans [get]
func (h *Handler) ListPlansHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := h.Queries.ListPlans(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	plans := make([]PlanSummary, len(rows))
	for i, row := range rows {
		plans[i] = PlanSummary{ID: int(row.ID), Name: row.Name}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plans)
}

// GetPlanHandler godoc
//
//	@Summary	Get a plan by ID
//	@Tags		plans
//	@Produce	json
//	@Param		id	path		int	true	"Plan ID"
//	@Success	200	{object}	models.Plan
//	@Failure	400	{string}	string	"invalid id"
//	@Failure	404	{string}	string	"plan not found"
//	@Router		/api/plans/{id} [get]
func (h *Handler) GetPlanHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	plan, err := h.Queries.GetPlan(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "plan not found", http.StatusNotFound)
		return
	}

	mPlan := models.Plan{
		ID:              int(plan.ID),
		Name:            plan.Name,
		DevourerLevel:   int(plan.DevourerLevel.Int32),
		FeatTiers:       int(plan.FeatTiers.Int32),
		OtherMultiplier: plan.OtherMultiplier.Float64,
		GroupBonusCount: int(plan.GroupBonusCount.Int32),
		LeftoverShards:  int(plan.LeftoverShards.Int32),
		Notes:           plan.Notes.String,
	}

	json.Unmarshal(plan.LegendaryCounts.RawMessage, &mPlan.LegendaryCounts)
	json.Unmarshal(plan.ExperimentLevels.RawMessage, &mPlan.ExperimentLevels)
	json.Unmarshal(plan.PossessedRunes.RawMessage, &mPlan.PossessedRunes)
	json.Unmarshal(plan.PossessedLegendaries.RawMessage, &mPlan.PossessedLegendaries)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mPlan)
}

// SavePlanHandler godoc
//
//	@Summary		Create or update a plan
//	@Description	Creates a new plan when id is 0 or omitted, updates an existing plan when id > 0
//	@Tags			plans
//	@Accept			json
//	@Produce		json
//	@Param			plan	body		models.Plan	true	"Plan data"
//	@Success		200		{object}	models.Plan
//	@Failure		400		{string}	string	"invalid request body"
//	@Failure		500		{string}	string	"internal server error"
//	@Router			/api/plans [post]
func (h *Handler) SavePlanHandler(w http.ResponseWriter, r *http.Request) {
	var plan models.Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	legCounts, _ := json.Marshal(plan.LegendaryCounts)
	expLevels, _ := json.Marshal(plan.ExperimentLevels)
	posRunes, _ := json.Marshal(plan.PossessedRunes)
	posLegs, _ := json.Marshal(plan.PossessedLegendaries)

	if plan.ID > 0 {
		err := h.Queries.UpdatePlan(r.Context(), sqlc.UpdatePlanParams{
			ID:                   int32(plan.ID),
			Name:                 plan.Name,
			DevourerLevel:        sql.NullInt32{Int32: int32(plan.DevourerLevel), Valid: true},
			FeatTiers:            sql.NullInt32{Int32: int32(plan.FeatTiers), Valid: true},
			OtherMultiplier:      sql.NullFloat64{Float64: plan.OtherMultiplier, Valid: true},
			GroupBonusCount:      sql.NullInt32{Int32: int32(plan.GroupBonusCount), Valid: true},
			LeftoverShards:       sql.NullInt32{Int32: int32(plan.LeftoverShards), Valid: true},
			LegendaryCounts:      pqtype.NullRawMessage{RawMessage: legCounts, Valid: true},
			ExperimentLevels:     pqtype.NullRawMessage{RawMessage: expLevels, Valid: true},
			PossessedRunes:       pqtype.NullRawMessage{RawMessage: posRunes, Valid: true},
			PossessedLegendaries: pqtype.NullRawMessage{RawMessage: posLegs, Valid: true},
			Notes:                sql.NullString{String: plan.Notes, Valid: true},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		id, err := h.Queries.CreatePlan(r.Context(), sqlc.CreatePlanParams{
			Name:                 plan.Name,
			DevourerLevel:        sql.NullInt32{Int32: int32(plan.DevourerLevel), Valid: true},
			FeatTiers:            sql.NullInt32{Int32: int32(plan.FeatTiers), Valid: true},
			OtherMultiplier:      sql.NullFloat64{Float64: plan.OtherMultiplier, Valid: true},
			GroupBonusCount:      sql.NullInt32{Int32: int32(plan.GroupBonusCount), Valid: true},
			LeftoverShards:       sql.NullInt32{Int32: int32(plan.LeftoverShards), Valid: true},
			LegendaryCounts:      pqtype.NullRawMessage{RawMessage: legCounts, Valid: true},
			ExperimentLevels:     pqtype.NullRawMessage{RawMessage: expLevels, Valid: true},
			PossessedRunes:       pqtype.NullRawMessage{RawMessage: posRunes, Valid: true},
			PossessedLegendaries: pqtype.NullRawMessage{RawMessage: posLegs, Valid: true},
			Notes:                sql.NullString{String: plan.Notes, Valid: true},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		plan.ID = int(id)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plan)
}

// DeletePlanHandler godoc
//
//	@Summary	Delete a plan
//	@Tags		plans
//	@Param		id	path		int	true	"Plan ID"
//	@Success	204	{string}	string	"no content"
//	@Failure	400	{string}	string	"invalid id"
//	@Failure	500	{string}	string	"internal server error"
//	@Router		/api/plans/{id} [delete]
func (h *Handler) DeletePlanHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.Queries.DeletePlan(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RecalculateHandler godoc
//
//	@Summary		Calculate Time Shards and resource costs for a plan
//	@Description	Given a plan, returns total Time Shards earned at prestige, experiment costs, rune requirements for planned legendaries, and per-experiment summaries
//	@Tags			calculator
//	@Accept			json
//	@Produce		json
//	@Param			plan	body		models.Plan			true	"Plan to calculate"
//	@Success		200		{object}	RecalculateResponse
//	@Failure		400		{string}	string	"invalid request body"
//	@Router			/api/recalculate [post]
func RecalculateHandler(w http.ResponseWriter, r *http.Request) {
	var plan models.Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shards := calculator.CalculateTimeShards(plan)
	baseShards := data.DevourerBaseShards[plan.DevourerLevel]
	featMultiplier := 1.0 + (float64(plan.FeatTiers) * 0.10)

	expCost := calculator.CalculateExperimentCost(plan)
	totalRunes, neededRunes := calculator.CalculateTotalRunes(plan)

	expSummaries := make(map[models.ExperimentID]ExperimentSummary)
	for _, exp := range data.Experiments {
		currentLevel := plan.ExperimentLevels[exp.ID]
		summary := ExperimentSummary{ID: exp.ID, CurrentLevel: currentLevel}

		if currentLevel > 0 {
			if currentLevel-1 < len(exp.Levels) {
				currentLv := exp.Levels[currentLevel-1]
				summary.CurrentLevelCost = shared.FormatLargeNumber(currentLv.Cost)
				summary.CurrentLevelValue = fmt.Sprintf("%s -> %s",
					shared.FormatExperimentValue(exp.ID, exp.Tier, currentLv.PrevValue),
					shared.FormatExperimentValue(exp.ID, exp.Tier, currentLv.Value))
			}
		} else {
			summary.CurrentLevelValue = shared.FormatExperimentValue(exp.ID, exp.Tier, 0)
			summary.CurrentLevelCost = "0"
		}

		if currentLevel < len(exp.Levels) {
			nextLevel := exp.Levels[currentLevel]
			summary.NextLevelCost = shared.FormatLargeNumber(nextLevel.Cost)
			summary.NextLevelValue = shared.FormatExperimentValue(exp.ID, exp.Tier, nextLevel.Value)
		} else {
			summary.MaxLevel = true
		}
		expSummaries[exp.ID] = summary
	}

	legRunes := make(map[models.LegendaryID]models.RuneCosts)
	for id := range data.LegendaryRecipes {
		legRunes[id] = calculator.GetLegendaryRuneCost(id)
	}

	resp := RecalculateResponse{
		TotalShards:      shards,
		BaseShards:       baseShards,
		FeatMultiplier:   featMultiplier,
		LegendMultiplier: calculator.CalculateLegendMultiplier(plan),
		OtherMultiplier:  1.0 + plan.OtherMultiplier,
		ExperimentCost:   expCost,
		Remaining:        shards - expCost,
		Experiments:      expSummaries,
		RuneTotal:        totalRunes,
		RuneNeeded:       neededRunes,
		LegendaryRunes:   legRunes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// ResourceCapHandler godoc
//
//	@Summary		Calculate resource caps and station requirements for a Feat threshold
//	@Description	Returns current Mana/Slime/Darkness caps, whether the requested Feat threshold is met, and (if not) how many stations of each level are needed per resource to close the gap. Percentage inputs are decimals (0.06 = 6%).
//	@Tags			calculator
//	@Accept			json
//	@Produce		json
//	@Param			threshold	path		string				true	"Combined storage threshold"	Enums(200k, 400k, 600k, 800k)
//	@Param			request		body		ResourceCapRequest	true	"Player state"
//	@Success		200			{object}	ResourceCapResponse
//	@Failure		400			{string}	string	"invalid threshold or request body"
//	@Router			/api/resource-cap/{threshold} [post]
func ResourceCapHandler(w http.ResponseWriter, r *http.Request) {
	thresholdKey := chi.URLParam(r, "threshold")
	threshold, ok := ValidThresholds[thresholdKey]
	if !ok {
		http.Error(w, "invalid threshold: must be one of 200k, 400k, 600k, 800k", http.StatusBadRequest)
		return
	}

	var req ResourceCapRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input := mapToCalculatorInput(req)
	result := calculator.CalculateResourceCaps(input)
	combined := result.Mana + result.Slime + result.Darkness

	resp := ResourceCapResponse{
		Mana:      result.Mana,
		Slime:     result.Slime,
		Darkness:  result.Darkness,
		Combined:  combined,
		Threshold: threshold,
		Met:       combined >= threshold,
		Delta:     combined - threshold,
	}

	if !resp.Met {
		manaTarget, slimeTarget, darkTarget := calculator.ResourceTargets(threshold, req.FixedTargets)

		manaGap := max(0, manaTarget-result.Mana)
		slimeGap := max(0, slimeTarget-result.Slime)
		darkGap := max(0, darkTarget-result.Darkness)

		resp.GapAnalysis = map[models.ResourceType]ResourceGapAnalysis{
			models.ResourceMana: {
				Current: result.Mana,
				Target:  manaTarget,
				Gap:     manaGap,
				Options: toAPIOptions(calculator.ManaStationOptions(manaGap, result.ManaMulti)),
			},
			models.ResourceSlime: {
				Current: result.Slime,
				Target:  slimeTarget,
				Gap:     slimeGap,
				Options: toAPIOptions(calculator.SlimeStationOptions(slimeGap, result.SlimeMulti)),
			},
			models.ResourceDarkness: {
				Current: result.Darkness,
				Target:  darkTarget,
				Gap:     darkGap,
				Options: toAPIOptions(calculator.DarknessStationOptions(darkGap, result.DarkMulti)),
			},
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func toAPIOptions(src []calculator.StationOptionResult) []StationOption {
	if src == nil {
		return nil
	}
	out := make([]StationOption, len(src))
	for i, s := range src {
		runeCost := make(map[models.RuneType]int, len(s.RuneCost))
		for k, v := range s.RuneCost {
			runeCost[k] = v
		}
		out[i] = StationOption{Level: s.Level, Count: s.Count, RuneCost: runeCost}
	}
	return out
}

func mapToCalculatorInput(req ResourceCapRequest) calculator.ResourceCapInput {
	servOResource := models.ResourceType("")
	servOUpgraded := false
	if req.ServO != nil {
		servOResource = req.ServO.Resource
		servOUpgraded = req.ServO.Upgraded
	}

	pearl := req.PearlBonus
	if pearl == nil {
		pearl = map[models.ResourceType]float64{}
	}

	return calculator.ResourceCapInput{
		ManaPools:  req.ManaPools.toSlice(),
		SlimeVats:  req.SlimeVats.toSlice(),
		DarkStores: req.DarkStores.toSlice(),

		ServOResource: servOResource,
		ServOUpgraded: servOUpgraded,

		WizardSkin:   req.Skins.Wizard,
		OozingSkin:   req.Skins.Oozing,
		SidSkin:      req.Skins.Sid,
		SantaSkin:    req.Skins.Santa,
		BirthdaySkin: req.Skins.Birthday,
		WitchSkin:    req.Skins.Witch,
		GoodSkin:     req.Skins.Good,
		RoyaltySkin:  req.Skins.Royalty,

		GoldenBoosts: req.GoldenBoosts,

		ManaSpell:         req.Spells.Mana,
		SlimeSpell:        req.Spells.Slime,
		DarknessSpell:     req.Spells.Darkness,
		AllResourcesSpell: req.Spells.AllResources,

		ManaRelic:         req.Relics.Mana,
		SlimeRelic:        req.Relics.Slime,
		DarknessRelic:     req.Relics.Darkness,
		AllResourcesRelic: req.Relics.AllResources,

		CapacityExp1: req.CapacityExp1,
		PearlBonus:   pearl,
		CapacityExp2: req.CapacityExp2,
	}
}
