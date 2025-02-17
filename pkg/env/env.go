package env

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func ReadEnv(key string) (string, error) {
	str := os.Getenv(key)
	if str == "" {
		msg := fmt.Sprintf("Error in reading %s", key)
		log.Fatalln(msg)
		return msg, errors.New(msg)
	}

	return str, nil
}
