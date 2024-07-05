CREATE TABLE people 
(
    id serial PRIMARY KEY,
    passport_serie int NOT NULL,
    passport_number int NOT NULL UNIQUE,
    surname varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    patronymic varchar(255) NOT NULL,
    address varchar(255) NOT NULL
);

CREATE TABLE task
(
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    name varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    start_time time,
    end_time time,
    diff_time time,
    done bool NOT NULL DEFAULT false
);
