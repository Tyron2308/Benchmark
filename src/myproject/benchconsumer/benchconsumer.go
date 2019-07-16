package benchconsumer

import (
    "fmt"
    dcode "myproject/decodeurtest"
    bench "myproject/abenchtest"
    "os/exec"
    "strconv"
    "sync"
)

type BenchConsumer struct {
}

func (e BenchConsumer) Run(cfgTest dcode.DecodeurTest, wg *sync.WaitGroup) bool {
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
    output, err := exec.Command("/usr/bin/make", "-C", ".", "run-kafka-consumer-perf", result).Output()

    bench.DeferWrtiing(err, output, "BenchConsumer.log")
    defer wg.Done()
    return true

}
