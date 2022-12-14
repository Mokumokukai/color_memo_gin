
BEGIN;

CREATE TABLE IF NOT EXISTS memos(
    id CHAR (7) UNIQUE NOT NULL,
    owner_id CHAR (7) NOT NULL,
    creater_id CHAR (7),
    color1 CHAR (8) NOT NULL,
    color2 CHAR (8) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
); 

COMMIT;