-- DB設計

-- +migrate Up
CREATE TABLE IF NOT EXISTS users(
    id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    line_id VARCHAR(64) NOT NULL,
    office_code INT NOT NULL,
    area_code INT NOT NULL
);

-- +migrate Down ユーザーがいたときに何もしない
DROP TABLE IF EXISTS users;