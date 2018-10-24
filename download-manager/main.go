package main

import (
	"github.com/cheggaaa/pb"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func main () {
	MakeRequest()
}

func MakeRequest() {

	userInput := os.Args[1]
	fileName := path.Base(userInput)


	out, err := os.Create(fileName)
	defer out.Close()
	resp, err := http.Get(userInput)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// Progress bar
	filesize := resp.ContentLength

	go func() {
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	countSize := int(filesize / 1000)
	bar := pb.StartNew(countSize)
	var fi os.FileInfo
	for fi == nil || fi.Size() < filesize {
		fi, _ = out.Stat()
		bar.Set(int(fi.Size() / 1000))
		time.Sleep(time.Millisecond)
	}

	bar.FinishPrint("All finished")

}





