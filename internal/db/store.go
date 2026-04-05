package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/sMteX/necro-prestige-planner/internal/db/sqlc"
	"github.com/sMteX/necro-prestige-planner/internal/models"
	"github.com/sqlc-dev/pqtype"
)

type Store struct {
	db      *sql.DB
	queries sqlc.Querier
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		queries: sqlc.New(db),
	}
}

func (s *Store) SavePlan(ctx context.Context, p *models.Plan) error {
	legJSON, _ := json.Marshal(p.LegendaryCounts)
	expJSON, _ := json.Marshal(p.ExperimentLevels)
	runesJSON, _ := json.Marshal(p.PossessedRunes)
	posLegJSON, _ := json.Marshal(p.PossessedLegendaries)

	if p.ID == 0 {
		params := sqlc.CreatePlanParams{
			Name:                 p.Name,
			DevourerLevel:        sql.NullInt32{Int32: int32(p.DevourerLevel), Valid: true},
			FeatTiers:            sql.NullInt32{Int32: int32(p.FeatTiers), Valid: true},
			OtherMultiplier:      sql.NullFloat64{Float64: p.OtherMultiplier, Valid: true},
			GroupBonusCount:      sql.NullInt32{Int32: int32(p.GroupBonusCount), Valid: true},
			LeftoverShards:       sql.NullInt32{Int32: int32(p.LeftoverShards), Valid: true},
			LegendaryCounts:      pqtype.NullRawMessage{RawMessage: legJSON, Valid: true},
			ExperimentLevels:     pqtype.NullRawMessage{RawMessage: expJSON, Valid: true},
			PossessedRunes:       pqtype.NullRawMessage{RawMessage: runesJSON, Valid: true},
			PossessedLegendaries: pqtype.NullRawMessage{RawMessage: posLegJSON, Valid: true},
			Notes:                sql.NullString{String: p.Notes, Valid: true},
		}
		id, err := s.queries.CreatePlan(ctx, params)
		if err != nil {
			return err
		}
		p.ID = int(id)
		return nil
	}

	params := sqlc.UpdatePlanParams{
		ID:                   int32(p.ID),
		Name:                 p.Name,
		DevourerLevel:        sql.NullInt32{Int32: int32(p.DevourerLevel), Valid: true},
		FeatTiers:            sql.NullInt32{Int32: int32(p.FeatTiers), Valid: true},
		OtherMultiplier:      sql.NullFloat64{Float64: p.OtherMultiplier, Valid: true},
		GroupBonusCount:      sql.NullInt32{Int32: int32(p.GroupBonusCount), Valid: true},
		LeftoverShards:       sql.NullInt32{Int32: int32(p.LeftoverShards), Valid: true},
		LegendaryCounts:      pqtype.NullRawMessage{RawMessage: legJSON, Valid: true},
		ExperimentLevels:     pqtype.NullRawMessage{RawMessage: expJSON, Valid: true},
		PossessedRunes:       pqtype.NullRawMessage{RawMessage: runesJSON, Valid: true},
		PossessedLegendaries: pqtype.NullRawMessage{RawMessage: posLegJSON, Valid: true},
		Notes:                sql.NullString{String: p.Notes, Valid: true},
	}
	return s.queries.UpdatePlan(ctx, params)
}

func (s *Store) GetPlan(ctx context.Context, id int) (*models.Plan, error) {
	row, err := s.queries.GetPlan(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	p := &models.Plan{
		ID:              int(row.ID),
		Name:            row.Name,
		DevourerLevel:   int(row.DevourerLevel.Int32),
		FeatTiers:       int(row.FeatTiers.Int32),
		OtherMultiplier: row.OtherMultiplier.Float64,
		GroupBonusCount: int(row.GroupBonusCount.Int32),
		LeftoverShards:  int(row.LeftoverShards.Int32),
		Notes:           row.Notes.String,
	}

	if row.LegendaryCounts.Valid {
		json.Unmarshal(row.LegendaryCounts.RawMessage, &p.LegendaryCounts)
	}
	if row.ExperimentLevels.Valid {
		json.Unmarshal(row.ExperimentLevels.RawMessage, &p.ExperimentLevels)
	}
	if row.PossessedRunes.Valid {
		json.Unmarshal(row.PossessedRunes.RawMessage, &p.PossessedRunes)
	}
	if row.PossessedLegendaries.Valid {
		json.Unmarshal(row.PossessedLegendaries.RawMessage, &p.PossessedLegendaries)
	}

	return p, nil
}

func (s *Store) ListPlans(ctx context.Context) ([]models.Plan, error) {
	rows, err := s.queries.ListPlans(ctx)
	if err != nil {
		return nil, err
	}

	plans := make([]models.Plan, len(rows))
	for i, row := range rows {
		plans[i] = models.Plan{
			ID:   int(row.ID),
			Name: row.Name,
		}
	}
	return plans, nil
}
