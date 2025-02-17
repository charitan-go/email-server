package env

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func ReadEnv(email string) (string, error) {
	str := os.Getenv(email)
	if str == "" {
		msg := fmt.Sprintf("Error in reading %s", email)
		log.Fatalln(msg)
		return msg, errors.New(msg)
	}

	return str, nil
}
