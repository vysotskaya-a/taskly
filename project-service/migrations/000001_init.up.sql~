CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    users UUID[] NOT NULL,
    admin_id UUID NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
)