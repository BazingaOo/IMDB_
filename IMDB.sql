/*
 Navicat Premium Data Transfer

 Source Server         : IMDB
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : imdb

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 19/04/2022 17:03:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cast
-- ----------------------------
DROP TABLE IF EXISTS `cast`;
CREATE TABLE `cast`  (
  `cast_id` int NOT NULL AUTO_INCREMENT,
  `cast_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `cast_description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `cast_image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`cast_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cast
-- ----------------------------
INSERT INTO `cast` VALUES (1, 'Tao', 'master of actors', 'l');
INSERT INTO `cast` VALUES (2, 'zhang', 'PHD', NULL);
INSERT INTO `cast` VALUES (3, 'test', 'She is a good actress.', 'Tom.png');
INSERT INTO `cast` VALUES (4, 'Leonardo DiCaprio', 'One of the most famous actor', 'leonardo.png');
INSERT INTO `cast` VALUES (5, 'Brie Larson', 'Brie Larson has built an impressive career as an acclaimed television actress, rising feature film star and emerging recording artist.', 'brie-larson.png');

-- ----------------------------
-- Table structure for genre
-- ----------------------------
DROP TABLE IF EXISTS `genre`;
CREATE TABLE `genre`  (
  `genre_id` int NOT NULL AUTO_INCREMENT,
  `genre` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`genre_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of genre
-- ----------------------------

-- ----------------------------
-- Table structure for movie
-- ----------------------------
DROP TABLE IF EXISTS `movie`;
CREATE TABLE `movie`  (
  `movie_id` int NOT NULL AUTO_INCREMENT,
  `movie_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `year` int NULL DEFAULT NULL,
  `grade` float NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`movie_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of movie
-- ----------------------------
INSERT INTO `movie` VALUES (1, 'Avengers-Endgame', 2019, 84, 'After the devastating events of Avengers: Infinity War (2018), the universe is in ruins. With the help of remaining allies, the Avengers assemble once more in order to reverse Thanos\' actions and restore balance to the universe.', 'Avengers-Endgame.jpg');
INSERT INTO `movie` VALUES (2, 'Joker', 2019, 84, 'A mentally troubled stand-up comedian embarks on a downward spiral that leads to the creation of an iconic villain.', 'Joker.jpg');
INSERT INTO `movie` VALUES (3, 'Black Panther', 2018, 73, 'T\'Challa, heir to the hidden but advanced kingdom of Wakanda, must step forward to lead his people into a new future and must confront a challenger from his country\'s past.', 'black-panther.jpg');
INSERT INTO `movie` VALUES (4, 'test', 2001, 98, 'awesome!', 'logo.png');
INSERT INTO `movie` VALUES (5, 'test1', 2222, 10, '111111', 'Avengers-Endgame.jpg');
INSERT INTO `movie` VALUES (6, 'Captain Marvel', 2019, 68, 'Carol Danvers becomes one of the universe\'s most powerful heroes when Earth is caught in the middle of a galactic war between two alien races.', 'Captain-Marvel.jpg');
INSERT INTO `movie` VALUES (7, '\r\nDjango Unchained', 2012, 84, 'With the help of a German bounty-hunter, a freed slave sets out to rescue his wife from a brutal plantation-owner in Mississippi.', 'Django.jpg');
INSERT INTO `movie` VALUES (8, 'CODA', 2021, 80, 'As a CODA (Child of Deaf Adults) Ruby is the only hearing person in her deaf family. When the family\'s fishing business is threatened, Ruby finds herself torn between pursuing her passion at Berklee College of Music and her fear of abandoning her parents.', 'CODA.jpg');
INSERT INTO `movie` VALUES (9, 'Forrest Gump', 1984, 88, 'The presidencies of Kennedy and Johnson, the Vietnam War, the Watergate scandal and other historical events unfold from the perspective of an Alabama man with an IQ of 75, whose only desire is to be reunited with his childhood sweetheart.', 'forrest.jpg');
INSERT INTO `movie` VALUES (10, 'Inception', 2012, 88, 'A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a C.E.O., but his tragic past may doom the project and his team to disaster.', 'inception.jpg');
INSERT INTO `movie` VALUES (11, 'Interstellar', 2014, 86, 'A team of explorers travel through a wormhole in space in an attempt to ensure humanity\'s survival.', 'interstellar.jpg');
INSERT INTO `movie` VALUES (12, 'Pulp Fiction', 1994, 89, 'The lives of two mob hitmen, a boxer, a gangster and his wife, and a pair of diner bandits intertwine in four tales of violence and redemption.', 'Pulp-Fiction1.jpg');
INSERT INTO `movie` VALUES (13, 'Schindler\'s List', 1993, 90, 'In German-occupied Poland during World War II, industrialist Oskar Schindler gradually becomes concerned for his Jewish workforce after witnessing their persecution by the Nazis.', 'SchindlerList.jpg');
INSERT INTO `movie` VALUES (14, 'Spider-Man: No Way Home', 2021, 84, 'With Spider-Man\'s identity now revealed, Peter asks Doctor Strange for help. When a spell goes wrong, dangerous foes from other worlds start to appear, forcing Peter to discover what it truly means to be Spider-Man.', 'Spider-Man3.jpg');
INSERT INTO `movie` VALUES (15, 'The Dark Knight', 2008, 90, 'When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests of his ability to fight injustice.', 'TheDarkKnight.jpg');
INSERT INTO `movie` VALUES (16, 'The Godfather', 1972, 92, 'The aging patriarch of an organized crime dynasty in postwar New York City transfers control of his clandestine empire to his reluctant youngest son.', 'TheGodfather.jpg');
INSERT INTO `movie` VALUES (17, 'The Lord of the Rings: The Return of the King', 2003, 90, 'Gandalf and Aragorn lead the World of Men against Sauron\'s army to draw his gaze from Frodo and Sam as they approach Mount Doom with the One Ring.', 'TheLord.jpg');
INSERT INTO `movie` VALUES (18, 'The Shawshank Redemption', 1994, 93, 'Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.', 'theShawshankRedemption.jpg');
INSERT INTO `movie` VALUES (19, 'Titanic', 1997, 79, 'A seventeen-year-old aristocrat falls in love with a kind but poor artist aboard the luxurious, ill-fated R.M.S. Titanic.', 'titanic.jpg');
INSERT INTO `movie` VALUES (20, 'The Wolf of Wall Street', 2013, 82, 'Based on the true story of Jordan Belfort, from his rise to a wealthy stock-broker living the high life to his fall involving crime, corruption and the federal government.', 'wolf.jpg');

-- ----------------------------
-- Table structure for movie_cast
-- ----------------------------
DROP TABLE IF EXISTS `movie_cast`;
CREATE TABLE `movie_cast`  (
  `cast_id` int NULL DEFAULT NULL,
  `movie_id` int NULL DEFAULT NULL,
  INDEX `cast_id`(`cast_id` ASC) USING BTREE,
  INDEX `movie_id`(`movie_id` ASC) USING BTREE,
  CONSTRAINT `movie_cast_ibfk_1` FOREIGN KEY (`cast_id`) REFERENCES `cast` (`cast_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `movie_cast_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movie` (`movie_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of movie_cast
-- ----------------------------
INSERT INTO `movie_cast` VALUES (1, 1);
INSERT INTO `movie_cast` VALUES (2, 1);
INSERT INTO `movie_cast` VALUES (3, 5);

-- ----------------------------
-- Table structure for movie_genre
-- ----------------------------
DROP TABLE IF EXISTS `movie_genre`;
CREATE TABLE `movie_genre`  (
  `movie_id` int NULL DEFAULT NULL,
  `genre_id` int NULL DEFAULT NULL,
  INDEX `movie_id`(`movie_id` ASC) USING BTREE,
  INDEX `genre_id`(`genre_id` ASC) USING BTREE,
  CONSTRAINT `movie_genre_ibfk_1` FOREIGN KEY (`movie_id`) REFERENCES `movie` (`movie_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `movie_genre_ibfk_2` FOREIGN KEY (`genre_id`) REFERENCES `genre` (`genre_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of movie_genre
-- ----------------------------

-- ----------------------------
-- Table structure for rating
-- ----------------------------
DROP TABLE IF EXISTS `rating`;
CREATE TABLE `rating`  (
  `user_id` int NULL DEFAULT NULL,
  `movie_id` int NULL DEFAULT NULL,
  `score` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `movie_id`(`movie_id` ASC) USING BTREE,
  CONSTRAINT `rating_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `rating_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movie` (`movie_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of rating
-- ----------------------------
INSERT INTO `rating` VALUES (NULL, 1, '5');

-- ----------------------------
-- Table structure for review
-- ----------------------------
DROP TABLE IF EXISTS `review`;
CREATE TABLE `review`  (
  `review_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NULL DEFAULT NULL,
  `movie_id` int NULL DEFAULT NULL,
  `review_content` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`review_id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `movie_id`(`movie_id` ASC) USING BTREE,
  CONSTRAINT `review_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `review_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movie` (`movie_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of review
-- ----------------------------
INSERT INTO `review` VALUES (1, NULL, 1, 'This movie is good.');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_type` int NULL DEFAULT NULL,
  `gender` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `age` int NULL DEFAULT NULL,
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'qwe', '123', NULL, NULL, NULL, NULL);
INSERT INTO `user` VALUES (2, 'qwe2', '123456', NULL, 'female', 12, '123@qwe.com');

-- ----------------------------
-- Table structure for watch_list
-- ----------------------------
DROP TABLE IF EXISTS `watch_list`;
CREATE TABLE `watch_list`  (
  `user_id` int NULL DEFAULT NULL,
  `movie_id` int NULL DEFAULT NULL,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `movie_id`(`movie_id` ASC) USING BTREE,
  CONSTRAINT `watch_list_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `watch_list_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movie` (`movie_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of watch_list
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
