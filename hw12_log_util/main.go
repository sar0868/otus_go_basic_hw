package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/joho/godotenv"
)

// 2025-01-09 12:00:00 [TRACE, DEBUG, INFO, WARN, ERROR, FATAL] [module=1...] text
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
	flag.StringVar(&file, "file", "", "[required] path log file")
	flag.StringVar(&level, "level", os.Getenv("LOG_ANALYZER_LEVEL"), "[optional] level for analysis")
	flag.StringVar(&output, "output", os.Getenv("LOG_ANALYZER_OUTPUT"), "[optional] path for output file")

	flag.Parse()
	if file == "" {
		file = os.Getenv("LOG_ANALYZER_FILE")
		fmt.Println("flag -file required")
	}

	inChan := make(chan string)
	outChan := make(chan Statistic)
	go CalcStatistic(inChan, level, outChan)
	if err2 := ReadFileInChan(file, inChan); err2 != nil {
		log.Fatalf("read file: %s", err2)
	}
	result := <-outChan
	close(outChan)

	if errWF := WriteFile(result, level, output); errWF != nil {
		log.Fatalf("write file error: %s", errWF)
	}
}

func ReadFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read all in file: %w", err)
	}
	result := strings.Split(string(data), "\n")
	return result, nil
}

func ReadFileInChan(path string, inChan chan string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("don't open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inChan <- scanner.Text()
	}
	close(inChan)
	if err2 := scanner.Err(); err2 != nil {
		return fmt.Errorf("unable to read all in file: %w", err2)
	}
	return nil
}

func CalcStatistic(inChan chan string, level string, outChan chan Statistic) {
	modulesLevel := map[string]int{}
	modulesLevelSum := map[string]int{}
	totalCount := 0
	levels := []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	for el := range inChan {
		arr := strings.Split(el, " ")
		if len(arr) < 4 {
			continue
		}
		if !slices.Contains(levels, arr[2]) {
			continue
		}
		modulesLevelSum[arr[3]]++
		totalCount++
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
	outChan <- Statistic{level: level, modules: modules, all: 100 * sumLevel / totalCount}
}

func WriteFile(data Statistic, level string, output string) error {
	modules := ""
	for m, v := range data.modules {
		modules += fmt.Sprintf("\tpercentage level=%s in messages by module=%s: %d%%\n", level, m, v)
	}
	sep := "\t====================\n"
	all := fmt.Sprintf("\tpercentage level=%s in all messages: %d%%\n", level, data.all)
	text := fmt.Sprintf("level %s\n%s%s%s%s", level, modules, sep, all, sep)

	if output != "" {
		info := []byte(text)
		if err := os.WriteFile(output, info, 0o600); err != nil {
			return fmt.Errorf("unable to write file: %w", err)
		}
		return nil
	}
	fmt.Print(text)

	return nil
}
