package main

import (
	"bufio"
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLogData(t *testing.T) {
	t.Parallel()

	var fileObj *os.File
	defer fileObj.Close()
	// Открываем файл
	fileObj, err := os.Open("test_files/log_file.log")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Файл не найден")
		}
	}
	assert.Equal(t, nil, err)
	assert.Equal(t, path.Ext(fileObj.Name()), ".log")

	want := "2024-03-17 15:15:00 INFO Тест Тест Тест"

	// Читаем все из файла
	scanner := bufio.NewScanner(fileObj)

	for scanner.Scan() {
		assert.Equal(t, want, scanner.Text())
	}
}
