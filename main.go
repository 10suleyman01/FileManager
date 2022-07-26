package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/////////////////////////////////////
//  !    файловый менеджер !       //
// 1.      Создать файл            //
// 2.      Удалить файл            //
// 3.      Копировать файл         //
// 4. Редактировать название файла //
// 5.       Создать папку          //
// 6.       Перемещение            //
/////////////////////////////////////

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
	DELETE: "Удалить файл",
	COPY:   "Копировать файл",
	RENAME: "Редактировать файл",
	MKDIR:  "Создать папку", // TODO
	MOVE:   "Перемещение",   // TODO
	EXIT:   "Выход",
}

func main() {

	fm := FileManager{}
	reader := bufio.NewReader(os.Stdin)

	for {
		printListOptions()

		input, _ := reader.ReadString('\n')
		res, _ := strconv.Atoi(strings.TrimSpace(input))

		fmt.Println("Введите имя файла: ")
		inputFile, _ := reader.ReadString('\n')
		fileName := strings.TrimSpace(inputFile)

		switch res {
		case CREATE:
			fm.CreateFile(fileName)
			break
		case DELETE:
			fm.DeleteFile(fileName)
			break
		case COPY:
			fm.CopyFile(fileName)
			break
		case RENAME:
			fmt.Println("Введите новое имя файла: ")
			inputFile, _ := reader.ReadString('\n')
			newFileName := strings.TrimSpace(inputFile)
			fm.RenameFile(fileName, newFileName)
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
