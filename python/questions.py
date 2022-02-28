class Questions:
    def __init__(self):
        self.historyQuestions = {
            "name": "history",
            "questions": [
                {"prompt": "When was the Paris Peace Conference?", "answers": ["1919"]},
                {"prompt": "Which major power was not invited to the Paris Peace Conference?", "answers": ["germany"]},
                {"prompt": "Which central power was Treaty of Versailles with?", "answers": ["germany"]},
                {"prompt": "Which central power was Treaty of Saint Germain with?", "answers": ["austria"]},
                {"prompt": "Which central power was Treaty of Neuilly with?", "answers": ["bulgaria"]},
                {"prompt": "Which central power was Treaty of Sevres with?", "answers": ["turkey"]},
                {"prompt": "Which central power was Treaty of Lausanne with?", "answers": ["turkey"]},
                {"prompt": "Who was the leader of Britain in the Paris peace conference?", "answers": ["david lloyd george", "lloyd george"]},
                {"prompt": "Who was the leader of France in the Paris peace conference?", "answers": ["georges clemenceau", "clemenceau"]},
                {"prompt": "Who was the leader of the USA in the Paris peace conference?", "answers": ["woodrow wilson", "wilson"]},
            ]
        }
        self.scienceQuestions = {
            "name": "science",
            "questions": [
                {"prompt": "What biological process removes carbon dioxide from the atmosphere?", "answers": ["photosynthesis"]},
                {"prompt": "How is energy lost along a food chain?", "answers": ["respiration", "urine", "faeces", "not all biomass eaten"]},
            ]
        }
        self.subjectList = [self.historyQuestions, self.scienceQuestions]
