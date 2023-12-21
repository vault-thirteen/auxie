package mysql

import (
	"testing"

	_ "github.com/go-sql-driver/mysql" // MySQL driver.
	"github.com/vault-thirteen/auxie/SQL/common/test"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_IdentifierIsGood(t *testing.T) {
	var aTest *tester.Test = tester.New(t)

	// Test #1.
	var identifierName string = "xB_9"
	var result bool
	var err error
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

func Test_TableNameIsGood(t *testing.T) {
	var aTest *tester.Test = tester.New(t)

	// Test #1.
	var identifierName string = "xB_9"
	var result bool
	var err error
	result, err = TableNameIsGood(identifierName)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, true)

	// Test #2.
	identifierName = "xB_9куку"
	result, err = TableNameIsGood(identifierName)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, false)

	// Test #3.
	identifierName = "xB_9!@"
	result, err = TableNameIsGood(identifierName)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, false)

	// Test #4.
	identifierName = "DROP TABLE xyz;"
	result, err = TableNameIsGood(identifierName)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, false)
}

func Test_ScreenSingleBacktickQuotes(t *testing.T) {

	var aTest *tester.Test
	var dst string
	var dstExpected string
	var src string

	aTest = tester.New(t)

	// Test #1.
	src = "John"
	dstExpected = "John"
	dst = ScreenSingleBacktickQuotes(src)
	aTest.MustBeEqual(dst, dstExpected)

	// Test #2.
	src = "D'Artagnan, " +
		`D"Artagnan, ` +
		"D`Artagnan, " +
		"D``Artagnan."
	dstExpected = "D'Artagnan, " +
		`D"Artagnan, ` +
		"D``Artagnan, " +
		"D````Artagnan."
	dst = ScreenSingleBacktickQuotes(src)
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
			Name: "test_GetTableColumnNames",
			Func: test_GetTableColumnNames,
		},
	}

	var ts *ct.TestSuite
	ts, err = ct.NewTestSuite(
		t,
		EnvVarNamePrefix,
		[]ct.IntroFunction{_createTestTable},
		testedFunctions,
		[]ct.OutroFunction{_deleteTestTable},
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
	dbName := sd.ts.GetSchema()

	// Test #1. Table exists.
	var tableExists bool
	tableExists, err = TableExists(sqlConnection, dbName, TableNameExistent)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(tableExists, true)

	// Test #2. Table does not exist.
	tableExists, err = TableExists(sqlConnection, dbName, TableNameNotExistent)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(tableExists, false)
}

// This test depends on the test environment.
// Please ensure that all the parameters are correct before using it.
func test_GetTableColumnNames(t *testing.T) {
	var aTest = tester.New(t)
	var err error

	sqlConnection := sd.ts.GetHandle()
	dbName := sd.ts.GetSchema()

	// Test #1. Table exists.
	var columnNames []string
	columnNames, err = GetTableColumnNames(sqlConnection, dbName, TableNameExistent)
	aTest.MustBeNoError(err)
	expectedColumnNames := []string{"Id", "Name", "Number"}
	aTest.MustBeEqual(columnNames, expectedColumnNames)

	// Test #2. Table does not exist.
	columnNames, err = GetTableColumnNames(sqlConnection, dbName, TableNameNotExistent)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(columnNames, []string(nil))
}
