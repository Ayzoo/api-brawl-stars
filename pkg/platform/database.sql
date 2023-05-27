CREATE DATABASE bralwstars;

CREATE TABLE brawlers(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    image VARCHAR(255),
    type VARCHAR(255),
    category VARCHAR(255),
    stellar INT,
    gadget INT,
    super INT,
    FOREIGN KEY (stellar) REFERENCES brawlers_stellar(id),
    FOREIGN KEY (gadget) REFERENCES brawlers_gadget(id),
    FOREIGN KEY (super) REFERENCES super(id)
);

CREATE TABLE brawlers_stellar(
    id INT AUTO_INCREMENT PRIMARY KEY,
    brawler_id INT,
    primary_stellar VARCHAR(255),
    second_stellar VARCHAR(255)
);


CREATE TABLE brawlers_gadget(
    id INT AUTO_INCREMENT PRIMARY KEY,
    brawler_id INT,
    primary_gadget VARCHAR(255),
    second_gadget VARCHAR(255)
);


CREATE TABLE super(
    id INT AUTO_INCREMENT PRIMARY KEY,
    brawler_id INT,
    name VARCHAR(255),
    description VARCHAR(255),
    damage VARCHAR(255),
    scope VARCHAR(255)
);