package testrunner

import (
    dcode "myproject/decodeurtest"
    benchC "myproject/benchconsumer"
    benchP "myproject/benchproducer"
    topic "myproject/topic"
)

type mapConfigTest struct {
    m map[string]interface{}
}

func (this *mapConfigTest) init() *mapConfigTest {
    this = new (mapConfigTest)
    this.m = make(map[string]interface{})
    this.AddNewTest("BenchConsumer", createRoutine("BenchConsumer"))
    this.AddNewTest("BenchProducer", createRoutine("BenchProducer"))
    return this
}

func InitMapConfigTest() *mapConfigTest {
    tmp := mapConfigTest{}
    return tmp.init()
}

func (this *mapConfigTest) AddNewTest(key string, value func(dcode.DecodeurTest) bool) {
    if _, ok := this.m[key]; !ok {
        this.m[key] = value
    }
}

func (this mapConfigTest) CallFunctorStored(cfg dcode.DecodeurTest) {
    switch cfg.Type {
        case "BenchConsumer": this.m["BenchConsumer"].(func(dcode.DecodeurTest) bool)(cfg)
        case "BenchProducer": this.m["BenchProducer"].(func(dcode.DecodeurTest) bool)(cfg)
    }
}

func createRoutine(str string) func(cfg dcode.DecodeurTest) bool {

    kAdmin := topic.KafkaClient{}
    switch str {
        case "BenchConsumer":
            return func(cfg dcode.DecodeurTest) bool {
                tmp := new (benchC.BenchConsumer)
                return tmp.RunConcreteTest(cfg)
            }
        case "BenchProducer":
            return func(cfg dcode.DecodeurTest) bool {

                   tmp := new (benchP.BenchProducer)
                   tmp.Admin = kAdmin.CreateClientKafka([]string {"localhost:9092"})

                   return tmp.RunConcreteTest(cfg)
            }
        default:
            return func(cfg dcode.DecodeurTest) bool {
                return false
            }
    }
}
