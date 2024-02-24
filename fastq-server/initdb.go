package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/gocraft/dbr/v2"
)

const mysqlDatabaseName = "fastq_check1"

func setupSQLAutomatic(ctx context.Context, db *sql.DB, folderPath string) error {
	res, err := DoesDatabaseExist(ctx, db)
	if err != nil {
		log.Println("error occured while setting mysql", err)
		return err
	}
	log.Println(res)
	err = CreateDB(ctx, db)
	if err != nil {
		log.Println("error occured while creating DB", err)
		return err
	}
	return nil
}
func CreateDB(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", mysqlDatabaseName))
	return err
}

type SchemaRes struct {
	Name      dbr.NullString `json:"name"`
	TableName dbr.NullString `json:"tableName"`
}

func DoesDatabaseExist(ctx context.Context, db *sql.DB) (bool, error) {
	// Construct the MySQL command to check if the database exists
	var result SchemaRes
	err := db.QueryRowContext(ctx, "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '\""+mysqlDatabaseName+"\"';").Scan(&result)
	if err != nil {
		log.Println("error occured while checking for DB", err)
		return false, err
	}
	// defer rows.Close()
	// var items []*models.ManageBranch
	log.Println("DB checked", result)
	return true, nil
}
func ExecuteSQLScripts(ctx context.Context, db *sql.DB, folderPath string) error {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			filePath := filepath.Join(folderPath, file.Name())
			sqlContent, err := ioutil.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("error reading SQL file %s: %w", filePath, err)
			}

			fmt.Printf("Executing SQL file: %s\n", filePath)

			command := string(sqlContent)
			_, err = db.ExecContext(ctx, command)
			if err != nil {
				log.Println("error occured while executing script", err)
			}
			// return err
			fmt.Printf("SQL file %s executed successfully\n", filePath)
		}
	}

	return nil
}
