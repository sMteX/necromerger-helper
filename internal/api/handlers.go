package api

import (
	"encoding/json"
	"net/http"

	"github.com/sMteX/necro-prestige-planner/internal/calculator"
	"github.com/sMteX/necro-prestige-planner/internal/models"
)

type RecalculateResponse struct {
	TotalShards    int                     `json:"totalShards"`
	ExperimentCost int                     `json:"experimentCost"`
	Remaining      int                     `json:"remaining"`
	RuneTotal      map[models.RuneType]int `json:"runeTotal"`
	RuneNeeded     map[models.RuneType]int `json:"runeNeeded"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func RecalculateHandler(w http.ResponseWriter, r *http.Request) {
	var plan models.Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shards := calculator.CalculateTimeShards(plan)
	expCost := calculator.CalculateExperimentCost(plan)
	totalRunes, neededRunes := calculator.CalculateTotalRunes(plan)

	resp := RecalculateResponse{
		TotalShards:    shards,
		ExperimentCost: expCost,
		Remaining:      shards - expCost,
		RuneTotal:      totalRunes,
		RuneNeeded:     neededRunes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
