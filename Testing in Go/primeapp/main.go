package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//n := 7

	//_, msg := isPrime(n)
	//fmt.Println(msg)

	intro()

	doneChan := make(chan bool)

	go readUserInput(os.Stdin, doneChan)

	<-doneChan

	close(doneChan)


	fmt.Println("Bye")
}

func readUserInput(in io.Reader, doneChan chan bool){
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumber(scanner)

		if done {
			doneChan <- true
			return
		}
	
		fmt.Println(res)
		prompt()
	}
}

func checkNumber(scanner *bufio.Scanner) (string, bool){

	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numberToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please Enter a number!", false
	}

	_, msg := isPrime(numberToCheck)

	return msg, false

}

func intro(){
	fmt.Println("Welcome to prime app!")
	fmt.Println("Enter a whole number")
	prompt()
}

func prompt(){
	fmt.Print("-> ")
}
func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}