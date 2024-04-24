-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	user_name text NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT uni_users_user_name UNIQUE (user_name),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_users_user_name ON public.users USING btree (user_name);
