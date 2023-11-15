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
    isPlay := true

    for {
        if !isPlay {
            break
        }

        remainGuesses := 6
        wordHint := ""
        word, err := selectRandomWord()

        if err != nil {
            log.Fatalln(err)
        }

        for range []rune(word) {
            wordHint += "_"
        }

        welcomeMsg()

        for {
            gameStatus(wordHint, remainGuesses)

            guess := ""

            for {
                guess, err = getUserInput()
                if err != nil {
                    fmt.Println(err)
                } else {
                    break
                }
            }

            if contains, err := isUserGuessCorrect(&word, &wordHint, &guess); err != nil {
                fmt.Println(err)
            } else {
                if !contains {
                    remainGuesses -= 1
                }
            }

            if isGameOver(&word, &wordHint, remainGuesses) {
                playAgainOrExit(&isPlay)
                break
            }
        }
    }
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


func gameOverMsg(remain int, word *string) {
    fmt.Print(`
            ############### 
            |  Game Over  |
            ############### 
       ____________
       |     | 
       |     | 
       |   \_O_/    Remain : `  + strconv.Itoa(remain) + `
       |     |       Word  : `  + *word + `
       |    / \
       |_____________
       |             |
       | You Lost ): |
       |_____________|
    `)
}


func gameWinMsg(remain int, word *string) {
    fmt.Print(`
            ############### 
            |    Winner   |
            ############### 
       ____________
       |      
       |      
       |   \_O_/    Remain : `  + strconv.Itoa(remain) + `
       |     |       Word  : `  + *word + `
       |    / \
       |______________
       |              |
       | Thank You (: |
       |______________|
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


func isUserGuessCorrect(word *string, wordHint *string, guess *string) (bool, error) {
    if strings.Contains(*wordHint, *guess) {
        return true, errors.New("You Already Guess that value, guess another one")
    }

    newHintWordSlice := []rune(*wordHint) 
    contains := false
    for i, w := range []rune(*word) {
        if string(w) == *guess {
            contains = true
            newHintWordSlice[i] = w
        }
    }

    *wordHint = string(newHintWordSlice)

    return contains, nil
}


func isGameOver(word *string, wordHint *string, remainGuesses int) bool {
    if *word == *wordHint {
        gameWinMsg(remainGuesses, word)
        return true;
    } else if (remainGuesses == 0) {
        gameOverMsg(remainGuesses, word)
        return true;
    }
    return false;
}


func playAgainOrExit(isPlay *bool) {
    scan := bufio.NewScanner(os.Stdin)
    fmt.Print("\nDo you want to play again? ")
    scan.Scan()

    input := []rune(scan.Text())

    if len(input) == 0 {
        *isPlay = false
    } else if strings.ToUpper(string(input[0])) != "Y" {
        *isPlay = false
    }
}
