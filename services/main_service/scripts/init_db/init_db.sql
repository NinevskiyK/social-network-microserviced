CREATE TABLE users_creds (
    id VARCHAR(36) PRIMARY KEY,
    login VARCHAR(256),
    password VARCHAR(255)
);

CREATE TABLE users_info (
    id VARCHAR(36) PRIMARY KEY,
    first_name VARCHAR(256),
    last_name VARCHAR(256),
    user_birthday VARCHAR(16),
    user_email VARCHAR(256),
    user_phone VARCHAR(64)
);