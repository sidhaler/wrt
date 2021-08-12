package CopyContent

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/sidhaler/wrt/ReadFile"
)

func isErr(err error) {
	if err != nil {
		log.Fatal("We got unexpected Error while copying")
	}
}

func CopytoFile(pathto1 string, pathto2 string, wg *sync.WaitGroup, c chan []byte) {
	wg.Done()
	dd := <-c
	perms := ReadFile.ReadPerm(pathto2)
	ioutil.WriteFile(pathto2, dd, perms)
}

func MoveFile(src *string, dst *string) {
	err := os.Rename(*src, *dst)
	isErr(err)
}
