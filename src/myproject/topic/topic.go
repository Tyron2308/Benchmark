package topic

import (
    dcode "myproject/decodeurtest"
    s     "github.com/Shopify/sarama"
          "log"
          "fmt"
)

type Topic struct {
    TopicName string
    NumPartitions int32
    ReplicationFactor int16
}

type KafkaClient struct {
    brokers []string
    Admin s.ClusterAdmin
    topics []Topic
}

func (k *KafkaClient) CreateClientKafka(brokerAddrs []string) *KafkaClient {
    config := s.NewConfig()
    config.Version = s.V2_1_0_0
    k.brokers = brokerAddrs
    tmp, err := s.NewClusterAdmin(k.brokers, config)
    if err == nil {
        k.Admin = tmp
        return k
    }
    return nil
}

func (k KafkaClient) CreateTopics(cfgTest dcode.DecodeurTest) {
    topics, _ := k.ListTopicCluster()
    topiclist := make(map[string]bool, len(cfgTest.ConfigTopic.Ctopic))
    for k := range cfgTest.ConfigTopic.Ctopic {
            fmt.Println("ffff===>", k)
            keyName := cfgTest.ConfigTopic.Ctopic[k].Name
            if _, ok := topics[keyName]; ok {
                topiclist[keyName] = true
            } else {
                topiclist[keyName] = false
            }
    }
    //for key, v := range topiclist {
        //if v == false {
            //k.createTopic(Topic{
                    //TopicName: key,
                    //NumPartitions: topics[key].NumPartitions,
                    //ReplicationFactor: topics[key].ReplicationFactor,
                //},
            //)
        //}
    //}
}

func (k KafkaClient) ListTopicCluster() (map[string]s.TopicDetail, error) {
    return k.Admin.ListTopics()
}

func (k *KafkaClient) createTopic(c Topic) bool {
    defer func() { _ = k.Admin.Close() }()
    err := k.Admin.CreateTopic(c.TopicName, &s.TopicDetail{
        NumPartitions:     c.NumPartitions,
        ReplicationFactor: c.ReplicationFactor,
    }, false)
    if err != nil {
        log.Fatal("Error while creating topic: ", err.Error())
        return false
    }
    return true
}
