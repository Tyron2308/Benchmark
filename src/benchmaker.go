package main

import (
   "fmt"
   topic "myproject/topic"
 //  testR  "myproject/testrunner"
   //dcode  "myproject/decodeurtest"
)
   //initt   "myproject/initbench"

func main() {

//    benchtest := testR.InitMapConfigTest()
   // cfg := dcode.GenerateBenchmark("/Users/tyron/Desktop/Work/sncf/Benchmark/testtorun.yml")

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
    //fmt.Println("len ===", len(cfg.Cfgs))
    //fmt.Println("len ===", len(cfg.Cfgs[0].Topic))
    //fmt.Println("is kAdmin null?", kAdmin)
    //fmt.Println("topic to create", cfg.Cfgs[0].Topic[0].Name)
    //kAdmin.CreateTopics(cfg.Cfgs[0])
    topics, _ := kAdmin.ListTopicCluster()

    for key, _ := range topics {
        fmt.Println("name topic", key)
    }
}
