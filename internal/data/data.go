package data

import "github.com/sMteX/necro-prestige-planner/internal/models"

// TODO: deduplicate
var Experiments = []models.Experiment{
	{
		ID: models.ExpSeasoning, Name: "Seasoning Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25, Value: 0.10, PrevValue: 0.0},
			{Level: 2, Cost: 250, Value: 0.20, PrevValue: 0.10},
			{Level: 3, Cost: 750, Value: 0.30, PrevValue: 0.20},
			{Level: 4, Cost: 1500, Value: 0.40, PrevValue: 0.30},
			{Level: 5, Cost: 2500, Value: 0.50, PrevValue: 0.40},
			{Level: 6, Cost: 10000, Value: 0.60, PrevValue: 0.50},
			{Level: 7, Cost: 25000, Value: 0.80, PrevValue: 0.60},
			{Level: 8, Cost: 75000, Value: 1.00, PrevValue: 0.80},
			{Level: 9, Cost: 250000, Value: 1.50, PrevValue: 1.00},
		},
	},
	{
		ID: models.ExpStrength, Name: "Strength Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25, Value: 0.10, PrevValue: 0.0},
			{Level: 2, Cost: 250, Value: 0.20, PrevValue: 0.10},
			{Level: 3, Cost: 750, Value: 0.30, PrevValue: 0.20},
			{Level: 4, Cost: 1500, Value: 0.40, PrevValue: 0.30},
			{Level: 5, Cost: 2500, Value: 0.50, PrevValue: 0.40},
			{Level: 6, Cost: 10000, Value: 0.60, PrevValue: 0.50},
			{Level: 7, Cost: 25000, Value: 0.80, PrevValue: 0.60},
			{Level: 8, Cost: 75000, Value: 1.00, PrevValue: 0.80},
			{Level: 9, Cost: 250000, Value: 1.50, PrevValue: 1.00},
		},
	},
	{
		ID: models.ExpTaste, Name: "Taste Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 50, Value: 0.20, PrevValue: 0.0},
			{Level: 2, Cost: 500, Value: 0.40, PrevValue: 0.20},
			{Level: 3, Cost: 1000, Value: 0.60, PrevValue: 0.40},
			{Level: 4, Cost: 2000, Value: 0.80, PrevValue: 0.60},
			{Level: 5, Cost: 4000, Value: 1.00, PrevValue: 0.80},
			{Level: 6, Cost: 20000, Value: 1.25, PrevValue: 1.00},
			{Level: 7, Cost: 50000, Value: 1.50, PrevValue: 1.25},
			{Level: 8, Cost: 100000, Value: 2.00, PrevValue: 1.50},
			{Level: 9, Cost: 500000, Value: 3.00, PrevValue: 2.00},
		},
	},
	{
		ID: models.ExpCapacity, Name: "Capacity Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 100, Value: 0.05, PrevValue: 0.0},
			{Level: 2, Cost: 750, Value: 0.10, PrevValue: 0.05},
			{Level: 3, Cost: 1500, Value: 0.15, PrevValue: 0.10},
			{Level: 4, Cost: 2500, Value: 0.20, PrevValue: 0.15},
			{Level: 5, Cost: 5000, Value: 0.25, PrevValue: 0.20},
			{Level: 6, Cost: 25000, Value: 0.30, PrevValue: 0.25},
			{Level: 7, Cost: 75000, Value: 0.35, PrevValue: 0.30},
			{Level: 8, Cost: 150000, Value: 0.40, PrevValue: 0.35},
			{Level: 9, Cost: 750000, Value: 0.50, PrevValue: 0.40},
		},
	},
	{
		ID: models.ExpBodySnatcher, Name: "Body Snatcher", Tier: models.TierPre100, IsSpecial: true,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 50, Value: 1.0, PrevValue: 1.0},
		},
	},
	{
		ID: models.ExpWeakening, Name: "Weakening Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 750, Value: 2.50, PrevValue: 5.00},
			{Level: 2, Cost: 15000, Value: 1.75, PrevValue: 2.50},
			{Level: 3, Cost: 25000, Value: 1.50, PrevValue: 1.75},
			{Level: 4, Cost: 50000, Value: 1.40, PrevValue: 1.50},
			{Level: 5, Cost: 500000, Value: 1.35, PrevValue: 1.40},
			{Level: 6, Cost: 5000000, Value: 1.30, PrevValue: 1.35},
			{Level: 7, Cost: 25000000, Value: 1.27, PrevValue: 1.30},
			{Level: 8, Cost: 100000000, Value: 1.25, PrevValue: 1.27},
		},
	},
	{
		ID: models.ExpDamageCap, Name: "Damage Cap Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25000, Value: 0.65, PrevValue: 0.50},
			{Level: 2, Cost: 1000000, Value: 0.75, PrevValue: 0.65},
			{Level: 3, Cost: 50000000, Value: 0.80, PrevValue: 0.75},
		},
	},
	{
		ID: models.ExpIceChest, Name: "Ice Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 10, Value: 1, PrevValue: 0}, {Level: 2, Cost: 500, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 1500, Value: 3, PrevValue: 2}, {Level: 4, Cost: 5000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 10000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 50000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 100000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 500000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 1000000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpPoisonChest, Name: "Poison Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25, Value: 1, PrevValue: 0}, {Level: 2, Cost: 1000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 3000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 7500, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 12500, Value: 5, PrevValue: 4}, {Level: 6, Cost: 75000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 150000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 500000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 1500000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpBloodChest, Name: "Blood Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 50, Value: 1, PrevValue: 0}, {Level: 2, Cost: 2500, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 5000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 10000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 50000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 100000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 250000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 750000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 2500000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpMoonChest, Name: "Moon Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 250, Value: 1, PrevValue: 0}, {Level: 2, Cost: 3000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 7500, Value: 3, PrevValue: 2}, {Level: 4, Cost: 15000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 75000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 150000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 500000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 1000000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 5000000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpDeathChest, Name: "Death Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 500, Value: 1, PrevValue: 0}, {Level: 2, Cost: 5000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 10000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 50000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 150000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 500000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 1000000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 5000000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 10000000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpCosmicChest, Name: "Cosmic Chest Experiment", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 5000, Value: 1, PrevValue: 0}, {Level: 2, Cost: 10000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 15000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 75000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 250000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 1000000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 5000000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 10000000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 25000000, Value: 9, PrevValue: 8},
		},
	},
	{
		ID: models.ExpSeasoning2, Name: "Seasoning Experiment II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 100000, Value: 5, PrevValue: 1},
			{Level: 2, Cost: 500000, Value: 10, PrevValue: 5},
			{Level: 3, Cost: 1000000, Value: 15, PrevValue: 10},
			{Level: 4, Cost: 2000000, Value: 20, PrevValue: 15},
			{Level: 5, Cost: 3000000, Value: 25, PrevValue: 20},
			{Level: 6, Cost: 5000000, Value: 30, PrevValue: 25},
			{Level: 7, Cost: 10000000, Value: 35, PrevValue: 30},
			{Level: 8, Cost: 20000000, Value: 40, PrevValue: 35},
			{Level: 9, Cost: 50000000, Value: 50, PrevValue: 40},
		},
	},
	{
		ID: models.ExpStrength2, Name: "Strength Experiment II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 100000, Value: 1.2, PrevValue: 1},
			{Level: 2, Cost: 500000, Value: 1.4, PrevValue: 1.2},
			{Level: 3, Cost: 1000000, Value: 1.6, PrevValue: 1.4},
			{Level: 4, Cost: 2000000, Value: 1.8, PrevValue: 1.6},
			{Level: 5, Cost: 3000000, Value: 2.0, PrevValue: 1.8},
			{Level: 6, Cost: 5000000, Value: 2.2, PrevValue: 2.0},
			{Level: 7, Cost: 10000000, Value: 2.4, PrevValue: 2.2},
			{Level: 8, Cost: 20000000, Value: 2.6, PrevValue: 2.4},
			{Level: 9, Cost: 50000000, Value: 3.0, PrevValue: 2.6},
		},
	},
	{
		ID: models.ExpTaste2, Name: "Taste Experiment II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 200000, Value: 2, PrevValue: 1},
			{Level: 2, Cost: 1000000, Value: 3, PrevValue: 2},
			{Level: 3, Cost: 3000000, Value: 4, PrevValue: 3},
			{Level: 4, Cost: 5000000, Value: 5, PrevValue: 4},
			{Level: 5, Cost: 10000000, Value: 6, PrevValue: 5},
			{Level: 6, Cost: 20000000, Value: 7, PrevValue: 6},
			{Level: 7, Cost: 30000000, Value: 8, PrevValue: 7},
			{Level: 8, Cost: 50000000, Value: 9, PrevValue: 8},
			{Level: 9, Cost: 100000000, Value: 10, PrevValue: 9},
		},
	},
	{
		ID: models.ExpCapacity2, Name: "Capacity Experiment II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 200000, Value: 1.1, PrevValue: 1},
			{Level: 2, Cost: 1000000, Value: 1.2, PrevValue: 1.1},
			{Level: 3, Cost: 3000000, Value: 1.3, PrevValue: 1.2},
			{Level: 4, Cost: 5000000, Value: 1.4, PrevValue: 1.3},
			{Level: 5, Cost: 10000000, Value: 1.5, PrevValue: 1.4},
			{Level: 6, Cost: 20000000, Value: 1.6, PrevValue: 1.5},
			{Level: 7, Cost: 30000000, Value: 1.8, PrevValue: 1.6},
			{Level: 8, Cost: 50000000, Value: 1.9, PrevValue: 1.8},
			{Level: 9, Cost: 100000000, Value: 2, PrevValue: 1.9},
		},
	},
}

