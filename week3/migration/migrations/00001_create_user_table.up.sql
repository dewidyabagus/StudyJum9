CREATE TABLE IF NOT EXISTS `users` (
    `id`            bigint unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `email`         varchar(100) NOT NULL,
    `first_name`    varchar(100) NOT NULL,
    `last_name`     varchar(100) NOT NULL,
    `address`       varchar(200) NOT NULL,
    `password`      varchar(128) NOT NULL,
    `created_at`    datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    datetime NULL ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`    datetime
) ENGINE=InnoDB AUTO_INCREMENT=1;
