-- ----------------------------------
--  Table structure for subscriptions
-- ----------------------------------
DROP TABLE IF EXISTS "subscriptions";
CREATE TABLE "subscriptions" (
	id serial primary key,
	name varchar(255) NOT NULL,
	email varchar(100) NOT NULL
)