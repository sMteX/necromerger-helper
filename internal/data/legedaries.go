package data

import "github.com/sMteX/necromerger-helper/internal/models"

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

var (
	LegendaryLich        = models.Legendary{ID: models.Lich, Name: "Lich", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05}
	LegendaryGorgon      = models.Legendary{ID: models.Gorgon, Name: "Gorgon", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05}
	LegendaryHarpy       = models.Legendary{ID: models.Harpy, Name: "Harpy", Group: models.Group1, FirstBonus: 0.10, Subsequent: 0.05}
	LegendaryReaper      = models.Legendary{ID: models.Reaper, Name: "Reaper", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10}
	LegendaryCyclops     = models.Legendary{ID: models.Cyclops, Name: "Cyclops", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10}
	LegendaryArchdemon   = models.Legendary{ID: models.Archdemon, Name: "Archdemon", Group: models.Group2, FirstBonus: 0.20, Subsequent: 0.10, MaxInstances: 4}
	LegendaryTheCursed   = models.Legendary{ID: models.TheCursed, Name: "The Cursed", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1}
	LegendaryTheColossus = models.Legendary{ID: models.TheColossus, Name: "The Colossus", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1}
	LegendaryTheInfernal = models.Legendary{ID: models.TheInfernal, Name: "The Infernal", Group: models.Group3, FirstBonus: 0.40, MaxInstances: 1}
	LegendaryRoboChicken = models.Legendary{ID: models.RoboChicken, Name: "Robo Chicken", Group: models.Group4, FirstBonus: 0.20, Subsequent: 0.10, MaxInstances: 4}
	LegendaryShieldBot   = models.Legendary{ID: models.ShieldBot, Name: "Shield Bot", Group: models.Group4, FirstBonus: 0.30, Subsequent: 0.15, MaxInstances: 4}
	LegendarySoulStalker = models.Legendary{ID: models.SoulStalker, Name: "Soul Stalker", Group: models.Group4, FirstBonus: 0.40, Subsequent: 0.20}
)
var Legendaries = []models.Legendary{
	LegendaryLich,
	LegendaryGorgon,
	LegendaryHarpy,
	LegendaryReaper,
	LegendaryCyclops,
	LegendaryArchdemon,
	LegendaryTheCursed,
	LegendaryTheColossus,
	LegendaryTheInfernal,
	LegendaryRoboChicken,
	LegendaryShieldBot,
	LegendarySoulStalker,
}
var LegendariesById = map[models.LegendaryID]models.Legendary{
	models.Lich:        LegendaryLich,
	models.Gorgon:      LegendaryGorgon,
	models.Harpy:       LegendaryHarpy,
	models.Reaper:      LegendaryReaper,
	models.Cyclops:     LegendaryCyclops,
	models.Archdemon:   LegendaryArchdemon,
	models.TheCursed:   LegendaryTheCursed,
	models.TheColossus: LegendaryTheColossus,
	models.TheInfernal: LegendaryTheInfernal,
	models.RoboChicken: LegendaryRoboChicken,
	models.ShieldBot:   LegendaryShieldBot,
	models.SoulStalker: LegendarySoulStalker,
}
