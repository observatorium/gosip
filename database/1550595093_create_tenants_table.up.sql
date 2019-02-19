CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE tenants (
  id         UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
  name       VARCHAR(100) NOT NULL,
  prometheus VARCHAR(256) NOT NULL,
  jaeger     VARCHAR(256) NOT NULL
);
