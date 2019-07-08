package benchproducer

import (
    "fmt"
    dcode "myproject/decodeurtest"
    "os/exec"
    "strconv"
    topic "myproject/topic"
)

type BenchProducer struct {
    channel chan interface{}
    Admin *topic.KafkaClient
}

func (this BenchProducer) RunConcreteTest(cfgTest dcode.DecodeurTest) bool {
    fmt.Println("======> BenchProducer test <=======")
    num_record, err_n := strconv.Atoi(cfgTest.NumRecord)
    throughput, err_t := strconv.Atoi(cfgTest.Throughput)
    recordsize, err_r := strconv.Atoi(cfgTest.RecordSize)
	if (err_n != nil || err_t != nil || err_r != nil) {
        return false
    }
    this.Admin.CreateTopics(cfgTest)
    command := "docker exec -it %s kafka-producer-perf-test --topic %s --num-records %d --throughput %d --producer.config=%s --record-size %d --payload-file %s  --producer-props acks=1 bootstrap.servers=kafka_1:9092 buffer.memory=67108864 batch.size=8196 "

    result := fmt.Sprintf(command, "kafka_1" , cfgTest.ConfigTopic.Ctopic[0].Name, num_record,
                          throughput, cfgTest.ConfigPath, recordsize, cfgTest.Payload)
    output, err := exec.Command(result).Output()
    if err!=nil {
        fmt.Println(err.Error())
    }
    fmt.Println(string(output))
    return true
}


