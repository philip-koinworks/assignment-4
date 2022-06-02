CREATE DATABASE assignment4;

CREATE TABLE Users (
    id integer PRIMARY KEY,
    username text,
    email text,
    password text,
    age numeric
);

CREATE TABLE Photos (
    id numeric PRIMARY KEY,
    title text,
    caption VARCHAR(50),
    photo_url VARCHAR(50),
    user_id integer REFERENCES Users (id)
);

CREATE TABLE Comments (
    id numeric PRIMARY KEY,
    user_id integer REFERENCES Users (id),
    photo_id integer REFERENCES Photos (id),
    message VARCHAR(255)
);

CREATE TABLE SocialMedias (
    id numeric PRIMARY KEY,
    user_id integer REFERENCES Users (id),
    name text,
    social_media_url text
);
