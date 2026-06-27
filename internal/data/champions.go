package data

import "github.com/sMteX/necromerger-helper/internal/models"

// ── Creatures ────────────────────────────────────────────────────────────────

var Creatures = map[models.CreatureID]models.CreatureType{
	models.CreatureSkeleton: {
		ID:          models.CreatureSkeleton,
		Name:        "Skeleton",
		MergePoints: []int{1000, 1200, 1500, 2000, 2500, 3000, 4000},
	},
	models.CreatureZombie: {
		ID:          models.CreatureZombie,
		Name:        "Zombie",
		MergePoints: []int{3000, 3500, 4000, 4500, 5000, 6000},
	},
	models.CreatureMummy: {
		ID:          models.CreatureMummy,
		Name:        "Mummy",
		MergePoints: []int{4000, 4500, 5000, 5500, 6000},
	},
	models.CreatureEyeMonster: {
		ID:          models.CreatureEyeMonster,
		Name:        "Eye Monster",
		MergePoints: []int{1250, 1500, 2000, 2500, 3000, 3500},
	},
	models.CreatureWerewolf: {
		ID:          models.CreatureWerewolf,
		Name:        "Werewolf",
		MergePoints: []int{1750, 2000, 2500, 3000, 4000},
	},
	models.CreatureBat: {
		ID:          models.CreatureBat,
		Name:        "Bat",
		MergePoints: []int{4000, 4500, 5000, 6000},
	},
	models.CreatureShade: {
		ID:          models.CreatureShade,
		Name:        "Shade",
		MergePoints: []int{1500, 2000, 2500, 3000},
	},
	models.CreatureBanshee: {
		ID:          models.CreatureBanshee,
		Name:        "Banshee",
		MergePoints: []int{1500, 2000, 2500, 3000},
	},
	models.CreatureGhoul: {
		ID:          models.CreatureGhoul,
		Name:        "Ghoul",
		MergePoints: []int{1000, 1200, 1500, 2000},
	},
	models.CreatureImp: {
		ID:          models.CreatureImp,
		Name:        "Imp",
		MergePoints: []int{1300, 1600, 1900, 2200},
	},
	models.CreatureAbomination: {
		ID:          models.CreatureAbomination,
		Name:        "Abomination",
		MergePoints: []int{1000, 1200, 1500},
	},
	models.CreatureDemon: {
		ID:          models.CreatureDemon,
		Name:        "Demon",
		MergePoints: []int{1300, 1600, 1900},
	},
}

// ── Champions ─────────────────────────────────────────────────────────────────

var Champions = []models.ChampionType{
	{
		ID:        models.ChampionPeasant,
		Name:      "Peasant",
		Threshold: 150_000,
		Creatures: []models.CreatureID{models.CreatureSkeleton, models.CreatureZombie, models.CreatureMummy},
	},
	{
		ID:        models.ChampionKnight,
		Name:      "Knight",
		Threshold: 200_000,
		Creatures: []models.CreatureID{models.CreatureEyeMonster, models.CreatureWerewolf},
	},
	{
		ID:        models.ChampionCleric,
		Name:      "Cleric",
		Threshold: 250_000,
		Creatures: []models.CreatureID{models.CreatureBat, models.CreatureShade},
	},
	{
		ID:        models.ChampionPaladin,
		Name:      "Paladin",
		Threshold: 200_000,
		Creatures: []models.CreatureID{models.CreatureBanshee, models.CreatureGhoul, models.CreatureImp},
	},
	{
		ID:        models.ChampionRival,
		Name:      "Rival",
		Threshold: 100_000,
		Creatures: []models.CreatureID{models.CreatureAbomination, models.CreatureDemon},
	},
}

var ChampionsById = map[models.ChampionID]models.ChampionType{
	models.ChampionPeasant: Champions[0],
	models.ChampionKnight:  Champions[1],
	models.ChampionCleric:  Champions[2],
	models.ChampionPaladin: Champions[3],
	models.ChampionRival:   Champions[4],
}

// ── Summoning stations ────────────────────────────────────────────────────────

