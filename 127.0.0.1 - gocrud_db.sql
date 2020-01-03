-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               5.6.45 - MySQL Community Server (GPL)
-- Server OS:                    Linux
-- HeidiSQL Version:             9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for gocrud_db
DROP DATABASE IF EXISTS `gocrud_db`;
CREATE DATABASE IF NOT EXISTS `gocrud_db` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `gocrud_db`;

-- Dumping structure for table gocrud_db.guestbook
DROP TABLE IF EXISTS `guestbook`;
CREATE TABLE IF NOT EXISTS `guestbook` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Name` varchar(50) DEFAULT NULL,
  `Message` varchar(140) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=latin1;

-- Dumping data for table gocrud_db.guestbook: ~1 rows (approximately)
DELETE FROM `guestbook`;
/*!40000 ALTER TABLE `guestbook` DISABLE KEYS */;
INSERT INTO `guestbook` (`id`, `Name`, `Message`, `created_at`) VALUES
	(62, 'Kurniawan', 'Lorem Ipsum dolor sit amet', '2020-01-03 13:14:44'),
	(63, 'Kurniawan', 'Lorem Ipsum dolor sit amet', '2020-01-03 13:14:45');
/*!40000 ALTER TABLE `guestbook` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
