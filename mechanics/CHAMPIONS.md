# Summoning champions
- There are 5 champions that we can farm for runes and for which the efficiencies we are interested in
- each champion has a different point threshold for summoning
- summoning points are earned by merging parts or creatures to create another creatures (e.g. 2x L2 part -> L1 creature, or 2x L1 creature -> L2 creature etc.)
- each champion is summoned by different creatures, and some give more points than others (also further ranks give more points per merge than lower ranks)
- there are also various in-game mechanics that together produce effects that improve the "summoning speed" (likely making each merge give more points than the base) - e.g. "150% speed" = 1.5x the points

## Champions
### The Peasant
- gives Ice Chest upon killing it
- can be summoned by merging Skeletons, Zombies and Mummies
- threshold = 150000 points
- summoning table:
    
| Monster \ Level | 1    | 2    | 3    | 4    | 5    | 6    | 7    |
|-----------------|------|------|------|------|------|------|------|
| Skeleton        | 1000 | 1200 | 1500 | 2000 | 2500 | 3000 | 4000 |
| Zombie          | 3000 | 3500 | 4000 | 4500 | 5000 | 6000 |      |
| Mummy           | 4000 | 4500 | 5000 | 5500 | 6000 |      |      |

- stations:
  - Skeletons and Zombies are summoned by tapping Graves (costing Mana)
  - Mummies are summoned by tapping Supply Cupboard (costing Slime)

### The Knight
- gives Poison Chest upon killing it
- can be summoned by merging Eye Monsters, (Spiders) and Werewolves
  - Spiders are being excluded as they are time-gated resource, not really farmable
- threshold = 200000 points
- summoning table:

| Monster \ Level | 1    | 2    | 3    | 4    | 5    | 6    |
|-----------------|------|------|------|------|------|------|
| Eye Monster     | 1250 | 1500 | 2000 | 2500 | 3000 | 3500 |
| Werewolf        | 1750 | 2000 | 2500 | 3000 | 4000 | -    |

- stations:
  - Eye Monsters are summoned by tapping Supply Cupboard (costing Slime)
  - Werewolves are summoned by tapping Altar (costing Darkness)

### The Cleric
- gives Blood Chest upon killing it
- can be summoned by merging (Snakes), Bats and Shades
    - Snakes are being excluded as they are time-gated resource, not really farmable
    - Shades might come as obvious losers, but they are special in the sense that feeding them gives back Darkness
- threshold = 250000 points
- summoning table:

| Monster \ Level | 1    | 2    | 3    | 4    |
|-----------------|------|------|------|------|
| Bat             | 4000 | 4500 | 5000 | 6000 |
| Shades          | 1500 | 2000 | 2500 | 3000 |

- stations:
    - Bats are summoned by tapping Altar (costing Darkness):
    - Shades are summoned by tapping Lectern (costing Mana):
        - there is no L1 part for Shades, so for sake of consistency, let's label the part as L2 part (as in 2x part -> L1 Shade)

### The Paladin
- gives Moon Chest upon killing it
- can be summoned by merging Banshees, Ghouls and Imps
- threshold = 200000 points
- summoning table:

| Monster \ Level | 1    | 2    | 3    | 4    |
|-----------------|------|------|------|------|
| Banshee         | 1500 | 2000 | 2500 | 3000 |
| Ghoul           | 1000 | 1200 | 1500 | 2000 |
| Imp             | 1300 | 1600 | 1900 | 2200 |

- stations:
    - Banshees are summoned by tapping Lectern (costing Mana)
    - Ghouls are summoned by tapping Fridge (costing Slime)
    - Imps are summoned by tapping Portal (costing Darkness)

### The Rival
- gives Death Chest upon killing it
- can be summoned by merging (Golden Geese), Abominations and Demons
    - again, Golden Geese are being excluded as they are time-gated resource, not really farmable
- threshold = 100000 points
- summoning table:

| Monster \ Level | 1    | 2    | 3    |
|-----------------|------|------|------|
| Abomination     | 1000 | 1200 | 1500 |
| Demon           | 1300 | 1600 | 1900 |

- stations:
    - Abominations are summoned by tapping Fridge (costing Slime)
    - Demons are summoned by tapping Portal (costing Darkness)

## Station drop tables
### Grave (costs Mana)

| Level \ Drop %   | Cost |                         |                        |                      |                      |
|------------------|------|-------------------------|------------------------|----------------------|----------------------|
| Level 1          | 250  | Skeleton L1 part (100%) |                        |                      |                      |
| Level 2          | 300  | Skeleton L1 part (60%)  | Skeleton L2 part (40%) |                      |                      |
| Level 2 (hacked) | 400  | Skeleton L2 part (100%) |                        |                      |                      |
| Level 3          | 400  | Skeleton L1 part (40%)  | Skeleton L2 part (30%) | Zombie L1 part (30%) |                      |
| Level 4          | 500  | Skeleton L1 part (25%)  | Skeleton L2 part (25%) | Zombie L1 part (30%) | Zombie L2 part (20%) |
| Level 5          | 700  | Skeleton L1 part (10%)  | Skeleton L2 part (10%) | Zombie L1 part (40%) | Zombie L2 part (40%) |
| Level 5 (hacked) | 1000 | Skeleton L1 part (10%)  | Zombie L2 part (90%)   |                      |                      |

### Supply Cupboard (costs Slime)

