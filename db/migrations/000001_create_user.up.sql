CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    user_id    BIGINT UNIQUE,
    email      VARCHAR(255) UNIQUE,
    salt       VARCHAR(255)
    );

-- 创建索引以提高根据salt字段查询的性能
CREATE INDEX IF NOT EXISTS idx_users_salt ON users (salt);