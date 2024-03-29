CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
	updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);

CREATE TRIGGER IF NOT EXISTS trg_users_upd AFTER UPDATE ON users
BEGIN
    UPDATE users
    SET updated_at = DATETIME('now', 'localtime') 
    WHERE rowid == NEW.rowid;
END;