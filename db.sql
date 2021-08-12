SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP DATABASE IF EXISTS `lbc`;
CREATE DATABASE `lbc` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `lbc`;

DROP TABLE IF EXISTS `ads`;
CREATE TABLE `ads` (
                       `id` int(11) NOT NULL AUTO_INCREMENT,
                       `title` tinytext NOT NULL,
                       `content` text NOT NULL,
                       `car` int(11) DEFAULT NULL,
                       PRIMARY KEY (`id`),
                       KEY `car` (`car`),
                       CONSTRAINT `ads_ibfk_2` FOREIGN KEY (`car`) REFERENCES `cars` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `brands`;
CREATE TABLE `brands` (
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `name` tinytext COLLATE utf8mb4_unicode_ci NOT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `brands` (`id`, `name`) VALUES
                                        (1,	'Audi'),
                                        (2,	'BMW'),
                                        (3,	'Citroen');

DROP TABLE IF EXISTS `cars`;
CREATE TABLE `cars` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `brand` int(11) NOT NULL,
                        `model` tinytext COLLATE utf8mb4_unicode_ci NOT NULL,
                        PRIMARY KEY (`id`),
                        KEY `brand` (`brand`),
                        FULLTEXT KEY `model_FT` (`model`),
                        CONSTRAINT `cars_ibfk_1` FOREIGN KEY (`brand`) REFERENCES `brands` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `cars` (`id`, `brand`, `model`) VALUES
                                                (1,	1,	'Cabriolet'),
                                                (2,	1,	'Q2'),
                                                (3,	1,	'Q3'),
                                                (4,	1,	'Q5'),
                                                (5,	1,	'Q7'),
                                                (6,	1,	'Q8'),
                                                (7,	1,	'R8'),
                                                (8,	1,	'Rs3'),
                                                (9,	1,	'Rs4'),
                                                (10,	1,	'Rs5'),
                                                (11,	1,	'Rs7'),
                                                (12,	1,	'S3'),
                                                (13,	1,	'S4'),
                                                (14,	1,	'S4 Avant'),
                                                (15,	1,	'S4 Cabriolet'),
                                                (16,	1,	'S5'),
                                                (17,	1,	'S7'),
                                                (18,	1,	'S8'),
                                                (19,	1,	'SQ5'),
                                                (20,	1,	'SQ7'),
                                                (21,	1,	'Tt'),
                                                (22,	1,	'Tts'),
                                                (23,	1,	'V8'),
                                                (24,	2,	'M3'),
                                                (25,	2,	'M4'),
                                                (26,	2,	'M5'),
                                                (27,	2,	'M535'),
                                                (28,	2,	'M6'),
                                                (29,	2,	'M635'),
                                                (30,	2,	'Serie 1'),
                                                (31,	2,	'Serie 2'),
                                                (32,	2,	'Serie 3'),
                                                (33,	2,	'Serie 4'),
                                                (34,	2,	'Serie 5'),
                                                (35,	2,	'Serie 6'),
                                                (36,	2,	'Serie 7'),
                                                (37,	2,	'Serie 8'),
                                                (38,	3,	'C1'),
                                                (39,	3,	'C15'),
                                                (40,	3,	'C2'),
                                                (41,	3,	'C25'),
                                                (42,	3,	'C25D'),
                                                (43,	3,	'C25E'),
                                                (44,	3,	'C25TD'),
                                                (45,	3,	'C3'),
                                                (46,	3,	'C3 Aircross'),
                                                (47,	3,	'C3 Picasso'),
                                                (48,	3,	'C4'),
                                                (49,	3,	'C4 Picasso'),
                                                (50,	3,	'C5'),
                                                (51,	3,	'C6'),
                                                (52,	3,	'C8'),
                                                (53,	3,	'Ds3'),
                                                (54,	3,	'Ds4'),
                                                (55,	3,	'Ds5');