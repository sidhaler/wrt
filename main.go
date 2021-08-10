package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/sidhaler/wrt/CopyContent"
	"github.com/sidhaler/wrt/ReadFile"
)

var args struct {
	FilePath          string `arg:"positional, required" help:"path to file"`
	ContentorNextFile string `arg:"positional" help"Content to write into file"`
	Sc                bool   `arg:"-s" default:"false" help:"Show content of file then exit"`
	CopyContent       bool   `arg:"-c" default:"false" help:"Copy content from one source to another"`
}

func Iserr(err error) {
	if err != nil {
		fmt.Println("We got unexpected error while reading file")
		os.Exit(1)
	}
}

func main() {
	arg.MustParse(&args)
	if args.Sc == true {
		ReadWorker(args.FilePath)
		os.Exit(00)
	}
	//ReadFile.IsFilleAv(args.FilePath)
	if args.CopyContent == true {
		go func() {
			CopyWorker(args.FilePath, args.ContentorNextFile)
		}()
	}
}

func ReadWorker(p string) {
	go ReadFile.Readfromfile(p)
	time.Sleep(time.Millisecond * 20)
}

func CopyWorker(src string, dst string) {
	CopyContent.CopytoFile(src, dst)
	time.Sleep(time.Millisecond * 200)
}
