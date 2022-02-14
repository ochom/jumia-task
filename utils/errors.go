package utils

import "log"

// FailOnError terminates the server on error
func FailOnError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
