DROP TABLE IF EXISTS `counter`;
CREATE TABLE `counter` (
  `counter_id` int NOT NULL,
  `isActive` tinyint(1) DEFAULT NULL,
  `branch_id` varchar(500) DEFAULT NULL,
  `counter_name` varchar(500) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `cpuId` varchar(255) NOT NULL,
  `diskId` varchar(255) NOT NULL,
  PRIMARY KEY (`counter_id`)
);