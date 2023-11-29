CREATE DATABASE `storage` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  `code_value` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_published` tinyint DEFAULT NULL,
  `expiration` date DEFAULT NULL,
  `price` decimal(10,0) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `seller` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `cod_seller` varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_authorization` tinyint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `sale` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(45) COLLATE utf8mb4_general_ci NOT NULL,
  `id_seller` int NOT NULL,
  `id_product` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id_seller_idx` (`id_seller`),
  KEY `id_product_idx` (`id_product`),
  CONSTRAINT `id_product` FOREIGN KEY (`id_product`) REFERENCES `product` (`id`),
  CONSTRAINT `id_seller` FOREIGN KEY (`id_seller`) REFERENCES `seller` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
