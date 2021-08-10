package CopyContent

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/sidhaler/wrt/ReadFile"
)

func CopytoFile(pathto1 string, pathto2 string) {
	var dd []byte
	//res := make(chan []byte, 1)
	ReadFile.Readfromfile(pathto1)
	time.Sleep(time.Second * 5)
	fmt.Println(string(dd))
	// go WriteTo(pathto2, dd)
}

func WriteTo(path string, c []byte) {
	hello := ReadFile.ReadPerm(path)
	//defer close(c)
	ioutil.WriteFile(path, c, hello)
}
