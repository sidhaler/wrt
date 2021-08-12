package Miniworkers

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type FakeWorkers struct {
	dst     *string
	wg      *sync.WaitGroup
	content *string
}

func (l *FakeWorkers) WriteWorker() {
	l.wg.Done()
	st, err := os.Stat(*l.dst)
	if err != nil {
		log.Fatal("We got unexpected error while reading file perms")
	}
	perms := st.Mode()
	siema := []byte(*l.content)
	ioutil.WriteFile(*l.dst, siema, perms)
	l.wg.Wait()
}

func (l *FakeWorkers) New(destt *string, cont *string, wg *sync.WaitGroup) FakeWorkers {
	return FakeWorkers{
		dst:     destt,
		content: cont,
		wg:      wg,
	}
}
