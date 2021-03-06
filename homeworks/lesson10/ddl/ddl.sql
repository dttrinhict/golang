#Postgress
DROP TABLE IF EXISTS golang.users;
CREATE TABLE golang.users (
                              id text PRIMARY KEY,
                              name text NOT NULL,
                              email text,
                              mobile text
);

DROP TABLE IF EXISTS golang.roles;
CREATE TABLE golang.roles (
                              id text PRIMARY KEY,
                              name text NOT NULL UNIQUE
);
INSERT INTO golang.roles (id, name) VALUES ('1', 'ADMIN');
INSERT INTO golang.roles (id, name) VALUES ('2', 'STUDENT');
INSERT INTO golang.roles (id, name) VALUES ('3', 'TRAINER');
INSERT INTO golang.roles (id, name) VALUES ('4', 'SALE');
INSERT INTO golang.roles (id, name) VALUES ('5', 'EMPLOYER');
INSERT INTO golang.roles (id, name) VALUES ('6', 'AUTHOR');
INSERT INTO golang.roles (id, name) VALUES ('7', 'EDITOR');
INSERT INTO golang.roles (id, name) VALUES ('8', 'MAINTAINER');

DROP TABLE IF EXISTS golang.user_role;
CREATE TABLE golang.user_role (
                                  user_id text REFERENCES golang.users(id),
                                  role_id text REFERENCES golang.roles(id)
);

DROP INDEX user_idx;
CREATE UNIQUE INDEX user_idx ON golang.user_role (user_id, role_id);





####MySQL

USE golang;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` VARCHAR(40) NOT NULL,
                         `name` VARCHAR(200) NOT NULL,
                         `email` VARCHAR(200) NOT NULL,
                         `mobile` VARCHAR(20) NOT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE INDEX `email_UNIQUE` (`email` ASC),
                         UNIQUE INDEX `mobile_UNIQUE` (`mobile` ASC)
);


DROP TABLE IF EXISTS `members`;
CREATE TABLE `members` (
                         `id` VARCHAR(40) NOT NULL,
                         `name` VARCHAR(200) NOT NULL,
                         `email` VARCHAR(200) NOT NULL,
                         `mobile` VARCHAR(20) NOT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE INDEX `email_UNIQUE` (`email` ASC),
                         UNIQUE INDEX `mobile_UNIQUE` (`mobile` ASC)
);

DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
                         `id` VARCHAR(40) NOT NULL,
                         `name` VARCHAR(200) NOT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE INDEX `name_UNIQUE` (`name` ASC)
);
INSERT INTO `roles` (`id`, `name`) VALUES ('1', 'ADMIN');
INSERT INTO `roles` (`id`, `name`) VALUES ('2', 'STUDENT');
INSERT INTO `roles` (`id`, `name`) VALUES ('3', 'TRAINER');
INSERT INTO `roles` (`id`, `name`) VALUES ('4', 'SALE');
INSERT INTO `roles` (`id`, `name`) VALUES ('5', 'EMPLOYER');
INSERT INTO `roles` (`id`, `name`) VALUES ('6', 'AUTHOR');
INSERT INTO `roles` (`id`, `name`) VALUES ('7', 'EDITOR');
INSERT INTO `roles` (`id`, `name`) VALUES ('8', 'MAINTAINER');

DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
     `user_id` VARCHAR(40) NOT NULL,
     `role_id` VARCHAR(40) NOT NULL,
     PRIMARY KEY (`user_id`, `role_id`)
);

DROP INDEX user_idx;
CREATE UNIQUE INDEX `user_idx` ON `user_role` (`user_id`, `role_id`);


DROP TABLE IF EXISTS `clubs`;
CREATE TABLE `clubs` (
                         `id` VARCHAR(40) NOT NULL,
                         `name` VARCHAR(200) NOT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE INDEX `name_UNIQUE` (`name` ASC)
);

DROP TABLE IF EXISTS `user_club`;
CREATE TABLE `user_club` (
                             `user_id` VARCHAR(40) NOT NULL,
                             `club_id` VARCHAR(40) NOT NULL,
                             PRIMARY KEY (`user_id`, `club_id`)
);

DROP INDEX user_club_idx;
CREATE UNIQUE INDEX `user_club_idx` ON `user_club` (`user_id`, `club_id`);

DROP TABLE IF EXISTS `member_club`;
CREATE TABLE `member_club` (
                             `member_id` VARCHAR(40) NOT NULL,
                             `club_id` VARCHAR(40) NOT NULL,
                             PRIMARY KEY (`member_id`, `club_id`)
);

DROP INDEX member_club_idx;
CREATE UNIQUE INDEX `member_club_idx` ON `member_club` (`member_id`, `club_id`);