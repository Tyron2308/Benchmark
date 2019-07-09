package decodeurtest

import (
    "log"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

//type configsTopic struct {
    //Ctopic []cfgTopic `yaml:topic`
//}

//type Topic struct {
        //Name string `yaml:name`
        //Partitions string `yaml:partitions`
        //Replications string `yaml:replication`
//}

type DecodeurTest struct {
    Type string `yaml:"type"`
//    ConfigTopic configsTopic `yaml:topic`
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


