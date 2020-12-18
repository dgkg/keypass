CREATE TABLE conteners (
 id text NULL,
 user_id text NULL,
 title text NULL,
 secret text NULL,
 creation_date datetime NULL
);

CREATE TABLE users (
 id text NULL,
 first_name text NULL,
 last_name text NULL,
 email text NULL,
 password text NULL,
 creation_date datetime NULL
);

CREATE TABLE cards (
 id text NULL,
 contener_id text NULL,
 url text,
 pic blob,
 activated tinyint(1) NULL,
 creation_date datetime NULL
);
