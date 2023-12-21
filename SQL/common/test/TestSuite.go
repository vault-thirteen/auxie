package ct

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/go-sql-driver/mysql" // MySQL driver.
	_ "github.com/lib/pq"            // PostgreSQL driver.
)

const (
	ErrTestedFunctionsAreNotSet = "tested functions are not set"
	ErrUnknownSqlDriver         = "unknown SQL driver: %v"
	ErrConnectionIsNull         = "connection is null"
)

// List of supported SQL drivers.
const (
	DriverMySQL      = "mysql"
	DriverPostgreSQL = "postgres"
)

// TestSuite is a suite of several tests. To run the tests create a test suite
// and use the 'Run' method.
//
// Test suite automatically connects to and disconnects from the test database,
// runs an "intro" function, then runs all the tests (tested functions), and
// then runs the "outro" function. Intro-function is used to prepare tested
// database objects, e.g. to create tables and data. Outro-function is used to
// clean up the mess in the database after the tests, e.g. to delete data and
// tables.
type TestSuite struct {
	t            *testing.T
	dbParameters *TestDbParameters
	dbDsn        string
	dbConnection *sql.DB

	intro           []IntroFunction
	testedFunctions []*TestedFunction
	outro           []OutroFunction
}

func NewTestSuite(
	t *testing.T,
	envVarNamePrefix string, // Prefix of environment variable names.
	intro []IntroFunction, // Function which prepares the database.
	testedFunctions []*TestedFunction, // Tested functions.
	outro []OutroFunction, // Function which cleans the database.
) (ts *TestSuite, err error) {
	if (testedFunctions == nil) || (len(testedFunctions) == 0) {
		return nil, errors.New(ErrTestedFunctionsAreNotSet)
	}

	var dbParameters *TestDbParameters
	dbParameters, err = NewTestDbParameters(envVarNamePrefix)
	if err != nil {
		return nil, err
	}

	// DSN.
	var dbDsn string
	switch dbParameters.Driver {
	case DriverMySQL:
		dbDsn = makeMySQLDatabaseDsn(dbParameters)
	case DriverPostgreSQL:
		dbDsn = makePostgreSQLDatabaseDsn(dbParameters)
	default:
		return nil, fmt.Errorf(ErrUnknownSqlDriver, dbParameters.Driver)
	}

	ts = &TestSuite{
		t:            t,
		dbParameters: dbParameters,
		dbDsn:        dbDsn,

		intro:           intro,
		testedFunctions: testedFunctions,
		outro:           outro,
	}

	return ts, nil
}

func makeMySQLDatabaseDsn(dbp *TestDbParameters) (dsn string) {
	// Pay attention how the MySQL driver provides a method for obtaining a DSN.
	// Check out how many options does the 'mysql.Config' type support.
	// And compare it with a "community supported" PostgreSQL below ...
	var configuration = mysql.Config{
		Net:  dbp.NetworkProtocol,
		Addr: net.JoinHostPort(dbp.Host, dbp.Port),
		Params: map[string]string{
			"allowNativePasswords": "true",
		},
		DBName: dbp.DB,
		User:   dbp.User,
		Passwd: dbp.Pwd,
	}

	return configuration.FormatDSN()
}

func makePostgreSQLDatabaseDsn(dbp *TestDbParameters) (dsn string) {
	// Unfortunately, the creator of this PostgreSQL driver decided not to make
	// a method for DSN composition. This shows the real difference in quality
	// between MySQL and PostgreSQL database drivers.

	dsn = "postgresql://" + dbp.User + ":" + dbp.Pwd + "@" +
		net.JoinHostPort(dbp.Host, dbp.Port) + "/" + dbp.DB +
		"?" + "sslmode=disable"

	return dsn
}

// Run runs all the tests (tested functions).
func (ts *TestSuite) Run() {
	var err error
	err = ts.connectToDb()
	if err != nil {
		ts.t.Fatal(err)
		return
	}

	defer func() {
		derr := ts.disconnectFromDb()
		if derr != nil {
			ts.t.Fatal(derr)
			return
		}
	}()

	for _, introFn := range ts.intro {
		err = introFn(ts)
		if err != nil {
			ts.t.Fatal(err)
			return
		}
	}

	defer func() {
		for _, outroFn := range ts.outro {
			err = outroFn(ts)
			if err != nil {
				ts.t.Fatal(err)
				return
			}
		}
	}()

	var ok bool
	for _, testedFunction := range ts.testedFunctions {
		ok = ts.t.Run(testedFunction.Name, testedFunction.Func)
		if !ok {
			ts.t.Fail()
		}
	}
}

func (ts *TestSuite) connectToDb() (err error) {
	ts.dbConnection, err = sql.Open(ts.dbParameters.Driver, ts.dbDsn)
	if err != nil {
		return err
	}

	err = ts.dbConnection.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (ts *TestSuite) disconnectFromDb() (err error) {
	if ts.dbConnection == nil {
		return errors.New(ErrConnectionIsNull)
	}

	return ts.dbConnection.Close()
}

// GetHandle returns the SQL database handle.
func (ts *TestSuite) GetHandle() (db *sql.DB) {
	return ts.dbConnection
}

// GetSchema returns the database schema.
func (ts *TestSuite) GetSchema() (schema string) {
	return ts.dbParameters.DB
}
