package main

import (
	"fmt"
	"strconv"
	"os"
	"io"
	"bufio"
	"errors"
)

func main() {
	endOp := false
	for endOp == false{
		fmt.Printf("Please input the operation ( 1 is read, 2 is write, 3 to exit )\n")
		var input string
		fmt.Scanln(&input)
		opType, err := strconv.Atoi(input)

		for{
			if err == nil {
				break
			} else if opType == 1 || opType == 2 {
				break
			}
			fmt.Printf("The input is wrong, please re-input\n")
			_, err = fmt.Scanf("%d", &opType)
		}

		switch (opType) {
		case 1:
			f, err := os.Open("file.txt")
			defer f.Close()
			reader := bufio.NewReader(f)
			if errors.Is(err, os.ErrNotExist) {
				fmt.Println("File doesn't exist yet, please write one ( 2 to write )")
				break
			}
			buf := make([]byte, 2 * 1024)
			for {
				n, err := reader.Read(buf)
				if err != nil && !errors.Is(err, io.EOF) {
					fmt.Println("an error has occured")
					panic(err)	
				}
				if n == 0 {
					break
				}
				fmt.Println(string(buf[:n]))
			}
		case 2:
			f, err := os.OpenFile("file.txt", os.O_APPEND | os.O_RDWR | os.O_CREATE, 0666)
			defer f.Close()
			writer := bufio.NewWriter(f)
			fmt.Println("What do you want to write to the file?")
			var newInput string
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				newInput = scanner.Text()
			}
			_, err = writer.WriteString(newInput + "\n")
			if err != nil {
				panic(err)
			}
			writer.Flush()
		case 3:
			fmt.Println("See you again later")
			endOp = true
		}
	}
}