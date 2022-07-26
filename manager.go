package main

import (
	"fmt"
	"io/ioutil"
	"log"
	os "os"
)

const READ int = os.O_RDONLY
const WRITE int = os.O_WRONLY

type FileManager struct{}

func (f *FileManager) CreateFile(name string) *os.File {
	defer fmt.Println("File has been created")
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	errClose := file.Close()
	if errClose != nil {
		log.Fatal(errClose)
	}
	return file
}

func (f *FileManager) DeleteFile(name string) {
	defer fmt.Println("File deleted")
	if len(name) <= 0 {
		return
	}
	err := os.Remove(name)
	if err != nil {
		log.Fatal(err)
	}
}

func (f *FileManager) RenameFile(name string, newName string) {
	if len(name) <= 0 && len(newName) <= 0 && name != newName {
		return
	}
	defer fmt.Println("File renamed")
	file := f.OpenFile(name, READ)
	data := f.ReadDataFromFile(file.Name())
	err := file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	f.DeleteFile(file.Name())
	f.CreateFile(newName)
	f.WriteDataToFile(newName, data)
}

func (f *FileManager) OpenFile(name string, flag int) *os.File {
	if len(name) <= 0 {
		return nil
	}
	file, err := os.OpenFile(name, flag, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func (f *FileManager) CopyFile(name string) {
	if len(name) <= 0 {
		return
	}
	data := f.ReadDataFromFile(name)
	f.WriteDataToFile("c_"+name, data)
}

func (f *FileManager) WriteDataToFile(name string, data []byte) {
	errWrite := ioutil.WriteFile(name, data, 0666)
	if errWrite != nil {
		log.Fatal(errWrite)
		return
	}
}

func (f *FileManager) ReadDataFromFile(name string) []byte {
	data, errRead := ioutil.ReadFile(name)
	if errRead != nil {
		log.Fatal(errRead)
		return nil
	}
	return data
}
