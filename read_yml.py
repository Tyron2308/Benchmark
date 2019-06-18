import yaml
import yamlordereddictloader
import sys
from enums import enum
from collections import OrderedDict
import subprocess

class BenchPerformer():
      def getValueFromDict(ymlData):
          pass

      def runTestFromYml(BenchConfig, ymlData):
          popen = subprocess.Popen(ymlData, stdout=subprocess.PIPE)
          popen.wait()
          output = popen.stdout.read()
          print('output subprocess:')
          print output
          print output
          print output
          print('< ==== output subprocess:')
          return


class BenchProducer(BenchPerformer):
      self.BenchCfg = enum('NBSMS','THROUGHPUT', 'CONFIG', 'RECORDSIZE', 'PAYLOAD-FILE')
      def getValueFromDict(ymlData):
          return (ymlData[BenchCfg.NBSMS],
                  ymlData[BenchCfg.THROUGHPUT],
                  ymlData[BenchCfg.CONFIG],
                  ymlData[BenchCfg.RECORDSIZE],
                  ymlData[BenchCfg.PAYLOADFILE])

class BenchConsumer(BenchPerformer):
      self.BenchCfg = enum('TOPIC', 'NBSMS',
                                 'THROUGHPUT', 'CONFIG', 'RECORDSIZE', 'PAYLOAD-FILE')
      def getValueFromDict(ymlData):
          return (ymlData[BenchCfg.TOPIC],
                  ymlData[BenchCfg.NBSMS],
                  ymlData[BenchCfg.THROUGHPUT],
                  ymlData[BenchCfg.CONFIG],
                  ymlData[BenchCfg.RECORDSIZE],
                  ymlData[BenchCfg.PAYLOADFILE])


if __name__ == "__main__":
    # data = OrderedDict([
        # ('key1', 'val1'),
        # ('key2', OrderedDict([('key21', 'val21'), ('key22', 'val22')])),
        # ('key3', ["test1", "test2"])
        # ])
    # yaml.dump(data, open('myfile.yml', 'w'), 
                    # Dumper=yamlordereddictloader.Dumper, default_flow_style=False)

    with open('myfile.yml') as f:
        benchProd = BenchProducer

        yamlData = yaml.load(f, Loader=yamlordereddictloader.Loader)
        benchProd.runTestFromYml(benchProd.getValueFromDict(yamlData), )
        decrypt_yml(yaml_data)

