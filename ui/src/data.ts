import type { ExperimentID, LegendaryID, RuneType } from './types/game';

export const experimentDescriptions: Record<ExperimentID, string> = {
  seasoning: "Increases the amount of Food from feeding the Devourer.",
  strength: "Increases all Damage dealt to Champions and enemies.",
  taste: "Increases the Food reward for completing Cravings.",
  capacity: "Increases Mana, Slime, and Darkness storage capacity.",
  body_snatcher: "Unlocks The Body Snatcher to repeat Champion fights.",
  weakening: "Reduces the health scale of The Protector.",
  damage_cap: "Increases the damage cap scale for The Mech.",
  ice_chest: "Increases the number of taps for Ice Chests.",
  poison_chest: "Increases the number of taps for Poison Chests.",
  blood_chest: "Increases the number of taps for Blood Chests.",
  moon_chest: "Increases the number of taps for Moon Chests.",
  death_chest: "Increases the number of taps for Death Chests.",
  cosmic_chest: "Increases the number of taps for Cosmic Chests.",
  seasoning_2: "Multiplicative Food boost (Post-100 only).",
  strength_2: "Multiplicative Damage boost (Post-100 only).",
  taste_2: "Multiplicative Craving boost (Post-100 only).",
  capacity_2: "Multiplicative Capacity boost (Post-100 only).",
};

export const experimentMaxLevels: Record<ExperimentID, number> = {
  seasoning: 9,
  strength: 9,
  taste: 9,
  capacity: 9,
  body_snatcher: 1,
  weakening: 8,
  damage_cap: 3,
  ice_chest: 9,
  poison_chest: 9,
  blood_chest: 9,
  moon_chest: 9,
  death_chest: 9,
  cosmic_chest: 9,
  seasoning_2: 9,
  strength_2: 9,
  taste_2: 9,
  capacity_2: 9,
};

export interface LegendaryInfo {
  id: LegendaryID;
  name: string;
  bonus: string;
  subsequent: string;
  max?: number;
}

export const legendaries: LegendaryInfo[] = [
  { id: 'lich', name: 'Lich', bonus: '10%', subsequent: '5%' },
  { id: 'gorgon', name: 'Gorgon', bonus: '10%', subsequent: '5%' },
  { id: 'harpy', name: 'Harpy', bonus: '10%', subsequent: '5%' },
  { id: 'reaper', name: 'Reaper', bonus: '20%', subsequent: '10%' },
  { id: 'cyclops', name: 'Cyclops', bonus: '20%', subsequent: '10%' },
  { id: 'archdemon', name: 'Archdemon', bonus: '20%', subsequent: '10%', max: 4 },
  { id: 'the_cursed', name: 'The Cursed', bonus: '40%', subsequent: '0%', max: 1 },
  { id: 'the_colossus', name: 'The Colossus', bonus: '40%', subsequent: '0%', max: 1 },
  { id: 'the_infernal', name: 'The Infernal', bonus: '40%', subsequent: '0%', max: 1 },
  { id: 'robo_chicken', name: 'Robo Chicken', bonus: '20%', subsequent: '10%', max: 4 },
  { id: 'shield_bot', name: 'Shield Bot', bonus: '30%', subsequent: '15%', max: 4 },
  { id: 'soul_stalker', name: 'Soul Stalker', bonus: '40%', subsequent: '20%' },
];

export const runeTypes: RuneType[] = ['ice', 'poison', 'blood', 'moon', 'death', 'cosmic'];

export const devourerLevels = [
  35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100, 150, 200, 300, 400, 500, 600, 700, 800, 900, 1000
];
