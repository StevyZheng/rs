package common

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"io/ioutil"
)

func ListFiles(path string) (fileArr arraylist.List, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, fi := range files {
		fileArr.Add(fi.Name())
	}
	return
}
