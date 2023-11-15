package main

import (
	"errors"
	"os"
	"testing"
)


func TestSelectRandomWord(t *testing.T) {
    word, err := selectRandomWord()

    switch {
    case word != "" && err != nil:
        t.Errorf("error and also return a word, word: %v, err: %v", word, err)
    case word == "" && err == nil:
        t.Errorf("no error also no word, word: %v, err: %v", word, err)
    case len(word) < 6:
        t.Errorf("word length is less than 6, word: %v, err: %v", word, err)
    }
}


func TestGetUserInput(t *testing.T) {
    testInputs := []string{"d", "s", "dilush", "\n", "", " ", "23", "2323"}

    for _, word := range testInputs {
        r, w, err := os.Pipe()

        if err != nil {
            t.Error(err)
        }

        if _, err := w.Write([]byte(word)); err != nil {
            t.Error(err)
        }
        w.Close()

        defer func (v *os.File) { os.Stdin = v } (os.Stdin)

        os.Stdin = r

        char, err := getUserInput()

        switch {
        case (char == "" && err != nil) || (char != "" && err == nil):
        default:
            t.Error(err)
        }
    }
}


func TestIsGameOver(t *testing.T) {
    testValues := []struct{
        word string
        wordHint string
        remain int
        want bool
    }{
        {"word", "word", 0, true},
        {"word", "word", 5, true},
        {"dilush", "hashen", 5, false},
        {"dilush", "hashen", 0, true},
    }

    for _, v := range testValues {
        if current := isGameOver(&v.word, &v.wordHint, v.remain); current != v.want {
            t.Errorf("word: %v, wordHint: %v, remain: %v, expect: %v, current: %v", v.word, v.wordHint, v.remain, v.want, current)
        }
    }
}

func TestPlayAgainOrExit(t *testing.T) {
    testInputs := []struct{
        input string
        want bool
    }{
        {"s", false},
        {"", false},
        {"2323", false},
        {"n", false},
        {"y", true},
        {"Y", true},
    }

    for _, testCase := range testInputs {
        var isPlay = true
        r, w, err := os.Pipe()

        if err != nil {
            t.Error(err)
        }

        if _, err := w.Write([]byte(testCase.input)); err != nil {
            t.Error(err)
        }
        w.Close()

        defer func (v *os.File) { os.Stdin = v } (os.Stdin)

        os.Stdin = r

        playAgainOrExit(&isPlay)

        switch {
        case testCase.want == isPlay:
        default:
            t.Errorf("input: %v, want: %v, value: %v", testCase.input, testCase.want, isPlay)
        }
    }
}


func TestIsUserGuessCorrect(t *testing.T) {
    testCases := []struct{
        word string
        wordHint string
        guess string
        want bool
        err error
    }{
        {"dilush", "______", "d", true, nil},
        {"dilush", "d_____", "d", true, errors.New("err")},
        {"dilush", "______", "g", false, nil},
        {"dilush", "_i____", "i", true, errors.New("err")},
        {"dilush", "_i____", "k", false, nil},
    }

    for _, v := range testCases {
        if b, err := isUserGuessCorrect(&v.word, &v.wordHint, &v.guess); 
            (b != v.want) || 
            (err == nil && v.err != nil) || 
            (err != nil && v.err == nil) {
            t.Errorf("word: %v, wordHint: %v, guess: %v, want_err: %v, want_bool: %v, current_err: %v, current_bool: %v", v.word, v.wordHint, v.guess, v.err, v.want, err, b)
        }
    }
}
