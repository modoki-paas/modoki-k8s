package testutil

import (
	"os"
	"strings"
)

// SetTemporalEnv sets passed envs, returns a function to rollback env
func SetTemporalEnv(envs ...string) func() {
	rollback := []func(){}
	for i := range envs {
		spl := strings.SplitN(envs[i], "=", 2)

		before, ok := os.LookupEnv(spl[0])

		if ok {
			rollback = append(rollback, func() {
				os.Setenv(spl[0], before)
			})
		} else {
			rollback = append(rollback, func() {
				os.Unsetenv(spl[0])
			})
		}

		os.Setenv(spl[0], spl[1])
	}

	return func() {
		for i := range rollback {
			rollback[i]()
		}
	}
}
