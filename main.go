package main

import (
	"fmt"
	"strings"
)

const originalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func  hashLetterFn(key int, letter string)(result string){
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter) - key : len(letter)])
	leftOverLetters := string(runes[0 : len(letter)-key])
	return fmt.Sprintf(`%s%s`, lastLetterKey, leftOverLetters)
}

func encrypt(key int, plainText string)(result string){

	hashLetter := hashLetterFn(key, originalLetter)
	var hashedString = ""
	findOne := func(r rune) rune{
		
	}



	strings.map(findOne, plainText)
}

func decrypt(key int, encryptedText string)(result string){

}

func main(){
	plainText := "HELLOWORLD"
	fmt.Println("Plain text:", plainText)
	encrypted := encrypt(5, plainText) //encryption key, here is 5 could be any other number
	fmt.Println("Encrypted text:", encrypted)
	decrypted := decrypt(5, encrypted)// use same decryption key.
	fmt.Println("Decrypted text:", decrypted)
}