CREATE DATABASE IF NOT EXISTS go_graphql

CREATE TABLE IF NOT EXISTS Users(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Username VARCHAR (127) NOT NULL UNIQUE,
    Password VARCHAR (127) NOT NULL,
    PRIMARY KEY (ID)
)

CREATE TABLE IF NOT EXISTS Links(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Title VARCHAR (255),
    Address VARCHAR (255),
    UserID INT,
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    PRIMARY KEY (ID)
)