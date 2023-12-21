package postgresql

import (
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver.
	ct "github.com/vault-thirteen/auxie/SQL/common/test"
)

const (
	// EnvVarNamePrefix is a prefix for names of environment variables storing
	// parameters of test database.
	EnvVarNamePrefix = "POSTGRESQL_TEST_"

	// SchemaCommon is a common schema.
	SchemaCommon = "public"

	TableNameExistent        = "TableA"
	TableNameNotExistent     = "xxxxxxxxx"
	ProcedureNameExistent    = "procedure_simulator"
	ProcedureNameNotExistent = "xxxxxxxxx"
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
	const QueryfCreateTable = `CREATE TABLE IF NOT EXISTS %v 
(
    "Id" serial 
);`

	sqlConnection := ts.GetHandle()

	tableName := fmt.Sprintf(`%s."%s"`, SchemaCommon, TableNameExistent)

	var ok bool
	ok, err = TableNameIsGood(SchemaCommon)
	if (err != nil) || !ok {
		return err
	}
	ok, err = TableNameIsGood(TableNameExistent)
	if (err != nil) || !ok {
		return err
	}

	// Create the Table.
	query := fmt.Sprintf(QueryfCreateTable, tableName)
	_, err = sqlConnection.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func _deleteTestTable(ts *ct.TestSuite) (err error) {
	const QueryfDeleteTestTableA = `DROP TABLE IF EXISTS %v;`

	sqlConnection := ts.GetHandle()

	tableName := fmt.Sprintf(`%s."%s"`, SchemaCommon, TableNameExistent)

	var ok bool
	ok, err = TableNameIsGood(SchemaCommon)
	if (err != nil) || !ok {
		return err
	}
	ok, err = TableNameIsGood(TableNameExistent)
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

func _createTestProcedure(ts *ct.TestSuite) (err error) {
	const QueryfCreateProcedure = `CREATE OR REPLACE PROCEDURE %v()
LANGUAGE SQL
AS
$$
	SELECT 123
$$;`

	sqlConnection := ts.GetHandle()

	procedureName := ProcedureNameExistent

	var ok bool
	ok, err = ProcedureNameIsGood(procedureName)
	if (err != nil) || !ok {
		return err
	}

	// Create the procedure.
	query := fmt.Sprintf(QueryfCreateProcedure, procedureName)

	_, err = sqlConnection.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func _deleteTestProcedure(ts *ct.TestSuite) (err error) {
	const QueryfDeleteProcedure = `DROP PROCEDURE IF EXISTS %v;`

	sqlConnection := ts.GetHandle()

	procedureName := ProcedureNameExistent

	var ok bool
	ok, err = ProcedureNameIsGood(procedureName)
	if (err != nil) || !ok {
		return err
	}

	// Create the procedure.
	query := fmt.Sprintf(QueryfDeleteProcedure, procedureName)

	_, err = sqlConnection.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
