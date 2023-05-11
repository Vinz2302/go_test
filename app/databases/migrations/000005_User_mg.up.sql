CREATE TABLE users (
	id int NOT NULL,
	"name" text NOT NULL,
	email varchar(256) NOT NULL,
	role_id int8 NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_name_key UNIQUE (name),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

ALTER TABLE public.users ADD CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES roles(id);