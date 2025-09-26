package internal

import (
	"os"

	"github.com/pkg/errors"
)

// reads fileName arg
func ReadFileNameArg() (string, error) {
	args := os.Args

	if len(args) < 2 {
		return "", errors.New("Parameters count mismatch.")
	}

	return args[1], nil
}
