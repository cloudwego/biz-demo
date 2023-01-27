create table `t_user`
(
    `id`         bigint unsigned auto_increment,
    `created_at` datetime(3) NULL,
    `updated_at` datetime(3) NULL,
    `deleted_at` datetime(3) NULL,
    `user_name`  varchar(128) NOT NULL DEFAULT '',
    `password`   varchar(128) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    KEY        `idx_username` (`user_name`) COMMENT 'username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='user account table';

create table `t_product`
(
    `id`          bigint unsigned auto_increment,
    `created_at`  datetime(3) NULL,
    `updated_at`  datetime(3) NULL,
    `deleted_at`  datetime(3) NULL,
    `product_id`  bigint(20) NOT NULL,
    `name`        varchar(255) NOT NULL DEFAULT '',
    `pic`         varchar(255) NOT NULL DEFAULT '',
    `description` text NULL,
    `isbn`        varchar(255) NOT NULL DEFAULT '',
    `spu_name`    varchar(255) NOT NULL DEFAULT '',
    `spu_price`   int(11) NOT NULL DEFAULT '0',
    `price`       int(11) NOT NULL DEFAULT '0',
    `stock`       int(11) NOT NULL DEFAULT '0',
    `status`      tinyint(4) NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY         `idx_product_id` (`product_id`) COMMENT 'product_id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='product table';

create table `t_order`
(
    `id`               bigint unsigned auto_increment,
    `created_at`       datetime(3) NULL,
    `updated_at`       datetime(3) NULL,
    `deleted_at`       datetime(3) NULL,
    `order_id`         bigint(20) NOT NULL,
    `user_id`          bigint NOT NULL,
    `address`          text NULL,
    `product_id`       bigint(20) NOT NULL,
    `stock_num`        int(11) NOT NULL DEFAULT '0',
    `product_snapshot` longtext NULL,
    `status`           tinyint(4) NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY              `idx_order_id` (`order_id`) COMMENT 'order_id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='order table';
