package questions

import (
    "fmt"
)

type Quiz struct {
    SubjectList []Subject
}

type Subject struct {
    Name string
    Questions []Question
}

type Question struct {
    Prompt string
    Answers []string
}

type QuestionsTemplate struct {
    HistoryQuestions []map[string][]string
    ScienceQuestions []map[string][]string
    QuestionList [][]map[string][]string
}

var historyQuestions = []map[string][]string {
    {"When was the Paris Peace Conference?" : {"1919"}},
    {"Which central power was the Treaty of Versailles with?" : {"germany"}},
    {"Which central power was the Treaty of Saint Germain with?" : {"austria"}},
    {"Which central power was the Treaty of Neuilly with?" : {"bulgaria"}},
    {"Which central power was the Treaty of Sevres with?" : {"turkey"}},
    {"Which central power was the Treaty of Lausanne with?" : {"turkey"}},
    {"Who was the leader of Britain in the Paris peace conference?" : {"david lloyd george", "lloyd george"}},
    {"Who was the leader of France in the Paris peace conference?" : {"georges clemenceau", "clemenceau"}},
    {"Who was the leader of the USA in the Paris peace conference?" : {"woodrow wilson", "wilson"}},
}

var scienceQuestions = []map[string][]string {
    {"What biological process removes carbon dioxide from the atmosphere?" : {"photosynthesis"}},
    {"How is energy lost along a food chain?" : {"respiration", "urine", "faeces", "not all biomass eaten", "not all biomass consumed"}},
}

var QuestionList = [][]map[string][]string {scienceQuestions, historyQuestions}

var maths = map[string]interface{} {
    "name": "maths",
    "questions": []map[string]interface{} {
        {"square root 9": []interface{} {3, -3}},
    },
}

var english = map[string]interface{} {
    "name": "english",
    "questions": []map[string]interface{} {
        {"how long on an AO2": []interface{} {"12", "12 mins", "12 minutes", 12}},
    },
}

var SubjectList = []map[string]interface{} {maths, english}

func MakeQuestions() QuestionsTemplate {
    questions := QuestionsTemplate{historyQuestions, scienceQuestions, QuestionList}
    return questions
}

func OtherMakeQuiz(subjects []map[string]interface{}) {
    var subjectList []Subject

    fmt.Println("\n\n\n", subjects)

    for _, subject := range subjects {
        name := subject["name"]
        rawQuestions := subject["questions"]

        fmt.Println("\n", subject)
        fmt.Println("\n", name)
        fmt.Println("\n", rawQuestions)

        var questionList []Question

        fmt.Println("\n\n\n\n", name, questionList)

        questions := rawQuestions.([]map[string]interface{})

        for _, question := range questions {
            fmt.Println("\n question:", question)
            for prompt, answers := range question {
                fmt.Println("\nprompt", prompt, "\nanswers", answers)
                answers = []string(answers)
                fmt.Println("\nprompt", prompt, "\nanswers", answers)
                //fmtQuestion := Question{prompt, answers}
                //questionList = append(questionList, fmtQuestion)
            }
        }
    }
    fmt.Println(subjectList)
    return
}

func MakeQuiz(subjects [][]map[string][]string) bool {
    fmt.Println(subjects)
    var subjectList []Subject

    for _, subject := range subjects {
        //name := subject["name"]
        var questionList []Question

        for _, questionMap := range subject {

            for prompt, answers := range questionMap {
                question := Question{prompt, answers}
                questionList = append(questionList, question)
            }
        }
        fmt.Println(questionList)
        newSubject := Subject{"", questionList}
        subjectList = append(subjectList, newSubject)
    }
    fmt.Println(subjectList)
    return true
}
