-- +goose Up
CREATE TABLE IF NOT EXISTS plans (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    devourer_level INTEGER DEFAULT 35,
    feat_tiers INTEGER DEFAULT 0,
    other_multiplier DOUBLE PRECISION DEFAULT 0.0,
    group_bonus_count INTEGER DEFAULT 1,
    leftover_shards INTEGER DEFAULT 0,
    legendary_counts JSONB DEFAULT '{}',
    experiment_levels JSONB DEFAULT '{}',
    possessed_runes JSONB DEFAULT '{}',
    possessed_legendaries JSONB DEFAULT '{}',
    notes TEXT DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS plans;
