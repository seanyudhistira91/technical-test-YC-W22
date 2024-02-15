CREATE TABLE IF NOT EXISTS profiles(
   id bigserial PRIMARY KEY,
   user_id bigint not null,
   name VARCHAR NOT NULL,
   created_at timestamptz not null default now(),
	updated_at timestamptz not null default now(),
	deleted_at timestamptz DEFAULT null
);

ALTER TABLE profiles
	ADD CONSTRAINT user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id) ON
	UPDATE
		CASCADE ON DELETE NO action;