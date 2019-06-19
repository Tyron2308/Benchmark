from enums import Option
import subprocess

class BenchPerformer():
      def __init__(self):
          pass
      def _getValueFromDict(self, ymlData):
          pass
      def _runTestFromYml(self, ymlData):
          popen = subprocess.Popen(ymlData, stdout=subprocess.PIPE)
          popen.wait()
          output = popen.stdout.read()
          return output
      def __functorBenchmark(self, ArgPrgm):
          return [self.prgBinary] + ArgPrgm
      def createTestFromScenario(ymlData):
          arrayOpt = benchProd.getValueFromDict(yamlData['configuration'])
          arrayOpt_tmp = benchProd.functorBenchmark(arrayOpt)
          print(arrayOpt_tmp)
          benchProd.runTestFromYml(arrayOpt_tmp)

class BenchProducer(BenchPerformer):

      def __init__(self):
          self.prgBinary = '/bin/echo'
         # self.prgBinary = './usr/bin/kafka-producer-perf-test'
      def __getValueFromDict(self, ymlData):
            return [ymlData[Option.NBRECORD.value],
                   ymlData[Option.THROUGHPUT.value],
                   ymlData[Option.CONFIG.value],
                   ymlData[Option.RECORDSIZE.value],
                   ymlData[Option.PAYLOADFILE.value]]


class BenchConsumer(BenchPerformer):

      def __init__(self):
         self.prgBinary = './usr/bin/kafka-consumer-perf-test'

      def __getValueFromDict(self, ymlData):
            return  [ymlData[Option.TOPIC.value],
                    ymlData[Option.NBRECORD.value],
                    ymlData[Option.THROUGHPUT.value],
                    ymlData[Option.CONFIG.value],
                    ymlData[Option.RECORDSIZE.value],
                    ymlData[Option.PAYLOADFILE.value]]
