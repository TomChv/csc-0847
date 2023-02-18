package utils

import (
	"fmt"
	"log"
	"os"
)

// ForceGetEnv retrieve an environment variable
// or exit fatal if the variable is missing.
func ForceGetEnv(name string) string {
	value, exist := os.LookupEnv(name)
	if !exist {
		log.Fatalln(fmt.Errorf("missing %s in environment", name))
	}

	return value
}
