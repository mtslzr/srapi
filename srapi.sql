PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "sports" (
	"id"	char(2) NOT NULL,
	"name"	char(20) NOT NULL,
	"host"	char(50) NOT NULL,
	"standings"	char(50) NOT NULL,
	"teams"	char(50) NOT NULL,
	"years"	char(50) NOT NULL,
	PRIMARY KEY("id")
);
INSERT INTO sports VALUES('bs','Baseball','baseball-reference.com','leagues/MLB/{year}-standings.shtml','teams','leagues');
COMMIT;
