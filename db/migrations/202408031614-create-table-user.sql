-- +migrate Up
CREATE TABLE "users" (
   user_id UUID PRIMARY KEY,
   username VARCHAR(255) NOT NULL,
   password VARCHAR(255) NOT NULL,
   email VARCHAR(255) NOT NULL UNIQUE,
   phone_number VARCHAR(255) NOT NULL UNIQUE,
   avatar_link VARCHAR(255),
   role_code INT NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   is_deleted BOOLEAN DEFAULT FALSE
);

-- +migrate Down
DROP TABLE "users";