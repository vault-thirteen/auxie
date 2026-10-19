package env

import (
	"os"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Errf_EnvEmpty = "environment variable '%v' is empty"
)

func GetEnv(variableName string) (envValue string, err error) {
	envValue = os.Getenv(variableName)

	if len(envValue) == 0 {
		err = errors.ErrorFromFormattedText(Errf_EnvEmpty, variableName)
		return envValue, err
	}

	return envValue, nil
}
