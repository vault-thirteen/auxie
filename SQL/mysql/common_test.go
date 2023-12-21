package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver.
	"github.com/vault-thirteen/auxie/SQL/common/test"
)

const (
	// EnvVarNamePrefix is a prefix for names of environment variables storing
	// parameters of test database.
	EnvVarNamePrefix = "MYSQL_TEST_"

	TableNameExistent    = "TableA"
	TableNameNotExistent = "xxxxxxxxx"
)

// SharedData is a shared object for tested functions.
// This shared object is used to pass arguments to tests.
//
// In Golang, tested functions are required to have a strict signature having a
// single input parameter of the '*testing.T' type. Such restriction makes it
// difficult to pass arguments to tested functions. So, we are using a shared
// object to pass arguments to tests.
type SharedData struct {
	ts *ct.TestSuite
}

var sd *SharedData

func _createTestTable(ts *ct.TestSuite) (err error) {
	const QueryfCreateTestTableA = `CREATE TABLE IF NOT EXISTS %v
(
	Id serial,
	Name varchar(255),
	Number int
);`

	sqlConnection := ts.GetHandle()

	tableName := TableNameExistent

	var ok bool
	ok, err = TableNameIsGood(tableName)
	if (err != nil) || !ok {
		return err
	}

	// Create the table.
	query := fmt.Sprintf(QueryfCreateTestTableA, tableName)
	_, err = sqlConnection.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func _deleteTestTable(ts *ct.TestSuite) (err error) {
	const QueryfDeleteTestTableA = `DROP TABLE IF EXISTS %v;`

	sqlConnection := ts.GetHandle()

	tableName := TableNameExistent

	var ok bool
	ok, err = TableNameIsGood(tableName)
	if (err != nil) || !ok {
		return err
	}

	// Delete the table.
	query := fmt.Sprintf(QueryfDeleteTestTableA, tableName)
	_, err = sqlConnection.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
