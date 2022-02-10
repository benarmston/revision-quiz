package main

import (
    "fmt"
    "os"
    "github.com/crus4d3/revision-quiz/go/quiz"
    //"github.com/crus4d3/revision/quiz"
    //"github.com/crus4d3/revision-quiz"
//    "programming/revision-quiz/questions.go"
)

func main() {

    progName := os.Args[0]
    args := os.Args[1:]

    if len(args) > 0 {
        for i := 0; i < len(args); i ++ {
            if args[i] == "--help" {
                help(progName)
            } else if args[i] == "--version" {
                version()
            } else {
                unkownArg(args[i])
            }
        }
    } else {
        runQuiz()
    }
}

func help (name string) {
    // fmt.Println("Howdy cowboy!")
    // fmt.Println("Hope that helped!!")
    fmt.Printf("usage: %s [--help] [--version]\n\n", name)
    fmt.Println("Revision quiz to test knowledge on various subjects\n")
    fmt.Println("optional arguments:")
    fmt.Println("  --help     show this message and exit")
    fmt.Println("  --version  output version information and exit")
}

func version () {
    fmt.Println("Version 0.4.0")
    fmt.Println("This project is still in active development")
    fmt.Println("Visit https://github.com/crus4d3/revision-quiz for the latest version")
}

func unkownArg (arg string) {
    fmt.Printf("Unkown arg '%s'\n", arg)
}

func runQuiz () {
    fmt.Println(quiz.Questions)
    response := ""
    score := 0
    fmt.Println("Please enter all answers without any trialling spaces and as concisely as possible.")
    fmt.Println("Type 'quit' to quit at any time.")
    fmt.Printf("Response %s score %d\n", response, score)
}

