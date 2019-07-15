package benchconsumer

import (
    "fmt"
    "myproject/decodeurtest"
    "os/exec"
    "strconv"
    "sync"
)

type BenchConsumer struct {
    channel chan interface{}
}

func (e BenchConsumer) Run(cfgTest decodeurtest.DecodeurTest, channel chan string, wg *sync.WaitGroup) bool {
    fmt.Println("------> BenchConsumer <-------")
    fmt.Println("struct ===> ", cfgTest)
    nb_mss, err_n := strconv.Atoi(cfgTest.Message)
    thrs, err_r := strconv.Atoi(cfgTest.Threads)
    if (err_n != nil || err_r != nil) {
        return false
    }
    command := " --broker-list %s --messages %d --topic %s --threads %d"

    result := fmt.Sprintf(command, cfgTest.Zookeeper, nb_mss, cfgTest.Topic[0].Name, thrs)
    fmt.Println("RESUKT", result)
    result = fmt.Sprintf("ARGS=%s",result)

    fmt.Println("execute kakfa-perf-consumer")
    output, _ := exec.Command("/usr/bin/make", "-C", ".", "run-kafka-consumer-perf", result).Output()
    channel <- string(output)

    wg.Done()
    return true
}
