package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/sidhaler/wrt/CopyContent"
	"github.com/sidhaler/wrt/ReadFile"
	"github.com/sidhaler/wrt/WrittingContent"
)

var args struct {
	FilePath          string `arg:"positional, required" help:"path to file"`
	ContentorNextFile string `arg:"positional" help"Content to write into file"`
	Sc                bool   `arg:"-s" default:"false" help:"Show content of file then exit"`
	CopyContent       bool   `arg:"-c" default:"false" help:"Copy content from one source to another"`
	Movefile          bool   `arg:"-m" default:"false" help:"Move file to destination"`
}

func Iserr(err error) {
	if err != nil {
		fmt.Println("We got unexpected error while reading file")
		os.Exit(1)
	}
}

func main() {
	arg.MustParse(&args)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	ReadFile.IsFilleAv(&args.FilePath, &args.ContentorNextFile, wg)

	// YEP
	if args.CopyContent == true {
		datachannel := make(chan []byte, 4)
		wg.Add(3)
		go CopyWorker(args.FilePath, args.ContentorNextFile, wg, datachannel)
		go CopyWorker(args.FilePath, args.ContentorNextFile, wg, datachannel)

		wg.Wait()
		defer close(datachannel)
	}
	// XD
	if args.Sc == true {
		wg.Add(1)
		go ReadWorker(args.FilePath, wg)
		wg.Wait()
	}
	wg.Add(1)
	go WrittingWorker(&args.FilePath, wg, &args.ContentorNextFile)
	wg.Wait()
}

//READING AND DISPLAYING CONTENT OF FILE
func ReadWorker(p string, wg *sync.WaitGroup) {
	wg.Done()
	ReadFile.Readfromfile(p)
}

//MOST INSANE THING I EVER DID IN GO, and I don't even know how it works :0
func CopyWorker(src string, dst string, wg *sync.WaitGroup, datach chan []byte) {
	wg.Done()
	go ReadFile.FastRead(src, wg, datach)
	go CopyContent.CopytoFile(src, dst, wg, datach)
	wg.Wait()
	time.Sleep(time.Millisecond * 200)
}

//WRITE CONTENT
func WrittingWorker(target *string, wg *sync.WaitGroup, content *string) {
	wg.Done()
	go WrittingContent.Write(content, target)
	//fmt.Println("Content has been succesfully written into file")
	wg.Wait()
}
