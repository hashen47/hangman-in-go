package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)


const wordListFileName string = "wordlist.txt"


func main() {
    word, err := selectRandomWord()
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(word)
}


func selectRandomWord() (string, error) {
    if _, err := os.Stat(wordListFileName); err != nil {
        if os.IsNotExist(err) {
            return "", fmt.Errorf("%v file is not exists", wordListFileName)
        }
        return "", err
    }

    data, err := os.ReadFile(wordListFileName)

    if err != nil {
        return "", err
    }

    words := strings.Split(strings.Trim(string(data), " \n"), " ")

    if len(words) == 1 {
        return "", fmt.Errorf("%v file doesn't content enough words", wordListFileName)
    }

    word := words[rand.Intn(len(words))]

    return string(word), nil
}

