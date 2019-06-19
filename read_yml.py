import yaml
import yamlordereddictloader
import sys
from collections import OrderedDict
from BenchClass import *
from BenchMarkSchenario import BenchMarkSchenario

if __name__ == "__main__":
    # data = OrderedDict([
        # ('key1', 'val1'),
        # ('key2', OrderedDict([('key21', 'val21'), ('key22', 'val22')])),
        # ('key3', ["test1", "test2"])
        # ])
    # yaml.dump(data, open('myfile.yml', 'w'), 
                    # Dumper=yamlordereddictloader.Dumper, default_flow_style=False)
    with open('myfile.yml') as f:
        benchProd = BenchProducer()
        yamlData = yaml.load(f, Loader=yamlordereddictloader.Loader)

        benchMaster = BenchMarkSchenario()

        benchMaster.initScenario(yamlData['configuration'])
