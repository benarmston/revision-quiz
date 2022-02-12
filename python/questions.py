class Questions:
    def __init__(self):
        self.historyQuestions = {
            "When was the Paris Peace Conference?\n" : ["1919"],
            "Which major power was not invited to the Paris Peace Conference?\n" : ["germany"],
            "Who was the Treaty of Versailles with?\n" : ["germany"],
            "Who was the Treaty of Saint Germain with?\n" : ["austria"],
            "Who was the Treaty of Neuilly with?\n" : ["bulgaria"],
            "Who was the Treaty of Sevres with?\n" : ["turkey"],
            "Who was the Treaty of Lausanne with?\n" : ["turkey"],
            "Who was the leader of Britain in the Paris peace conference?\n" : ["david lloyd george", "lloyd george"],
            "Who was the leader of France in the Paris peace conference?\n" : ["georges clemenceau", "clemenceau"],
            "Who was the leader of the USA in the Paris peace conference?\n" : ["woodrow wilson", "wilson"],
        }
        self.scienceQuestions = {
            "What process removes carbon dioxide from the atmosphere?\n" : ["photosynthesis"],
            "How is energy lost along a food chain?\n" : ["respiration", "urine", "faeces", "not all biomass eaten"],
        }
        self.questionList = [self.historyQuestions, self.scienceQuestions]
