
CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password', 
    `follow_count` int unsigned Not NULL DEFAULT 0 COMMENT 'User follow count' ,
    `follower_count` int unsigned Not NULL DEFAULT 0 COMMENT 'User follower count',
    `avatar`   varchar(128) Not NULL DEFAULT '' COMMENT 'Avatar',
    `background_image`   varchar(128) NOT NULL DEFAULT '' COMMENT 'background image',
    `signature`   varchar(128) NOT NULL DEFAULT '' COMMENT 'signature',
    `total_favorited`  int unsigned Not NULL DEFAULT 0 COMMENT 'total favorited',
    `work_count` int unsigned Not NULL DEFAULT 0 COMMENT 'User work count',
    `favorite_count` int unsigned Not NULL DEFAULT 0 COMMENT 'User favorite count',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `video`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `author`      bigint unsigned NOT NULL COMMENT 'author id',
    `play_url`    varchar(128) NOT NULL DEFAULT '' COMMENT 'video play url',
    `cover_url`   varchar(128) NOT NULL DEFAULT '' COMMENT 'vidoe cover url',
    `title`       varchar(128) NOT NULL DEFAULT '' COMMENT 'video title',
    `favorite_count` int unsigned Not NULL DEFAULT 0 COMMENT 'video favorite count' ,
    `comment_count` int unsigned Not NULL DEFAULT 0 COMMENT 'video comment count' ,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'video upload time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'vidoe update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'video delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_author_id` (`author`) COMMENT 'Author id index',
    UNIQUE KEY   (`created_at`, `id`),
    CONSTRAINT   `author_id` FOREIGN KEY (`author`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video information table';

CREATE TABLE `comment`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `video` bigint unsigned NOT NULL COMMENT 'video id',
    `user` bigint unsigned NOT NULL COMMENT 'user id',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'comment upload time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'comment update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'comment delete time',
    `content` varchar(200) NOT NULL DEFAULT '' COMMENT 'comment content',
    `date` varchar(5) NOT NULL DEFAULT '01-01' COMMENT 'comment create date',
    PRIMARY KEY (`id`),
    KEY          `idx_video_id` (`video`) COMMENT 'Video id index',
    CONSTRAINT   `video_id1` FOREIGN KEY (`video`) REFERENCES `video` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT   `user_id1` FOREIGN KEY (`user`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Comment information table';

CREATE TABLE `relation`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `follow_id`   bigint unsigned NOT NULL COMMENT '被关注的用户ID',
    `follower_id` bigint unsigned NOT NULL COMMENT '粉丝的用户ID',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'video upload time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'vidoe update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY   `idx_follow_id_follower_id` (`follow_id`, `follower_id`) COMMENT 'follow id and follower id index',
    KEY          (`follower_id`) COMMENT 'follower id index',
    CONSTRAINT   `follow` FOREIGN KEY (`follow_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT   `follower` FOREIGN KEY (`follower_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User relation table';

CREATE TABLE `favorite`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`   bigint unsigned NOT NULL COMMENT 'user id',
    `video_id`   bigint unsigned NOT NULL COMMENT 'video id',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'video upload time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'vidoe update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY          `idx_user_id_video_id` (`user_id`, `video_id`) COMMENT 'User id and Video id index',
    CONSTRAINT   `user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT   `video_id` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favorite information table';

CREATE TABLE `message`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `from_user_id`   bigint unsigned NOT NULL COMMENT 'from_user_id',
    `to_user_id`   bigint unsigned NOT NULL COMMENT 'to_user_id',
    `content`   varchar(128) NOT NULL DEFAULT '' COMMENT 'content',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'video upload time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id_video_id` (`from_user_id`, `to_user_id`) COMMENT 'User id and Video id index',
    KEY (`created_at`),
    CONSTRAINT   `from_user_id` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT   `to_user_id` FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Messsage information table';