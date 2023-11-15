package main

import (
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
    welcomeMsg()
    gameStatus("H__l_", 6)
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


func gameOverMsg(word string) {
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
