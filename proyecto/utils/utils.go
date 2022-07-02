package utils

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetPWD() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

// Filename is the __filename equivalent
func Filename() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal(errors.New("unable to get the current filename"))
	}
	return filename
}

// Dirname is the __dirname equivalent
func Dirname() string {
	filename := Filename()
	return filepath.Dir(filename)
}
