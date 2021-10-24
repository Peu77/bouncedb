package file

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func mkdir(path string) {
	_ = os.MkdirAll(path, os.ModePerm)
}

func InitFiles() {
	mkdir("files/databases")
}

func WriteInFile(path string, content string) {
	file, err := os.Create(path)
	check(err)

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	check(err)
	err = writer.Flush()
	check(err)
	err = file.Close()
	check(err)
}

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	check(err)
	return string(data)
}

func ExistFile(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}
