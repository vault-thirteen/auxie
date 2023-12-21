package ct

import (
	"github.com/vault-thirteen/auxie/env"
)

const (
	EnvVarNameBase_NetProtocol = "PROTO"  // E.g. "tcp".
	EnvVarNameBase_Driver      = "DRIVER" // E.g. "mysql".
	EnvVarNameBase_Host        = "HOST"   // E.g. "localhost".
	EnvVarNameBase_Port        = "PORT"   // E.g. "3306".
	EnvVarNameBase_Database    = "DB"
	EnvVarNameBase_User        = "USER"
	EnvVarNameBase_Password    = "PWD"
)

// TestDbParameters are parameters of a test database.
type TestDbParameters struct {
	NetworkProtocol string // E.g. "tcp".
	Driver          string // E.g. "mysql".
	Host            string // E.g. "localhost".
	Port            string // E.g. "3306".
	DB              string
	User            string
	Pwd             string
}

func NewTestDbParameters(envVarNamePrefix string) (p *TestDbParameters, err error) {
	postfixes := []string{
		EnvVarNameBase_NetProtocol,
		EnvVarNameBase_Driver,
		EnvVarNameBase_Host,
		EnvVarNameBase_Port,
		EnvVarNameBase_Database,
		EnvVarNameBase_User,
		EnvVarNameBase_Password,
	}

	envVarNames := make([]string, 0, len(postfixes))
	for _, postfix := range postfixes {
		envVarName := envVarNamePrefix + postfix
		envVarNames = append(envVarNames, envVarName)
	}

	envVarValues := make([]string, 0, len(envVarNames))
	var envVarValue string
	for _, envVarName := range envVarNames {
		envVarValue, err = env.GetEnv(envVarName)
		if err != nil {
			return nil, err
		}

		envVarValues = append(envVarValues, envVarValue)
	}

	p = &TestDbParameters{
		NetworkProtocol: envVarValues[0],
		Driver:          envVarValues[1],
		Host:            envVarValues[2],
		Port:            envVarValues[3],
		DB:              envVarValues[4],
		User:            envVarValues[5],
		Pwd:             envVarValues[6],
	}

	return p, nil
}
