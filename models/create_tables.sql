CREATE TABLE users (
                       id INT PRIMARY KEY AUTO_INCREMENT,
                       user_id BIGINT UNIQUE NOT NULL,
                       username VARCHAR(64) UNIQUE NOT NULL,
                       password VARCHAR(64) NOT NULL,
                       is_admin BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
CREATE INDEX idx_is_admin ON users(is_admin);


CREATE TABLE `scripts` (
                           `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '脚本ID', -- 加UNSIGNED，规范主键类型
                           `name` VARCHAR(64) NOT NULL COMMENT '脚本名称',
                           `description` VARCHAR(500) DEFAULT '' COMMENT '脚本描述',
                           `owner_id` BIGINT NOT NULL COMMENT '所有者用户ID（关联users.user_id）',
                           `content` LONGTEXT NOT NULL COMMENT '脚本内容',
                           `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',


                           PRIMARY KEY (`id`),
                           UNIQUE KEY `uk_script_name` (`name`) COMMENT '脚本名称唯一，避免重复', -- 新增唯一约束
                           CONSTRAINT `fk_scripts_owner` FOREIGN KEY (`owner_id`)
                               REFERENCES `users` (`user_id`)
                               ON DELETE RESTRICT
                               ON UPDATE CASCADE,

                           INDEX `idx_owner_id` (`owner_id`),
                           INDEX `idx_created_at` (`created_at`),
                           INDEX `idx_updated_at` (`updated_at`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='脚本表'; -- 优化排序规则

CREATE TABLE `tasks` (
                         `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '任务唯一ID',
                         `script_id` int unsigned NOT NULL COMMENT '关联的脚本ID（外键）',
                         `task_name` varchar(64) DEFAULT NULL COMMENT '任务名称（如“脚本执行-20251205”）',
                         `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '任务状态：0=待执行，1=执行中，2=执行成功，3=执行失败',
                         `log_content` longtext DEFAULT NULL COMMENT '任务执行日志（存储输出、错误信息等）',
                         `user_id` bigint NOT NULL COMMENT '任务创建者（关联用户表user_id）',
                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '任务创建时间',
                         `executed_at` timestamp NULL DEFAULT NULL COMMENT '任务执行完成时间（成功/失败后填充）',
                         PRIMARY KEY (`id`),
    -- 外键关联：确保任务关联的脚本和用户必须存在
                         CONSTRAINT `fk_tasks_script_id` FOREIGN KEY (`script_id`) REFERENCES `scripts` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
                         CONSTRAINT `fk_tasks_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
    -- 索引：优化查询（按脚本查任务、按用户查任务、按状态查任务）
                         INDEX `idx_script_id` (`script_id`),
                         INDEX `idx_created_by` (`user_id`),
                         INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='脚本任务表';