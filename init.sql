-- データベースへの接続
\c sample_db;

-- テーブルの作成
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);

-- 初期データの挿入
INSERT INTO users (name, email) VALUES
('John Doe', 'john.doe@example.com'),
('Jane Doe2', 'jane2.doe@example.com'),
('Jane Doe3', 'jane3.doe@example.com');
