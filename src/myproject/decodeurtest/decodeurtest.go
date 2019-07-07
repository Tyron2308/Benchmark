package decodeurtest

import (
    "log"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type DecodeurTest struct {
    Type string `yaml:"type"`
    Topic []string
    NumRecord string  `yaml:num-record, string`
    Throughput string `yaml:throughput`
    ConfigPath string `yaml:config-path`
    RecordSize string `yaml:record-size`
    PayloadFile string `yaml:payload-file-path`
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


