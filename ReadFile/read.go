package ReadFile

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func Iserr(err error) {
	if err != nil {
		fmt.Println("We got unexpected error while reading file")
		os.Exit(1)
	}
}

func Readfromfile(hey string) string {
	data, err := ioutil.ReadFile(hey)
	Iserr(err)
	return string(data)
}

func FastRead(target string, wg *sync.WaitGroup, c chan []byte) {
	wg.Done()
	data, err := ioutil.ReadFile(target)
	Iserr(err)
	c <- data
}

func ReadPerm(pathtofile string) os.FileMode {
	inf, err := os.Stat(pathtofile)
	Iserr(err)
	mode := inf.Mode()
	return mode
}
