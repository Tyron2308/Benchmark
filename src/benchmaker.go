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
    cfg := dcode.GenerateBenchmark(yaml)
    var paths []string

    fmt.Println("====> initiated cluster <====")
    benchtest := testR.InitMapConfigTest(&wg)
    fmt.Println("size cfgs", len(cfg.Cfgs))
    wg.Add(len(cfg.Cfgs))
    for _, configBench := range cfg.Cfgs {
         benchtest.MapBenchtestFunctor(configBench, &wg)
         paths = append(paths, configBench.Outputfile)

    }
    wg.Wait()
    fmt.Println("sleep berfore killing programs")
    time.Sleep(15 * time.Second)
}

