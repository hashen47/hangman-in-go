package main


import "testing"


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
