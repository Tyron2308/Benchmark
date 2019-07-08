package main
import (
   testR  "myproject/testrunner"
   //initt   "myproject/initbench"
   dcode  "myproject/decodeurtest"
)

func main() {

    benchtest := testR.InitMapConfigTest()
    cfg := dcode.GenerateBenchmark("/Users/tyron/Desktop/Work/sncf/Benchmark/testtorun.yml")

    for _, configBench := range cfg.Cfgs {
     benchtest.CallFunctorStored(configBench)
    }
}
