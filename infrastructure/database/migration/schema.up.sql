CREATE TABLE IF NOT EXISTS wallets (
    id INT(32) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    balance INT(32) NOT NULL DEFAULT 0,
    user_id INT(32) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);
