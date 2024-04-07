-- MySQL dump 10.13  Distrib 8.3.0, for Linux (x86_64)
--
-- Host: localhost    Database: postit
-- ------------------------------------------------------
-- Server version	8.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `boards`
--

DROP TABLE IF EXISTS `boards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `boards` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `thumbnail` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `creator_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `creator_id` (`creator_id`),
  CONSTRAINT `boards_ibfk_1` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `boards`
--

LOCK TABLES `boards` WRITE;
/*!40000 ALTER TABLE `boards` DISABLE KEYS */;
INSERT INTO `boards` VALUES (3,'memes','Memes! A way of describing cultural information being shared.','https://styles.redditmedia.com/t5_2qjpg/styles/bannerBackgroundImage_b0ojh4ocr2241.png','2024-04-07 19:56:12','77ae5cb2-9a65-4792-9bf2-823a8089e8df'),(4,'ProgrammerHumor','memesAndJokesAboutEverythingProgrammingAndCS\nFor anything funny related to programming and software development. ','https://styles.redditmedia.com/t5_2tex6/styles/bannerBackgroundImage_z6ejudtdz7p41.png','2024-04-07 20:21:24','c8caaf6a-f1bc-4cef-9ca7-90c8f2f8a668'),(5,'Schnitzelverbrechen','Tunke auf die Panier? Nudeln als Beilage? Diese und andere Schnitzelverbrechen sollen hier gesammelt werden.',NULL,'2024-04-07 20:27:18','9969f4e8-590a-4b47-9617-82205dbbe86a');
/*!40000 ALTER TABLE `boards` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `posts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `thumbnail` text,
  `like_count` int NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `author_id` varchar(255) NOT NULL,
  `board_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `author_id` (`author_id`),
  KEY `board_id` (`board_id`),
  CONSTRAINT `posts_ibfk_1` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`),
  CONSTRAINT `posts_ibfk_2` FOREIGN KEY (`board_id`) REFERENCES `boards` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
INSERT INTO `posts` VALUES (4,'A short story','.','https://preview.redd.it/70zdjiilli351.jpg?auto=webp&s=4050c6cb744a212502a81fda36780386c76553b7',16,'2024-04-07 19:57:24','77ae5cb2-9a65-4792-9bf2-823a8089e8df',3),(5,'This house ain’t big enough for the six of us','-','https://preview.redd.it/x4s2e60m7vm41.jpg?auto=webp&s=1dc499bf53f94585d5a1b50aa6982eae99ce3d16',11,'2024-04-07 20:01:23','0ea82a15-301f-4509-a7c3-e571770967c8',3),(6,'They what???','-','https://external-preview.redd.it/wHP3JyFg8SVTsvZJCi36tMWUCYj_2f30LriRuVvD7wc.png?auto=webp&s=9360475a62bcb46ced4e1b2fb0c62de91613e3b9',6,'2024-04-07 20:19:23','0ea82a15-301f-4509-a7c3-e571770967c8',3),(7,'Why aren’t you playing by the rules of the game!','-','https://preview.redd.it/bq54d2yu1ss91.jpg?auto=webp&s=9897b3faccede787a04e7887a6b70aaa2773495f',5,'2024-04-07 20:22:16','c8caaf6a-f1bc-4cef-9ca7-90c8f2f8a668',4),(8,'it\'s the most important skill','-','https://preview.redd.it/its-the-most-important-skill-v0-g37eb3qivuv81.jpg?auto=webp&s=37a3639ba946144fca6199cc4b401bff0a86a884',10,'2024-04-07 20:23:11','c8caaf6a-f1bc-4cef-9ca7-90c8f2f8a668',4),(9,'Steal what is stolen','-','https://preview.redd.it/steal-what-is-stolen-v0-hwurhp7crzf81.png?auto=webp&s=e73425aa7637794627ee62c94ef2c4b608cc9d76',13,'2024-04-07 20:24:30','5ef71b3d-32cc-49d2-b310-35187db0e6fe',4),(10,'Schnitzelverbrechen gegen die Menschlichkeit','-','https://preview.redd.it/schnitzelverbrechen-gegen-die-menschlichkeit-v0-xwg5weaessza1.png?auto=webp&s=1cc9218f2e29161c4d7f312741ec79c7119c51ad',7,'2024-04-07 20:27:43','9969f4e8-590a-4b47-9617-82205dbbe86a',5),(11,'Wittichenau, Fasching für 4,50. Größer als mein Kopf, völlig abstruser Preis ','-','https://preview.redd.it/wittichenau-fasching-f%C3%BCr-4-50-gr%C3%B6%C3%9Fer-als-mein-kopf-v%C3%B6llig-v0-yrgkmjr7e7ic1.jpg?width=3000&format=pjpg&auto=webp&s=243aa8d81ff868d8ca203a4ea0c687d4d5dccf6c',5,'2024-04-07 20:28:13','9969f4e8-590a-4b47-9617-82205dbbe86a',5),(12,'Krapfen mit Schnitzel.','.','https://preview.redd.it/krapfen-mit-schnitzel-v0-sulq280dqrcc1.jpeg?auto=webp&s=dcc5e735db5432cb96e128cf1e896bae7bdcae60',9,'2024-04-07 20:29:30','62deb71f-4cb3-41e1-995f-bbaaf0602d02',5),(14,'As an Australian I would like to apologise ','-','https://preview.redd.it/as-an-australian-i-would-like-to-apologise-v0-tppam1heamwa1.jpg?width=1080&crop=smart&auto=webp&s=f8879b5cd353abd9d4f13d223c73738778e2f036',25,'2024-04-07 20:34:49','abddcddf-4580-4aab-8b55-27dc788ec4e7',5);
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `avatar` text,
  `bio` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('0ea82a15-301f-4509-a7c3-e571770967c8','theonlymexicanman','$2a$10$M3Oe6l79qndFdsfvdVne5eyx7hDwrQp0B.td7sS/7amaUHl1X0aU6','2024-04-07 20:00:07','https://i.kym-cdn.com/entries/icons/facebook/000/027/475/Screen_Shot_2018-10-25_at_11.02.15_AM.jpg',NULL),('5ef71b3d-32cc-49d2-b310-35187db0e6fe','LongLiveGOSR','$2a$10$DlQYMd9l55rpkCMF2ziM3uI87aOwVkgLyyKI.qkmSSLQEDqjXFO5S','2024-04-07 20:23:52','https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Tux.svg/800px-Tux.svg.png',NULL),('62deb71f-4cb3-41e1-995f-bbaaf0602d02','johnsmart42','$2a$10$tbuUYBY8HPK4MaNHY3VZ6ev0mweGPNjqKCyeg13I7.azVBmXNtdiG','2024-04-07 20:28:51','https://upload.wikimedia.org/wikipedia/commons/thumb/0/03/Wien_3_Wappen.svg/1200px-Wien_3_Wappen.svg.png',NULL),('77ae5cb2-9a65-4792-9bf2-823a8089e8df','rextraneous','$2a$10$Libr/g/YVIHNFzoWfHlRbO6CdLqoYrTcmG3nkm.EJdb8vtIv7zC.S','2024-04-07 19:55:42','https://styles.redditmedia.com/t5_2qjpg/styles/communityIcon_pyzoe71124sc1.png',NULL),('9969f4e8-590a-4b47-9617-82205dbbe86a','SpeedyLeone','$2a$10$Fbl/HbMcecFmnJy8uGJGAOsUb5oTkCNuzBKMNC3jEn.6CizG637E2','2024-04-07 20:25:56','https://cdn.countryflags.com/thumbs/austria/flag-square-250.png',NULL),('abddcddf-4580-4aab-8b55-27dc788ec4e7','fddfgs','$2a$10$W3UdWd1Ulc4hBtPsKZcSl.81KIubAILf4dYro3zbIXAkMNrl4NHSW','2024-04-07 20:34:29','https://images.eatsmarter.de/sites/default/files/styles/max_size/public/pizza-hawaii-118301.jpg',NULL),('c8caaf6a-f1bc-4cef-9ca7-90c8f2f8a668','Mika','$2a$10$GLQDpc9JmQltcMlTNOoLeO0IJx/.zr2D8kUUHF1ajNo2KSJM9BmJm','2024-04-07 18:59:55','https://cdn.discordapp.com/avatars/452415473687068672/1c6bad1e46a6612c40d20d1ac2f61c7e.webp?size=1024',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-04-07 20:58:19
