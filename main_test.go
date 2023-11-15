package main

import (
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
