# Wordle Guesser
This program is to guess against a daily wordle.

## Project Structure

```
wordle-guesser
|--cmd
|   |--main.go        # Entry point for the program
|--internal
|   |--constant
|   |   |--common.go  # Common constants used throughout the program
|   |--guesser
|       |--guesser.go # The guesser
|--go.mod             # Module definition
|--README.md          # This file
```

## Getting Started
To run the program, follow these steps:

1. Clone the repository:
   ```
   git clone https://github.com/doankimhuy-it/wordle-guesser
   cd wordle-guesser
   ```
2. Install neccessary dependencies:
   ```
   go mod tidy
   ```
3. Run the program:
   ```
   go run cmd/main.go
   ```
## Contributing
Feel free to submit issues or pull requests if you have any suggestions or improvements.