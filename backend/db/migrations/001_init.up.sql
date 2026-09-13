CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username TEXT NOT NULL UNIQUE CHECK (username <> ''),
	password_hash TEXT NOT NULL CHECK (password_hash <> '')
);

CREATE TABLE tasks (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id),
	description TEXT NOT NULL CHECK (description <> '')
);

CREATE TABLE user_sessions (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id),
	session_string TEXT NOT NULL UNIQUE CHECK(session_string <> '')
)
