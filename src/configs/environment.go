package configs

import (
	"errors"
	"os"
)

func InitEnv() {
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(dirname + "/.env"); errors.Is(err, os.ErrNotExist) {
		env := []byte("DB_HOST=\nDB_PORT=\nDB_USER=\nDB_PASS=\nDB_NAME=\nAPP_KEY=")
		err := os.WriteFile(dirname+"/.env", env, 0644)
		if err != nil {
			panic(err)
		}
	}
}
