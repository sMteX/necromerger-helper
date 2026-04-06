export type RuneType = 'ice' | 'poison' | 'blood' | 'moon' | 'death' | 'cosmic';

export type LegendaryID = 'lich' | 'gorgon' | 'harpy' | 'reaper' | 'cyclops' | 'archdemon' | 
                          'the_cursed' | 'the_colossus' | 'the_infernal' | 'robo_chicken' | 
                          'shield_bot' | 'soul_stalker';

export type ExperimentID = 'seasoning' | 'strength' | 'taste' | 'capacity' | 'body_snatcher' | 
                           'weakening' | 'damage_cap' | 'ice_chest' | 'poison_chest' | 
                           'blood_chest' | 'moon_chest' | 'death_chest' | 'cosmic_chest' | 
                           'seasoning_2' | 'strength_2' | 'taste_2' | 'capacity_2';

export interface Plan {
  id: number;
  name: string;
  devourerLevel: number;
  featTiers: number;
  otherMultiplier: number;
  groupBonusCount: number;
  leftoverShards: number;
  legendaryCounts: Record<LegendaryID, number>;
  experimentLevels: Record<ExperimentID, number>;
  possessedRunes: Record<RuneType, number>;
  possessedLegendaries: Record<LegendaryID, number>;
  notes: string;
}

export interface PlanSummary {
  id: number;
  name: string;
}

export interface ExperimentSummary {
  id: ExperimentID;
  currentLevel: number;
  currentLevelValue: string;
  currentLevelCost: string;
  nextLevelCost: string;
  nextLevelValue: string;
  maxLevel: boolean;
}

export interface CalculationResponse {
  totalShards: number;
  baseShards: number;
  featMultiplier: number;
  legendMultiplier: number;
  otherMultiplier: number;
  experimentCost: number;
  remaining: number;
  experiments: Record<ExperimentID, ExperimentSummary>;
  runeTotal: Record<RuneType, number>;
  runeNeeded: Record<RuneType, number>;
  legendaryRunes: Record<LegendaryID, Record<RuneType, number>>;
}
