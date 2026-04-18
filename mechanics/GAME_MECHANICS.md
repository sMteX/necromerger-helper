# Game Mechanics: Necromerger Prestige Planner

This document serves as a knowledge base for the game mechanics and planning logic.

## Overview

### Base gameplay facts
- the game is a merger type
- the goal of the game is to "feed The Devourer", which means dragging things to an entity in the middle of the screen
  - this gives The Devourer a certain amount of `Food`, which is used to level up The Devourer (with the Food requirements rising each level)
- user operates on a grid ("lair") of variable size (starts small, grows as The Devourer levels up - thresholds for lair space growth are pre-set)
- the user has 3 types of resources - `Mana`, `Slime` and `Darkness`
- the user can earn another types of resources called `Runes`, which are of different color/type - `Ice` (blue), `Poison` (green), `Blood` (red), `Moon` (yellow), `Death` (dark blue) and `Cosmic` (purple/pink)
- the user can buy `Stations` for runes (types and costs are pre-set)
  - these stations occupy lair space, and do one or more of the following:
    - increase the resource cap
    - can be tapped by the player to use resources to produce creatures ("minions") or items
- there are about 25 creatures in the game, majority of which have levels
  - stations produce items ("parts") that the player merges together to form either higher level parts, or minions
  - merging minions increases their level
  - player can feed almost anything to The Devourer, even parts give Food
  - minions can have multiple of characteristics:
    - they can produce resources (on a set timer, e.g. 10 Mana every 10 seconds)
    - they can be dragged onto enemies to damage/kill them
  - there are minions who are specialized at different things:
    - some are better for damage
    - some are better for resource generation
    - some are better for feeding The Devourer
    - some are completely special and have special effects/abilities (mid-game and later)
- merging certain types of creatures also help summon `Champions`
  - these are special NPCs with various amount of HP that need to be killed (or fed to the Devourer)
  - upon killing a Champion, it drops a `Chest` of certain type
  - this Chest can be then tapped by the player to produce `Rune Piles`
    - detailed info can be found in [CHESTS.md](./CHESTS.md)
- stations also have levels and can be merged to increase their level
  - higher level stations produce higher level parts but also have higher resource cost per tap
- The Devourer also has Cravings:
  - it picks a certain amount of minions it wants the player to feed it
  - when a player completes the Craving, they get a substantial amount of Food as reward
- to unlock new stations and to give player some game goals, there are `Feats`
  - these are like quests for the player to complete
  - they are divided into `Tiers` (e.g. Tier 22 Feats)
  - completing a Feat gives a preset reward
  - completing all Feats in a Tier completes the Tier and unlocks new stations
- to help the player make the game easier, the game introduces `Spells`
  - these are divided into `Spell Pages` which unlock at specific Devourer levels
  - there are 34 spells divided into 6 pages (last page is incomplete yet)
  - each spell has multiple levels with predefined specific Rune costs
  - when player buys a Spell for Runes, it's a passive effect that is in effect for the rest of the game

# Prestiging
- details to this game mechanic is in [PRESTIGE.md](./PRESTIGE.md)

## Game details

### Stations
- not all stations listed here are relevant to the tasks this app is designed for
  - if needed, they will be added later
- almost all stations have 5 levels (with exception of `Soul Grinder` which only has 2 levels)
  - this means that Level 5 = 2x Level 4 = 4x Level 3 = 8x Level 2 = 16x Level 1
- stations:
  - `Grave` - costs 20 Ice Runes
  - `Supply Cupboard` - costs 20 Poison Runes
  - `Foul Chicken` - costs 30 Ice Runes and 15 Poison Runes
  - `Altar` - costs 20 Blood Runes
  - `Lectern` - costs 50 Ice Runes and 20 Moon Runes
  - `Fridge` - costs 50 Poison Runes and 20 Moon Runes
  - `Portal` - costs 30 Blood Runes and 30 Death Runes
  - `Crashed Saucer` - costs 20 Cosmic Runes
  - `Soul Grinder` - costs 200 Death Runes and 200 Cosmic Runes

### Legendaries
- here are again the individual legendaries and how they are created:
- legendaries:
  - `Group 1`:
    - `Lich` - created by merging 2x Level 5 `Grave`, gives back 1x Level 1 `Grave` 
    - `Gorgon` - created by merging 2x Level 5 `Supply Cupboard`, gives back 1x Level 1 `Supply Cupboard`
    - `Harpy` - created by merging 2x Level 5 `Altar`, gives back 1x Level 1 `Altar`
  - `Group 2`:
    - `Reaper` - created by merging 2x Level 5 `Lectern`, gives back 1x Level 1 `Lectern`
    - `Cyclops` - created by merging 2x Level 5 `Fridge`, gives back 1x Level 1 `Fridge`
    - `Archdemon` - created by merging 2x Level 5 `Portal`, gives back 1x Level 1 `Portal`
  - `Group 3`:
    - `The Cursed` - created by merging `Lich` and `Reaper`, **doesn't give back anything**
    - `The Colossus` - created by merging `Gorgon` and `Cyclops`, **doesn't give back anything**
    - `The Infernal` - created by merging `Harpy` and `Archdemon`, **doesn't give back anything**
  - `Group 4`:
    - `Robo Chicken` - created by merging 2x Level 5 `Foul Chicken`, **doesn't give back anything**
    - `Shield Bot` - created by merging 2x Level 5 `Crashed Saucer`, gives back 1x Level 1 `Crashed Saucer`
    - `Soul Grinder` - created by merging 2x Level 5 `Soul Grinder`, gives back 1x Level 1 `Soul Grinder`
- from this, you can deduce the individual cost of these legendaries
