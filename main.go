package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)


const wordListFileName string = "wordlist.txt"


func main() {
    word, err := selectRandomWord()

    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(word)

    guess, err := getUserInput()

    if err != nil {
        fmt.Printf("err: %v\n", err)
    }

    fmt.Println(guess)
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


func welcomeMsg() {
    fmt.Print(`
            ############### 
            |   Hangman   |
            ############### 
    `)
}


func gameOverMsg() {
    fmt.Print(`
            ############### 
            |  Game Over  |
            ############### 
    `)
}


func gameStatus(word string, remain int) {
    switch remain {
    case 0:
        fmt.Print(`
       ____________
       |     | 
       |     | 
       |   \_O_/    Remain : `  + strconv.Itoa(remain) + `
       |     |       Word  : `  + word + `
       |    / \
       |___________
       |           |
       |  Help Me! |
       |___________|
        `)
    case 1:
        fmt.Print(`
       ____________
       |      
       |      
       |   \_O_/    Remain : `  + strconv.Itoa(remain) + `
       |     |       Word  : `  + word + `
       |    / \
       |___________
       |           |
       |  Help Me! |
       |___________|
        `)
    case 2:
        fmt.Print(`
       ____________
       |      
       |      
       |   \_O_/    Remain : `  + strconv.Itoa(remain) + `
       |     |       Word  : `  + word + `
       |    / 
       |___________
       |           |
       |  Help Me! |
       |___________|
        `)
    case 3:
        fmt.Print(`
       ____________
       |      
       |      
       |   \_O_/    Remain : `  + strconv.Itoa(remain) + `
       |     |       Word  : `  + word + `
       |     
       |___________
       |           |
       |  Help Me! |
       |___________|
        `)
    case 4:
        fmt.Print(`
       ____________
       |      
       |      
       |     O_/    Remain : `  + strconv.Itoa(remain) + `
       |     |       Word  : `  + word + `
       |     
       |___________
       |           |
       |  Help Me! |
       |___________|
        `)
    case 5:
        fmt.Print(`
       ____________
       |      
       |      
       |     O_/    Remain : `  + strconv.Itoa(remain) + `
       |     |       Word  : `  + word + `
       |     
       |___________
       |           |
       |  Help Me! |
       |___________|
        `)
    case 6:
        fmt.Print(`
       ____________
       |      
       |      
       |     O      Remain : `  + strconv.Itoa(remain) + `
       |     |       Word  : `  + word + `
       |     
       |___________
       |           |
       |  Help Me! |
       |___________|
        `)
    default:
        log.Fatal("invalid guess remain value")
    }
}


func getUserInput() (string, error) {
    scan := bufio.NewScanner(os.Stdin)
    fmt.Print("\nGuess: ")
    scan.Scan()

    charArr := []rune(scan.Text())

    if len(charArr) == 0 {
        return "", errors.New("User input required")
    }

    input := strings.TrimSpace(string(charArr[0]))

    if input == "" {
        return "", errors.New("User input is empty") 
    }

    return input, nil
}
