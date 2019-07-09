package testrunner

import (
    dcode "myproject/decodeurtest"
    benchC "myproject/benchconsumer"
    benchP "myproject/benchproducer"
    topic "myproject/topic"
    ini     "myproject/initbench"
    "os/exec"
    "fmt"
    "strconv"
    "net"
    "time"
)

type mapConfigTest struct {
    m map[string]interface{}
}

func (this *mapConfigTest) init() *mapConfigTest {
    this = new (mapConfigTest)
    this.m = make(map[string]interface{})
    this.AddNewTest("BenchConsumer", createRoutine("", "BenchConsumer"))
    this.AddNewTest("BenchProducer", createRoutine("", "BenchProducer"))
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

func (this mapConfigTest) MapBenchtestFunctor(cfg dcode.DecodeurTest) {
    switch cfg.Type {
        case "BenchConsumer": this.m["BenchConsumer"].(func(dcode.DecodeurTest) bool)(cfg)
        case "BenchProducer": this.m["BenchProducer"].(func(dcode.DecodeurTest) bool)(cfg)
    }
}

func createRoutine(broker string, str string) func(cfg dcode.DecodeurTest) bool {

    kAdmin := topic.KafkaClient{}
    switch str {
        case "BenchConsumer":
            return func(cfg dcode.DecodeurTest) bool {
                tmp := new (benchC.BenchConsumer)
                return tmp.Run(cfg)
            }
        case "BenchProducer":
            return func(cfg dcode.DecodeurTest) bool {
                   tmp := new (benchP.BenchProducer)
                   tmp.Admin = kAdmin.CreateClientKafka([]string {"localhost:9092"})
                   return tmp.Run(cfg)
            }
        default:
            return func(cfg dcode.DecodeurTest) bool {
                return false
            }
    }
}

func DestroyCluster() bool {
        ini.CreateContainer(".", "destroy-archi")
        return true
}

func InitCluster() bool {
    output, err := exec.Command("/bin/sh", "script/find_cluster_alive.sh").Output()
    if err!=nil {
        fmt.Println("err ==> alive", err.Error())
        return false
    }
    x, err := strconv.Atoi(string(output[:len(output) - 1]))
    if err != nil {
        return false
    }
    fmt.Println("value of x", x)
    if x <= 1  {
        fmt.Println(string(output))
        ini.CreateContainer(".", "build-kafka")
    }
    time.Sleep(9 * time.Second)
    waitOnCluster("localhost:9092")
    return true
}

func  waitOnCluster(host string) bool{
    i := 0
    for {
        conn, erro := net.Dial("tcp", host)
        i++
        fmt.Println("wait for cluster to be ready....", conn, erro)
        if (i % 10000 == 0){
        fmt.Println("wait for cluster to be ready....", conn, erro)
        output, err := exec.Command("docker", "ps", "-a").Output()
            if err!=nil {
                fmt.Println("err ==> alive", err.Error())
                return false
            }
        fmt.Println("output docker ps -a", string(output))
        }
        if conn != nil {
            conn.Close()
            fmt.Println("broker okay")
            break
        }
    }
    return true
}


