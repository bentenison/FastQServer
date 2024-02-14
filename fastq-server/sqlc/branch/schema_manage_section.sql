CREATE TABLE `manage_section` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `bench_wait` int(255) DEFAULT NULL,
  `bench_process` int(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `company_code` varchar(255) NOT NULL,
  `branch_code` varchar(255) NOT NULL,
  `company_name` varchar(255) DEFAULT NULL,
  `brach_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`,`company_code`,`branch_code`)
) 