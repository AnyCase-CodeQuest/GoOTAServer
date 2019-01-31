package ESP

import (
	"log"
	"os"
)

const BinFilePath = "./HTTPServer/ESP/scetch/index.ino.generic.bin"

func CanBeUpdated(device *ESP) bool {
	fileSize := getFileSize()
	if device.FreeSpace() >= fileSize {
		return true
	}
	log.Printf("File size is: %d", fileSize)
	return false
}

func getFileSize() int64 {
	fi, err := os.Stat(BinFilePath)

	if err != nil {
		log.Println(err)
	}

	return fi.Size()
}
