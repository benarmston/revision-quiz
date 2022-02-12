#!/usr/bin/env python3

import random
import argparse

from questions import Questions

def parseArgs():
    parser = argparse.ArgumentParser(
        description='Revision quiz to test knowledge on various subjects.'
    )
    parser.add_argument('--version', action='store_true', help='output version information and exit')
    parser.add_argument('--history', action='store_true', help='include history questions')
    parser.add_argument('--science', '-s', action='store_true', help='inlude science questions')
    return parser.parse_args()

def version():
    print("Version {}".format('0.4.0'))
    print("This project is still in active development")
    print("Visit https://github.com/crus4d3/revision-quiz for the latest version")

def getSubjects(questions):
    done = False
    subjectList = []
    while done == False:
        print("Which subjects do you want to revise 'history', 'science' or 'all'?")
        subject = input("Type 'done' to continue\n")
        if subject in ['h', 'history']:
            print("History chosen!")
            subjectList.append(questions.historyQuestions)
        elif subject in ['s', 'science']:
            print("Science chosen")
            subjectList.append(questions.scienceQuestions)
        elif subject in ['a', 'all']:
            print("All chosen!")
            done = True
        elif subject in ['d', 'done']:
            if subjectList != []:
                done = True
                print("Subjects chosen!")
            else:
                print("You must chose at least one subject")
        elif subject in ['q', 'quit']:
            quit()
        else:
            print("Unkown option")
    return subjectList

def quiz(questions):
    response = ''
    score = 0
    total = 0
    print("\nPlease enter all answers without any trialling spaces and as concisely as possible.")
    print("Type 'quit' to quit at any time.")
    keys = list(questions.keys())
    random.shuffle(keys)
    for i in keys:
        question = i
        answers = questions[i]
        response = input(question)
        if response.lower() in ['q', 'quit']:
            break
        elif response.lower() in answers:
            print('Correct!')
            score += 1
        else:
            print('Incorrect')
            print('The correct answers were', answers)
        total += 1
    print('You got:', score, 'out of', total, 'questions correct.')

def main():
    args = parseArgs()
    questionList = []
    questionDict = {}
    questions = Questions()
    if args.version:
        version()
        quit()
    if args.history:
        questionList.append(questions.historyQuestions)
    if args.science:
        questionList.append(questions.scienceQuestions)
    elif questionList == []:
        subjectList = getSubjects(questions)
        if subjectList == []:
            for question in questions.questionList:
                questionList.append(question)
        else:
            for question in subjectList:
                questionList.append(question)
    for d in questionList:
        questionDict.update(d)
    quiz(questionDict)

if __name__ == '__main__':
    main()
