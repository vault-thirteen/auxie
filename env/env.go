// env.go.

package env

import (
	"fmt"
	"os"
)

const ErrfEnvEmpty = "environment variable '%v' is empty"

func GetEnv(
	variableName string,
) (envValue string, err error) {
	envValue = os.Getenv(variableName)

	if len(envValue) == 0 {
		err = fmt.Errorf(ErrfEnvEmpty, variableName)
		return "", err
	}

	return envValue, nil
}
