-- Copyright (c) 2020 Vorotynsky Maxim

-- MySQL
-- database schema for microAuth microservice.

USE microAuthDB;

CREATE TABLE IF NOT exists users (
    userId INT NOT NULL AUTO_INCREMENT,
    userName VARCHAR(35) NOT NULL,
    passwordHash BINARY(60) NOT NULL,

    PRIMARY KEY(userId)
);
