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