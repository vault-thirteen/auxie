package mysql

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql" // MySQL driver.

	"github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/auxie/unicode"
)

const (
	Errf_BadSymbol = "bad symbol: '%v'"
)

// SQL queries.
const (
	QueryfTableExists = `SELECT COUNT(*)
FROM information_schema.tables
WHERE
	(table_schema = ?) AND
	(table_name = ?);`

	QueryfTableColumnNames = `SELECT
	COLUMN_NAME AS col
FROM INFORMATION_SCHEMA.COLUMNS
WHERE
	(table_schema = ?) AND
    (table_name = ?)
ORDER BY col ASC;`
)

// Symbols.
const (
	SingleBacktickQuote = "`"
	DoubleBacktickQuote = SingleBacktickQuote + SingleBacktickQuote
	Underscore          = '_'
)

func IdentifierIsGood(identifierName string) (ok bool, err error) {
	for _, letter := range identifierName {
		if (!unicode.SymbolIsLatLetter(letter)) &&
			(!unicode.SymbolIsNumber(letter)) &&
			(letter != Underscore) {
			return false, errors.ErrorFromFormattedText(Errf_BadSymbol, string(letter))
		}
	}

	return true, nil
}

func TableNameIsGood(tableName string) (ok bool, err error) {
	return IdentifierIsGood(tableName)
}

// ScreenSingleBacktickQuotes does the single quotes screening.
func ScreenSingleBacktickQuotes(src string) (dst string) {
	return strings.ReplaceAll(src, SingleBacktickQuote, DoubleBacktickQuote)
}

// TableExists checks whether the specified table exists or not.
func TableExists(connection *sql.DB, schemaName string, tableName string) (tableExists bool, err error) {
	var statement *sql.Stmt
	statement, err = connection.Prepare(QueryfTableExists)
	if err != nil {
		return tableExists, err
	}

	defer func() {
		derr := statement.Close()
		err = errors.Combine(err, derr)
	}()

	row := statement.QueryRow(schemaName, tableName)
	var tablesCount int
	err = row.Scan(&tablesCount)
	if err != nil {
		return tableExists, err
	}

	if tablesCount == 1 {
		return true, nil
	}

	return false, nil
}

// GetTableColumnNames lists the table's column names sorted alphabetically
// (from A to Z).
func GetTableColumnNames(connection *sql.DB, schemaName string, tableName string) (columnNames []string, err error) {
	var statement *sql.Stmt
	statement, err = connection.Prepare(QueryfTableColumnNames)
	if err != nil {
		return columnNames, err
	}

	defer func() {
		derr := statement.Close()
		err = errors.Combine(err, derr)
	}()

	var rows *sql.Rows
	rows, err = statement.Query(schemaName, tableName)
	if err != nil {
		return columnNames, err
	}

	defer func() {
		derr := rows.Close()
		err = errors.Combine(err, derr)
	}()

	var columnName string
	for rows.Next() {
		err = rows.Scan(&columnName)
		if err != nil {
			return columnNames, err
		}

		columnNames = append(columnNames, columnName)
	}

	return columnNames, nil
}
