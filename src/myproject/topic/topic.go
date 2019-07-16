package topic

import (
    dcode "myproject/decodeurtest"
    s     "github.com/Shopify/sarama"
          "log"
          "strconv"
)

type Request struct {
    flag bool
    Partitions int32
    Replications int16
}

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

    //k.waitOnCluster(brokerAddrs[0])
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

func (this KafkaClient) CreateTopics(cfgTest dcode.DecodeurTest) {
    topics, _ := this.ListTopicCluster()
    topiclist := make(map[string]Request, len(cfgTest.Topic))
    for k := range cfgTest.Topic {
            keyName := cfgTest.Topic[k].Name
            p, _ := strconv.Atoi(cfgTest.Topic[k].Partitions)
            r, _ := strconv.Atoi(cfgTest.Topic[k].Replications)
            if _, ok := topics[keyName]; ok {
               topiclist[keyName] = Request{
                    flag: true,
                    Partitions:  int32(p),
                    Replications: int16(r),
                }
           } else {
                topiclist[keyName] = Request{
                    flag: false,
                    Partitions:  int32(p),
                    Replications: int16(r),
                }
            }
    }
    for key, structValue := range topiclist {
        if structValue.flag == false {
           this.createTopic(Topic{ key, structValue.Partitions, structValue.Replications, },)
        }
    }
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
