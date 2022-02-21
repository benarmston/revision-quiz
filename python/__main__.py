#!/usr/bin/env python3

import argparse

from questions import Questions
from utils import version, getSubjects, getQuestions
from quiz import runQuiz

def parseArgs():
    parser = argparse.ArgumentParser(
        description="Revision quiz to test knowledge on various subjects.",
        add_help=False
    )
    parser.add_argument(
        "--help",
        action="help",
        help="show this help message and exit",
    )
    parser.add_argument(
        "--version",
        action="store_true",
        help="output version information and exit"
    )
    parser.add_argument(
        "--history",
        "-h",
        action="store_true",
        help="include history questions"
    )
    parser.add_argument(
        "--science",
        "-s",
        action="store_true",
        help="inlude science questions"
    )
    return parser.parse_args()

def main():
    args = parseArgs()
    quiz = getQuestions()
    questionsList = []
    questionDict = {}
    questions = Questions()
    if args.version:
        version()
        quit()
    if args.science:
        for subject in quiz.subjectList:
            if subject.name == "science":
                for question in subject.questions:
                    questionsList.append(question)
    if args.history:
        for subject in quiz.subjectList:
            if subject.name == "history":
                for question in subject.questions:
                    questionsList.append(question)
    elif questionsList == []:
        subjectList = getSubjects(quiz)
        if subjectList == []:
            for subject in quiz.subjectList:
                for question in subject.questions:
                    questionsList.append(question)
        else:
            for subject in subjectList:
                for question in subject.questions:
                    questionsList.append(question)
    runQuiz(questionsList)

if __name__ == "__main__":
    main()
