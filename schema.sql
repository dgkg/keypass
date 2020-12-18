CREATE TABLE `cards` (
  `id` varchar(36) DEFAULT NULL,
  `user_id` varchar(36) DEFAULT NULL,
  `url` text,
  `user_account_login` varchar(30) DEFAULT NULL,
  `user_account_password` varchar(30) DEFAULT NULL,
  `creation_date` datetime DEFAULT NULL,
  `activated` tinyint(1) DEFAULT NULL,
  `pic` blob
);

CREATE TABLE `users` (
  `id` varchar(36) DEFAULT NULL,
  `first_name` varchar(100) DEFAULT NULL,
  `last_name` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `creation_date` datetime DEFAULT NULL
);