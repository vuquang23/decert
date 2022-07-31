CREATE TABLE `certs` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `description` VARCHAR(255) NOT NULL,
    `issued_at` DATETIME NOT NULL,
    `expired_at` DATETIME,
    `collection_id` int NOT NULL,
    `cert_nft_id` int NOT NULL,
    `data` JSON NOT NULL,
    `revoked_at` DATETIME,
    `revoked_reason` VARCHAR(255),
    `receiver` VARCHAR(128) NOT NULL
);

CREATE TABLE `collections` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `symbol` VARCHAR(128) NOT NULL,
    `address` VARCHAR(128) NOT NULL,
    `total_issued` int NOT NULL,
    `total_revoked` int NOT NULL,
    `issuer` VARCHAR(128) NOT NULL,
    `created_at` DATETIME NOT NULL
);

ALTER TABLE
    `certs`
ADD
    FOREIGN KEY (`collection_id`) REFERENCES `collections` (`id`) ON DELETE CASCADE;