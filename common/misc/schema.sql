CREATE TABLE public.users (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	user_name text NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT uni_users_user_name UNIQUE (user_name),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_users_user_name ON public.users USING btree (user_name);

CREATE TABLE public.follow (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  follower uuid NOT NULL,
  followee uuid NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT uni_follow_follower UNIQUE (follower),
	CONSTRAINT uni_follow_followee UNIQUE (followee)
);
CREATE INDEX idx_follow_follower ON public.follow USING btree (follower);
CREATE INDEX idx_follow_followee ON public.follow USING btree (followee);