var ExperimentsById = map[models.ExperimentID]models.Experiment{

	models.ExpSeasoning: {
		ID: models.ExpSeasoning, Name: "Seasoning I", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25, Value: 0.10, PrevValue: 0.0},
			{Level: 2, Cost: 250, Value: 0.20, PrevValue: 0.10},
			{Level: 3, Cost: 750, Value: 0.30, PrevValue: 0.20},
			{Level: 4, Cost: 1500, Value: 0.40, PrevValue: 0.30},
			{Level: 5, Cost: 2500, Value: 0.50, PrevValue: 0.40},
			{Level: 6, Cost: 10000, Value: 0.60, PrevValue: 0.50},
			{Level: 7, Cost: 25000, Value: 0.80, PrevValue: 0.60},
			{Level: 8, Cost: 75000, Value: 1.00, PrevValue: 0.80},
			{Level: 9, Cost: 250000, Value: 1.50, PrevValue: 1.00},
		},
	},
	models.ExpStrength: {
		ID: models.ExpStrength, Name: "Strength I", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25, Value: 0.10, PrevValue: 0.0},
			{Level: 2, Cost: 250, Value: 0.20, PrevValue: 0.10},
			{Level: 3, Cost: 750, Value: 0.30, PrevValue: 0.20},
			{Level: 4, Cost: 1500, Value: 0.40, PrevValue: 0.30},
			{Level: 5, Cost: 2500, Value: 0.50, PrevValue: 0.40},
			{Level: 6, Cost: 10000, Value: 0.60, PrevValue: 0.50},
			{Level: 7, Cost: 25000, Value: 0.80, PrevValue: 0.60},
			{Level: 8, Cost: 75000, Value: 1.00, PrevValue: 0.80},
			{Level: 9, Cost: 250000, Value: 1.50, PrevValue: 1.00},
		},
	},
	models.ExpTaste: {
		ID: models.ExpTaste, Name: "Taste I", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 50, Value: 0.20, PrevValue: 0.0},
			{Level: 2, Cost: 500, Value: 0.40, PrevValue: 0.20},
			{Level: 3, Cost: 1000, Value: 0.60, PrevValue: 0.40},
			{Level: 4, Cost: 2000, Value: 0.80, PrevValue: 0.60},
			{Level: 5, Cost: 4000, Value: 1.00, PrevValue: 0.80},
			{Level: 6, Cost: 20000, Value: 1.25, PrevValue: 1.00},
			{Level: 7, Cost: 50000, Value: 1.50, PrevValue: 1.25},
			{Level: 8, Cost: 100000, Value: 2.00, PrevValue: 1.50},
			{Level: 9, Cost: 500000, Value: 3.00, PrevValue: 2.00},
		},
	},
	models.ExpCapacity: {
		ID: models.ExpCapacity, Name: "Capacity I", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 100, Value: 0.05, PrevValue: 0.0},
			{Level: 2, Cost: 750, Value: 0.10, PrevValue: 0.05},
			{Level: 3, Cost: 1500, Value: 0.15, PrevValue: 0.10},
			{Level: 4, Cost: 2500, Value: 0.20, PrevValue: 0.15},
			{Level: 5, Cost: 5000, Value: 0.25, PrevValue: 0.20},
			{Level: 6, Cost: 25000, Value: 0.30, PrevValue: 0.25},
			{Level: 7, Cost: 75000, Value: 0.35, PrevValue: 0.30},
			{Level: 8, Cost: 150000, Value: 0.40, PrevValue: 0.35},
			{Level: 9, Cost: 750000, Value: 0.50, PrevValue: 0.40},
		},
	},
	models.ExpBodySnatcher: {
		ID: models.ExpBodySnatcher, Name: "Body Snatcher", Tier: models.TierPre100, IsSpecial: true,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 50, Value: 1.0, PrevValue: 0},
		},
	},
	models.ExpWeakening: {
		ID: models.ExpWeakening, Name: "Weakening", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 750, Value: 2.50, PrevValue: 5.00},
			{Level: 2, Cost: 15000, Value: 1.75, PrevValue: 2.50},
			{Level: 3, Cost: 25000, Value: 1.50, PrevValue: 1.75},
			{Level: 4, Cost: 50000, Value: 1.40, PrevValue: 1.50},
			{Level: 5, Cost: 500000, Value: 1.35, PrevValue: 1.40},
			{Level: 6, Cost: 5000000, Value: 1.30, PrevValue: 1.35},
			{Level: 7, Cost: 25000000, Value: 1.27, PrevValue: 1.30},
			{Level: 8, Cost: 100000000, Value: 1.25, PrevValue: 1.27},
		},
	},
	models.ExpDamageCap: {
		ID: models.ExpDamageCap, Name: "Mech Damage Cap", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25000, Value: 0.65, PrevValue: 0.50},
			{Level: 2, Cost: 1000000, Value: 0.75, PrevValue: 0.65},
			{Level: 3, Cost: 50000000, Value: 0.80, PrevValue: 0.75},
		},
	},
	models.ExpIceChest: {
		ID: models.ExpIceChest, Name: "Ice Chest", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 10, Value: 1, PrevValue: 0}, {Level: 2, Cost: 500, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 1500, Value: 3, PrevValue: 2}, {Level: 4, Cost: 5000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 10000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 50000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 100000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 500000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 1000000, Value: 9, PrevValue: 8},
		},
	},
	models.ExpPoisonChest: {
		ID: models.ExpPoisonChest, Name: "Poison Chest", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 25, Value: 1, PrevValue: 0}, {Level: 2, Cost: 1000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 3000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 7500, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 12500, Value: 5, PrevValue: 4}, {Level: 6, Cost: 75000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 150000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 500000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 1500000, Value: 9, PrevValue: 8},
		},
	},
	models.ExpBloodChest: {
		ID: models.ExpBloodChest, Name: "Blood Chest", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 50, Value: 1, PrevValue: 0}, {Level: 2, Cost: 2500, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 5000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 10000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 50000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 100000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 250000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 750000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 2500000, Value: 9, PrevValue: 8},
		},
	},
	models.ExpMoonChest: {
		ID: models.ExpMoonChest, Name: "Moon Chest", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 250, Value: 1, PrevValue: 0}, {Level: 2, Cost: 3000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 7500, Value: 3, PrevValue: 2}, {Level: 4, Cost: 15000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 75000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 150000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 500000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 1000000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 5000000, Value: 9, PrevValue: 8},
		},
	},
	models.ExpDeathChest: {
		ID: models.ExpDeathChest, Name: "Death Chest", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 500, Value: 1, PrevValue: 0}, {Level: 2, Cost: 5000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 10000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 50000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 150000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 500000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 1000000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 5000000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 10000000, Value: 9, PrevValue: 8},
		},
	},
	models.ExpCosmicChest: {
		ID: models.ExpCosmicChest, Name: "Cosmic Chest", Tier: models.TierPre100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 5000, Value: 1, PrevValue: 0}, {Level: 2, Cost: 10000, Value: 2, PrevValue: 1},
			{Level: 3, Cost: 15000, Value: 3, PrevValue: 2}, {Level: 4, Cost: 75000, Value: 4, PrevValue: 3},
			{Level: 5, Cost: 250000, Value: 5, PrevValue: 4}, {Level: 6, Cost: 1000000, Value: 6, PrevValue: 5},
			{Level: 7, Cost: 5000000, Value: 7, PrevValue: 6}, {Level: 8, Cost: 10000000, Value: 8, PrevValue: 7},
			{Level: 9, Cost: 25000000, Value: 9, PrevValue: 8},
		},
	},
	models.ExpSeasoning2: {
		ID: models.ExpSeasoning2, Name: "Seasoning II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 100000, Value: 5, PrevValue: 1},
			{Level: 2, Cost: 500000, Value: 10, PrevValue: 5},
			{Level: 3, Cost: 1000000, Value: 15, PrevValue: 10},
			{Level: 4, Cost: 2000000, Value: 20, PrevValue: 15},
			{Level: 5, Cost: 3000000, Value: 25, PrevValue: 20},
			{Level: 6, Cost: 5000000, Value: 30, PrevValue: 25},
			{Level: 7, Cost: 10000000, Value: 35, PrevValue: 30},
			{Level: 8, Cost: 20000000, Value: 40, PrevValue: 35},
			{Level: 9, Cost: 50000000, Value: 50, PrevValue: 40},
		},
	},
	models.ExpStrength2: {
		ID: models.ExpStrength2, Name: "Strength II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 100000, Value: 1.2, PrevValue: 1},
			{Level: 2, Cost: 500000, Value: 1.4, PrevValue: 1.2},
			{Level: 3, Cost: 1000000, Value: 1.6, PrevValue: 1.4},
			{Level: 4, Cost: 2000000, Value: 1.8, PrevValue: 1.6},
			{Level: 5, Cost: 3000000, Value: 2.0, PrevValue: 1.8},
			{Level: 6, Cost: 5000000, Value: 2.2, PrevValue: 2.0},
			{Level: 7, Cost: 10000000, Value: 2.4, PrevValue: 2.2},
			{Level: 8, Cost: 20000000, Value: 2.6, PrevValue: 2.4},
			{Level: 9, Cost: 50000000, Value: 3.0, PrevValue: 2.6},
		},
	},
	models.ExpTaste2: {
		ID: models.ExpTaste2, Name: "Taste II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 200000, Value: 2, PrevValue: 1},
			{Level: 2, Cost: 1000000, Value: 3, PrevValue: 2},
			{Level: 3, Cost: 3000000, Value: 4, PrevValue: 3},
			{Level: 4, Cost: 5000000, Value: 5, PrevValue: 4},
			{Level: 5, Cost: 10000000, Value: 6, PrevValue: 5},
			{Level: 6, Cost: 20000000, Value: 7, PrevValue: 6},
			{Level: 7, Cost: 30000000, Value: 8, PrevValue: 7},
			{Level: 8, Cost: 50000000, Value: 9, PrevValue: 8},
			{Level: 9, Cost: 100000000, Value: 10, PrevValue: 9},
		},
	},
	models.ExpCapacity2: {
		ID: models.ExpCapacity2, Name: "Capacity II", Tier: models.TierPost100,
		Levels: []models.ExperimentLevel{
			{Level: 1, Cost: 200000, Value: 1.1, PrevValue: 1},
			{Level: 2, Cost: 1000000, Value: 1.2, PrevValue: 1.1},
			{Level: 3, Cost: 3000000, Value: 1.3, PrevValue: 1.2},
			{Level: 4, Cost: 5000000, Value: 1.4, PrevValue: 1.3},
			{Level: 5, Cost: 10000000, Value: 1.5, PrevValue: 1.4},
			{Level: 6, Cost: 20000000, Value: 1.6, PrevValue: 1.5},
			{Level: 7, Cost: 30000000, Value: 1.8, PrevValue: 1.6},
			{Level: 8, Cost: 50000000, Value: 1.9, PrevValue: 1.8},
			{Level: 9, Cost: 100000000, Value: 2, PrevValue: 1.9},
		},
	},
}
var LegendaryRecipes = map[models.LegendaryID]models.LegendaryRecipe{
	models.Lich:        {StationID: models.StationGrave, Levels: 5, ReturnsL1: true},
	models.Gorgon:      {StationID: models.StationSupplyCupboard, Levels: 5, ReturnsL1: true},
	models.Harpy:       {StationID: models.StationAltar, Levels: 5, ReturnsL1: true},
	models.Reaper:      {StationID: models.StationLectern, Levels: 5, ReturnsL1: true},
	models.Cyclops:     {StationID: models.StationFridge, Levels: 5, ReturnsL1: true},
	models.Archdemon:   {StationID: models.StationPortal, Levels: 5, ReturnsL1: true},
	models.RoboChicken: {StationID: models.StationFoulChicken, Levels: 5, ReturnsL1: false},
	models.ShieldBot:   {StationID: models.StationCrashedSaucer, Levels: 5, ReturnsL1: true},
	models.SoulStalker: {StationID: models.StationSoulGrinder, Levels: 2, ReturnsL1: true},

	models.TheCursed:   {Requires: []models.LegendaryID{models.Lich, models.Reaper}},
	models.TheColossus: {Requires: []models.LegendaryID{models.Gorgon, models.Cyclops}},
	models.TheInfernal: {Requires: []models.LegendaryID{models.Harpy, models.Archdemon}},
}

