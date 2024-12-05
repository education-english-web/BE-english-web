-- +migrate Up
CREATE TABLE user_tokens (
   id SERIAL PRIMARY KEY,
   user_id UUID UNIQUE,
   refresh_token TEXT,
   expired_at INT,
   created_at TIMESTAMP(3),
   updated_at TIMESTAMP(3),
   CONSTRAINT fk_user_tokens_user_id FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE user_tokens;
