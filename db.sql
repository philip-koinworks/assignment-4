CREATE DATABASE assignment4;

CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username text,
    email text,
    password text,
    age numeric
);

CREATE TABLE Photos (
    id SERIAL PRIMARY KEY,
    title text,
    caption VARCHAR(50),
    photo_url VARCHAR(50),
    user_id integer REFERENCES Users (id)
);

CREATE TABLE Comments (
    id SERIAL PRIMARY KEY,
    user_id integer REFERENCES Users (id),
    photo_id integer REFERENCES Photos (id),
    message VARCHAR(255)
);

CREATE TABLE SocialMedias (
    id SERIAL PRIMARY KEY,
    user_id integer REFERENCES Users (id),
    name text,
    social_media_url text
);
