CREATE DATABASE besthabit;

USE besthabit;

-- Bảng User
DROP TABLE IF EXISTS users;

CREATE TABLE `users` (
    `id` int NOT NULL AUTO_INCREMENT,
    `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `name` varchar(100) NOT NULL,
    `fb_id` varchar(255) DEFAULT NULL,
    `gg_id` varchar(255) DEFAULT NULL,
    `salt` varchar(255) DEFAULT NULL,
    `avatar` json DEFAULT NULL,
    `level` int DEFAULT '1',
    `experience` int DEFAULT '0',
    `settings` json DEFAULT NULL,
    `role` enum('user', 'admin') DEFAULT 'user',
    `habit_count` int DEFAULT '0',
    `task_count` int DEFAULT '0',
    `challenge_count` int DEFAULT '0',
    `status` int DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci -- Bảng Habits
DROP TABLE IF EXISTS habits;

CREATE TABLE `habits` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `name` varchar(100) NOT NULL,
    `description` text,
    `start_date` datetime DEFAULT NULL,
    `end_date` datetime DEFAULT NULL,
    `type` enum(
        'health',
        'work_and_study',
        'spiritually_and_psychologically',
        'social_and_relational',
        'personal'
    ) DEFAULT 'personal',
    `days` json DEFAULT NULL,
    `is_count_based` tinyint(1) DEFAULT '1',
    `reminder` time DEFAULT NULL,
    `status` int DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci -- Bảng Tasks
DROP TABLE IF EXISTS tasks;

CREATE TABLE `tasks` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `name` varchar(100) NOT NULL,
    `description` text,
    `deadline` datetime DEFAULT NULL,
    `reminder` timestamp NULL DEFAULT NULL,
    `status` enum('pending', 'completed', 'overdue', 'deleted') DEFAULT 'pending',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 8 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci -- Bảng habit_logs
DROP TABLE IF EXISTS habit_logs;

CREATE TABLE habit_logs (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `habit_id` INT,
    `complete_day` json,
    -- lưu lại các ngày hoàn thành dưới dạng json {'d1':'timestamp string', 'd2':'timestamp string', 'd3', 'd4',...}
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Bảng Challenges
DROP TABLE IF EXISTS challenges;

CREATE TABLE `challenges` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `description` text,
    `start_date` date DEFAULT NULL,
    `end_date` date DEFAULT NULL,
    `experience_point` int DEFAULT NULL,
    `status` int DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci -- Bảng Participants
DROP TABLE IF EXISTS participants;

CREATE TABLE participants (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `challenge_id` INT NOT NULL,
    `status` enum('joined', 'completed', 'failed') DEFAULT 'joined',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Bảng images
DROP TABLE IF EXISTS images;

CREATE TABLE `images` (
    `id` int NOT NULL AUTO_INCREMENT,
    `url` varchar(255) DEFAULT NULL,
    `width` int DEFAULT NULL,
    `height` int DEFAULT NULL,
    `cloud_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `extension` varchar(255) DEFAULT NULL,
    `created_by` int NOT NULL,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 4 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci