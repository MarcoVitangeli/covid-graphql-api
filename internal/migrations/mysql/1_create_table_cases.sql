CREATE TABLE cases(
    id varchar(255) PRIMARY KEY,
    province VARCHAR(255) NOT NULL,
    gender VARCHAR(255) NOT NULL,
    neighborhood VARCHAR(255) NOT NULL,
    age SMALLINT NOT NULL,
    stage VARCHAR(255) NOT NULL,
    dead ENUM('YES','NO'),
    inserted_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);