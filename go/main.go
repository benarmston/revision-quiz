package main

import (
	"flag"
	"math/rand"
	"os"
	"time"
	"github.com/crus4d3/revision-quiz/go/cmd"
	"github.com/crus4d3/revision-quiz/go/pkg/questions"
	"github.com/crus4d3/revision-quiz/go/pkg/tests"
	"github.com/crus4d3/revision-quiz/go/pkg/utils"
)

func main() {
	var questionsList []map[string][]string
	var totalQuestions map[string][]string
	questions := questions.MakeQuestions()

	version := flag.Bool("version", false, "output version information and exit")
	test := flag.Bool("test", false, "run tests")
	history := flag.Bool("history", false, "include history questions")
	historyS := flag.Bool("h", false, "include history questions")
	science := flag.Bool("science", false, "include science questions")
	scienceS := flag.Bool("s", false, "include science questions")
	flag.Parse()

	if len(os.Args) > 1 {
		if *version {
			utils.Version()
			return
		}
		if *test {
			tests.Test()
			return
		}
		if *science || *scienceS {
			questionsList = append(questionsList, questions.ScienceQuestions)
		}
		if *history || *historyS {
			questionsList = append(questionsList, questions.HistoryQuestions)
		}
	} else if questionsList == nil {
		subjectList := utils.GetSubjects(questions)
		if subjectList == nil {
			for question := range questions.QuestionList {
				questionsList = append(questionsList, questions.QuestionList[question])
			}
		} else {
			for question := range subjectList {
				questionsList = append(questionsList, subjectList[question])
			}
		}
	}
	for m := range questionsList {
		totalQuestions = utils.Combine(questionsList[m], totalQuestions)
	}
	rand.Seed(time.Now().UnixNano())
	quiz.Quiz(totalQuestions)
}
