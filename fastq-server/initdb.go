package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const mysqlDatabaseName = "fastq"
const createTableOne = `CREATE TABLE announce_text (
	id varchar(255) NOT NULL,
	announce_text varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
	speed int DEFAULT '0',
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) DEFAULT NULL,
	branch_code varchar(255) DEFAULT NULL,
	company_code varchar(255) DEFAULT NULL,
	isSelected tinyint(1) DEFAULT NULL,
	PRIMARY KEY (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableTwo = `CREATE TABLE audio_conf (
	id varchar(255) NOT NULL,
	hide_date tinyint(1) DEFAULT NULL,
	message varchar(1024) DEFAULT NULL,
	tts_lang varchar(255) DEFAULT NULL,
	repeat_call int DEFAULT NULL,
	branch_code varchar(255) DEFAULT NULL,
	branch_name varchar(255) DEFAULT NULL,
	company_code varchar(255) DEFAULT NULL,
	company_id varchar(255) DEFAULT NULL,
	company_name varchar(255) DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) DEFAULT NULL,
	bell_name varchar(255) DEFAULT NULL,
	starting_bell tinyint(1) DEFAULT NULL,
	ending_bell tinyint(1) DEFAULT NULL,
	PRIMARY KEY (id)
  );`

const createTableThree = `CREATE TABLE branch_admin (
  id varchar(255) NOT NULL,
  branch_name varchar(255) DEFAULT NULL,
  branch_code varchar(255) NOT NULL,
  company_code varchar(255) NOT NULL,
  company_name varchar(255) DEFAULT NULL,
  username varchar(255) DEFAULT NULL,
  firstname varchar(255) DEFAULT NULL,
  lastname varchar(255) DEFAULT NULL,
  email varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  created_by varchar(255) DEFAULT NULL,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_by varchar(255) DEFAULT NULL,
  password varchar(500) DEFAULT NULL,
  isdeleted tinyint(1) DEFAULT NULL,
  PRIMARY KEY (id,branch_code,company_code)
);`

const createTableFour = `CREATE TABLE company_info (
  id varchar(255) NOT NULL,
  email varchar(255) DEFAULT NULL,
  firstname varchar(255) DEFAULT NULL,
  lastname varchar(255) DEFAULT NULL,
  company_code varchar(255) NOT NULL,
  company varchar(255) DEFAULT NULL,
  country varchar(255) DEFAULT NULL,
  timezone varchar(255) DEFAULT NULL,
  password varchar(255) DEFAULT NULL,
  updated_by varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id,company_code)
);`

const createTableFive = `CREATE TABLE counter_emp (
  id int NOT NULL,
  counter_id int DEFAULT NULL,
  uname varchar(500) DEFAULT NULL,
  psw varchar(1000) DEFAULT NULL,
  emp_name varchar(500) DEFAULT NULL,
  logged_in timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  isdeleted tinyint(1) DEFAULT NULL,
  branch_id varchar(500) DEFAULT NULL,
  PRIMARY KEY (id)
);`
const createTableSix = `CREATE TABLE counter_login_log (
	id int NOT NULL AUTO_INCREMENT,
	counter_id int NOT NULL,
	logged_in_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	logged_out_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (counter_id),
	KEY id (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableSeven = `CREATE TABLE counter (
	counter_id int NOT NULL,
	isActive tinyint(1) DEFAULT NULL,
	branch_id varchar(500) DEFAULT NULL,
	counter_name varchar(500) DEFAULT NULL,
	is_deleted tinyint(1) DEFAULT NULL,
	cpuId varchar(255) NOT NULL,
	diskId varchar(255) NOT NULL,
	PRIMARY KEY (counter_id)
  );`
const createTableEignt = `CREATE TABLE counterservices (
	CounterServiceID int NOT NULL AUTO_INCREMENT,
	CounterID varchar(255) DEFAULT NULL,
	ServiceID varchar(500) DEFAULT NULL,
	PRIMARY KEY (CounterServiceID)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableNine = `CREATE TABLE country (
	country_code char(2) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin DEFAULT NULL,
	country_name varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin DEFAULT NULL,
	KEY idx_country_code (country_code)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableTen = `CREATE TABLE customer (
	id int NOT NULL,
	name varchar(255) DEFAULT NULL,
	name_ar varchar(255) DEFAULT NULL,
	contact_number varchar(255) DEFAULT NULL,
	ticket_id int DEFAULT NULL,
	PRIMARY KEY (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableEleven = `CREATE TABLE display_conf (
	id varchar(255) DEFAULT NULL,
	branch_code varchar(255) DEFAULT NULL,
	branch_name varchar(255) DEFAULT NULL,
	company_code varchar(255) DEFAULT NULL,
	company_name varchar(255) DEFAULT NULL,
	template varchar(255) DEFAULT NULL,
	retain_q tinyint(1) DEFAULT NULL,
	confirmation_delay tinyint(1) DEFAULT NULL,
	show_form tinyint(1) DEFAULT NULL,
	show_scroll tinyint(1) DEFAULT NULL,
	show_dt tinyint(1) DEFAULT NULL,
	volume int DEFAULT NULL,
	form_conf varchar(255) DEFAULT NULL,
	online_form varchar(255) DEFAULT NULL,
	show_url tinyint(1) DEFAULT NULL,
	url varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableTwelve = `CREATE TABLE form_conf (
	name tinyint(1) DEFAULT NULL,
	phone tinyint(1) DEFAULT NULL,
	email tinyint(1) DEFAULT NULL,
	send_mail tinyint(1) DEFAULT NULL,
	date tinyint(1) DEFAULT NULL,
	data_1 varchar(255) DEFAULT NULL,
	data_2 varchar(255) DEFAULT NULL,
	data_3 varchar(255) DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	created_by varchar(255) DEFAULT NULL,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) DEFAULT NULL,
	company_code varchar(255) DEFAULT NULL,
	company_name varchar(255) DEFAULT NULL,
	branch_code varchar(255) DEFAULT NULL,
	branch_name varchar(255) DEFAULT NULL,
	id varchar(255) NOT NULL,
	PRIMARY KEY (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableThirteen = `CREATE TABLE manage_branch (
	id varchar(255) NOT NULL,
	code varchar(255) NOT NULL,
	name varchar(255) DEFAULT NULL,
	license varchar(255) DEFAULT NULL,
	services varchar(1024) DEFAULT NULL,
	address varchar(255) DEFAULT NULL,
	phone varchar(255) DEFAULT NULL,
	license_key varchar(1024) DEFAULT NULL,
	check_in_url varchar(255) DEFAULT NULL,
	ticket_page_url varchar(255) DEFAULT NULL,
	display_url varchar(255) DEFAULT NULL,
	company_code varchar(255) NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	created_by varchar(255) DEFAULT NULL,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) DEFAULT NULL,
	printer_details varchar(1024) DEFAULT NULL,
	email varchar(255) NOT NULL,
	password varchar(255) DEFAULT NULL,
	PRIMARY KEY (id,code,company_code)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableFourteen = `CREATE TABLE manage_counters (
	id varchar(255) NOT NULL,
	section varchar(255) DEFAULT NULL,
	sub_section varchar(255) DEFAULT NULL,
	user_id varchar(255) NOT NULL,
	counter_number varchar(255) DEFAULT NULL,
	counter_name varchar(255) DEFAULT NULL,
	override_fifo tinyint(1) DEFAULT NULL,
	transfer_q tinyint(1) DEFAULT NULL,
	asssign_user tinyint(1) DEFAULT NULL,
	assign_service tinyint(1) DEFAULT NULL,
	transfer_priority tinyint(1) DEFAULT NULL,
	cancel_q tinyint(1) DEFAULT NULL,
	activated tinyint(1) DEFAULT NULL,
	diskId varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
	cpuId varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
	branch_code varchar(255) NOT NULL,
	branch_name varchar(255) DEFAULT NULL,
	company_name varchar(255) DEFAULT NULL,
	company_code varchar(255) NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by varchar(255) DEFAULT NULL,
	updated_by varchar(255) DEFAULT NULL
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableFifteen = `CREATE TABLE manage_service (
	id varchar(255) NOT NULL,
	name varchar(255) DEFAULT NULL,
	code varchar(255) NOT NULL,
	prefix varchar(255) DEFAULT NULL,
	number_starts int DEFAULT NULL,
	number_ends int DEFAULT NULL,
	hide tinyint(1) DEFAULT NULL,
	show_display tinyint(1) DEFAULT NULL,
	description varchar(1024) DEFAULT NULL,
	start_time datetime DEFAULT NULL,
	end_time datetime DEFAULT NULL,
	default_time int DEFAULT NULL,
	workflow varchar(255) DEFAULT NULL,
	branch_code varchar(255) NOT NULL,
	branch_name varchar(255) DEFAULT NULL,
	company_code varchar(255) NOT NULL,
	company_name varchar(255) DEFAULT NULL,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) DEFAULT NULL
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableSixteen = `CREATE TABLE manage_user (
	id varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	firstname varchar(255) DEFAULT NULL,
	lastname varchar(255) DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	user_type varchar(255) DEFAULT NULL,
	suspended tinyint(1) DEFAULT NULL,
	branch_name varchar(255) DEFAULT NULL,
	branch_code varchar(255) NOT NULL,
	company_name varchar(255) DEFAULT NULL,
	company_code varchar(255) NOT NULL,
	title varchar(255) DEFAULT NULL,
	password varchar(255) DEFAULT NULL,
	image_url varchar(255) DEFAULT NULL,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by varchar(255) DEFAULT NULL,
	updated_by varchar(255) DEFAULT NULL
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableSeventeen = `CREATE TABLE server_details (
	server_ip varchar(100) DEFAULT NULL,
	server_cpu varchar(100) DEFAULT NULL,
	server_disk_id varchar(100) DEFAULT NULL,
	id varchar(100) DEFAULT NULL,
	type varchar(100) DEFAULT NULL
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableEighteen = `CREATE TABLE systems (
  server_ip varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  server_cpu varchar(255) DEFAULT NULL,
  server_disk_id varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  client_ip varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  client_cpu varchar(100) DEFAULT NULL,
  client_disk_id varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  is_active tinyint(1) DEFAULT NULL,
  is_deleted tinyint(1) DEFAULT NULL,
  counter_name varchar(100) DEFAULT NULL,
  counter_id varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableNineteen = `CREATE TABLE ticket_conf (
	image_header tinyint(1) DEFAULT NULL,
	image_url varchar(255) DEFAULT NULL,
	no_show_exp int DEFAULT NULL,
	text_header tinyint(1) DEFAULT NULL,
	header_text varchar(255) DEFAULT NULL,
	hide_time tinyint(1) DEFAULT NULL,
	print_duplicate tinyint(1) DEFAULT NULL,
	date_format varchar(255) DEFAULT NULL,
	print_position tinyint(1) DEFAULT NULL,
	number_digits int DEFAULT NULL,
	estimated_time tinyint(1) DEFAULT NULL,
	pop_up_delay tinyint(1) DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) DEFAULT NULL,
	id varchar(255) NOT NULL,
	branch_code varchar(255) DEFAULT NULL,
	branch_name varchar(255) DEFAULT NULL,
	company_code varchar(255) DEFAULT NULL,
	company_name varchar(255) DEFAULT NULL,
	PRIMARY KEY (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableTwenty = `CREATE TABLE ticket_status (
	id int NOT NULL AUTO_INCREMENT,
	text varchar(255) DEFAULT NULL,
	PRIMARY KEY (id)
  ) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;`

const createTableTwentyOne = `CREATE TABLE ticket (
	id varchar(255) NOT NULL,
	service varchar(255) DEFAULT NULL,
	ticket_status varchar(255) DEFAULT NULL,
	counter_id varchar(255) DEFAULT NULL,
	transfered_to varchar(255) DEFAULT NULL,
	transfered_by varchar(255) DEFAULT NULL,
	customer_id varchar(255) DEFAULT NULL,
	ticket_number varchar(255) DEFAULT NULL,
	created_at timestamp NULL DEFAULT NULL,
	updated_at timestamp NULL DEFAULT NULL,
	updated_by varchar(255) DEFAULT NULL,
	started_serving_at timestamp NULL DEFAULT NULL,
	end_serving_at timestamp NULL DEFAULT NULL,
	ticket_name varchar(255) DEFAULT NULL,
	company_name varchar(255) DEFAULT NULL,
	company_code varchar(255) NOT NULL,
	branch_code varchar(255) NOT NULL,
	branch_name varchar(255) DEFAULT NULL,
	serving_time varchar(100) DEFAULT NULL,
	waiting_time varchar(100) DEFAULT NULL
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableTwentyTwo = `CREATE TABLE time_zone (
	zone_name varchar(35) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL,
	country_code char(2) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL,
	abbreviation varchar(6) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL,
	time_start decimal(11,0) NOT NULL,
	gmt_offset int NOT NULL,
	dst char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL,
	KEY idx_zone_name (zone_name),
	KEY idx_country_code (country_code),
	KEY idx_time_start (time_start)
  ) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_bin;`

const createTableTwentyThree = `CREATE TABLE video_schedular (
	id varchar(255) NOT NULL,
	video_id varchar(255) DEFAULT NULL,
	day varchar(255) DEFAULT NULL,
	ordinality int DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	created_by varchar(255) DEFAULT NULL,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) DEFAULT NULL,
	branch_code varchar(255) DEFAULT NULL,
	company_code varchar(255) DEFAULT NULL,
	PRIMARY KEY (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableTwentyFour = `CREATE TABLE videos (
	id varchar(255) NOT NULL,
	file_name varchar(255) DEFAULT NULL,
	type varchar(255) DEFAULT NULL,
	size varchar(255) DEFAULT NULL,
	modified_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by varchar(255) DEFAULT NULL,
	branch_code varchar(255) DEFAULT NULL,
	company_code varchar(255) DEFAULT NULL,
	company_name varchar(255) DEFAULT NULL,
	branch_name varchar(255) DEFAULT NULL,
	isdeleted tinyint(1) DEFAULT NULL,
	day varchar(100) DEFAULT NULL,
	ordinality varchar(100) DEFAULT NULL,
	PRIMARY KEY (id)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

const createTableTwentyFive = `INSERT INTO country VALUES ('AF','Afghanistan'),('AX','Åland Islands'),('AL','Albania'),('DZ','Algeria'),('AS','American Samoa'),('AD','Andorra'),('AO','Angola'),('AI','Anguilla'),('AQ','Antarctica'),('AG','Antigua and Barbuda'),('AR','Argentina'),('AM','Armenia'),('AW','Aruba'),('AU','Australia'),('AT','Austria'),('AZ','Azerbaijan'),('BS','Bahamas'),('BH','Bahrain'),('BD','Bangladesh'),('BB','Barbados'),('BY','Belarus'),('BE','Belgium'),('BZ','Belize'),('BJ','Benin'),('BM','Bermuda'),('BT','Bhutan'),('BO','Bolivia, Plurinational State of'),('BQ','Bonaire, Sint Eustatius and Saba'),('BA','Bosnia and Herzegovina'),('BW','Botswana'),('BV','Bouvet Island'),('BR','Brazil'),('IO','British Indian Ocean Territory'),('BN','Brunei Darussalam'),('BG','Bulgaria'),('BF','Burkina Faso'),('BI','Burundi'),('KH','Cambodia'),('CM','Cameroon'),('CA','Canada'),('CV','Cape Verde'),('KY','Cayman Islands'),('CF','Central African Republic'),('TD','Chad'),('CL','Chile'),('CN','China'),('CX','Christmas Island'),('CC','Cocos (Keeling) Islands'),('CO','Colombia'),('KM','Comoros'),('CG','Congo'),('CD','Congo, the Democratic Republic of the'),('CK','Cook Islands'),('CR','Costa Rica'),('CI','Côte d\'Ivoire'),('HR','Croatia'),('CU','Cuba'),('CW','Curaçao'),('CY','Cyprus'),('CZ','Czech Republic'),('DK','Denmark'),('DJ','Djibouti'),('DM','Dominica'),('DO','Dominican Republic'),('EC','Ecuador'),('EG','Egypt'),('SV','El Salvador'),('GQ','Equatorial Guinea'),('ER','Eritrea'),('EE','Estonia'),('ET','Ethiopia'),('FK','Falkland Islands (Malvinas)'),('FO','Faroe Islands'),('FJ','Fiji'),('FI','Finland'),('FR','France'),('GF','French Guiana'),('PF','French Polynesia'),('TF','French Southern Territories'),('GA','Gabon'),('GM','Gambia'),('GE','Georgia'),('DE','Germany'),('GH','Ghana'),('GI','Gibraltar'),('GR','Greece'),('GL','Greenland'),('GD','Grenada'),('GP','Guadeloupe'),('GU','Guam'),('GT','Guatemala'),('GG','Guernsey'),('GN','Guinea'),('GW','Guinea-Bissau'),('GY','Guyana'),('HT','Haiti'),('HM','Heard Island and McDonald Islands'),('VA','Holy See (Vatican City State)'),('HN','Honduras'),('HK','Hong Kong'),('HU','Hungary'),('IS','Iceland'),('IN','India'),('ID','Indonesia'),('IR','Iran, Islamic Republic of'),('IQ','Iraq'),('IE','Ireland'),('IM','Isle of Man'),('IL','Israel'),('IT','Italy'),('JM','Jamaica'),('JP','Japan'),('JE','Jersey'),('JO','Jordan'),('KZ','Kazakhstan'),('KE','Kenya'),('KI','Kiribati'),('KP','Korea, Democratic People\'s Republic of'),('KR','Korea, Republic of'),('KW','Kuwait'),('KG','Kyrgyzstan'),('LA','Lao People\'s Democratic Republic'),('LV','Latvia'),('LB','Lebanon'),('LS','Lesotho'),('LR','Liberia'),('LY','Libya'),('LI','Liechtenstein'),('LT','Lithuania'),('LU','Luxembourg'),('MO','Macao'),('MK','Macedonia, the Former Yugoslav Republic of'),('MG','Madagascar'),('MW','Malawi'),('MY','Malaysia'),('MV','Maldives'),('ML','Mali'),('MT','Malta'),('MH','Marshall Islands'),('MQ','Martinique'),('MR','Mauritania'),('MU','Mauritius'),('YT','Mayotte'),('MX','Mexico'),('FM','Micronesia, Federated States of'),('MD','Moldova, Republic of'),('MC','Monaco'),('MN','Mongolia'),('ME','Montenegro'),('MS','Montserrat'),('MA','Morocco'),('MZ','Mozambique'),('MM','Myanmar'),('NA','Namibia'),('NR','Nauru'),('NP','Nepal'),('NL','Netherlands'),('NC','New Caledonia'),('NZ','New Zealand'),('NI','Nicaragua'),('NE','Niger'),('NG','Nigeria'),('NU','Niue'),('NF','Norfolk Island'),('MP','Northern Mariana Islands'),('NO','Norway'),('OM','Oman'),('PK','Pakistan'),('PW','Palau'),('PS','Palestine, State of'),('PA','Panama'),('PG','Papua New Guinea'),('PY','Paraguay'),('PE','Peru'),('PH','Philippines'),('PN','Pitcairn'),('PL','Poland'),('PT','Portugal'),('PR','Puerto Rico'),('QA','Qatar'),('RE','Réunion'),('RO','Romania'),('RU','Russian Federation'),('RW','Rwanda'),('BL','Saint Barthélemy'),('SH','Saint Helena, Ascension and Tristan da Cunha'),('KN','Saint Kitts and Nevis'),('LC','Saint Lucia'),('MF','Saint Martin (French part)'),('PM','Saint Pierre and Miquelon'),('VC','Saint Vincent and the Grenadines'),('WS','Samoa'),('SM','San Marino'),('ST','Sao Tome and Principe'),('SA','Saudi Arabia'),('SN','Senegal'),('RS','Serbia'),('SC','Seychelles'),('SL','Sierra Leone'),('SG','Singapore'),('SX','Sint Maarten (Dutch part)'),('SK','Slovakia'),('SI','Slovenia'),('SB','Solomon Islands'),('SO','Somalia'),('ZA','South Africa'),('GS','South Georgia and the South Sandwich Islands'),('SS','South Sudan'),('ES','Spain'),('LK','Sri Lanka'),('SD','Sudan'),('SR','Suriname'),('SJ','Svalbard and Jan Mayen'),('SZ','Swaziland'),('SE','Sweden'),('CH','Switzerland'),('SY','Syrian Arab Republic'),('TW','Taiwan, Province of China'),('TJ','Tajikistan'),('TZ','Tanzania, United Republic of'),('TH','Thailand'),('TL','Timor-Leste'),('TG','Togo'),('TK','Tokelau'),('TO','Tonga'),('TT','Trinidad and Tobago'),('TN','Tunisia'),('TR','Turkey'),('TM','Turkmenistan'),('TC','Turks and Caicos Islands'),('TV','Tuvalu'),('UG','Uganda'),('UA','Ukraine'),('AE','United Arab Emirates'),('GB','United Kingdom'),('US','United States'),('UM','United States Minor Outlying Islands'),('UY','Uruguay'),('UZ','Uzbekistan'),('VU','Vanuatu'),('VE','Venezuela, Bolivarian Republic of'),('VN','Viet Nam'),('VG','Virgin Islands, British'),('VI','Virgin Islands, U.S.'),('WF','Wallis and Futuna'),('EH','Western Sahara'),('YE','Yemen'),('ZM','Zambia'),('ZW','Zimbabwe'),('AF','Afghanistan'),('AX','Åland Islands'),('AL','Albania'),('DZ','Algeria'),('AS','American Samoa'),('AD','Andorra'),('AO','Angola'),('AI','Anguilla'),('AQ','Antarctica'),('AG','Antigua and Barbuda'),('AR','Argentina'),('AM','Armenia'),('AW','Aruba'),('AU','Australia'),('AT','Austria'),('AZ','Azerbaijan'),('BS','Bahamas'),('BH','Bahrain'),('BD','Bangladesh'),('BB','Barbados'),('BY','Belarus'),('BE','Belgium'),('BZ','Belize'),('BJ','Benin'),('BM','Bermuda'),('BT','Bhutan'),('BO','Bolivia, Plurinational State of'),('BQ','Bonaire, Sint Eustatius and Saba'),('BA','Bosnia and Herzegovina'),('BW','Botswana'),('BV','Bouvet Island'),('BR','Brazil'),('IO','British Indian Ocean Territory'),('BN','Brunei Darussalam'),('BG','Bulgaria'),('BF','Burkina Faso'),('BI','Burundi'),('KH','Cambodia'),('CM','Cameroon'),('CA','Canada'),('CV','Cape Verde'),('KY','Cayman Islands'),('CF','Central African Republic'),('TD','Chad'),('CL','Chile'),('CN','China'),('CX','Christmas Island'),('CC','Cocos (Keeling) Islands'),('CO','Colombia'),('KM','Comoros'),('CG','Congo'),('CD','Congo, the Democratic Republic of the'),('CK','Cook Islands'),('CR','Costa Rica'),('CI','Côte d\'Ivoire'),('HR','Croatia'),('CU','Cuba'),('CW','Curaçao'),('CY','Cyprus'),('CZ','Czech Republic'),('DK','Denmark'),('DJ','Djibouti'),('DM','Dominica'),('DO','Dominican Republic'),('EC','Ecuador'),('EG','Egypt'),('SV','El Salvador'),('GQ','Equatorial Guinea'),('ER','Eritrea'),('EE','Estonia'),('ET','Ethiopia'),('FK','Falkland Islands (Malvinas)'),('FO','Faroe Islands'),('FJ','Fiji'),('FI','Finland'),('FR','France'),('GF','French Guiana'),('PF','French Polynesia'),('TF','French Southern Territories'),('GA','Gabon'),('GM','Gambia'),('GE','Georgia'),('DE','Germany'),('GH','Ghana'),('GI','Gibraltar'),('GR','Greece'),('GL','Greenland'),('GD','Grenada'),('GP','Guadeloupe'),('GU','Guam'),('GT','Guatemala'),('GG','Guernsey'),('GN','Guinea'),('GW','Guinea-Bissau'),('GY','Guyana'),('HT','Haiti'),('HM','Heard Island and McDonald Islands'),('VA','Holy See (Vatican City State)'),('HN','Honduras'),('HK','Hong Kong'),('HU','Hungary'),('IS','Iceland'),('IN','India'),('ID','Indonesia'),('IR','Iran, Islamic Republic of'),('IQ','Iraq'),('IE','Ireland'),('IM','Isle of Man'),('IL','Israel'),('IT','Italy'),('JM','Jamaica'),('JP','Japan'),('JE','Jersey'),('JO','Jordan'),('KZ','Kazakhstan'),('KE','Kenya'),('KI','Kiribati'),('KP','Korea, Democratic People\'s Republic of'),('KR','Korea, Republic of'),('KW','Kuwait'),('KG','Kyrgyzstan'),('LA','Lao People\'s Democratic Republic'),('LV','Latvia'),('LB','Lebanon'),('LS','Lesotho'),('LR','Liberia'),('LY','Libya'),('LI','Liechtenstein'),('LT','Lithuania'),('LU','Luxembourg'),('MO','Macao'),('MK','Macedonia, the Former Yugoslav Republic of'),('MG','Madagascar'),('MW','Malawi'),('MY','Malaysia'),('MV','Maldives'),('ML','Mali'),('MT','Malta'),('MH','Marshall Islands'),('MQ','Martinique'),('MR','Mauritania'),('MU','Mauritius'),('YT','Mayotte'),('MX','Mexico'),('FM','Micronesia, Federated States of'),('MD','Moldova, Republic of'),('MC','Monaco'),('MN','Mongolia'),('ME','Montenegro'),('MS','Montserrat'),('MA','Morocco'),('MZ','Mozambique'),('MM','Myanmar'),('NA','Namibia'),('NR','Nauru'),('NP','Nepal'),('NL','Netherlands'),('NC','New Caledonia'),('NZ','New Zealand'),('NI','Nicaragua'),('NE','Niger'),('NG','Nigeria'),('NU','Niue'),('NF','Norfolk Island'),('MP','Northern Mariana Islands'),('NO','Norway'),('OM','Oman'),('PK','Pakistan'),('PW','Palau'),('PS','Palestine, State of'),('PA','Panama'),('PG','Papua New Guinea'),('PY','Paraguay'),('PE','Peru'),('PH','Philippines'),('PN','Pitcairn'),('PL','Poland'),('PT','Portugal'),('PR','Puerto Rico'),('QA','Qatar'),('RE','Réunion'),('RO','Romania'),('RU','Russian Federation'),('RW','Rwanda'),('BL','Saint Barthélemy'),('SH','Saint Helena, Ascension and Tristan da Cunha'),('KN','Saint Kitts and Nevis'),('LC','Saint Lucia'),('MF','Saint Martin (French part)'),('PM','Saint Pierre and Miquelon'),('VC','Saint Vincent and the Grenadines'),('WS','Samoa'),('SM','San Marino'),('ST','Sao Tome and Principe'),('SA','Saudi Arabia'),('SN','Senegal'),('RS','Serbia'),('SC','Seychelles'),('SL','Sierra Leone'),('SG','Singapore'),('SX','Sint Maarten (Dutch part)'),('SK','Slovakia'),('SI','Slovenia'),('SB','Solomon Islands'),('SO','Somalia'),('ZA','South Africa'),('GS','South Georgia and the South Sandwich Islands'),('SS','South Sudan'),('ES','Spain'),('LK','Sri Lanka'),('SD','Sudan'),('SR','Suriname'),('SJ','Svalbard and Jan Mayen'),('SZ','Swaziland'),('SE','Sweden'),('CH','Switzerland'),('SY','Syrian Arab Republic'),('TW','Taiwan, Province of China'),('TJ','Tajikistan'),('TZ','Tanzania, United Republic of'),('TH','Thailand'),('TL','Timor-Leste'),('TG','Togo'),('TK','Tokelau'),('TO','Tonga'),('TT','Trinidad and Tobago'),('TN','Tunisia'),('TR','Turkey'),('TM','Turkmenistan'),('TC','Turks and Caicos Islands'),('TV','Tuvalu'),('UG','Uganda'),('UA','Ukraine'),('AE','United Arab Emirates'),('GB','United Kingdom'),('US','United States'),('UM','United States Minor Outlying Islands'),('UY','Uruguay'),('UZ','Uzbekistan'),('VU','Vanuatu'),('VE','Venezuela, Bolivarian Republic of'),('VN','Viet Nam'),('VG','Virgin Islands, British'),('VI','Virgin Islands, U.S.'),('WF','Wallis and Futuna'),('EH','Western Sahara'),('YE','Yemen'),('ZM','Zambia'),('ZW','Zimbabwe');`

func setupSQLAutomatic(ctx context.Context, db *sql.DB, folderPath string) error {
	res, err := DoesDatabaseExist(ctx, db)
	if err != nil {
		log.Println("error occured while setting mysql", err)
		return err
	}
	log.Println(res)
	if !res {
		err = CreateDB(ctx, db)
		if err != nil {
			log.Println("error occured while creating DB", err)
			return err
		}
		err := ExecuteSQLScripts(context.Background(), db, folderPath)
		if err != nil {
			log.Println("error occured while creating DB", err)
			return err
		}
	}
	return nil
}
func CreateDB(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", mysqlDatabaseName))
	return err
}

// type SchemaRes struct {
// 	Name dbr.NullString db:"SCHEMA_NAME"
// 	// TableName dbr.NullString json:"tableName"
// }

func DoesDatabaseExist(ctx context.Context, db *sql.DB) (bool, error) {
	// Construct the MySQL command to check if the database exists
	var result string
	query := fmt.Sprintf("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '%s';", mysqlDatabaseName)
	// log.Println(" '\\" + mysqlDatabaseName + "\\';")
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterating over the query result
	for rows.Next() {

		err := rows.Scan(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Schema Name:", result)
	}

	// Checking for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	if result == "" {
		return false, nil
	}
	log.Println("DB checked", result)
	return true, nil
}
func ExecuteSQLScripts(ctx context.Context, db *sql.DB, folderPath string) error {
	// files, err := ioutil.ReadDir(folderPath)
	// if err != nil {
	// 	return fmt.Errorf("error reading directory: %w", err)
	// }

	// for _, file := range files {
	// 	if strings.HasSuffix(file.Name(), ".sql") {
	// 		filePath := filepath.Join(folderPath, file.Name())
	// 		sqlContent, err := ioutil.ReadFile(filePath)
	// 		if err != nil {
	// 			return fmt.Errorf("error reading SQL file %s: %w", filePath, err)
	// 		}
	_, err := db.ExecContext(ctx, fmt.Sprintf("USE %s", mysqlDatabaseName))
	if err != nil {
		log.Println("error occured while executing script", err)
	}

	_, err = db.ExecContext(ctx, createTableOne)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableTwo)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableThree)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableFour)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableFive)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableSix)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableSeven)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableEignt)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableNine)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableTen)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableEleven)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableTwelve)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableThirteen)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableFourteen)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableFifteen)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableSixteen)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableSeventeen)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableEighteen)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableNineteen)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableTwenty)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableTwentyOne)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableTwentyTwo)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableTwentyThree)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableTwentyFour)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	_, err = db.ExecContext(ctx, createTableTwentyFive)
	if err != nil {
		log.Println("error occured while executing script", err)
	}
	// return err
	// fmt.Printf("SQL file %s executed successfully\n", filePath)
	// 	}
	// }

	return nil
}
