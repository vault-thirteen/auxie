package main

import (
	"database/sql"
	"errors"
	"net"
	"os"

	"github.com/go-sql-driver/mysql"
)

func mustBeNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := getScriptFilePath()
	mustBeNoError(err)

	dsn := createDsn()
	db, err := connectToDb(dsn)
	mustBeNoError(err)
	defer func() {
		derr := disconnectFromDb(db)
		mustBeNoError(derr)
	}()

	err = execQueryFromFile(file, db)
	mustBeNoError(err)
}

func getScriptFilePath() (sfp string, err error) {
	if len(os.Args) < 2 {
		return "", errors.New("not enough arguments")
	}

	return os.Args[1], nil
}

func createDsn() (dsn string) {
	mc := mysql.Config{
		Net:                  "tcp",
		Addr:                 net.JoinHostPort("localhost", "3306"),
		DBName:               "test",
		User:                 "test",
		Passwd:               "test",
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		MaxAllowedPacket:     1_000_000_000,
		Params:               map[string]string{},
	}

	return mc.FormatDSN()
}

func connectToDb(dsn string) (db *sql.DB, err error) {
	return sql.Open("mysql", dsn)
}

func disconnectFromDb(db *sql.DB) (err error) {
	return db.Close()
}

func execQueryFromFile(queryFilePath string, db *sql.DB) (err error) {
	buf, err := os.ReadFile(queryFilePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(buf))
	if err != nil {
		return err
	}

	return nil
}
