package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

func main() {
	args := sanitizeArgLength(1, 1)
	sanitizeFile(args[0], "image/jpeg", "image/png")
	prepareConfigFolder("wallpaper")
}

// Verify that the program execution arguments are not less than min, and not greater than max.
// Not including the program name 0th argument. This will fatally kill the program if the arguments provided
// do not fit the parameter requirements. Otherwise it will return the sanitized args.
func sanitizeArgLength(min, max int) []string {
	args := os.Args
	if len(args) < min+1 {
		log.Fatal("No arguments provided.")
	} else if len(args) > max+1 {
		log.Fatal("Too many arguments provided.")
	}
	return args[1:]
}

// Create a config folder of a specified name if one does not exist already.
func prepareConfigFolder(folderName string) {
	if ucd, err := os.UserConfigDir(); err == nil {
		dir := fmt.Sprintf("%s/%s", ucd, folderName)
		os.Mkdir(dir, os.FileMode(0755))
	}
}

// Will kill the program fatally if the file does not exist or if it does not fit into one of the specified mimes.
func sanitizeFile(path string, fileMimes ...string) {
	mt, err := mimetype.DetectFile(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, fm := range fileMimes {
		if mt.Is(fm) {
			return
		}
	}
	log.Fatalf("%s is a non-valid file type: %s", path, mt.String())
}
