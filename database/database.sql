-- Copyright (c) 2020 Vorotynsky Maxim

-- MySQL
-- database schema for microAuth microservice.

USE microAuthDB;

CREATE TABLE IF NOT exists users (
    userId INT NOT NULL AUTO_INCREMENT,
    userName VARCHAR(35) NOT NULL,
    passwordHash INT NOT NULL,

    PRIMARY KEY(userId)
);
