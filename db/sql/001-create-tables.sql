---- create ----
create table IF not exists `users`
(
 `id`               INT(20) AUTO_INCREMENT,
 `first_name`       VARCHAR(50) NOT NULL,
 `last_name`        VARCHAR(50) NOT NULL,
 `email`            VARCHAR(50) NOT NULL,
 `date_created`     Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
