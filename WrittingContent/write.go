package WrittingContent

import (
	"io/ioutil"

	"github.com/sidhaler/wrt/ReadFile"
)

func Write(content *string, target *string) {
	perms := ReadFile.ReadPerm(*target)
	cn := []byte(*content)
	ioutil.WriteFile(*target, cn, perms)
}

//https://jogendra.dev/import-cycles-in-golang-and-how-to-deal-with-them
