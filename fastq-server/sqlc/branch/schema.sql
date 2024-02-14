CREATE TABLE `branch_admin` (
  `id` varchar(255) NOT NULL,
  `branch_name` varchar(255) DEFAULT NULL,
  `branch_code` varchar(255) NOT NULL,
  `company_code` varchar(255) NOT NULL,
  `company_name` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `firstname` varchar(255) DEFAULT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`,`branch_code`,`company_code`)
)