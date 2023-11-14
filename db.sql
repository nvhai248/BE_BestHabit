CREATE DATABASE besthabit;

USE besthabit;

-- Bảng User
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `email` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(20) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `name` VARCHAR(100) NOT NULL,
    `fb_id` VARCHAR(255),
    `gg_id` VARCHAR(255),
    `salt` VARCHAR(255),
    `avatar` json,
    `level` INT DEFAULT 1,
    `experience` INT DEFAULT 0,
    `settings` json,
    `role` enum('user', 'admin') DEFAULT 'user',
    `habit_count` INT DEFAULT 0,
    `task_count` INT DEFAULT 0,
    `challenge_count` INT DEFAULT 0,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Bảng Habits
DROP TABLE IF EXISTS habits;

CREATE TABLE habits (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `name` VARCHAR(100) NOT NULL,
    `description` TEXT,
    `start_date` DATETIME,
    `end_date` DATETIME,
    `type` enum(
        'health',
        'work_and_study',
        'spiritually_and_psychologically',
        'social_and_relational',
        'personal'
    ) DEFAULT 'personal',
    `days` json,
    -- ghi lại các ngày làm thói quen này theo kiểu {"monday":1, "tuesday":0,...}
    `is_count_based` BOOLEAN DEFAULT TRUE,
    -- kiểm tra xem là loại số lượng hay thời gian
    `reminder` TIME,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Bảng Tasks
DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `name` VARCHAR(100) NOT NULL,
    `description` TEXT,
    `deadline` DATETIME,
    `reminder` TIME,
    -- thời gian nhắc nhở, báo thức
    `status` enum('pending', 'completed', 'overdue') DEFAULT 'pending',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Bảng habit_logs
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

CREATE TABLE challenges (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL,
    `description` TEXT,
    `start_date` DATE,
    `end_date` DATE,
    `experience_point` INT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Bảng Participants
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

CREATE TABLE images (
    `id` int NOT NULL AUTO_INCREMENT,
    `url` varchar(255) DEFAULT NULL,
    `width` int DEFAULT NULL,
    `height` int DEFAULT NULL,
    `cloud_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `extension` varchar(255) DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 4 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci