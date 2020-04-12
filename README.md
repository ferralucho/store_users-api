# store_users-api
user api for store in golang

set environment variables:

mysql_users_username youruser
mysql_users_password yourpassword
mysql_users_host 127.0.0.1
mysql_users_schema users_db

CREATE TABLE `users_db`.`users` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `email` VARCHAR(45) NOT NULL,
  `date_created` VARCHAR(45) NULL,
  `password` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC));


SELECT * FROM users_db.users;

POSTMAN REST:

CREATE USER

curl --location --request POST 'localhost:8080/users' \
--header 'Content-Type: text/plain' \
--data-raw '{
	"first_name": "Luciano",
	"last_name": "Ferrari",
	"email": "lucianodarioferrari@gmail.com",
	"password": "password1"
}'
