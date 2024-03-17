DROP TABLE IF EXISTS announce_text;
CREATE TABLE announce_text (
    id varchar(255) NOT NULL, announce_text varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL, speed int DEFAULT '0', created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_by varchar(255) DEFAULT NULL, branch_code varchar(255) DEFAULT NULL, company_code varchar(255) DEFAULT NULL, isSelected tinyint(1) DEFAULT NULL, PRIMARY KEY (id)
);