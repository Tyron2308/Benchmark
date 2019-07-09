package benchconsumer

import (
    "fmt"
    "myproject/decodeurtest"
    "os/exec"
)

type BenchConsumer struct {
    channel chan interface{}
}

func (e BenchConsumer) Run(cfgTest decodeurtest.DecodeurTest) bool {
    fmt.Println("=====> BenchConsumer test <======")
    output, err := exec.Command("/bin/ls").Output()
    if err!=nil {
        fmt.Println(err.Error())
    }
    fmt.Println(string(output))
    return true
}
