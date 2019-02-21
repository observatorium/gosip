CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users
(
  id       UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username VARCHAR(100) NOT NULL,
  password TEXT         NOT NULL,
  token    UUID             DEFAULT gen_random_uuid()
);

CREATE UNIQUE INDEX users_username_uniq_idx ON users (username);
CREATE UNIQUE INDEX users_token_uniq_idx ON users (token);
