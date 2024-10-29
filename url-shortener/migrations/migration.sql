-- Create Database
CREATE DATABASE IF NOT EXISTS url_mapping_db;

-- Use the newly created database
USE url_mapping_db;

-- Create url_mapping Table
CREATE TABLE IF NOT EXISTS url_mapping (
    short_url VARCHAR(7) NOT NULL PRIMARY KEY,
    long_url VARCHAR(400) DEFAULT NULL,
    user_id VARCHAR(28) DEFAULT NULL,
    active TINYINT(1) NOT NULL DEFAULT 1);

-- Differentiate from upper and lowercase ids
ALTER TABLE url_mapping
    MODIFY COLUMN short_url VARCHAR(7) COLLATE utf8mb4_bin NOT NULL;
