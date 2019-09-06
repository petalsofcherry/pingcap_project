package util

import (
	"bufio"
	"io"
	"os"
)

func AppendFileArrayItem(filename string, item string) error {
	var f *os.File
	var err error
	if checkFileIsExist(filename) { //如果文件存在
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	} else {
		f, err = os.Create(filename) //创建文件
	}
	defer f.Close()

	if err != nil {
		return err
	}

	item += "\n"
	_, err = io.WriteString(f, item)
	if err != nil {
		return err
	}

	return nil
}

func CheckItemExist(filename string, item string) (int, bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	count := 1
	for scanner.Scan() {
		count += 1
		if item == scanner.Text() {
			return count, true, nil
		}
	}

	return count, false, nil
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}