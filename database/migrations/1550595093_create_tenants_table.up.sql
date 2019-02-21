CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE tenants
(
  id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  prometheus TEXT NOT NULL,
  jaeger     TEXT NOT NULL
);
