package util

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(fn string) {
	f, err := os.Open(fn, os.O_RDWR|os.O_CREATE|os.O_APPEND, 066)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)

}
