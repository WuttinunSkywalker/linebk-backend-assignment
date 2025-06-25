CREATE TABLE `users` (
    `user_id` varchar(50) NOT NULL,
    `name` varchar(100) NOT NULL,
    `image` varchar(255) DEFAULT NULL,
    `password_hash` varchar(255) NOT NULL,
    `pin_hash` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `user_greetings` (
    `user_id` varchar(50) NOT NULL,
    `greeting` text NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `banners` (
    `banner_id` varchar(50) NOT NULL,
    `user_id` varchar(50) DEFAULT NULL,
    `title` varchar(255) NOT NULL,
    `description` text,
    `image` varchar(255) DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`banner_id`),
    KEY `idx_user_id` (`user_id`),
    CONSTRAINT `fk_banners_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `debit_cards` (
    `card_id` varchar(50) NOT NULL,
    `user_id` varchar(50) DEFAULT NULL,
    `name` varchar(100) DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`card_id`),
    KEY `idx_user_id_created_at` (`user_id`, `created_at`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `debit_card_status` (
    `card_id` varchar(50) NOT NULL,
    `status` varchar(20) DEFAULT NULL,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`card_id`),
    FOREIGN KEY (`card_id`) REFERENCES `debit_cards` (`card_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `debit_card_details` (
    `card_id` varchar(50) NOT NULL,
    `issuer` varchar(100) DEFAULT NULL,
    `number` varchar(25) DEFAULT NULL,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`card_id`),
    FOREIGN KEY (`card_id`) REFERENCES `debit_cards` (`card_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `debit_card_design` (
    `card_id` varchar(50) NOT NULL,
    `color` varchar(10) DEFAULT NULL,
    `border_color` varchar(10) DEFAULT NULL,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`card_id`),
    FOREIGN KEY (`card_id`) REFERENCES `debit_cards` (`card_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `accounts` (
    `account_id` varchar(50) NOT NULL,
    `user_id` varchar(50) NOT NULL,
    `name` varchar(100) NOT NULL,
    `type` varchar(50) NOT NULL,
    `currency` varchar(10) NOT NULL,
    `account_number` varchar(20) DEFAULT NULL,
    `issuer` varchar(100) DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`account_id`),
    KEY `idx_user_id_created_at` (`user_id`, `created_at`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `account_balances` (
    `account_id` varchar(50) NOT NULL,
    `amount` decimal(15, 2) NOT NULL DEFAULT 0.00,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`account_id`),
    FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `account_details` (
    `account_id` varchar(50) NOT NULL,
    `color` varchar(10) DEFAULT NULL,
    `is_main_account` tinyint(1) NOT NULL DEFAULT 0,
    `progress` int DEFAULT NULL,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`account_id`),
    FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `account_flags` (
    `flag_id` int NOT NULL AUTO_INCREMENT,
    `account_id` varchar(50) NOT NULL,
    `flag_type` varchar(50) NOT NULL,
    `flag_value` varchar(30) NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`flag_id`),
    KEY `idx_account_id` (`account_id`),
    FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 6000001 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `transactions` (
    `transaction_id` varchar(50) NOT NULL,
    `user_id` varchar(50) NOT NULL,
    `name` varchar(100) DEFAULT NULL,
    `image` varchar(255) DEFAULT NULL,
    `is_bank` tinyint(1) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`transaction_id`),
    KEY `idx_user_id_created_at` (`user_id`, `created_at`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;