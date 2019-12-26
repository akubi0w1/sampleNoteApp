-- define Database
DROP DATABASE `note_app`;
CREATE DATABASE IF NOT EXISTS `note_app`;

-- define Tables
DROP TABLE IF EXISTS `note_app`.`users`;
CREATE TABLE IF NOT EXISTS `note_app`.`users`(
    `id` VARCHAR(36) NOT NULL,
    -- `x-token` VARCHAR(128) UNIQUE,
    `name` VARCHAR(36) NOT NULL,
    `mail` VARCHAR(128) NOT NULL UNIQUE,
    `password` VARCHAR(128) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
)
ENGINE = InnoDB;

DROP TABLE IF EXISTS `note_app`.`notes`;
CREATE TABLE IF NOT EXISTS `note_app`.`notes`(
    `id` VARCHAR(36) NOT NULL,
    `title` VARCHAR(36) NOT NULL,
    `content` TEXT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `user_id` VARCHAR(36),
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_notes_user`
        FOREIGN KEY (`user_id`)
        REFERENCES `note_app`.`users` (`id`)
        ON DELETE SET NULL
        ON UPDATE CASCADE
)
ENGINE = InnoDB;

-- DROP TABLE IF EXISTS `note_app`.`sessions`;
-- CREATE TABLE IF NOT EXISTS `note_app`.`sessions`(
--     `user_id` VARCHAR(128) NOT NULL,
--     `session_id` VARCHAR(128),
--     PRIMARY KEY(`user_id`),
--     CONSTRAINT `fk_sessions_users`
--         FOREIGN KEY (`user_id`)
--         REFERENCES `note_app`.`users` (`id`)
--         ON DELETE CASCADE
--         ON UPDATE CASCADE
-- )
-- ENGINE = InnoDB;

