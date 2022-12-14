
BEGIN;

CREATE TABLE IF NOT EXISTS tags(
    id CHAR (7) UNIQUE NOT NULL,
    name VARCHAR (15) UNIQUE NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
); 

CREATE TABLE IF NOT EXISTS memo_tags(
    tag_id CHAR (7) NOT NULL,
    memo_id CHAR (7) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (memo_id) REFERENCES memos (id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE,
    UNIQUE (tag_id, memo_id)

);

COMMIT;
