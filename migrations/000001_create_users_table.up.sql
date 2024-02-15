CREATE TABLE IF NOT EXISTS users(
   id bigserial PRIMARY KEY,
   hash_password VARCHAR NOT NULL,
   email VARCHAR UNIQUE NOT NULL,
   phone_number VARCHAR UNIQUE NOT NULL,
   is_premium boolean DEFAULT false,
   created_at timestamptz not null default now(),
	updated_at timestamptz not null default now(),
	deleted_at timestamptz DEFAULT null
);