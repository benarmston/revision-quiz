package quiz

import (
    "os"
    "fmt"
    "strings"
    "bufio"
    "github.com/crus4d3/revision-quiz/go/pkg/utils"
)

func Quiz(questions map[string][]string) {
    response := ""
    score := 0
    total := 0
    reader := bufio.NewReader(os.Stdin)
    quitMSG := []string{"q", "quit"}
    fmt.Println("Please enter all answers without any trialling spaces and as concisely as possible.")
    fmt.Println("Type 'quit' to quit at any time.")
    keys := utils.GetRandKeys(questions)
    for _, key := range keys {
        question := key
        answers := questions[key]
        fmt.Println(question)
        response, _ = reader.ReadString('\n')
        response = strings.TrimSuffix(response, "\n")
        if utils.Contains(strings.ToLower(response), quitMSG) {
            break
        }
        if utils.Contains(strings.ToLower(response), answers){
            fmt.Println("Correct!")
            score += 1
        } else {
            fmt.Println("Incorrect")
            fmt.Printf("The correct answers were %+q\n", answers)
        }
        total += 1
    }
    fmt.Printf("You got: %d out of %d questions correct.\n", score, total)
}
