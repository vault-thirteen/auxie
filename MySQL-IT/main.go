package main

import (
	"database/sql"
	"net"
	"os"

	"github.com/go-sql-driver/mysql"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Err_NotEnoughArguments = "not enough arguments"
)

const (
	DriverName_MySQL = "mysql"
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
	var db *sql.DB
	db, err = connectToDb(dsn)
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
		return "", errors.ErrorFromText(Err_NotEnoughArguments)
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
	return sql.Open(DriverName_MySQL, dsn)
}

func disconnectFromDb(db *sql.DB) (err error) {
	return db.Close()
}

func execQueryFromFile(queryFilePath string, db *sql.DB) (err error) {
	var buf []byte
	buf, err = os.ReadFile(queryFilePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(buf))
	if err != nil {
		return err
	}

	return nil
}
