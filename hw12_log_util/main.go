package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// 2025-01-09 12:00:00 [TRACE, DEBUG, INFO, WARN, ERROR, FATAL] [module=1...5] text
// статистика модуль уровень=(%от всех сообщений по модулю count/sum(count)),
// уровень=(% от всех записей c данным уровнем в файле [count/sum(level)])

type Statistic struct {
	level   string
	modules map[string]int
	all     int
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	var file string
	var level string
	var output string
	flag.StringVar(&file, "file", os.Getenv("LOG_ANALYZER_FILE"), "path log file")
	flag.StringVar(&level, "level", os.Getenv("LOG_ANALYZER_LEVEL"), "level for analysis")
	flag.StringVar(&output, "output", os.Getenv("LOG_ANALYZER_OUTPUT"), "path for output file")
	flag.Parse()
	data, err := ReadFile(file)
	if err != nil {
		log.Fatalf("read file: %s", err)
	}
	statistic := CalcStatistics(data, level)

	fmt.Println(file, level, output)
	fmt.Println(statistic)
}

func ReadFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read all in file: %w", err)
	}
	result := strings.Split(string(data), "\n")
	return result, nil
}

func CalcStatistics(data []string, level string) Statistic {
	modulesLevel := map[string]int{}
	modulesLevelSum := map[string]int{}
	for _, el := range data {
		arr := strings.Split(el, " ")
		modulesLevelSum[arr[3]]++
		if arr[2] == level {
			modulesLevel[arr[3]]++
		}
	}
	modules := map[string]int{}
	sumLevel := 0

	for module, item := range modulesLevel {
		modules[module] = 100 * item / modulesLevelSum[module]
		sumLevel += modulesLevel[module]
	}
	return Statistic{level: level, modules: modules, all: 100 * sumLevel / len(data)}
}

func WriteFile(data map[string]int) error {
	fmt.Println(data)
	return nil
}
