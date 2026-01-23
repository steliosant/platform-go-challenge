-- CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE DATABASE dashboard;

-- USERS
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

-- ASSETS
CREATE TYPE asset_type AS ENUM ('chart', 'insight', 'audience');

CREATE TABLE assets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type asset_type NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    data JSONB,
    created_at TIMESTAMP DEFAULT now()
);

-- FAVORITES
CREATE TABLE favorites (
    user_id TEXT REFERENCES users(id) ON DELETE CASCADE,
    asset_id UUID REFERENCES assets(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now(),
    PRIMARY KEY (user_id, asset_id)
);

-- SEED DATA
INSERT INTO users (id, name) VALUES
('u1', 'Alice'),
('u2', 'Bob');

INSERT INTO assets (type, title, description, data) VALUES
(
  'insight',
  'Social Media Usage',
  '40% of millennials spend more than 3 hours on social media daily',
  '{"source":"survey"}'
),
(
  'chart',
  'Purchases per Age Group',
  'Monthly purchases',
  '{"x":["18-24","25-34"],"y":[5,9]}'
);
