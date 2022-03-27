package quiz

import (
    "bufio"
    "fmt"
    "os"
    "math/rand"
    "strings"
    "github.com/crus4d3/revision-quiz/go/pkg/utils"
    "github.com/crus4d3/revision-quiz/go/pkg/questions"
)

func runQuiz(questions []questions.Question) {
    //response := ""
    //score := 0
    //total := 0
    //correct := false
    //reader := bufio.NewReader(os.Stdin)
    //quitMSG := []string{"q", "quit"}

    //fmt.Println("Please enter all answers as concisely as possible.")
    //fmt.Println("Type 'quit' to quit at any time.")

    //rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })
    fmt.Println(questions)
}

func Quiz(questionList []map[string][]string, quiz bool) {
    var questions map[string][]string
    response := ""
    score := 0
    total := 0
    correct := false
    reader := bufio.NewReader(os.Stdin)
    quitMSG := []string{"q", "quit"}

    for _, question := range questionList {
        questions = utils.CombineMap(questions, question)
    }
    keys := utils.GetRandKeys(questions)

    //quizKeys := utils.GetRandKeys(quiz)

    fmt.Println("Please enter all answers as concisely as possible.")
    fmt.Println("Type 'quit' to quit at any time.")

    for _, question := range keys {
        answers := questions[question]
        fmt.Println(question)

        response, _ = reader.ReadString('\n')
        response = strings.TrimSuffix(response, "\n")
        response = strings.ReplaceAll(response, " ", "")
        response = strings.ReplaceAll(response, "\t", "")
        response = strings.ToLower(response)

        if utils.Contains(strings.ToLower(response), quitMSG) {
            break
        }
        for _, answer := range answers {
            answer = strings.ReplaceAll(answer, " ", "")
            answer = strings.ReplaceAll(answer, "\t", "")
            answer = strings.ToLower(answer)

            if response == answer {
                fmt.Println("Correct!")
                correct = true
                score += 1
            }
        }
        if correct == false {
            fmt.Println("Incorrect")
            fmt.Printf("The correct answers were %+q\n", answers)
        }
        total += 1
    }
    fmt.Printf("You got: %d out of %d questions correct.\n", score, total)
}
