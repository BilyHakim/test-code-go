package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Fungsi untuk mengecek bilangan prima
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Fungsi untuk mengecek 3 digit terakhir adalah bilangan prima yang sama
func lastThreeDigitSamePrime(num int) bool {
	lastDigit := num % 10
	secondLastDigit := (num / 10) % 10
	thirdLastDigit := (num / 100) % 10

	return isPrime(lastDigit) && isPrime(secondLastDigit) && isPrime(thirdLastDigit) &&
		lastDigit == secondLastDigit && secondLastDigit == thirdLastDigit
}

// FUngsi untuk mengecek 2 digit terakhir adalah 2 bilangan prima berurutan
func lastTwoDigitConsecutivePrime(num int) bool {
	lastDigit := num % 10
	secondLastDigit := (num / 10) % 10

	return isPrime(lastDigit) && isPrime(secondLastDigit) && (lastDigit-secondLastDigit == 1 || secondLastDigit-lastDigit == 1)
}

// Fungsi percabangan untuk menentukan posisi container
func determineContainerPosition(containerID string) string {
	num, err := strconv.Atoi(containerID)
	if err != nil {
		return "DEAD"
	}

	// Posisi Center
	if isPrime(num) && !strings.Contains(containerID, "0") && isPrime(num%1000) {
		return "CENTER"
	}

	// Posisi LEFT
	if isPrime(num) && !strings.Contains(containerID, "0") && lastTwoDigitConsecutivePrime(num%100) && (num%100 == (num/1000)*100) {
		return "LEFT"
	}

	// Posisi RIGHT
	if isPrime(num) && !strings.Contains(containerID, "0") && lastThreeDigitSamePrime(num%1000) {
		return "RIGHT"
	}

	// Posisi Reject
	return "DEAD"
}

func main() {

	// Sampel Input
	input := []string{
		"3137",
		"1367",
		"2333",
		"2001",
	}

	// Sampel Output
	fmt.Println("Sampel Inputan")
	for _, containerID := range input {
		fmt.Println(containerID)
	}

	fmt.Println("\nSampel Outputan")
	for _, containerID := range input {
		position := determineContainerPosition(containerID)
		fmt.Println(position)
	}
}
