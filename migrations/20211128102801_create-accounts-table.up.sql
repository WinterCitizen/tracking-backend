CREATE TABLE accounts (
    username VARCHAR(32) CONSTRAINT primarykey PRIMARY KEY,
    password VARCHAR(60) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);