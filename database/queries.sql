CREATE TABLE profiles (
    id serial,
    first_name varchar(30) NOT NULL,
    last_name varchar(30) NOT NULL,
    pass varchar(40) NOT NULL,
    email varchar(40) NOT NULL,
    birthdate TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    PRIMARY KEY(id) 
);

SELECT first_name, last_name, pass, email, birthdate FROM profiles;

INSERT INTO profiles (
    first_name, 
    last_name, 
    pass, 
    email, 
    birthdate
) VALUES (

)

CREATE TABLE yandex_profiles (
    id serial,
    first_name varchar(30) NOT NULL,
    last_name varchar(30) NOT NULL,
    password varchar(40) NOT NULL,
    email varchar(40) NOT NULL,
    question varchar(60) NOT NULL,
    answer varchar(30) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    PRIMARY KEY(id) 
);
