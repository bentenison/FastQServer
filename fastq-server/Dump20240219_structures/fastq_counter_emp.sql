

DROP TABLE IF EXISTS `counter_emp`;
CREATE TABLE `counter_emp` (
  `id` int NOT NULL,
  `counter_id` int DEFAULT NULL,
  `uname` varchar(500) DEFAULT NULL,
  `psw` varchar(1000) DEFAULT NULL,
  `emp_name` varchar(500) DEFAULT NULL,
  `logged_in` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `isdeleted` tinyint(1) DEFAULT NULL,
  `branch_id` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`)
);