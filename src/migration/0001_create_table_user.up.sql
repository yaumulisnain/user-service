DROP TABLE IF EXISTS public.user;

CREATE TABLE public.user (
	id bigserial NOT NULL,
	email varchar(50) NOT NULL UNIQUE,
	"password" varchar(255) NOT NULL,
	address text NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	CONSTRAINT user_pkey PRIMARY KEY (id)
);
