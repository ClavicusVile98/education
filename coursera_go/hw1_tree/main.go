package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func PrintTree(output io.Writer, path string, printFiles bool, prefix string) error {
	/* переменная для записи содержимого директорий
	 * с информацией о них */
	var dir []os.FileInfo

	/* чтение директории 
	 * если не удалось, то выдается ошибка */
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	// показать файлы директории
	for _, file := range files {
		if !printFiles && !file.IsDir() {
			continue
		}
		/* добавление файлов в массив */
		dir = append(dir, file)
	}

	/* кол-во вхождений зависит от
	 * кол-ва файлов */
	for i, file := range dir {
		fileName := file.Name()
		if i == len(dir) - 1 {
			_, err = fmt.Println(prefix + "|___" + fileName)
		} else {
			_, err = fmt.Println(prefix + "|---" + fileName)
		}
		if err != nil {
			return err
		}

		curPrefix := prefix
		if file.IsDir() {
			if i != len(dir) - 1 {
				curPrefix = curPrefix + "|\t"
			} else {
				curPrefix = curPrefix + "\t"
			}
			err = PrintTree(output, path + string(os.PathSeparator) + file.Name(), printFiles, curPrefix)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func dirTree(output io.Writer, path string, printFiles bool) error {
	return PrintTree(output, path, printFiles,"")
}

func main() {
	output := os.Stdout
	/* сколько считалось аргументов */
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("Please enter 'go run main.go . [-f] to start the program'")
	}
	/* первый аргумент */
	path := os.Args[1]
	// fmt.Println(os.Args)
	//path := "D:\\Postgres"
	if os.Args[2] != "-f"{
		panic("Wrong parameter")
	}else{
		printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
		/* если в директорию зайти не удалось,
		 * то выдается ошибка */
		err := dirTree(output, path, printFiles)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
