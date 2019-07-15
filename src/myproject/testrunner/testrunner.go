package testrunner

import (
    dcode "myproject/decodeurtest"
    benchC "myproject/benchconsumer"
    benchP "myproject/benchproducer"
    topic "myproject/topic"
    ini     "myproject/initbench"
    "os"
    "os/exec"
    "fmt"
    "strconv"
    "net"
    "time"
    "sync"
)


const (
    host = "localhost:9092"
)

type mapConfigTest struct {
    m map[string]interface{}
}

func (this *mapConfigTest) init(channel chan string, wg *sync.WaitGroup) *mapConfigTest {
    this = new (mapConfigTest)
    this.m = make(map[string]interface{})
    this.AddNewTest("BenchConsumer", createRoutine("BenchConsumer"))
    this.AddNewTest("BenchProducer", createRoutine("BenchProducer"))
    return this
}

func InitMapConfigTest(channel chan string, wg *sync.WaitGroup) *mapConfigTest {
    tmp := mapConfigTest{}
    return tmp.init(channel, wg)
}

func (this *mapConfigTest) AddNewTest(key string, value func(dcode.DecodeurTest, chan string, *sync.WaitGroup) bool) {
    if _, ok := this.m[key]; !ok {
        this.m[key] = value
    }
}

func (this mapConfigTest) MapBenchtestFunctor(cfg dcode.DecodeurTest, channel chan string, wg *sync.WaitGroup) {
    switch cfg.Type {
        case "BenchConsumer":
            go this.m["BenchConsumer"].(func(dcode.DecodeurTest, chan string, *sync.WaitGroup) bool)(cfg, channel, wg)
        case "BenchProducer":
            go this.m["BenchProducer"].(func(dcode.DecodeurTest, chan string, *sync.WaitGroup) bool)(cfg, channel, wg)
            fmt.Println("PRODUCER DONE")
    }
}

func createRoutine(str string) func(cfg dcode.DecodeurTest, channel chan string, wg *sync.WaitGroup) bool {

    kAdmin := topic.KafkaClient{}
    switch str {
        case "BenchConsumer":
            return func(cfg dcode.DecodeurTest, channel chan string, wg *sync.WaitGroup) bool {
                tmp := new (benchC.BenchConsumer)
                return tmp.Run(cfg, channel, wg)
            }
        case "BenchProducer":
            return func(cfg dcode.DecodeurTest, channel chan string, wg *sync.WaitGroup) bool {
                   tmp := new (benchP.BenchProducer)
                   tmp.Admin = kAdmin.CreateClientKafka([]string {host})
                   return tmp.Run(cfg, channel, wg)
            }
        default:
            return func(cfg dcode.DecodeurTest, channel chan string, wg *sync.WaitGroup) bool {
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

func isError(err error) bool {
    return (err != nil)
}

func FlushOutputRoutine(bufferedOutput []string, path []string) {
    //var wg sync.WaitGroup
  //k  wg.Add(len(bufferedOutput) - 1)
    if x := func(x, z int) bool {
        return (x == z)
    }(len(bufferedOutput), len(path)); !x {
        fmt.Println("bufferedOutput, path must have the same size", len(bufferedOutput), len(path))
    }
    for idx, buffer := range bufferedOutput {
        fmt.Println("bufferedOutput", len(buffer))
         func(buffer []byte) {
            file, err := os.OpenFile(path[idx], os.O_RDWR, 0644)
            if !isError(err) {
               fmt.Println("wasn't able to write buffer to file", path)
               //wg.Done()
               return
            }
            fmt.Println("contenu buffer ====> ", string(buffer))
            fmt.Println("write to file", path[idx])
            file.Write(buffer)
//            wg.Done()
            file.Close()
            return
        }([]byte(buffer))
    }
 //   wg.Wait()
}