| Level \ Drop %   | Cost |                            |                           |                     |                     |
|------------------|------|----------------------------|---------------------------|---------------------|---------------------|
| Level 1          | 250  | Eye Monster L1 part (100%) |                           |                     |                     |
| Level 2          | 300  | Eye Monster L1 part (60%)  | Eye Monster L2 part (40%) |                     |                     |
| Level 2 (hacked) | 400  | Eye Monster L2 part (100%) |                           |                     |                     |
| Level 3          | 400  | Eye Monster L1 part (40%)  | Eye Monster L2 part (30%) | Mummy L1 part (30%) |                     |
| Level 4          | 500  | Eye Monster L1 part (25%)  | Eye Monster L2 part (25%) | Mummy L1 part (30%) | Mummy L2 part (20%) |
| Level 5          | 700  | Eye Monster L1 part (10%)  | Eye Monster L2 part (10%) | Mummy L1 part (40%) | Mummy L2 part (40%) |
| Level 5 (hacked) | 1000 | Eye Monster L1 part (10%)  | Mummy L2 part (90%)       |                     |                     |

### Altar (costs Darkness)

| Level \ Drop %   | Cost |                         |                        |                   |                   |
|------------------|------|-------------------------|------------------------|-------------------|-------------------|
| Level 1          | 250  | Werewolf L1 part (100%) |                        |                   |                   |
| Level 2          | 300  | Werewolf L1 part (60%)  | Werewolf L2 part (40%) |                   |                   |
| Level 2 (hacked) | 400  | Werewolf L2 part (100%) |                        |                   |                   |
| Level 3          | 400  | Werewolf L1 part (40%)  | Werewolf L2 part (30%) | Bat L1 part (30%) |                   |
| Level 4          | 500  | Werewolf L1 part (25%)  | Werewolf L2 part (25%) | Bat L1 part (30%) | Bat L2 part (20%) |
| Level 5          | 700  | Werewolf L1 part (10%)  | Werewolf L2 part (10%) | Bat L1 part (40%) | Bat L2 part (40%) |
| Level 5 (hacked) | 1000 | Werewolf L1 part (10%)  | Bat L2 part (90%)      |                   |                   |

### Lectern (costs Mana)

| Level \ Drop %   | Cost |                        |                     |                       |                       |
|------------------|------|------------------------|---------------------|-----------------------|-----------------------|
| Level 1          | 500  | Skeleton L2 part (70%) | Shade L2 part (30%) |                       |                       |
| Level 2          | 750  | Skeleton L2 part (40%) | Shade L2 part (60%) |                       |                       |
| Level 2 (hacked) | 1000 | Skeleton L2 part (10%) | Shade L2 part (90%) |                       |                       |
| Level 3          | 1000 | Skeleton L2 part (30%) | Shade L2 part (50%) | Banshee L1 part (20%) |                       |
| Level 4          | 1250 | Skeleton L2 part (20%) | Shade L2 part (40%) | Banshee L1 part (30%) | Banshee L2 part (10%) |
| Level 5          | 1500 | Skeleton L2 part (10%) | Shade L2 part (30%) | Banshee L1 part (40%) | Banshee L2 part (40%) |
| Level 5 (hacked) | 2500 | Skeleton L2 part (10%) | Shade L2 part (20%) | Banshee L2 part (70%) |                       |

### Fridge (costs Slime)

| Level \ Drop %   | Cost |                           |                     |                           |                           |                           |
|------------------|------|---------------------------|---------------------|---------------------------|---------------------------|---------------------------|
| Level 1          | 500  | Eye Monster L2 part (70%) | Ghoul L1 part (30%) |                           |                           |                           |
| Level 2          | 750  | Eye Monster L2 part (40%) | Ghoul L1 part (60%) |                           |                           |                           |
| Level 3          | 1000 | Eye Monster L2 part (30%) | Ghoul L1 part (40%) | Ghoul L2 part (30%)       |                           |                           |
| Level 3 (hacked) | 1750 | Eye Monster L2 part (10%) | Ghoul L2 part (90%) |                           |                           |                           |
| Level 4          | 1250 | Eye Monster L2 part (20%) | Ghoul L1 part (30%) | Ghoul L2 part (25%)       | Abomination L1 part (25%) |                           |
| Level 5          | 1500 | Eye Monster L2 part (10%) | Ghoul L1 part (20%) | Ghoul L2 part (20%)       | Abomination L1 part (30%) | Abomination L2 part (20%) |
| Level 5 (hacked) | 2500 | Eye Monster L2 part (10%) | Ghoul L2 part (20%) | Abomination L2 part (70%) |                           |                           |

### Portal (costs Darkness)

| Level \ Drop %   | Cost |                        |                   |                     |                     |                     |
|------------------|------|------------------------|-------------------|---------------------|---------------------|---------------------|
| Level 1          | 500  | Werewolf L2 part (70%) | Imp L1 part (30%) |                     |                     |                     |
| Level 2          | 750  | Werewolf L2 part (40%) | Imp L1 part (60%) |                     |                     |                     |
| Level 3          | 1000 | Werewolf L2 part (30%) | Imp L1 part (40%) | Imp L2 part (30%)   |                     |                     |
| Level 3 (hacked) | 1750 | Werewolf L2 part (10%) | Imp L2 part (90%) |                     |                     |                     |
| Level 4          | 1250 | Werewolf L2 part (20%) | Imp L1 part (30%) | Imp L2 part (25%)   | Demon L1 part (25%) |                     |
| Level 5          | 1500 | Werewolf L2 part (10%) | Imp L1 part (20%) | Imp L2 part (20%)   | Demon L1 part (30%) | Demon L2 part (20%) |
| Level 5 (hacked) | 2500 | Werewolf L2 part (10%) | Imp L2 part (20%) | Demon L2 part (70%) |                     |                     |