var SummoningStations = []models.SummoningStation{
	{
		ID:       models.StationGrave,
		Name:     "Grave",
		Resource: models.ResourceMana,
		Variants: []models.StationVariant{
			{Level: 1, Hacked: false, Cost: 250, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelOne, Probability: 1.00},
			}},
			{Level: 2, Hacked: false, Cost: 300, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelOne, Probability: 0.60},
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 0.40},
			}},
			{Level: 2, Hacked: true, Cost: 400, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 1.00},
			}},
			{Level: 3, Hacked: false, Cost: 400, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelOne, Probability: 0.40},
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 0.30},
				{CreatureID: models.CreatureZombie, PartLevel: models.PartLevelOne, Probability: 0.30},
			}},
			{Level: 4, Hacked: false, Cost: 500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelOne, Probability: 0.25},
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 0.25},
				{CreatureID: models.CreatureZombie, PartLevel: models.PartLevelOne, Probability: 0.30},
				{CreatureID: models.CreatureZombie, PartLevel: models.PartLevelTwo, Probability: 0.20},
			}},
			{Level: 5, Hacked: false, Cost: 700, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelOne, Probability: 0.10},
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureZombie, PartLevel: models.PartLevelOne, Probability: 0.40},
				{CreatureID: models.CreatureZombie, PartLevel: models.PartLevelTwo, Probability: 0.40},
			}},
			{Level: 5, Hacked: true, Cost: 1000, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelOne, Probability: 0.10},
				{CreatureID: models.CreatureZombie, PartLevel: models.PartLevelTwo, Probability: 0.90},
			}},
		},
	},
	{
		ID:       models.StationSupplyCupboard,
		Name:     "Supply Cupboard",
		Resource: models.ResourceSlime,
		Variants: []models.StationVariant{
			{Level: 1, Hacked: false, Cost: 250, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelOne, Probability: 1.00},
			}},
			{Level: 2, Hacked: false, Cost: 300, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelOne, Probability: 0.60},
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 0.40},
			}},
			{Level: 2, Hacked: true, Cost: 400, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 1.00},
			}},
			{Level: 3, Hacked: false, Cost: 400, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelOne, Probability: 0.40},
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 0.30},
				{CreatureID: models.CreatureMummy, PartLevel: models.PartLevelOne, Probability: 0.30},
			}},
			{Level: 4, Hacked: false, Cost: 500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelOne, Probability: 0.25},
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 0.25},
				{CreatureID: models.CreatureMummy, PartLevel: models.PartLevelOne, Probability: 0.30},
				{CreatureID: models.CreatureMummy, PartLevel: models.PartLevelTwo, Probability: 0.20},
			}},
			{Level: 5, Hacked: false, Cost: 700, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelOne, Probability: 0.10},
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureMummy, PartLevel: models.PartLevelOne, Probability: 0.40},
				{CreatureID: models.CreatureMummy, PartLevel: models.PartLevelTwo, Probability: 0.40},
			}},
			{Level: 5, Hacked: true, Cost: 1000, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelOne, Probability: 0.10},
				{CreatureID: models.CreatureMummy, PartLevel: models.PartLevelTwo, Probability: 0.90},
			}},
		},
	},
	{
		ID:       models.StationAltar,
		Name:     "Altar",
		Resource: models.ResourceDarkness,
		Variants: []models.StationVariant{
			{Level: 1, Hacked: false, Cost: 250, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelOne, Probability: 1.00},
			}},
			{Level: 2, Hacked: false, Cost: 300, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelOne, Probability: 0.60},
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 0.40},
			}},
			{Level: 2, Hacked: true, Cost: 400, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 1.00},
			}},
			{Level: 3, Hacked: false, Cost: 400, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelOne, Probability: 0.40},
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 0.30},
				{CreatureID: models.CreatureBat, PartLevel: models.PartLevelOne, Probability: 0.30},
			}},
			{Level: 4, Hacked: false, Cost: 500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelOne, Probability: 0.25},
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 0.25},
				{CreatureID: models.CreatureBat, PartLevel: models.PartLevelOne, Probability: 0.30},
				{CreatureID: models.CreatureBat, PartLevel: models.PartLevelTwo, Probability: 0.20},
			}},
			{Level: 5, Hacked: false, Cost: 700, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelOne, Probability: 0.10},
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureBat, PartLevel: models.PartLevelOne, Probability: 0.40},
				{CreatureID: models.CreatureBat, PartLevel: models.PartLevelTwo, Probability: 0.40},
			}},
			{Level: 5, Hacked: true, Cost: 1000, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelOne, Probability: 0.10},
				{CreatureID: models.CreatureBat, PartLevel: models.PartLevelTwo, Probability: 0.90},
			}},
		},
	},
	{
		ID:       models.StationLectern,
		Name:     "Lectern",
		Resource: models.ResourceMana,
		Variants: []models.StationVariant{
			{Level: 1, Hacked: false, Cost: 500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 0.70},
				{CreatureID: models.CreatureShade, PartLevel: models.PartLevelTwo, Probability: 0.30},
			}},
			{Level: 2, Hacked: false, Cost: 750, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 0.40},
				{CreatureID: models.CreatureShade, PartLevel: models.PartLevelTwo, Probability: 0.60},
			}},
			{Level: 2, Hacked: true, Cost: 1000, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureShade, PartLevel: models.PartLevelTwo, Probability: 0.90},
			}},
			{Level: 3, Hacked: false, Cost: 1000, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 0.30},
				{CreatureID: models.CreatureShade, PartLevel: models.PartLevelTwo, Probability: 0.50},
				{CreatureID: models.CreatureBanshee, PartLevel: models.PartLevelOne, Probability: 0.20},
			}},
			{Level: 4, Hacked: false, Cost: 1250, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 0.20},
				{CreatureID: models.CreatureShade, PartLevel: models.PartLevelTwo, Probability: 0.40},
				{CreatureID: models.CreatureBanshee, PartLevel: models.PartLevelOne, Probability: 0.30},
				{CreatureID: models.CreatureBanshee, PartLevel: models.PartLevelTwo, Probability: 0.10},
			}},
			{Level: 5, Hacked: false, Cost: 1500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureShade, PartLevel: models.PartLevelTwo, Probability: 0.30},
				{CreatureID: models.CreatureBanshee, PartLevel: models.PartLevelOne, Probability: 0.40},
				{CreatureID: models.CreatureBanshee, PartLevel: models.PartLevelTwo, Probability: 0.40},
			}},
			{Level: 5, Hacked: true, Cost: 2500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureSkeleton, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureShade, PartLevel: models.PartLevelTwo, Probability: 0.20},
				{CreatureID: models.CreatureBanshee, PartLevel: models.PartLevelTwo, Probability: 0.70},
			}},
		},
	},
	{
		ID:       models.StationFridge,
		Name:     "Fridge",
		Resource: models.ResourceSlime,
		Variants: []models.StationVariant{
			{Level: 1, Hacked: false, Cost: 500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 0.70},
				{CreatureID: models.CreatureGhoul, PartLevel: models.PartLevelOne, Probability: 0.30},
			}},
			{Level: 2, Hacked: false, Cost: 750, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 0.40},
				{CreatureID: models.CreatureGhoul, PartLevel: models.PartLevelOne, Probability: 0.60},
			}},
			{Level: 3, Hacked: false, Cost: 1000, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 0.30},
				{CreatureID: models.CreatureGhoul, PartLevel: models.PartLevelOne, Probability: 0.40},
				{CreatureID: models.CreatureGhoul, PartLevel: models.PartLevelTwo, Probability: 0.30},
			}},
			{Level: 3, Hacked: true, Cost: 1750, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureGhoul, PartLevel: models.PartLevelTwo, Probability: 0.90},
			}},
			{Level: 4, Hacked: false, Cost: 1250, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 0.20},
				{CreatureID: models.CreatureGhoul, PartLevel: models.PartLevelOne, Probability: 0.30},
				{CreatureID: models.CreatureGhoul, PartLevel: models.PartLevelTwo, Probability: 0.25},
				{CreatureID: models.CreatureAbomination, PartLevel: models.PartLevelOne, Probability: 0.25},
			}},
			{Level: 5, Hacked: false, Cost: 1500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureGhoul, PartLevel: models.PartLevelOne, Probability: 0.20},
				{CreatureID: models.CreatureGhoul, PartLevel: models.PartLevelTwo, Probability: 0.20},
				{CreatureID: models.CreatureAbomination, PartLevel: models.PartLevelOne, Probability: 0.30},
				{CreatureID: models.CreatureAbomination, PartLevel: models.PartLevelTwo, Probability: 0.20},
			}},
			{Level: 5, Hacked: true, Cost: 2500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureEyeMonster, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureGhoul, PartLevel: models.PartLevelTwo, Probability: 0.20},
				{CreatureID: models.CreatureAbomination, PartLevel: models.PartLevelTwo, Probability: 0.70},
			}},
		},
	},
	{
		ID:       models.StationPortal,
		Name:     "Portal",
		Resource: models.ResourceDarkness,
		Variants: []models.StationVariant{
			{Level: 1, Hacked: false, Cost: 500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 0.70},
				{CreatureID: models.CreatureImp, PartLevel: models.PartLevelOne, Probability: 0.30},
			}},
			{Level: 2, Hacked: false, Cost: 750, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 0.40},
				{CreatureID: models.CreatureImp, PartLevel: models.PartLevelOne, Probability: 0.60},
			}},
			{Level: 3, Hacked: false, Cost: 1000, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 0.30},
				{CreatureID: models.CreatureImp, PartLevel: models.PartLevelOne, Probability: 0.40},
				{CreatureID: models.CreatureImp, PartLevel: models.PartLevelTwo, Probability: 0.30},
			}},
			{Level: 3, Hacked: true, Cost: 1750, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureImp, PartLevel: models.PartLevelTwo, Probability: 0.90},
			}},
			{Level: 4, Hacked: false, Cost: 1250, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 0.20},
				{CreatureID: models.CreatureImp, PartLevel: models.PartLevelOne, Probability: 0.30},
				{CreatureID: models.CreatureImp, PartLevel: models.PartLevelTwo, Probability: 0.25},
				{CreatureID: models.CreatureDemon, PartLevel: models.PartLevelOne, Probability: 0.25},
			}},
			{Level: 5, Hacked: false, Cost: 1500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureImp, PartLevel: models.PartLevelOne, Probability: 0.20},
				{CreatureID: models.CreatureImp, PartLevel: models.PartLevelTwo, Probability: 0.20},
				{CreatureID: models.CreatureDemon, PartLevel: models.PartLevelOne, Probability: 0.30},
				{CreatureID: models.CreatureDemon, PartLevel: models.PartLevelTwo, Probability: 0.20},
			}},
			{Level: 5, Hacked: true, Cost: 2500, Drops: []models.StationDrop{
				{CreatureID: models.CreatureWerewolf, PartLevel: models.PartLevelTwo, Probability: 0.10},
				{CreatureID: models.CreatureImp, PartLevel: models.PartLevelTwo, Probability: 0.20},
				{CreatureID: models.CreatureDemon, PartLevel: models.PartLevelTwo, Probability: 0.70},
			}},
		},
	},
}
