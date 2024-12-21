package guesser

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"strings"

	"github.com/doankimhuy-it/wordle-guesser/internal/constant"
)

type WordleResponse struct {
	Slot   int    `json:"slot"`
	Guess  string `json:"guess"`
	Result string `json:"result"`
}

type Guesser struct {
	// This map will keep track of the available words that are not used yet
	availableWords map[string]bool
}

func NewGuesser() *Guesser {
	availableWords := make(map[string]bool)
	for c := 'a'; c <= 'z'; c++ {
		availableWords[string(c)] = true
	}
	return &Guesser{availableWords: availableWords}
}

func (g *Guesser) Guess() {
	guess := "aeiou"
	for _, c := range guess {
		g.availableWords[string(c)] = false
	}

	// The main loop to guess the word, ends when the word is found or the program is terminated
	for {
		log.Printf("Guessing: %v\n", guess)

		res, err := http.Get("https://wordle.votee.dev:8000/daily?guess=" + guess)
		if err != nil {
			log.Fatalf("Failed to make request: %v", err)
		}
		defer res.Body.Close()

		var wordleResponse []WordleResponse
		if err := json.NewDecoder(res.Body).Decode(&wordleResponse); err != nil {
			log.Fatalf("Failed to decode response: %v", err)
		}

		// Init a new guess array, we will join this array to get the final guess later
		newGuess := make([]string, constant.DefaultWordLength)

		// Count the number of correct positions
		count := 0

		for _, response := range wordleResponse {
			if response.Result == "correct" {
				newGuess[response.Slot] = response.Guess
				count++
			}
		}

		if count == constant.DefaultWordLength {
			log.Printf("The word is %v", guess)
			return
		}

		// If we have not found the word yet, we will try to find the "present" characters
		for _, response := range wordleResponse {
			if response.Result == "present" {
				// Generate a new position for the character
				nPos := rand.IntN(constant.DefaultWordLength)
				// If the position is already taken, we will find a new position
				for nPos == response.Slot || newGuess[nPos] != "" {
					nPos = (nPos + 1) % constant.DefaultWordLength
				}

				newGuess[nPos] = response.Guess
			}
		}

		log.Printf("New guess: %v", newGuess)

		// Guess the missing characters
		for i, c := range newGuess {
			if c == "" {
				for k := range g.availableWords {
					if g.availableWords[k] {
						newGuess[i] = k
						g.availableWords[k] = false
						break
					}
				}
			}
		}

		guess = strings.Join(newGuess, "")

		if len(guess) != constant.DefaultWordLength {
			log.Fatalf("Could not guess that word")
		}
	}
}
