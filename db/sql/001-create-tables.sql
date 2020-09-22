---- create ----
create table IF not exists `users`
(
 `id`               BIGINT(20) AUTO_INCREMENT NOT NULL,
 `first_name`       VARCHAR(45),
 `last_name`        VARCHAR(45),
 `email`            VARCHAR(45) UNIQUE NOT NULL,
 `date_created`     VARCHAR(45),
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
