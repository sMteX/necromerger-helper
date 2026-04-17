package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sMteX/necro-prestige-planner/internal/calculator"
	"github.com/sMteX/necro-prestige-planner/internal/db/sqlc"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sqlc-dev/pqtype"
)

type Handler struct {
	Queries sqlc.Querier
}

type PlanSummary struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ExperimentSummary struct {
	ID                models.ExperimentID `json:"id"`
	CurrentLevel      int                 `json:"currentLevel"`
	CurrentLevelValue string              `json:"currentLevelValue"`
	CurrentLevelCost  string              `json:"currentLevelCost"`
	NextLevelCost     string              `json:"nextLevelCost"`
	NextLevelValue    string              `json:"nextLevelValue"`
	MaxLevel          bool                `json:"maxLevel"`
}

type RecalculateResponse struct {
	TotalShards      int                                         `json:"totalShards"`
	BaseShards       int                                         `json:"baseShards"`
	FeatMultiplier   float64                                     `json:"featMultiplier"`
	LegendMultiplier float64                                     `json:"legendMultiplier"`
	OtherMultiplier  float64                                     `json:"otherMultiplier"`
	ExperimentCost   int                                         `json:"experimentCost"`
	Remaining        int                                         `json:"remaining"`
	Experiments      map[models.ExperimentID]ExperimentSummary   `json:"experiments"`
	RuneTotal        map[models.RuneType]int                     `json:"runeTotal"`
	RuneNeeded       map[models.RuneType]int                     `json:"runeNeeded"`
	LegendaryRunes   map[models.LegendaryID]calculator.RuneCosts `json:"legendaryRunes"`
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

	// Convert sqlc.Plan to models.Plan
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
//	@Summary	Create or update a plan
//	@Description	Creates a new plan when id is 0 or omitted, updates an existing plan when id > 0
//	@Tags		plans
//	@Accept		json
//	@Produce	json
//	@Param		plan	body		models.Plan	true	"Plan data"
//	@Success	200		{object}	models.Plan
//	@Failure	400		{string}	string	"invalid request body"
//	@Failure	500		{string}	string	"internal server error"
//	@Router		/api/plans [post]
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
	baseShards := calculator.DevourerBaseShards[plan.DevourerLevel]
	featMultiplier := 1.0 + (float64(plan.FeatTiers) * 0.10)

	expCost := calculator.CalculateExperimentCost(plan)
	totalRunes, neededRunes := calculator.CalculateTotalRunes(plan)

	// Experiment summaries
	expSummaries := make(map[models.ExperimentID]ExperimentSummary)
	for _, exp := range calculator.Experiments {
		currentLevel := plan.ExperimentLevels[exp.ID]

		summary := ExperimentSummary{
			ID:           exp.ID,
			CurrentLevel: currentLevel,
		}

		if currentLevel > 0 {
			prevVal := 0.0
			if currentLevel-1 < len(exp.Levels) {
				currentLv := exp.Levels[currentLevel-1]
				prevVal = currentLv.PrevValue
				summary.CurrentLevelCost = calculator.FormatLargeNumber(currentLv.Cost)
				summary.CurrentLevelValue = fmt.Sprintf("%s -> %s",
					calculator.FormatExperimentValue(exp.ID, exp.Tier, prevVal),
					calculator.FormatExperimentValue(exp.ID, exp.Tier, currentLv.Value))
			}
		} else {
			summary.CurrentLevelValue = calculator.FormatExperimentValue(exp.ID, exp.Tier, 0)
			summary.CurrentLevelCost = "0"
		}

		if currentLevel < len(exp.Levels) {
			nextLevel := exp.Levels[currentLevel]
			summary.NextLevelCost = calculator.FormatLargeNumber(nextLevel.Cost)
			summary.NextLevelValue = calculator.FormatExperimentValue(exp.ID, exp.Tier, nextLevel.Value)
		} else {
			summary.MaxLevel = true
		}
		expSummaries[exp.ID] = summary
	}

	// Legendary rune costs
	legRunes := make(map[models.LegendaryID]calculator.RuneCosts)
	for id := range calculator.LegendaryRecipes {
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
