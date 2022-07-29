package manager

import (
	"io/ioutil"
	"log"
	"os"
)

const READ = os.O_RDONLY
const WRITE = os.O_WRONLY

type FileManager struct{}

func (f FileManager) CreateFile(name string) *os.File {
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

func (f FileManager) CreateFolder(name string) {
	if err := os.Mkdir(name, os.ModeDir); err != nil {
		log.Fatal(err)
	}
}

func (f FileManager) DeleteFile(name string) {
	if len(name) <= 0 {
		log.Fatal("filename len must be greater than 0")
		return
	}
	err := os.Remove(name)
	if err != nil {
		return
	}
}

func (f FileManager) RenameFile(name, newName string) {
	if len(name) <= 0 && len(newName) <= 0 && name != newName {
		return
	}
	file, isExists := f.OpenFile(name, READ)
	if !isExists {
		return
	}
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

func (f FileManager) OpenFile(name string, flag int) (file *os.File, isCreated bool) {
	if len(name) <= 0 {
		log.Fatal("filename len must be greater than 0")
		return nil, false
	}
	file, err := os.OpenFile(name, flag, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file, true
}

func (f FileManager) CopyFile(name string) {
	if len(name) <= 0 {
		log.Fatal("filename len must be greater than 0")
		return
	}
	data := f.ReadDataFromFile(name)
	f.WriteDataToFile("c_"+name, data)
}

func (f FileManager) WriteDataToFile(name string, data []byte) {
	errWrite := ioutil.WriteFile(name, data, 0666)
	if errWrite != nil {
		log.Fatal(errWrite)
		return
	}
}

func (f FileManager) ReadDataFromFile(name string) []byte {
	data, errRead := ioutil.ReadFile(name)
	if errRead != nil {
		log.Fatal(errRead)
		return nil
	}
	return data
}
