CREATE TABLE IF NOT EXISTS entries (
	id INT PRIMARY KEY NOT NULL,
	name VARCHAR(255) UNIQUE,
	ctime DATETIME NOT NULL,
	atime DATETIME NOT NULL,
	abbrs TEXT,
	caption TEXT
);

CREATE TABLE IF NOT EXISTS comments (
	id INT PRIMARY KEY NOT NULL,
	eid INT NOT NULL,
	ctime DATETIME NOT NULL,
	atime DATETIME NOT NULL,
	user VARCHAR(255),
	email VARCHAR(255),
	url TEXT,
	title TEXT,
	content TEXT
);