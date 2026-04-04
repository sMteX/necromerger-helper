-- name: GetPlan :one
SELECT * FROM plans
WHERE id = $1 LIMIT 1;

-- name: ListPlans :many
SELECT id, name FROM plans
ORDER BY updated_at DESC;

-- name: CreatePlan :one
INSERT INTO plans (
    name, devourer_level, feat_tiers, other_multiplier, group_bonus_count, legendary_counts, experiment_levels, notes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
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
    legendary_counts = $7,
    experiment_levels = $8,
    notes = $9,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeletePlan :exec
DELETE FROM plans
WHERE id = $1;
