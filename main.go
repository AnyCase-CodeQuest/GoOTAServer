package main

import (
	"awesomeProject1/HTTPServer/ESP"
	"crypto/md5"
	"io"
	"log"
	"net/http"
	"os"
)

const defaultLogfile string = "testlogfile"
const EnvLogVar string = "SRV_LOG"

func co2update(w http.ResponseWriter, req *http.Request) {
	var esp = ESP.NewESP(req)
	if ESP.CanBeUpdated(esp) {
		w.WriteHeader(200)
		file, err := os.OpenFile(ESP.BinFilePath, os.O_RDONLY, 0777)
		if err != nil {
			log.Printf("Openfile Error")
		}
		fileInfo, err := file.Stat()
		if err != nil {
			log.Printf("Stat file Error")
		}
		h := md5.New()
		if _, err := io.Copy(h, file); err != nil {
			log.Println(err)
		}
		w.Header().Add("x-MD5", string(h.Sum(nil)))
		w.Header().Add("Content-Length", string(fileInfo.Size()))
		written, err := io.Copy(w, file)
		if written != fileInfo.Size() {
			log.Printf("Sized different %d and %d", written, fileInfo.Size())
		}

	} else {
		w.WriteHeader(304)
	}
}

func setupLogger(filename string) {
	//rewrite to custom log place
	var envLogFile = os.Getenv(EnvLogVar)
	if envLogFile != "" {
		filename = envLogFile
	}

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer f.Close()

	log.SetOutput(f)
}

func main() {
	setupLogger(defaultLogfile)
	log.Println("Start GO server...")
	http.HandleFunc("/co2.bin", co2update)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
