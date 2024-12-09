create database hiholive;
use hiholive;

-- Create tables
DROP TABLE IF EXISTS `auths`;
CREATE TABLE `auths` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `user_id` int NOT NULL,
                         `auth_type` enum('email_password','gmail','facebook') DEFAULT 'email_password',
                         `email` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
                         `salt` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
                         `password` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
                         `facebook_id` varchar(35) CHARACTER SET utf8mb4 DEFAULT NULL,
                         `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `email` (`email`) USING BTREE,
                         KEY `user_id` (`user_id`) USING BTREE,
                         KEY `facebook_id` (`facebook_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `phone_number` VARCHAR(255),
                         `address` VARCHAR(255),
                         `first_name` VARCHAR(255),
                         `last_name` VARCHAR(255),
                         `display_name` VARCHAR(255),
                         `date_of_birth` date,
                         `email` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
                         `gender` enum('male','female','other') NOT NULL DEFAULT 'other',
                         `system_role` ENUM('admin','viewer','streamer','moderator'),
                         `avatar` JSON,
                         `status` INT DEFAULT 1,
                         `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`),
                         KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `channels`;
CREATE TABLE `channels` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `user_id` INT NOT NULL,
                            `panel` JSON,
                            `description` TEXT,
                            `url` VARCHAR(2083),
                            `contact` VARCHAR(255),
                            `status` INT DEFAULT 1,
                            `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            PRIMARY KEY (`id`),
                            KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `channel_analytics`;
CREATE TABLE `channel_analytics` (
                                     `id` int NOT NULL AUTO_INCREMENT,
                                     `channel_id` int NOT NULL,
                                     `total_subscribers` INT,
                                     `total_views` INT,
                                     `total_likes` INT,
                                     `total_dislikes` INT,
                                     `total_comments` INT,
                                     `total_shares` INT,
                                     `average_view_duration` FLOAT,
                                     `engagement_rate` FLOAT,
                                     `revenue` FLOAT,
                                     `top_performing_video_id` int NOT NULL,
                                     `least_performing_video_id` int NOT NULL,
                                     `total_videos` INT,
                                     `total_stories` INT,
                                     `total_viewing_time` INT,
                                     `status` INT DEFAULT 1,
                                     `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                     PRIMARY KEY (`id`),
                                     KEY `status` (`status`) USING BTREE,
                                     KEY `channel_id` (`channel_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `name` VARCHAR(255),
                              `description` TEXT,
                              `image` JSON,
                              `status` INT DEFAULT 1,
                              `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`),
                              KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `channel_id` int NOT NULL,
                          `title` VARCHAR(255),
                          `categoryId` INT NOT NULL,
                          `url` VARCHAR(2083),
                          `duration` FLOAT,
                          `total_view` INT,
                          `total_like` INT,
                          `total_dislike` INT,
                          `thumbnail` JSON,
                          `description` TEXT,
                          `is_from_livestream` BOOLEAN,
                          `original_livestream_id` int NOT NULL,
                          `status` INT DEFAULT 1,
                          `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`),
                          KEY `status` (`status`) USING BTREE,
                          KEY `channel_id` (`channel_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `live_streams`;
CREATE TABLE `live_streams` (
                                `id` int NOT NULL AUTO_INCREMENT,

                                `channel_id` int NOT NULL,
                                `title` VARCHAR(255),
                                `description` TEXT,
                                `notification` TEXT,
                                `categoryId` INT,
                                `is_rerun` boolean default false,
                                `scheduled_start_time` TIMESTAMP,

                                `actual_start_time` TIMESTAMP,
                                `actual_end_time` TIMESTAMP,
                                `peak_concurrent_view` INT,
                                `total_unique_viewers` INT,
                                `state` ENUM('scheduled','pending', 'running', 'ended'),
                                `stream_key` BINARY(16) default (UUID_TO_BIN(UUID())),
                                `status` INT DEFAULT 1,
                                `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                PRIMARY KEY (`id`),
                                KEY `status` (`status`) USING BTREE,
                                KEY `state` (`state`) USING BTREE,
                                KEY `stream_key` (`stream_key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `live_stream_metric`;
CREATE TABLE `live_stream_metric` (
                                      `id` int NOT NULL AUTO_INCREMENT,
                                      `live_stream_id` int NOT NULL,
                                      `current_view` INT,
                                      `status` INT DEFAULT 1,
                                      `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                      PRIMARY KEY (`id`),
                                      KEY `status` (`status`) USING BTREE,
                                      KEY `live_stream_id` (`live_stream_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `playlists`;
CREATE TABLE `playlists` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `user_id` int NOT NULL,
                             `name` VARCHAR(255),
                             `image` JSON,
                             `total_video` INT,
                             `state` ENUM('public', 'private'),
                             `status` INT DEFAULT 1,
                             `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             PRIMARY KEY (`id`),
                             KEY `status` (`status`) USING BTREE,
                             KEY `user_id` (`user_id`) USING BTREE,
                             KEY `state` (`state`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `playlist_details`;
CREATE TABLE `playlist_details` (
                                    `id` int NOT NULL AUTO_INCREMENT,
                                    `playlist_id` int NOT NULL,
                                    `video_id` int NOT NULL,
                                    `status` INT DEFAULT 1,
                                    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`id`),
                                    KEY `status` (`status`) USING BTREE,
                                    KEY `playlist_id` (`playlist_id`) USING BTREE,
                                    KEY `video_id` (`video_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `stories`;
CREATE TABLE `stories` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `channel_id` int NOT NULL,
                           `url` VARCHAR(2083),
                           `duration` FLOAT,
                           `total_view` INT,
                           `total_like` INT,
                           `total_dislike` INT,
                           `thumbnail` JSON,
                           `status` INT DEFAULT 1,
                           `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`),
                           KEY `status` (`status`) USING BTREE,
                           KEY `channel_id` (`channel_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `subscribers`;
CREATE TABLE `subscribers` (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `user_id` int NOT NULL,
                               `channel_id` int NOT NULL,
                               `status` INT DEFAULT 1,
                               `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               PRIMARY KEY (`id`),
                               KEY `status` (`status`) USING BTREE,
                               KEY `user_id` (`user_id`) USING BTREE,
                               KEY `channel_id` (`channel_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `user_dislike_videos`;
CREATE TABLE `user_dislike_videos` (
                                       `id` int NOT NULL AUTO_INCREMENT,
                                       `user_id` int NOT NULL,
                                       `video_id` int NOT NULL,
                                       `status` INT DEFAULT 1,
                                       `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       PRIMARY KEY (`id`),
                                       KEY `status` (`status`) USING BTREE,
                                       KEY `user_id` (`user_id`) USING BTREE,
                                       KEY `video_id` (`video_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `user_like_videos`;
CREATE TABLE `user_like_videos` (
                                    `id` int NOT NULL AUTO_INCREMENT,
                                    `user_id` int NOT NULL,
                                    `video_id` int NOT NULL,
                                    `status` INT DEFAULT 1,
                                    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`id`),
                                    KEY `status` (`status`) USING BTREE,
                                    KEY `user_id` (`user_id`) USING BTREE,
                                    KEY `video_id` (`video_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `user_watch_videos`;
CREATE TABLE `user_watch_videos` (
                                     `id` int NOT NULL AUTO_INCREMENT,
                                     `user_id` int NOT NULL,
                                     `video_id` int NOT NULL,
                                     `duration` FLOAT,
                                     `status` INT DEFAULT 1,
                                     `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                     PRIMARY KEY (`id`),
                                     KEY `status` (`status`) USING BTREE,
                                     KEY `user_id` (`user_id`) USING BTREE,
                                     KEY `video_id` (`video_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `user_block_lists`;
CREATE TABLE `user_block_lists` (
                                    `id` int NOT NULL AUTO_INCREMENT,
                                    `user_id` int NOT NULL,
                                    `channel_id` int NOT NULL,
                                    `status` INT DEFAULT 1,
                                    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`id`),
                                    KEY `status` (`status`) USING BTREE,
                                    KEY `user_id` (`user_id`) USING BTREE,
                                    KEY `channel_id` (`channel_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `name` VARCHAR(255) UNIQUE,
                        `status` INT DEFAULT 1,
                        `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `video_tags`;
CREATE TABLE `video_tags` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `video_id` int NOT NULL,
                              `tag_id` int NOT NULL,
                              `status` INT DEFAULT 1,
                              `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`),
                              KEY `status` (`status`) USING BTREE,
                              KEY `video_id` (`video_id`) USING BTREE,
                              KEY `tag_id` (`tag_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



DROP TABLE IF EXISTS `broadcast_schedule`;
CREATE TABLE `broadcast_schedule` (
                                      `id` int NOT NULL AUTO_INCREMENT,
                                      `channel_id` int NOT NULL,
                                      `title` VARCHAR(255),
                                      `description` TEXT,
                                      `scheduled_start` TIMESTAMP,
                                      `scheduled_end` TIMESTAMP,
                                      `actual_start` TIMESTAMP,
                                      `actual_end` TIMESTAMP,
                                      `state` ENUM('scheduled', 'live', 'ended'),
                                      `status` INT DEFAULT 1,
                                      `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                      PRIMARY KEY (`id`),
                                      KEY `status` (`status`) USING BTREE,
                                      KEY `channel_id` (`channel_id`) USING BTREE,
                                      KEY `state` (`state`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `stream_highlights`;
CREATE TABLE `stream_highlights` (
                                     `id` int NOT NULL AUTO_INCREMENT,
                                     `live_stream_id` int NOT NULL,
                                     `title` VARCHAR(255),
                                     `start_time` FLOAT,
                                     `end_time` FLOAT,
                                     `description` TEXT,
                                     `status` INT DEFAULT 1,
                                     `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                     PRIMARY KEY (`id`),
                                     KEY `status` (`status`) USING BTREE,
                                     KEY `live_stream_id` (`live_stream_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

# Dynamo & OpenSearch
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `parent_id` int NOT NULL NULL,
                            `user_id` int NOT NULL,
                            `video_id` int NOT NULL,
                            `content` TEXT,
                            `state` ENUM('published', 'pending', 'removed'),
                            `status` INT DEFAULT 1,
                            `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            PRIMARY KEY (`id`),
                            KEY `status` (`status`) USING BTREE,
                            KEY `parent_id` (`parent_id`) USING BTREE,
                            KEY `state` (`state`) USING  BTREE,
                            KEY `user_id` (`user_id`) USING BTREE,
                            KEY `video_id` (`video_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `search_history`;
CREATE TABLE `search_history` (
                                  `id` int NOT NULL AUTO_INCREMENT,
                                  `user_id` int NOT NULL,
                                  `query` VARCHAR(255),
                                  `timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`id`),
                                  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `comments_reply`;
CREATE TABLE `comments_reply` (
                                  `id` int NOT NULL AUTO_INCREMENT,
                                  `parent_comment_id` int NOT NULL,
                                  `child_comment_id` int NOT NULL,
                                  `status` INT DEFAULT 1,
                                  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`id`),
                                  KEY `status` (`status`) USING BTREE,
                                  KEY `parent_comment_id` (`parent_comment_id`) USING BTREE,
                                  KEY `child_comment_id` (`child_comment_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `chat_messages`;
CREATE TABLE `chat_messages` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `user_id` int NOT NULL,
                                 `live_stream_id` int NOT NULL,
                                 `content` TEXT,
                                 `state` ENUM('published', 'pending', 'removed'),
                                 `status` INT DEFAULT 1,
                                 `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`),
                                 KEY `status` (`status`) USING BTREE,
                                 KEY `user_id` (`user_id`) USING BTREE,
                                 KEY `state` (`state`) USING  BTREE,
                                 KEY `live_stream_id` (`live_stream_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;