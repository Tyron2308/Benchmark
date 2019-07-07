package benchproducer

import (
    "fmt"
    "myproject/decodeurtest"
    "os/exec"
)

type BenchProducer struct {
    channel chan interface{}
}

func (e BenchProducer) RunConcreteTest(cfgTest decodeurtest.DecodeurTest) bool {
    fmt.Println("======> BenchProducer test <=======")
    output, err := exec.Command("/bin/ls").Output()
    if err!=nil {
        fmt.Println(err.Error())
    }
    fmt.Println(string(output))
    return true
}


