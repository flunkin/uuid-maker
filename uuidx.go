package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\033[37mhow much uuids to make?\033[0m ")
	countStr, _ := reader.ReadString('\n')
	countStr = strings.TrimSpace(countStr)
	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		fmt.Println("\033[37mnot a real number\033[0m")
		return
	}

	fmt.Print("\033[37mput uuids in file? y/n \033[0m")
	saveStr, _ := reader.ReadString('\n')
	saveStr = strings.ToLower(strings.TrimSpace(saveStr))

	var file *os.File
	if saveStr == "y" || saveStr == "yes" {
		file, err = os.Create("uuids.txt")
		if err != nil {
			fmt.Println("\033[37mcouldnt make file for uuids :( \033[0m", err)
			return
		}
		defer file.Close()
	}

	colors := []string{
		"\033[34m", // blue
		"\033[31m", // red
		"\033[32m", // green
	}

	for i := 0; i < count; i++ {
		id := uuid.New().String()
		if file != nil {
			file.WriteString(id + "\n")
			fmt.Println(colors[i%len(colors)] + id + "\033[0m")
		} else {
			fmt.Println(colors[i%len(colors)] + id + "\033[0m")
		}
	}

	if file != nil {
		fmt.Println("\033[37muuid saved to file\033[0m")
	} else {
		fmt.Println("\033[37muuids are created\033[0m")
	}
}
