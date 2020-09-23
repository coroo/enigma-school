# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.31)
# Database: enigma-school
# Generation Time: 2020-09-23 05:22:10 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table mahasiswa
# ------------------------------------------------------------

DROP TABLE IF EXISTS `mahasiswa`;

CREATE TABLE `mahasiswa` (
  `nim` int(11) NOT NULL,
  `nama_mahasiswa` varchar(255) DEFAULT NULL,
  `jurusan_mahasiswa` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`nim`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `mahasiswa` WRITE;
/*!40000 ALTER TABLE `mahasiswa` DISABLE KEYS */;

INSERT INTO `mahasiswa` (`nim`, `nama_mahasiswa`, `jurusan_mahasiswa`, `created_at`, `updated_at`)
VALUES
	(1234,'Denny','Matematika','2020-09-23 09:45:12','2020-09-23 09:45:12');

/*!40000 ALTER TABLE `mahasiswa` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table m_pelajaran
# ------------------------------------------------------------

DROP TABLE IF EXISTS `m_pelajaran`;

CREATE TABLE `m_pelajaran` (
  `kode_pelajaran` char(36) NOT NULL DEFAULT '',
  `nama_pelajaran` varchar(255) DEFAULT NULL,
  `jumlah_sks` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`kode_pelajaran`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `m_pelajaran` WRITE;
/*!40000 ALTER TABLE `m_pelajaran` DISABLE KEYS */;

INSERT INTO `m_pelajaran` (`kode_pelajaran`, `nama_pelajaran`, `jumlah_sks`, `created_at`, `updated_at`)
VALUES
	('SCFI602116','Fisika Energi',3,'2020-09-23 10:18:11','2020-09-23 10:18:11'),
	('SCFI603116','Mekanika Kuantum 2',3,'2020-09-23 10:17:50','2020-09-23 10:17:50'),
	('SCFI603416','Fisika Komputasi Lanjut',5,'2020-09-23 10:16:22','2020-09-23 10:16:22');

/*!40000 ALTER TABLE `m_pelajaran` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table penempatan
# ------------------------------------------------------------

DROP TABLE IF EXISTS `penempatan`;

CREATE TABLE `penempatan` (
  `id` char(36) NOT NULL DEFAULT '',
  `nim` int(11) DEFAULT NULL,
  `kode_pelajaran` char(36) DEFAULT '',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `tahun_ajaran` varchar(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `penempatan` WRITE;
/*!40000 ALTER TABLE `penempatan` DISABLE KEYS */;

INSERT INTO `penempatan` (`id`, `nim`, `kode_pelajaran`, `created_at`, `updated_at`, `tahun_ajaran`)
VALUES
	('b1ab3fe6-800e-4782-aab7-44e87e2aac01',1234,'SCFI602116','2020-09-23 11:23:07','2020-09-23 11:23:07','2019'),
	('c787eaa0-cde2-4d1d-bc6c-e4fac58e9496',1234,'SCFI603416','2020-09-23 11:29:11','2020-09-23 11:29:11','2019');

/*!40000 ALTER TABLE `penempatan` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
