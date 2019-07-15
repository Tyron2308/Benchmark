package decodeurtest

import (
    "log"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type DecodeurTest struct {
    Type string `yaml:"type"`
    Topic []struct {
            Name string `yaml:name`
            Partitions string `yaml:partitions`
            Replications string `yaml:replications`
    }
    NumRecord string  `yaml:numrecord`
    Throughput string `yaml:throughput`
    ConfigPath string `yaml:configpath`
    RecordSize string `yaml:recordsize`
    Payload string `yaml:payload`
    Message string `yaml:message`
    Zookeeper string `yaml:zookeeper`
    Threads string `yaml:threads`
    Outputfile string `yaml:outputFile`
}

type configs struct {
    Cfgs []DecodeurTest `yaml:"configuration"`
}

func (d *configs) readFromFile(filename string) *configs {
    source, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    err = yaml.Unmarshal(source, d)
    if err != nil {
         log.Fatalf("error: %v", err)
    }
    return d
}

func GenerateBenchmark(filename string) *configs {
    tmp := new (configs)
    tmp.readFromFile(filename)
    return tmp
}


