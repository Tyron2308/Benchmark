package main

import (
  "fmt"
   testR  "myproject/testrunner"
   dcode  "myproject/decodeurtest"
   "sync"
   "time"
)

const (
    yaml = "/Users/tyron/Desktop/Work/sncf/Benchmark/testtorun.yml"
)

func main() {
    if ok := testR.InitCluster(); !ok {
        fmt.Println("cluster has not been able to initiate itself successfully")
        return
    }
    var wg sync.WaitGroup
    i := 0
    cfg := dcode.GenerateBenchmark(yaml)

    bufferedOutput  := make([]string, len(cfg.Cfgs))
    var paths []string
    channel := make(chan string, len(cfg.Cfgs))

    defer close(channel)

    fmt.Println("====> initiated cluster <====")
    benchtest := testR.InitMapConfigTest(channel, &wg)
    wg.Add(len(cfg.Cfgs))
    for _, configBench := range cfg.Cfgs {
         benchtest.MapBenchtestFunctor(configBench, channel, &wg)
         paths = append(paths, configBench.Outputfile)

    }
    c := <-channel
    fmt.Println("value buffer ===> ", string(c))
    bufferedOutput[i] = c
    i++
    wg.Wait()
    fmt.Println("====> destroying cluster test <====", c)
//    testR.FlushOutputRoutine(bufferedOutput, paths)
    fmt.Println("sleep berfore killing programs")
    time.Sleep(15 * time.Second)
}

