package environ

import (
	"os"
)

func Env() []string {
	env_var := os.Environ()

	return env_var

}
