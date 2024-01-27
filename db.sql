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
    `device_tokens` json DEFAULT NULL,
    `status` int DEFAULT '-2',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 17 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci -- Bảng Habits
DROP TABLE IF EXISTS habits;

CREATE TABLE `habits` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `name` varchar(100) NOT NULL,
    `description` text,
    `start_date` date DEFAULT NULL,
    `end_date` date DEFAULT NULL,
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
    `target` json DEFAULT NULL,
    `completed_dates` json DEFAULT NULL,
    `status` int DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 46 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci -- Bảng Tasks
DROP TABLE IF EXISTS tasks;

CREATE TABLE `tasks` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `name` varchar(100) NOT NULL,
    `description` text,
    `deadline` date DEFAULT NULL,
    `reminder` timestamp NULL DEFAULT NULL,
    `status` enum('pending', 'completed', 'overdue', 'deleted') DEFAULT 'pending',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 46 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci -- Bảng habit_logs
DROP TABLE IF EXISTS habit_logs;

-- Bảng Challenges
DROP TABLE IF EXISTS challenges;

CREATE TABLE `challenges` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `description` text,
    `start_date` date DEFAULT NULL,
    `end_date` date DEFAULT NULL,
    `experience_point` int DEFAULT '0',
    `count_user_joined` int DEFAULT '0',
    `status` tinyint(1) DEFAULT '1',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci -- Bảng Participants
DROP TABLE IF EXISTS participants;

CREATE TABLE `participants` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `challenge_id` int NOT NULL,
    `status` enum('joined', 'completed', 'failed', 'cancel') DEFAULT 'joined',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci -- Bảng images
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
) ENGINE = InnoDB AUTO_INCREMENT = 4 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci DROP TABLE IF EXISTS `cron_notice_tasks`;

create table `cron_notice_tasks` (
    `user_id` int not null,
    `entry_id` int not null,
    `task_id` int not null,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS `cron_notice_habits`;

create table `cron_notice_habits` (
    `user_id` int not null,
    `entry_id` int not null,
    `habit_id` int not null,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
)