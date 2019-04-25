-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Apr 25, 2019 at 04:29 PM
-- Server version: 10.1.35-MariaDB
-- PHP Version: 7.1.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang_project`
--

-- --------------------------------------------------------

--
-- Table structure for table `articles`
--

CREATE TABLE `articles` (
  `id` bigint(20) NOT NULL,
  `article_id` bigint(20) NOT NULL,
  `title` varchar(100) NOT NULL,
  `short_description` varchar(50) NOT NULL,
  `description` varchar(500) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_by` bigint(20) NOT NULL,
  `updated_by` bigint(20) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL COMMENT '1 => new,2 => unpublish,3 => publish'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `articles`
--

INSERT INTO `articles` (`id`, `article_id`, `title`, `short_description`, `description`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`, `status`) VALUES
(3, 1, 'a', 'aaaa', 'a', '2019-04-25 05:16:38', '2019-04-25 05:16:38', '2019-04-25 05:34:01', 1, 1, 0),
(4, 2, 'c', 'c', 'aaa', '2019-04-25 05:17:01', '2019-04-25 05:17:01', '2019-04-25 05:34:06', 1, 1, 0),
(5, 3, 'tes update', 'aaaaa', 'aaaaaaaaaaaa', '2019-04-25 05:19:05', '2019-04-25 14:23:41', NULL, 1, 1, 2),
(6, 4, 'ax', 'axxxx', 'aaaaaaaa', '2019-04-25 05:21:03', '2019-04-25 05:46:56', NULL, 1, 1, 3),
(7, 5, 'ax', 'axxxx', 'aaaaaaaa', '2019-04-25 05:33:56', '2019-04-25 05:33:56', '2019-04-25 05:33:58', 1, 1, 1);

-- --------------------------------------------------------

--
-- Table structure for table `messages`
--

CREATE TABLE `messages` (
  `id` bigint(20) NOT NULL,
  `messages_id` bigint(20) NOT NULL,
  `from_name` varchar(100) NOT NULL,
  `from_email` varchar(255) NOT NULL,
  `message` varchar(500) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_by` bigint(20) NOT NULL,
  `updated_by` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `messages`
--

INSERT INTO `messages` (`id`, `messages_id`, `from_name`, `from_email`, `message`, `created_at`, `updated_at`, `deleted_at`, `created_by`, `updated_by`) VALUES
(2, 1, 'agus', 'tiar.agisti@gmail.com', 'aaa', '2019-04-25 03:32:46', '2019-04-25 03:32:46', NULL, 1, 1),
(3, 2, 'agus2', 'tiar.agisti@gmail.com', 'aaa', '2019-04-25 03:33:00', '2019-04-25 03:33:00', NULL, 1, 1),
(4, 3, 'agus3', 'tiar.agisti@gmail.com', 'aaa', '2019-04-25 03:53:26', '2019-04-25 03:53:26', NULL, 1, 1),
(5, 4, 'agus3', 'tiar.agisti@gmail.com', 'aaa', '2019-04-25 03:53:26', '2019-04-25 03:53:26', NULL, 1, 1),
(6, 5, 'agus3', 'tiar.agisti@gmail.com', 'aaa', '2019-04-25 03:53:26', '2019-04-25 03:53:26', NULL, 1, 1),
(7, 6, 'agus3', 'tiar.agisti@gmail.com', 'aaa', '2019-04-25 03:57:45', '2019-04-25 03:57:45', NULL, 1, 1),
(8, 7, 'agus3', 'tiar.agisti@gmail.com', 'aaa', '2019-04-25 03:57:45', '2019-04-25 03:57:45', NULL, 1, 1),
(9, 8, 'agus3', 'tiar.agisti@gmail.com', 'aaa', '2019-04-25 03:57:45', '2019-04-25 03:57:45', NULL, 1, 1),
(10, 9, 'agus3', 'tiar.agisti@gmail.com', 'aaa', '2019-04-25 04:02:17', '2019-04-25 04:02:17', NULL, 1, 1),
(11, 10, 'test', 'Findcode19@gmail.com', 'aaaa', '2019-04-25 04:03:53', '2019-04-25 04:03:53', NULL, 1, 1);

-- --------------------------------------------------------

--
-- Table structure for table `user_infos`
--

CREATE TABLE `user_infos` (
  `user_id` bigint(20) NOT NULL,
  `username` varchar(100) NOT NULL,
  `user_password` varchar(500) NOT NULL,
  `user_roles` smallint(6) NOT NULL COMMENT '1 => admin,2 => user',
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_by` bigint(20) NOT NULL,
  `updated_by` bigint(20) DEFAULT NULL,
  `deleted_by` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Stand-in structure for view `vw_articles`
-- (See below for the actual view)
--
CREATE TABLE `vw_articles` (
`article_id` bigint(20)
,`title` varchar(100)
,`short_description` varchar(50)
,`description` varchar(500)
,`status` smallint(6)
,`status_desc` varchar(9)
);

-- --------------------------------------------------------

--
-- Stand-in structure for view `vw_last_ids`
-- (See below for the actual view)
--
CREATE TABLE `vw_last_ids` (
`table_id` bigint(20)
,`table_description` varchar(10)
,`last_id` bigint(20)
);

-- --------------------------------------------------------

--
-- Structure for view `vw_articles`
--
DROP TABLE IF EXISTS `vw_articles`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `vw_articles`  AS  select `articles`.`article_id` AS `article_id`,`articles`.`title` AS `title`,`articles`.`short_description` AS `short_description`,`articles`.`description` AS `description`,`articles`.`status` AS `status`,(case when (`articles`.`status` = 1) then 'New' when (`articles`.`status` = 2) then 'Unpublish' when (`articles`.`status` = 3) then 'Publish' else 'Undefined' end) AS `status_desc` from `articles` where isnull(`articles`.`deleted_at`) ;

-- --------------------------------------------------------

--
-- Structure for view `vw_last_ids`
--
DROP TABLE IF EXISTS `vw_last_ids`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `vw_last_ids`  AS  select 1 AS `table_id`,'messages' AS `table_description`,coalesce(max(`m`.`messages_id`),0) AS `last_id` from `messages` `m` union all select 2 AS `table_id`,'user_infos' AS `table_description`,coalesce(max(`u`.`user_id`),0) AS `last_id` from `user_infos` `u` union all select 3 AS `table_id`,'articles' AS `table_description`,coalesce(max(`a`.`article_id`),0) AS `last_id` from `articles` `a` ;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `articles`
--
ALTER TABLE `articles`
  ADD PRIMARY KEY (`id`,`article_id`);

--
-- Indexes for table `messages`
--
ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`,`messages_id`);

--
-- Indexes for table `user_infos`
--
ALTER TABLE `user_infos`
  ADD PRIMARY KEY (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `articles`
--
ALTER TABLE `articles`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `messages`
--
ALTER TABLE `messages`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
