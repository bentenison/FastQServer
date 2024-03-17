

DROP TABLE IF EXISTS `audio_conf`;
CREATE TABLE `audio_conf` (
  `id` varchar(255) NOT NULL,
  `hide_date` tinyint(1) DEFAULT NULL,
  `message` varchar(1024) DEFAULT NULL,
  `tts_lang` varchar(255) DEFAULT NULL,
  `repeat_call` int DEFAULT NULL,
  `branch_code` varchar(255) DEFAULT NULL,
  `branch_name` varchar(255) DEFAULT NULL,
  `company_code` varchar(255) DEFAULT NULL,
  `company_id` varchar(255) DEFAULT NULL,
  `company_name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `bell_name` varchar(255) DEFAULT NULL,
  `starting_bell` tinyint(1) DEFAULT NULL,
  `ending_bell` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
);