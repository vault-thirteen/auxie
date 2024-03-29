package postgresql

import (
	"testing"

	_ "github.com/lib/pq" // PostgreSQL driver.
	"github.com/vault-thirteen/auxie/SQL/common/test"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_MakeDsn(t *testing.T) {

	var dsnExpected string
	var dsnReceived string

	// Test #1. Full String.
	dsnExpected = "postgresql://vasya:pwd@localhost:1234/vasyadb?xyz=123"
	dsnReceived = MakeDsn(
		"localhost",
		"1234",
		"vasyadb",
		"vasya",
		"pwd",
		"xyz=123",
	)
	if dsnExpected != dsnReceived {
		t.Error("Full String")
		t.FailNow()
	}

	// Test #2. No Password.
	dsnExpected = "postgresql://vasya@localhost:1234/vasyadb?xyz=123"
	dsnReceived = MakeDsn(
		"localhost",
		"1234",
		"vasyadb",
		"vasya",
		"",
		"xyz=123",
	)
	if dsnExpected != dsnReceived {
		t.Error("No Password")
		t.FailNow()
	}

	// Test #3. No Username.
	dsnExpected = "postgresql://localhost:1234/vasyadb?xyz=123"
	dsnReceived = MakeDsn(
		"localhost",
		"1234",
		"vasyadb",
		"",
		"password-not-used",
		"xyz=123",
	)
	if dsnExpected != dsnReceived {
		t.Error("No Username")
		t.FailNow()
	}

	// Test #4. No Database.
	dsnExpected = "postgresql://localhost:1234?xyz=123"
	dsnReceived = MakeDsn(
		"localhost",
		"1234",
		"",
		"",
		"password-not-used",
		"xyz=123",
	)
	if dsnExpected != dsnReceived {
		t.Error("No Database")
		t.FailNow()
	}

	// Test #5. No Parameters.
	dsnExpected = "postgresql://localhost:1234"
	dsnReceived = MakeDsn(
		"localhost",
		"1234",
		"",
		"",
		"password-not-used",
		"",
	)
	if dsnExpected != dsnReceived {
		t.Error("No Parameters")
		t.FailNow()
	}
}

func Test_IdentifierIsGood(t *testing.T) {

	var aTest *tester.Test
	var err error
	var identifierName string
	var result bool

	aTest = tester.New(t)

	// Test #1.
	identifierName = "xB_9"
	result, err = IdentifierIsGood(identifierName)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, true)

	// Test #2.
	identifierName = "xB_9куку"
	result, err = IdentifierIsGood(identifierName)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, false)

	// Test #3.
	identifierName = "xB_9!@"
	result, err = IdentifierIsGood(identifierName)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, false)

	// Test #4.
	identifierName = "DROP TABLE xyz;"
	result, err = IdentifierIsGood(identifierName)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, false)
}

func Test_ScreenSingleQuotes(t *testing.T) {

	var aTest *tester.Test
	var dst string
	var dstExpected string
	var src string

	aTest = tester.New(t)

	// Test #1.
	src = `John`
	dstExpected = `John`
	dst = ScreenSingleQuotes(src)
	aTest.MustBeEqual(dst, dstExpected)

	// Test #2.
	src = `John's Car`
	dstExpected = `John''s Car`
	dst = ScreenSingleQuotes(src)
	aTest.MustBeEqual(dst, dstExpected)

	// Test #3.
	src = `John''x`
	dstExpected = `John''''x`
	dst = ScreenSingleQuotes(src)
	aTest.MustBeEqual(dst, dstExpected)
}

func Test_Suite(t *testing.T) {
	var aTest = tester.New(t)
	var err error

	var testedFunctions = []*ct.TestedFunction{
		{
			Name: "test_TableExists",
			Func: test_TableExists,
		},
		{
			Name: "test_ProcedureExists",
			Func: test_ProcedureExists,
		},
	}

	var ts *ct.TestSuite
	ts, err = ct.NewTestSuite(
		t,
		EnvVarNamePrefix,
		[]ct.IntroFunction{_createTestTable, _createTestProcedure},
		testedFunctions,
		[]ct.OutroFunction{_deleteTestTable, _deleteTestProcedure},
	)
	aTest.MustBeNoError(err)

	sd = &SharedData{
		ts: ts,
	}

	ts.Run()
}

// This test depends on the test environment.
// Please ensure that all the parameters are correct before using it.
func test_TableExists(t *testing.T) {
	var aTest = tester.New(t)
	var err error

	sqlConnection := sd.ts.GetHandle()

	// Test #1. Table exists.
	var tableExists bool
	tableExists, err = TableExists(sqlConnection, SchemaCommon, TableNameExistent)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(tableExists, true)

	// Test #2. Table does not exist.
	tableExists, err = TableExists(sqlConnection, SchemaCommon, TableNameNotExistent)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(tableExists, false)
}

// This test depends on the test environment.
// Please ensure that all the parameters are correct before using it.
func test_ProcedureExists(t *testing.T) {
	var aTest = tester.New(t)
	var err error

	sqlConnection := sd.ts.GetHandle()

	// Test #1. Procedure exists.
	var procedureExists bool
	procedureExists, err = ProcedureExists(sqlConnection, SchemaCommon, ProcedureNameExistent)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(procedureExists, true)

	// Test #2. Procedure does not exist.
	procedureExists, err = ProcedureExists(sqlConnection, SchemaCommon, ProcedureNameNotExistent)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(procedureExists, false)
}
