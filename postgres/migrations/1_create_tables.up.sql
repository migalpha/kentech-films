CREATE TABLE IF NOT EXISTS users(
	id serial PRIMARY KEY,
	username VARCHAR ( 75 ) UNIQUE NOT NULL,
	password VARCHAR ( 75 ) NOT NULL,
    is_active   bool    NOT null,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
    DEFAULT (current_timestamp AT TIME ZONE 'UTC')
);

CREATE TABLE IF NOT EXISTS films(
	id serial PRIMARY KEY,
	title VARCHAR ( 255 ) UNIQUE NOT NULL,
	director VARCHAR ( 75 ) NOT NULL,
	genre VARCHAR ( 75 ) NOT NULL,
	sypnosis text null,
	starring VARCHAR ( 255 ) NULL,
	release_year  smallint not null,
	created_by int REFERENCES users (id) not null,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
    DEFAULT (current_timestamp AT TIME ZONE 'UTC')
);

CREATE TABLE IF NOT EXISTS favourites(
	id serial PRIMARY KEY,
	film_id int REFERENCES films (id) not null,
	user_id int REFERENCES users (id) not null
);