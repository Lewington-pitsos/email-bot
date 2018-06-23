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

SELECT id, first_name, last_name, pass, email, birthdate FROM profiles;

INSERT INTO profiles (
    id, 
    first_name, 
    last_name, 
    pass, 
    email, 
    birthdate

)