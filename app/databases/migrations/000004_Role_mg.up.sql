CREATE TABLE roles (
	id bigserial NOT NULL,
	"name" varchar(256) NOT NULL,
	CONSTRAINT roles_pkey PRIMARY KEY (id)
);