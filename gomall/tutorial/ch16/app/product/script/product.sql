CREATE TABLE `category`
(
    `id`          int          NOT NULL AUTO_INCREMENT,
    `name`        varchar(50)  NOT NULL,
    `description` varchar(255) NOT NULL,
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `category`
VALUES (1, 'Clothes', 'Clothes', '2023-12-06 15:05:06', '2023-12-06 15:05:06'),
       (2, 'Other', 'Other', '2023-12-06 15:05:06', '2023-12-06 15:05:06');
CREATE TABLE `product`
(
    `id`          int            NOT NULL AUTO_INCREMENT,
    `name`        varchar(50)    NOT NULL,
    `description` varchar(255)   NOT NULL,
    `picture`     varchar(255)   NOT NULL,
    `price`       decimal(10, 2) NOT NULL,
    `created_at`  datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `product`
VALUES (1, 'Notebook',
        'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ',
        '/static/image/notebook.jpeg', 9.90, '2023-12-06 15:26:19', '2023-12-09 22:29:10'),
       (2, 'Mouse-Pad',
        'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ',
        '/static/image/mouse-pad.jpeg', 8.80, '2023-12-06 15:26:19', '2023-12-09 22:29:59'),
       (3, 'T-Shirt',
        'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.',
        '/static/image/t-shirt.jpeg', 6.60, '2023-12-06 15:26:19', '2023-12-09 22:31:20'),
       (4, 'T-Shirt',
        'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.',
        '/static/image/t-shirt-1.jpeg', 2.20, '2023-12-06 15:26:19', '2023-12-09 22:31:20'),
       (5, 'Sweatshirt',
        'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.',
        '/static/image/sweatshirt.jpeg', 1.10, '2023-12-06 15:26:19', '2023-12-09 22:32:35'),
       (6, 'T-Shirt',
        'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.',
        '/static/image/t-shirt-2.jpeg', 1.80, '2023-12-06 15:26:19', '2023-12-09 22:31:20'),
       (10, 'mascot',
        'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.',
        '/static/image/logo.jpg', 4.80, '2023-12-06 15:26:19', '2023-12-09 22:39:47');
CREATE TABLE `product_category`
(
    `id`          int      NOT NULL AUTO_INCREMENT,
    `product_id`  int      NOT NULL,
    `category_id` int      NOT NULL,
    `created_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `product_category`
VALUES (1, 1, 2, '2023-12-06 15:27:30', '2023-12-09 22:41:29'),
       (2, 2, 2, '2023-12-06 15:27:30', '2023-12-09 22:41:29'),
       (3, 3, 1, '2023-12-06 15:27:30', '2023-12-06 15:27:30'),
       (4, 4, 1, '2023-12-06 15:27:30', '2023-12-06 15:27:30'),
       (5, 5, 1, '2023-12-06 15:27:30', '2023-12-06 15:27:30'),
       (6, 6, 1, '2023-12-06 15:27:30', '2023-12-09 22:41:47'),
       (10, 10, 2, '2023-12-06 15:27:30', '2023-12-06 15:27:30');