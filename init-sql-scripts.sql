CREATE DATABASE `go-pizza-order-db`;

USE `go-pizza-order-db`;

CREATE TABLE `init_test_table` (
`id` int NOT NULL AUTO_INCREMENT,
`version` varchar(150) DEFAULT NULL,
`created_at` datetime DEFAULT CURRENT_TIMESTAMP,
`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `init_test_table` VALUES(
    '1',
    '1.0.0',
    NOW(),
    NOW()
);