
CREATE DATABASE classroom CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE classroom;

-- Người dùng
CREATE TABLE IF NOT EXISTS `users` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_name` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `phone_number` VARCHAR(20),
    `gender` TINYINT DEFAULT 0 COMMENT '0: unknown, 1: male, 2: female',
    `full_name` VARCHAR(255) NOT NULL,
    `avatar` VARCHAR(255),
    `is_verified` BOOLEAN NOT NULL DEFAULT FALSE,
    `verification_code` VARCHAR(20),
    `role` INT NOT NULL COMMENT '1: teacher, 2: student',
    `create_time` BIGINT NOT NULL,
    `update_time` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);

-- Lớp học
CREATE TABLE IF NOT EXISTS `classes` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `class_code` VARCHAR(10) NOT NULL UNIQUE,
    `class_name` VARCHAR(100) NOT NULL,
    `description` TEXT,
    `teacher_id` BIGINT NOT NULL,
    `create_time` BIGINT NOT NULL,
    `update_time` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);

-- Ghi danh
CREATE TABLE IF NOT EXISTS `enrollments` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `student_id` BIGINT NOT NULL,
    `class_id` BIGINT NOT NULL,
    `join_time` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);

-- Môn học
CREATE TABLE IF NOT EXISTS `subjects` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `class_id` BIGINT NOT NULL,
    `subject_name` VARCHAR(100) NOT NULL,
    `description` TEXT,
    `create_time` BIGINT NOT NULL,
    `update_time` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);

-- Bài học
CREATE TABLE IF NOT EXISTS `lessons` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `subject_id` BIGINT NOT NULL,
    `title` VARCHAR(200) NOT NULL,
    `content` TEXT,
    `file_url` VARCHAR(500),
    `upload_time` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);

-- Đầu điểm
CREATE TABLE IF NOT EXISTS `grade_components` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `subject_id` BIGINT NOT NULL,
    `component_name` VARCHAR(100) NOT NULL,
    `weight` DECIMAL(5,2) NOT NULL,
    PRIMARY KEY (`id`)
);

-- Điểm số
CREATE TABLE IF NOT EXISTS `grades` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `student_id` BIGINT NOT NULL,
    `component_id` BIGINT NOT NULL,
    `score` DECIMAL(5,2) NOT NULL,
    `grade_time` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);

-- Xác thực email
CREATE TABLE IF NOT EXISTS `email_confirmations` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL,
    `confirmation_code` VARCHAR(255) NOT NULL,
    `is_verified` BOOLEAN DEFAULT FALSE,
    `sent_time` BIGINT NOT NULL,
    `verified_time` BIGINT NOT NULL,
    PRIMARY KEY (`id`)
);
