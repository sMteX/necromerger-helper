# Resource Cap Calculator

## Goal

Calculate the combined resource storage (Mana + Slime + Darkness) a player currently has or can reach,
in order to complete Feats that require hitting a combined storage threshold. Given a player's boosts
and owned stations, the tool should:

- show current combined storage
- show how far they are from the next Feat threshold
- tell them what stations to buy/upgrade and what rune cost that implies

---

## Feat Thresholds

| Feat tier | Combined storage required |
|-----------|---------------------------|
| 18        | 200000                    |
| 21        | 400000                    |
| 26        | 600000                    |
| 31        | 800000                    |

---

## Inputs

### Stations

For all types, folowing is true:

- you can only buy either Level 1, or Level 6 (as a shortcut, for full price)
- you create higher level stations by merging 2 lower level stations (Level 6 = 2x Level 5 = 4x Level 4 = 8x Level 3 =
  16x Level 2 = 32x Level 1)
- resource cap per level goes down after Level 2 => lower level stations are more cost efficient, but higher level
  stations are space efficient (and space in the lair is a very limited commodity in late game)

Player inputs: how many of each station they own, and at what level.

Resource cap stations and their base contribution to each resource type:

#### Mana - Mana Pool

Costs 10 Ice Runes and 5 Poison Runes per Level 1.

| Level   | Base mana cap |
|---------|---------------|
| Level 1 | 2000          |
| Level 2 | 4000          |
| Level 3 | 6000          |
| Level 4 | 10000         |
| Level 5 | 15000         |
| Level 6 | 20000         |

#### Slime - Slime Vat

Costs 10 Poison Runes and 5 Blood Runes per Level 1.

| Level   | Base slime cap |
|---------|----------------|
| Level 1 | 1500           |
| Level 2 | 3000           |
| Level 3 | 5000           |
| Level 4 | 7500           |
| Level 5 | 10000          |
| Level 6 | 15000          |

#### Darkness - Dark Stores

Costs 10 Blood Runes and 5 Moon Runes per Level 1.

| Level   | Base darkness cap |
|---------|-------------------|
| Level 1 | 1000              |
| Level 2 | 2000              |
| Level 3 | 3000              |
| Level 4 | 5000              |
| Level 5 | 7000              |
| Level 6 | 10000             |

### Base boosts

These are effects that change the base value, before any multiplicative boosts.

1. Base - this is the base value in the game, without any stations or other boosts
    - Mana - 20000 Mana
    - Slime - 15000 Slime
    - Darkness - 10000 Darkness
2. Resource cap stations
    - input: individual station levels and their counts
3. Serv-O:
    - this is one of Inventor's creations
    - it's essentially a station that you unlock passively by spending Time Shards
    - you can only have one
    - you can select its "color" (resource type), and it will increase the base resource cap of it + provide a small
      resource generation
    - you can upgrade it once (permanently), raising the resource cap
    - its base stats are:
        - 6000 Mana, upgraded 9000 Mana
        - 4500 Slime, upgraded 7500 Slime
        - 3000 Darkness, upgraded 6000 Darkness
    - input: whether player has the Serv-O upgrade, and which resource he selected
4. Skins:
    - these are effects you can unlock in the game with various sources
    - they provide small bonuses that are active whether you currently use the skin or not
    - in this section, relevant skins are:
        - `Wizard` - 2000 Mana
        - `Oozing` - 2000 Slime
        - `Sid` - 2000 Darkness
    - input: boolean values, if player has the skins or not

### Multiplicative boosts

These are effects that multiply the base value. Individually, they add up (I include 100% of the previous base), and
together they multiply the base.

1. Base - 100%
2. Golden Boosts - 25%
    - these are in-game purchases for real money, permanent and fixed value
    - input: boolean value, if player has them or not
3. Spells
    - Mana (Spellbook page 1):
        - +5% per spell level, up to +25%
    - Slime (Spellbook page 2):
        - +5% per spell level, up to +25%
    - Darkness (Spellbook page 3):
        - +5% per spell level, up to +25%
    - All resources (Spellbook page 6):
        - +5% per spell level, up to +10%
    - input: spell levels for each of them
4. Skins:
    - `Santa` - +5% Mana
    - `Birthday` - +5% Slime
    - `Witch` - +5% Darkness
    - `Good` - +2% All
    - `Royalty` - +5% All
    - input: boolean values, if player has the skins or not
5. Relics:
    - Mana, Slime, Darkness each has its own Relic (that you can equip)
        - level 1 - +2%
        - level 2 - +4%
        - level 3 - +6%
        - level 4 - +8%
        - level 5 - +10%
        - level 6 - +13%
        - level 7 - +16%
        - level 8 - +20%
        - level 9 - +25%
        - level 10 - +30%
    - also, there's a relic for all resources
        - level 1 - +3%
        - level 2 - +6%
        - level 3 - +9%
        - level 4 - +12%
        - level 5 - +15%
    - input: relic levels for each equipped
6. Experiments
    - at this point, just the Capacity Experiment I (pre-100) - up to +50%
    - input: experiment level
7. Pearls
    - these are late game items that you get from various sources
    - upon using, it presents user with 3 random small efects that stack until the end of the prestige
    - there are bonuses for individual resources, and also for all resources
    - input: per-resource total % bonus as shown in the game statistics page (e.g. tapping on Mana cap info shows "Pearls +6%")
      - the game already combines the individual and all-resources pearl contributions into one number per resource, so we just take that

### Multiplicative boost II

There's only 1 known effect of this type. After base and the previous multiplicative boosts are applied, this effect is
multiplied on top of it.

1. `Capacity Experiment II` - up to `x2` all resource caps. Input: experiment level

---

### Others

Sometimes, it was useful for me to manually input a goal for a resource type, and let it calculate the rest

- e.g. I knew I wanted to have more Mana than the spreadsheet decided, so I set that as a "given", and let it lower the
  other 2 resources so I don't necessarily overcap

## Calculation model

```
game_base        = 20000 Mana/15000 Slime/10000 Darkness
stations         = sum of all owned station levels and their base caps
serv_o           = 6000 Mana/4500 Slime/3000 Darkness, depending on selected resource (+ 3000 if upgraded)
skins_base       = 2000 Mana/Slime/Darkness, if respective skin is unlocked

base_subtotal    = game_base + stations + serv_o + skins_base
---
multiplicative   = 100% + golden boosts + spells + skins + relics + capacity experiment I + pearls
subtotal         = base_subtotal * multiplicative
---
total            = subtotal * capacity experiment II
```

---

## Outputs

- Current combined storage given the player's inputs
- Delta to each Feat threshold (how much more storage is needed)
- For the next unmet threshold: suggested stations to buy/upgrade to close the gap, with total rune cost
    - per-resource: for each station type, show how many of each level are needed and what they cost in runes
    - algorithm: effective contribution of one station at level N = base_cap[N] × total_multiplier
      divide the per-resource gap by that to get count needed; rune cost = count × 2^(N-1) × L1 rune cost

- individual resource targets:
    - default ratio (no fixed targets): ~69% Mana, 16.8% Slime, 14.2% Darkness of the combined threshold
    - if Mana is fixed: split remaining combined target as Slime 2/3, Darkness 1/3
    - if Slime is fixed: split remaining combined target as Mana 3/4, Darkness 1/4
    - if Darkness is fixed: split remaining combined target as Mana 2/3, Slime 1/3

---

## Implementation notes

- Station rune costs can reuse the existing `StationCosts` map in `calculator/runes.go`
- Capacity Experiment levels/values are already in `calculator/experiments.go`
- New endpoint: `POST /api/resource-cap/{threshold}` — `{threshold}` is one of `200k|400k|600k|800k`; takes player state, returns current storage vs. that single threshold + gap analysis if not met
- persistence of these calculations (similarly to prestige plans) would be nice, but not required