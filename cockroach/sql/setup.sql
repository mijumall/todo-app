CREATE DATABASE app;
SET DATABASE = app;

CREATE TABLE todo (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	content TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,

	CONSTRAINT pkey PRIMARY KEY (id)
);