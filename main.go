package main

import (
	"bufio"
	"fmanager/manager"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	CREATE = 1
	DELETE = 2
	COPY   = 3
	RENAME = 4
	MKDIR  = 5
	MOVE   = 6
	EXIT   = 0
)

var options = map[int]string{
	CREATE: "Создать файл",
	DELETE: "Удалить",
	COPY:   "Копировать файл",
	RENAME: "Редактировать файл",
	MKDIR:  "Создать папку",
	MOVE:   "Перемещение", // TODO
	EXIT:   "Выход",
}

func main() {

	fm := manager.FileManager{}
	reader := bufio.NewReader(os.Stdin)

	printListOptions()

	for {
		input, _ := reader.ReadString('\n')
		res, _ := strconv.Atoi(strings.TrimSpace(input))

		if res == MKDIR {
			fmt.Println("Введите имя папки: ")
		} else if res == MOVE {
			fmt.Println("Введите путь к файлу: ")
		} else if res == DELETE {
			fmt.Println("Введите путь файла или папки")
		} else {
			fmt.Println("Введите имя файла: ")
		}

		inputFile, _ := reader.ReadString('\n')
		name := strings.TrimSpace(inputFile)

		switch res {
		case CREATE:
			fm.CreateFile(name)
			break
		case DELETE:
			fm.DeleteFile(name)
			break
		case COPY:
			fm.CopyFile(name)
			break
		case RENAME:
			fmt.Println("Введите новое имя файла: ")
			inputFile, _ := reader.ReadString('\n')
			newFileName := strings.TrimSpace(inputFile)
			fm.RenameFile(name, newFileName)
			break
		case MKDIR:
			fm.CreateFolder(name)
			break
		case MOVE:
			break
		case EXIT:
			os.Exit(1)
		}
	}
}

func printListOptions() {
	fmt.Println("Список функций: ")
	for index, option := range options {
		fmt.Printf("%d. %s\n", index, option)
	}
}
