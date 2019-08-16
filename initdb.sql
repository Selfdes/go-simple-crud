CREATE TABLE IF NOT EXISTS account
(
    id SERIAL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    api_token TEXT NOT NULL,
    CONSTRAINT account_pkey PRIMARY KEY (id)
);
