package utils

import (
    "bufio"
    "fmt"
    "gopkg.on/yaml.v3"
    "io/ioutil"
    "math/rand"
    "os"
    "strings"
    "github.com/crus4d3/revision-quiz/go/pkg/questions"
)

func Version() {
    fmt.Println("Version 0.4.0")
    fmt.Println("This project is still in active development")
    fmt.Println("Visit https://github.com/crus4d3/revision-quiz for the latest version")
}

func Contains(needle string, haystack []string) bool {
    for _, value := range haystack {
        if value == needle {
            return true
        }
    }
    return false
}

func Combine(map1 map[string][]string, map2 map[string][]string) map[string][]string {
    for k, v := range map2 {
        map1[k] = v
    }
    return map1
}

func GetRandKeys(questions map[string][]string) []string {
    keys := make([]string, len(questions))
    i := 0
    for k := range questions {
        keys[i] = k
        i++
    }
    rand.Shuffle(len(keys), func(i, j int) {keys[i], keys[j] = keys[j], keys[i]})
    return keys
}

func GetSubjects(questions questions.QuestionsTemplate) []map[string][]string {
    finished := false
    reader := bufio.NewReader(os.Stdin)

    history := []string{"h", "history"}
    science := []string{"s", "science"}
    all := []string{"a", "all"}
    done := []string{"d", "done"}
    quitMSG := []string{"q", "quit"}

    var subject string
    var subjectList []map[string][]string

    for finished == false {
        fmt.Println("Which subjects do you want to revise 'history', 'science' or 'all'?")
        fmt.Println("Type 'done' to continue")
        subject, _ = reader.ReadString('\n')
        subject = strings.TrimSuffix(subject, "\n")

        if Contains(strings.ToLower(subject), history) {
            fmt.Println("History chosen!")
            subjectList = append(subjectList, questions.HistoryQuestions)
        } else if Contains(strings.ToLower(subject), science) {
            fmt.Println("Science chosen!")
            subjectList = append(subjectList, questions.ScienceQuestions)
        } else if Contains(strings.ToLower(subject), all) {
            fmt.Println("All chosen!\n")
            finished = true
        } else if Contains(strings.ToLower(subject), done) {
            if subjectList != nil {
                fmt.Println("Subjects chosen!\n")
                finished = true
            } else {
                fmt.Println("You must choose at least one subject")
            }
        } else if Contains(strings.ToLower(subject), quitMSG) {
            os.Exit(0)
        } else {
            fmt.Println("Unkown option")
        }
    }
    return subjectList
}

func ReadYaml(quiz *questions.Quiz) *questions.Quiz {
}
