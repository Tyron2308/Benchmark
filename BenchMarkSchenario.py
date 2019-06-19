from BenchClass import *

class BenchMarkSchenario(object):
      def __init__(self):
          self.benchClass = {
                  "BenchProducer": (lambda : BenchProducer()),
                  "BenchConsumer": (lambda : BenchConsumer())
          }

      def initScenario(self, ylData):
        for ymlData in ylData:
           for config in ymlData:
                benchObj = self.benchClass[ymlData[config]['type']]()
                print(benchObj.prgBinary)
