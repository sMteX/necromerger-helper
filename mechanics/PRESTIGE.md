# Prestige Reset Mechanics
- starting from level 50, the player can choose to reset their game
- this effect looks at certain things ("what the player has accomplished") and based on that, give them a certain amount of `Time Shards`
- this act resets the players Runes, Devourer level, Spells, Feats, Stations, raises the Devourer level cap
- but keeps the player's `Gems` (which can be bought for real money or slowly accumulated through normal free to play gameplay), `Gold` (similarly to Gems) and few other things:
    - `Astro Coins`, `Energy Cubes`, `Wobular` and `Time Shards` (which are all currencies accumulated in the late game and through prestiging)
- upon confirmation, the game resets, the player is awarded a calculated amount of Time Shards

## Time Shard calculation
- there are four variables used to determine the amount of Time Shards (base amount, and then 3 percentage multipliers):
    - **Devourer level** - every level increases the base amount of Time Shards (predetermined amounts)
        - it is a must to max out the Devourer level before even considering to prestige
        - concrete values are already implemented in code
    - **Feats** - `100%` base, every **completed** Tier of Feats gives +10% Time Shards
        - partial completion of a Tier does not count (sometimes the player is blocked by a game mechanic, or a Devourer level cap)
    - **Legendaries** - `100%` base
        - `Legendaries` are special types of creatures that are produced when merging certain top level Stations (or even other Legendaries)
        - each Legendary has a certain +% Time Shard boost (for the first one, and usually a little less for each additional one)
        - there are 4 groups of Legendaries:
            - `Group 1`:
                - `Lich`, `Gorgon`, `Harpy`
                - each gives +10% Time Shards for the first one, and +5% for each additional one
                - when the player owns at least 1 of each, they "complete" the group and receive additional +20% Time Shards bonus
            - `Group 2`:
                - `Reaper`, `Cyclops`, `Archdemon`
                - each gives +20% Time Shards for the first one, and +10% for each additional one
                - +40% Time Shard bonus for completing this group
                - Archdemon is capped at 3 or 4
            - `Group 3`:
                - `The Cursed`, `The Colossus`, `The Infernal`
                - these are special and the player can only have 1 of each - +40% Time Shard bonus for each of them
                - +80% Time Shard bonus for completing this group
            - `Group 4`:
                - `Robo Chicken`, `Shield Bot`, `Soul Stalker`
                - Robo Chicken - +20% Time Shard (+10% for each additional one) - capped at 3 or 4
                - Shield Bot - +30% Time Shard (+15% for each additional one) - capped at 3 or 4
                - Soul Stalker - +40% Time Shard (+20% for each additional one) - uncapped
                - +60% Time Shard bonus for completing this group
        - **NOTE:** in the **very** late game, the player can unlock a feature that lets them claim the group bonus multiple times
            - e.g. at +1 group bonus, they can have 2 of each in Group 1 and receive +40% Time Shard bonus instead of +20%
    - **Other** - `100%` base
        - most notably skins and `Time Pieces` consumed
        - `Time Pieces` are items received from various sources in late game, that are again level 1/2/3 and which give +1/+2/+5% Time Shards
- due to the nature of these variables, it is recommended (prior to prestiging) to:
    - cap the Devourer to maximum
    - complete all available Feats
    - have at least 1 of each Legendary
- this leaves just a single way to adjust how many Time Shards the player gets - **farming additional Legendaries**

## The Inventor
- The Inventor is a special NPC that is used to spend the Time Shards for `Experiments`
- each experiment level has a predefined Time Shard cost that goes up roughly exponentially with each experiment level
- the experiments are the crux to make new repetitions of the game (also called "runs") easier and quicker and more fun
- there are experiments that increase the Food gain from feeding the Devourer, that increase all Damage, that increase the Craving rewards, that increase resource caps...
- there are also experiments that increase the amount of Chest taps (separated per each Chest type)
- and also a few special experiments that affect the key enemy types or other mechanics
- these experiments are effective **only until the player prestiges again** => they need to be repurchased again (the costs don't change)

### Experiments
- there are 2 "tiers" of experiments, pre-100 and post-100 (referring to the Devourer level, they unlock after reaching level 100 for the first time)
- pre-100 (usually percentage bonuses additive with other bonuses; but also bonus chest taps and other special effects; concrete values are already implemented in code):
    - **Seasoning Experiment** - increases the amount of Food from feeding the Devourer
    - **Strength Experiment** - increases all Damage
    - **Taste Experiment** - get more Food from Cravings
    - **Capacity Experiment** - increase Mana, Slime and Darkness caps
    - **Body Snatcher** (50 Time Shards) - unlocks The Body Snatcher
    - **Weakening Experiment** - Reduces the Protector's health scale (Protector is a special Champion that starts with a fixed amount of HP, but each subsequent summon scales his HP up by a percentage)
    - **Damage Cap Experiment** - Increase the Mech's Damage Cap scale (The Mech is a special Champion that has the same amount of HP each time, but has a damage cap = you can't deal more damage at one time than the cap value. The cap value gets smaller with each summon)
    - **Ice Chest Experiment** - increases the number of taps for Ice Chests
    - **Poison Chest Experiment** - increases the number of taps for Poison Chests
    - **Blood Chest Experiment** - increases the number of taps for Blood Chests
    - **Moon Chest Experiment** - increases the number of taps for Moon Chests
    - **Death Chest Experiment** - increases the number of taps for Death Chests
    - **Cosmic Chest Experiment** - increases the number of taps for Cosmic Chests
- post-100:
    - these experiments are multiplicative and are applied after usual additive bonuses, that's what makes them much more powerful and expensive
    - **Seasoning Experiment II** - get more Food when feeding the Devourer
    - **Strength Experiment II** - increases all Damage
    - **Taste Experiment II** - get more Food from Cravings
    - **Capacity Experiment II** - increase Mana, Slime and Darkness caps

## Optimizing The Inventor
- this is the goal of this application
- it is to plan what the player needs to do in the game, in order to maximize the amount of Time Shards received at the point of prestiging
- it's a balance between what experiments the user wants to buy/unlock and what they need to do to afford that (how many extra Legendaries to farm and how many Runes to do so)