package main

import (
	"fmt"
	"os"
)

func main() {
	destFs := &gcsFs{}

	err := Synchronize(destFs)
	fmt.Println(err)

}

type FileSystem interface {
	MkdirAll(name string, perm os.FileMode) error
}

type gcsFs struct{}

func (fs *gcsFs) MkdirAll(name string, _ os.FileMode) error {
	return nil
}

func Synchronize(destFs FileSystem) error {
	return nil
}
