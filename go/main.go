package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func main() {
	start := time.Now()
	var name string

	if len(os.Args) > 1 {
		// Берём первый аргумент (индекс 1)
		name = strings.TrimSpace(os.Args[1])
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Введите ваше имя: ")
		input, _ := reader.ReadString('\n')
		name = strings.TrimSpace(input)
	}

	currentTime := time.Now().Format("02.01.2006 15:04:05")

	var colorCode string
	// Важно: кириллическая буква "А"
	if strings.HasPrefix(strings.ToUpper(name), "А") {
		colorCode = Red
	} else {
		colorCode = Green
	}

	fmt.Printf("%sПривет, %s! Текущее время: %s%s\n", colorCode, name, currentTime, Reset)

	duration := time.Since(start)
	fmt.Printf("Время выполнения: %v\n", duration)
}
