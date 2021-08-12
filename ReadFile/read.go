package ReadFile

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	//"github.com/sidhaler/wrt/WrittingContent"
	Miniworkers "github.com/sidhaler/wrt/MiniWorkers"
)

func Iserr(err error) {
	if err != nil {
		fmt.Println("We got unexpected error while reading file")
		os.Exit(1)
	}
}

func IsFilleAv(path *string, con *string, wg *sync.WaitGroup) {
	_, err := os.Lstat(*path)
	if err != nil {
		fmt.Println("File isn't created yet, wanna do it rn ?[Y,y/N,n] : ")
		var ans string
		fmt.Scanln(&ans)
		switch ans {
		case "Y", "y":
			os.Create(*path)
		case "N", "n":
			os.Exit(00)
		}
		obj := new(Miniworkers.FakeWorkers)
		Wri
		//WrittingContent.FakeWorker(con, wg, path)
	}

}

func Readfromfile(hey string) {
	data, err := ioutil.ReadFile(hey)
	Iserr(err)
	fmt.Println(string(data))
}

func FastRead(target string, wg *sync.WaitGroup, c chan []byte) {
	defer wg.Done()
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
