package WrittingContent

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/sidhaler/wrt/ReadFile"
)

func WritetoOld(content string, target string, oldcontent string) {
	//var newline string = "\n"
	perms := ReadFile.ReadPerm(target)
	content = oldcontent + "\n" + content
	fmt.Println(content)
	cn := []byte(content)
	ioutil.WriteFile(target, cn, perms)
	fmt.Println("Success write into", target)
}

func IsFilleAv(path string, con string, wg *sync.WaitGroup) {
	_, err := os.Lstat(path)
	if err != nil {
		fmt.Println("File isn't created yet, wanna do it rn ?[Y,y/N,n] : ")
		var ans string
		fmt.Scanln(&ans)
		switch ans {
		case "Y", "y":
			os.Create(path)
		case "N", "n":
			os.Exit(00)
		}
		if len(con) > 0 {
			WriteNew(path, con)
			wg.Done()
			os.Exit(00)
		}
		os.Exit(00)
	}
	wg.Done()

}

func WriteNew(path string, con string) {
	perms := ReadFile.ReadPerm(path)
	cn := []byte(con)
	ioutil.WriteFile(path, cn, perms)
}
