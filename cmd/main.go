package main

import "github.com/doankimhuy-it/wordle-guesser/internal/guesser"

func main() {
	guesser := guesser.NewGuesser()
	guesser.Guess()
}