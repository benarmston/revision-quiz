package questions

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
    HistoryQuestions map[string][]string
    ScienceQuestions map[string][]string
    QuestionList []map[string][]string
}

var historyQuestions = map[string][]string {
    "When was the Paris Peace Conference?" : {"1919"},
    "Which central power was the Treaty of Versailles with?" : {"germany"},
    "Which central power was the Treaty of Saint Germain with?" : {"austria"},
    "Which central power was the Treaty of Neuilly with?" : {"bulgaria"},
    "Which central power was the Treaty of Sevres with?" : {"turkey"},
    "Which central power was the Treaty of Lausanne with?" : {"turkey"},
    "Who was the leader of Britain in the Paris peace conference?" : {"david lloyd george", "lloyd george"},
    "Who was the leader of France in the Paris peace conference?" : {"georges clemenceau", "clemenceau"},
    "Who was the leader of the USA in the Paris peace conference?" : {"woodrow wilson", "wilson"},
}

var scienceQuestions = map[string][]string {
    "What biological process removes carbon dioxide from the atmosphere?" : {"photosynthesis"},
    "How is energy lost along a food chain?" : {"respiration", "urine", "faeces", "not all biomass eaten", "not all biomass consumed"},
}

var questionList = []map[string][]string {scienceQuestions, historyQuestions}

func MakeQuestions() QuestionsTemplate {
    questions := QuestionsTemplate{historyQuestions, scienceQuestions, questionList}
    return questions
}