var StationCosts = map[models.StationID]models.RuneCosts{
	models.StationGrave:          {models.RuneIce: 20},
	models.StationSupplyCupboard: {models.RunePoison: 20},
	models.StationFoulChicken:    {models.RuneIce: 30, models.RunePoison: 15},
	models.StationAltar:          {models.RuneBlood: 20},
	models.StationLectern:        {models.RuneIce: 50, models.RuneMoon: 20},
	models.StationFridge:         {models.RunePoison: 50, models.RuneMoon: 20},
	models.StationPortal:         {models.RuneBlood: 30, models.RuneDeath: 30},
	models.StationCrashedSaucer:  {models.RuneCosmic: 20},
	models.StationSoulGrinder:    {models.RuneDeath: 200, models.RuneCosmic: 200},
}

var DevourerBaseShards = map[int]int{
	35: 150, 40: 275, 45: 500, 50: 750, 55: 1000, 60: 1500, 65: 2000, 70: 3250,
	75: 4500, 80: 5750, 85: 7500, 90: 10000, 95: 12500, 100: 15000, 150: 40000,
	200: 65000, 300: 150000, 400: 275000, 500: 450000, 600: 700000, 700: 1050000,
	800: 1550000, 900: 2250000, 1000: 3250000,
}

// TODO: cleanup, duplicated for now
var Legendaries = []models.Legendary{
	{ID: models.Lich, Name: "Lich", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05},
	{ID: models.Gorgon, Name: "Gorgon", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05},
	{ID: models.Harpy, Name: "Harpy", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05},
	{ID: models.Reaper, Name: "Reaper", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10},
	{ID: models.Cyclops, Name: "Cyclops", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10},
	{ID: models.Archdemon, Name: "Archdemon", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10, MaxInstances: 4},
	{ID: models.TheCursed, Name: "The Cursed", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1},
	{ID: models.TheColossus, Name: "The Colossus", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1},
	{ID: models.TheInfernal, Name: "The Infernal", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1},
	{ID: models.RoboChicken, Name: "Robo Chicken", Group: models.Group4, FirstBonus: 0.20, Subsequent: 0.10, MaxInstances: 4},
	{ID: models.ShieldBot, Name: "Shield Bot", Group: models.Group4, FirstBonus: 0.30, Subsequent: 0.15, MaxInstances: 4},
	{ID: models.SoulStalker, Name: "Soul Stalker", Group: models.Group4, FirstBonus: 0.40, Subsequent: 0.20},
}
var LegendariesById = map[models.LegendaryID]models.Legendary{
	models.Lich:        {ID: models.Lich, Name: "Lich", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05},
	models.Gorgon:      {ID: models.Gorgon, Name: "Gorgon", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05},
	models.Harpy:       {ID: models.Harpy, Name: "Harpy", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05},
	models.Reaper:      {ID: models.Reaper, Name: "Reaper", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10},
	models.Cyclops:     {ID: models.Cyclops, Name: "Cyclops", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10},
	models.Archdemon:   {ID: models.Archdemon, Name: "Archdemon", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10, MaxInstances: 4},
	models.TheCursed:   {ID: models.TheCursed, Name: "The Cursed", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1},
	models.TheColossus: {ID: models.TheColossus, Name: "The Colossus", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1},
	models.TheInfernal: {ID: models.TheInfernal, Name: "The Infernal", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1},
	models.RoboChicken: {ID: models.RoboChicken, Name: "Robo Chicken", Group: models.Group4, FirstBonus: 0.20, Subsequent: 0.10, MaxInstances: 4},
	models.ShieldBot:   {ID: models.ShieldBot, Name: "Shield Bot", Group: models.Group4, FirstBonus: 0.30, Subsequent: 0.15, MaxInstances: 4},
	models.SoulStalker: {ID: models.SoulStalker, Name: "Soul Stalker", Group: models.Group4, FirstBonus: 0.40, Subsequent: 0.20},
}
