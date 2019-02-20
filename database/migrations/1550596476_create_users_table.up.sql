CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
  id         UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
  name       VARCHAR(100) NOT NULL,
  password   TEXT         NOT NULL,
  token      TEXT         NOT NULL
);
