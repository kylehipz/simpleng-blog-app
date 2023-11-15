CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email VARCHAR(100) NOT NULL,
  username VARCHAR(30) NOT NULL
);

CREATE TABLE follow (
  follower UUID NOT NULL,
  followee UUID NOT NULL,
  CONSTRAINT fk_follower FOREIGN KEY(follower) REFERENCES users(id),
  CONSTRAINT fk_followee FOREIGN KEY(followee) REFERENCES users(id)
);

CREATE TABLE post (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(140) NOT NULL,
  content VARCHAR NOT NULL,
  created_by UUID NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  CONSTRAINT fk_created_by FOREIGN KEY(created_by) REFERENCES users(id)
);

CREATE TABLE attachment (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  post_id UUID NOT NULL,
  url VARCHAR NOT NULL,
  CONSTRAINT fk_post FOREIGN KEY(post_id) REFERENCES post(id)
);
