package utils

import (
	"log"
	"os"
)

func GetPWD() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}
