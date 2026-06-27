package data

import "github.com/sMteX/necromerger-helper/internal/models"

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
