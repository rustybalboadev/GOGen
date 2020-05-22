package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Password Generator")
	var integers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var symbols = []string{"&", "$", "#", "@", "*", "!"}
	var password string
	x := 0
	rand.Seed(time.Now().UnixNano())
	fmt.Print("Length of Password: ")
	reader := bufio.NewReader(os.Stdin)
	amount, _ := reader.ReadString('\n')
	amount = strings.TrimSuffix(amount, "\r\n")
	toInt, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print()
	for x < toInt {
		choice := rand.Intn(3)
		if choice == 0 {
			index := rand.Intn(len(integers))
			password += strconv.FormatInt(int64(integers[index]), 10)
		} else if choice == 1 {
			index := rand.Intn(len(letters))
			password += letters[index]
		} else if choice == 2 {
			index := rand.Intn(len(symbols))
			password += symbols[index]
		}
		x++
	}
	fmt.Println("Your Password Is " + password)
}
