CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     uid varchar(255) NOT NULL unique ,
                                     name TEXT NOT NULL
);
