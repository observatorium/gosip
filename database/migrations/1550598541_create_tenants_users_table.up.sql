CREATE TABLE tenants_users
(
  tenant_id UUID REFERENCES tenants ON UPDATE CASCADE ON DELETE CASCADE,
  user_id   UUID REFERENCES users ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT tenants_users_pkey PRIMARY KEY (tenant_id, user_id)
);
