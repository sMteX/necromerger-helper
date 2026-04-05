-- name: GetPlan :one
SELECT * FROM plans
WHERE id = $1 LIMIT 1;

-- name: ListPlans :many
SELECT id, name FROM plans
ORDER BY updated_at DESC;

-- name: CreatePlan :one
INSERT INTO plans (
    name, devourer_level, feat_tiers, other_multiplier, group_bonus_count, leftover_shards, legendary_counts, experiment_levels, possessed_runes, possessed_legendaries, notes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING id;

-- name: UpdatePlan :exec
UPDATE plans
SET 
    name = $2,
    devourer_level = $3,
    feat_tiers = $4,
    other_multiplier = $5,
    group_bonus_count = $6,
    leftover_shards = $7,
    legendary_counts = $8,
    experiment_levels = $9,
    possessed_runes = $10,
    possessed_legendaries = $11,
    notes = $12,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeletePlan :exec
DELETE FROM plans
WHERE id = $1;
