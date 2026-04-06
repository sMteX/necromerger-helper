function updatePlanner() {
    const form = document.getElementById('planner-form');
    if (!form) return;

    const formData = new FormData(form);
    const data = {
        devourerLevel: parseInt(formData.get('devourer_level')),
        featTiers: parseInt(formData.get('feat_tiers')),
        leftoverShards: parseInt(formData.get('leftover_shards') || 0),
        otherMultiplier: parseFloat(formData.get('other_multiplier') || 0) / 100.0,
        groupBonusCount: parseInt(formData.get('group_bonus_count') || 0) + 1,
        legendaryCounts: {},
        experimentLevels: {},
        possessedRunes: {},
        possessedLegendaries: {},
        notes: formData.get('notes')
    };

    for (let [key, value] of formData.entries()) {
        if (key.startsWith('leg_')) {
            data.legendaryCounts[key.replace('leg_', '')] = parseInt(value);
        } else if (key.startsWith('exp_')) {
            data.experimentLevels[key.replace('exp_', '')] = parseInt(value);
        } else if (key.startsWith('rune_')) {
            data.possessedRunes[key.replace('rune_', '')] = parseInt(value);
        } else if (key.startsWith('inv_leg_')) {
            data.possessedLegendaries[key.replace('inv_leg_', '')] = parseInt(value);
        }
    }

    fetch('/api/recalculate', {
        method: 'POST', headers: {'Content-Type': 'application/json'}, body: JSON.stringify(data)
    })
        .then(res => res.json())
        .then(res => {
            // Update summary
            const totalShardsStr = res.totalShards.toLocaleString();
            document.querySelectorAll('.js-total-shards').forEach(el => el.textContent = totalShardsStr);
            
            const expCostStr = '-' + res.experimentCost.toLocaleString();
            document.querySelectorAll('.js-exp-cost').forEach(el => el.textContent = expCostStr);

            const remaining = res.remaining;
            const remainingStr = remaining.toLocaleString();
            document.querySelectorAll('.js-remaining-shards').forEach(el => {
                el.textContent = remainingStr;
                el.classList.toggle('text-emerald-400', remaining >= 0);
                el.classList.toggle('text-red-500', remaining < 0);
            });

            // Update experiment values and next level info
            for (const [id, summary] of Object.entries(res.experiments)) {
                const expEl = document.querySelector(`[data-exp-id="${id}"]`);
                if (expEl) {
                    const valueEl = expEl.querySelector('.js-exp-value');
                    if (valueEl) valueEl.textContent = summary.currentLevelValue;

                    const nextLevelEl = expEl.querySelector('.js-exp-next-level');
                    const shardIcon = expEl.querySelector('img[src="/static/time_shard.png"]');
                    if (nextLevelEl) {
                        if (summary.maxLevel) {
                            nextLevelEl.textContent = 'MAXED';
                            nextLevelEl.classList.add('text-emerald-500');
                            if (shardIcon) shardIcon.classList.add('hidden');
                        } else {
                            nextLevelEl.textContent = summary.nextLevelCost;
                            nextLevelEl.classList.remove('text-emerald-500');
                            if (shardIcon) shardIcon.classList.remove('hidden');
                        }
                    }
                }
            }

            // Update rune deficiency (optimized)
            const runeContainer = document.getElementById('rune-deficiency-container');
            const emptyMsg = document.getElementById('rune-deficiency-empty');
            if (runeContainer) {
                let anyNeeded = false;
                for (const rune of ['ice', 'poison', 'blood', 'moon', 'death', 'cosmic']) {
                    const count = res.runeNeeded[rune] || 0;
                    const item = runeContainer.querySelector(`[data-rune-type="${rune}"]`);
                    if (item) {
                        if (count > 0) {
                            item.querySelector('.js-rune-count').textContent = count.toLocaleString();
                            item.classList.remove('hidden');
                            anyNeeded = true;
                        } else {
                            item.classList.add('hidden');
                        }
                    }
                }
                if (emptyMsg) emptyMsg.classList.toggle('hidden', anyNeeded);
            }

            // Update per-legendary costs in UI
            for (const [id, count] of Object.entries(data.legendaryCounts)) {
                const legEl = document.querySelector(`[data-leg-id="${id}"]`);
                if (legEl) {
                    legEl.querySelectorAll('.js-leg-rune-count').forEach(el => {
                        const baseAmount = parseInt(el.dataset.baseAmount);
                        el.textContent = (baseAmount * (count || 1)).toLocaleString();
                        el.classList.toggle('text-slate-400', count === 0);
                        el.classList.toggle('text-emerald-400', count > 0);
                    });
                }
            }
        });
}

document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('planner-form');
    if (form) {
        form.addEventListener('input', () => {
            clearTimeout(window.updateTimeout);
            window.updateTimeout = setTimeout(updatePlanner, 300);
        });
        updatePlanner();
    }
});
