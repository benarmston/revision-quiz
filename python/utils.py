#!/usr/bin/env python3

import yaml

def version():
    print("Version 0.9.0")
    print("This project is still in active development")
    print("Visit https://github.com/crus4d3/revision-quiz for the latest version")

def getSubjects(quiz):
    done = False
    chosenSubjects = []
    subjectNames = [s.name for s in quiz.subjectList]

    while done == False:
        print("Which subjects do you want to revise", *subjectNames, "or all?", sep=", ")
        inputSubject = input("Type 'done' to continue\n")
        inputSubject = inputSubject.strip()

        if inputSubject in ["q", "quit"]:
            quit()
        elif inputSubject in ["a", "all"]:
            print("All chosen!\n")
            done = True
        elif inputSubject in ["d", "done"]:
            if chosenSubjects != []:
                done = True
                print("Subjects chosen!\n")
            else:
                print("\nYou must choose at least one subject")
        else:
            found = False
            for subject in quiz.subjectList:
                if subject.name == inputSubject:
                    chosenSubjects.append(subject)
                    found = True
                    print()
                    print(inputSubject, "chosen!")
            if found == False:
                print("\nUnkown option:", inputSubject)

    return chosenSubjects

class Quiz():
    def __init__(self, subjectList):
        self.subjectList = self.createSubjects(subjectList)

    def createSubjects(self, subjectList):
        fmtSubjects = []
        for subject in subjectList:
            fmtSubject = Subject(subject)
            fmtSubjects.append(fmtSubject)
        return fmtSubjects

class Subject():
    def __init__(self, questions):
        self.name = questions["name"]
        questions.pop("name")
        self.questions = self.createQuestions(questions)

    def createQuestions(self, questionList):
        questionList = questionList["questions"]
        fmtQuestions = []
        for question in questionList:
            prompt = question["prompt"]
            answers = question["answers"]
            fmtQuestion = Question(prompt, answers)
            fmtQuestions.append(fmtQuestion)
        return fmtQuestions

class Question():
    def __init__(self, prompt, answers):
        self.prompt = prompt
        self.answers = answers

def getQuestions():
    with open("questions.yaml", "r") as questions:
        try:
            subjects = yaml.safe_load(questions)
            quiz = Quiz(subjects)
            return quiz
        except yaml.YAMLError as error:
            print(error)
