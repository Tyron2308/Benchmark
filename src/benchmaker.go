package main

import (
   "fmt"
   topic "myproject/topic"
 //  testR  "myproject/testrunner"
   dcode  "myproject/decodeurtest"
)
   //initt   "myproject/initbench"

func main() {

//    benchtest := testR.InitMapConfigTest()
    cfg := dcode.GenerateBenchmark("/Users/tyron/Desktop/Work/sncf/Benchmark/testtorun.yml")

    //for _, configBench := range cfg.Cfgs {
     //benchtest.CallFunctorStored(configBench)
    //}

    //t := topic.Topic{
        //TopicName: "test-123",
        //NumPartitions: 1,
        //ReplicationFactor: 1,
    //}

    kAdmin := topic.KafkaClient{}
    kAdmin.CreateClientKafka([]string{"localhost:9092"})

    kAdmin.CreateTopics(cfg.Cfgs[0])
    topics, _ := kAdmin.ListTopicCluster()
    for key, _ := range topics {
        fmt.Println("name topic", key)
    }
}
