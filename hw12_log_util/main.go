package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
)

// init is invoked before main().
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func getParams() (file, level, output string) {
	// считываем из .env
	logAnalyzerFile, _ := os.LookupEnv("LOG_ANALYZER_FILE")
	logAnalyzerLevel, _ := os.LookupEnv("LOG_ANALYZER_LEVEL")
	logAnalyzerOutput, _ := os.LookupEnv("LOG_ANALYZER_OUTPUT")

	// Когда указывают с именами: go run main.go -f=result.txt -l=INFO -o=hello.txt
	pflag.StringVarP(&file, "file", "f", "", "file path input")
	pflag.StringVarP(&level, "level", "l", "", "log level")
	pflag.StringVarP(&output, "output", "o", "", "file path output")

	// Когда указывают: go run main.go -f -l -o - подставить значения по умолчанию
	// Если просто укажут -o (без значения) - будет вывод в консоль
	pflag.Lookup("file").NoOptDefVal = "/"
	pflag.Lookup("level").NoOptDefVal = "INFO"
	pflag.Lookup("output").NoOptDefVal = "stdout"

	pflag.Parse()

	// Когда указывают: go run main.go - подставляем переменные окружения
	// -o будет равен значению из переменной окружения если его вообще никак не укажут
	if file == "" {
		file = logAnalyzerFile
	}
	if level == "" {
		level = logAnalyzerLevel
	}
	if output == "" {
		output = logAnalyzerOutput
	}
	return
}

func searchInLogFileByLevel(fileName, level string) (result string) {
	var fileObj *os.File
	var counterRowsByLevel int
	defer fileObj.Close()
	// Открываем файл
	fileObj, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Файл не найден")
		}
	}

	// Читаем все из файла
	scanner := bufio.NewScanner(fileObj)

	for scanner.Scan() {
		// Проверяем вхождение level в строку из лог файла - считаем в скольких строках вхождение
		if strings.Contains(scanner.Text(), level) {
			// fmt.Println(scanner.Text())
			counterRowsByLevel++
		}
	}

	return fmt.Sprintf("В лог-файле обнаружено %d строк с level=%s\n", counterRowsByLevel, level)
}

func writeResult(result, output string) {
	var fileOutput *os.File
	defer fileOutput.Close()
	// Создаем и открываем файл с полученным именем и печатаем в него
	fileOutput, err := os.Create(output)
	if err != nil {
		log.Println("Unable to create file:", err)
	} else {
		fileOutput.WriteString(result)
	}
}

func main() {
	file, level, output := getParams()

	result := searchInLogFileByLevel(file, level)

	if output == "stdout" {
		// печатаем в консоль (Просто указали -o (без значения))
		log.Println(result)
	} else {
		// печатаем в файл (указали -o=res.txt либо просто не указали -o)
		writeResult(result, output)
	}
}
