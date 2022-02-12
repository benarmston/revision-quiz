package main

import (
    "fmt"
    "strings"
    "os"
    "time"
    "math/rand"
    "github.com/crus4d3/revision-quiz/go/questions"
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
            } else if args[i] == "--test" {
                test()
            } else {
                unkownArg(args[i])
            }
        }
    } else {
        rand.Seed(time.Now().UnixNano())
        runQuiz()
    }
}

func help (name string) {
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

func test () bool {
    fmt.Println("Running tests")
    test1 := test1()
    if test1 {
        fmt.Println("Test 1 passed")
    } else {
        fmt.Println("Test 1 failed")
    }
    fmt.Println("Tests finished")
    return test1
}

func test1 () bool {
    slice1 := []string{"a"}
    contains(slice1)
    return true
}

func unkownArg (arg string) {
    fmt.Printf("Unkown arg '%s'\n", arg)
}

func runQuiz () {
    var response string
    score := 0
    quitMSG := []string{"q", "quit"}
    fmt.Println("Please enter all answers without any trialling spaces and as concisely as possible.")
    fmt.Println("Type 'quit' to quit at any time.")
    keys := getRandKeys(questions.Questions)
    for _, key := range keys {
        question := key
        answers := questions.Questions[key]
        fmt.Println(question)
        fmt.Scanln(&response)
        if contains(strings.ToLower(response), quitMSG) {
            break
        }
        if contains(strings.ToLower(response), answers){
            fmt.Println("Correct!")
            score += 1
        } else {
            fmt.Println("Incorrect")
            fmt.Printf("The correct answers were %#v\n", answers)
        }
    }
    fmt.Printf("You got: %d answers correct.\n", score)
}

func contains (search string, slice []string) bool {
    for _, value := range slice {
        if value == search {
            return true
        }
    }
    return false
}

func getRandKeys (questions map[string][]string) []string {
    keys := make([]string, len(questions))
    i := 0
    for k := range questions {
        keys[i] = k
        i++
    }
    rand.Shuffle(len(keys), func(i, j int) {keys[i], keys[j] = keys[j], keys[i]})
    return keys
}
