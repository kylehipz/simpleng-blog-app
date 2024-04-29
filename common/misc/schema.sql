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
	CONSTRAINT uni_follow_follow UNIQUE (follower, followee)
);
CREATE INDEX idx_follow_follower ON public.follow USING btree (follower);
CREATE INDEX idx_follow_followee ON public.follow USING btree (followee);

CREATE TABLE public.blogs (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  author uuid NOT NULL,
  title text NOT NULL,
  body text NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NULL
);
CREATE INDEX idx_blog_author ON public.blogs USING btree (author);
CREATE INDEX idx_blog_title ON public.blogs USING btree (title);
