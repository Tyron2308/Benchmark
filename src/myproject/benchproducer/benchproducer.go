package benchproducer

import (
    dcode "myproject/decodeurtest"
    topic "myproject/topic"
    bench "myproject/abenchtest"
    "strconv"
    "fmt"
    "sync"
    "os/exec"
)

type BenchProducer struct {
    Admin *topic.KafkaClient
}

func (this BenchProducer) Run(cfgTest dcode.DecodeurTest, wg *sync.WaitGroup) bool {
    fmt.Println("======> BenchProducer test <=======")
    num_record, err_n := strconv.Atoi(cfgTest.NumRecord)
    throughput, err_t := strconv.Atoi(cfgTest.Throughput)
    if (err_n != nil || err_t != nil ) {
        return false
    }
    this.Admin.CreateTopics(cfgTest)
    command := " --topic %s --num-record %d --throughput %d --producer.config=%s --payload-file=%s"

//--producer-props acks=1 bootstrap.servers=kafka_1:9092 buffer.memory=67108864 batch.size=8196
    result := fmt.Sprintf(command, cfgTest.Topic[0].Name, num_record,
                          throughput, cfgTest.ConfigPath, cfgTest.Payload)
    result = fmt.Sprintf("ARGS=%s",result)
    output, err := exec.Command("/usr/bin/make", "-C", ".", "run-kafka-producer-perf", result).Output()
    bench.DeferWrtiing(err, output, "Benchproducer")
    defer wg.Done()
    return true
}


