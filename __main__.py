#!/usr/bin/env python3

import random
import argparse

from questions import questions

def parseArgs():
    parser = argparse.ArgumentParser(description='Process command line arguments')
    parser.add_argument('--version', action='store_true', help='output version information and exit')
    return parser.parse_args()

def loadGame(args):
    def version():
        print("Version {}".format('0.4.0'))
        print("This project is still in active development")
        print("Visit https://github.com/crus4d3/revision-quiz for the latest version")

    def quiz():
        response = ''
        score = 0
        print("Please enter all answers without any trialling spaces and as concisely as possible.")
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
        print('You got:', score, 'answers correct.')

    def newQuiz():
        response = ''
        print("Please enter all answers without any trialling spaces and as concisely as possible.")
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
            else:
                print('Incorrect')
                print('The correct answers were', answers)

    if args.version:
        version()
    else:
        quiz()

def main():
    args = parseArgs()
    loadGame(args)

if __name__ == '__main__':
    main()
