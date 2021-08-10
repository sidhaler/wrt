package ReadFile

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Iserr(err error) {
	if err != nil {
		fmt.Println("We got unexpected error while reading file")
		os.Exit(1)
	}
}

func IsFilleAv(path string) {
	_, err := os.Lstat(path)
	if err != nil {
		fmt.Println("File isn't created yet, wanna do it rn ?[Y,y/N,n] : ")
		var ans string
		fmt.Scanln(&ans)
		switch ans {
		case "Y", "y":
			// bb := strings.Split(path, "/")
			// filename := bb[len(bb)-1]
			os.Create(path)
		case "N", "n":
			os.Exit(00)
		}
	}

}

func Readfromfile(hey string) {
	data, err := ioutil.ReadFile(hey)
	Iserr(err)
	fmt.Println(string(data))
}

func FastRead(target string) []byte {
	data, err := ioutil.ReadFile(target)
	Iserr(err)
	return data
}

func ReadPerm(pathtofile string) os.FileMode {
	inf, err := os.Stat(pathtofile)
	Iserr(err)
	fmt.Println("gg")
	mode := inf.Mode()
	return mode
}
