package main
 //testR  "myproject/testrunner"
    //dcode  "myproject/decodeurtest"
import (
   initt   "myproject/initbench"
)

func main() {

    initt.CreateClusterKafka()
    //benchtest := testR.InitMapConfigTest()
    //cfg := dcode.GenerateBenchmark("/Users/tyron/Desktop/Work/sncf/Benchmark/myfile.yml")

    //for _, configBench := range cfg.Cfgs {
     //benchtest.CallFunctorStored(configBench)
    //}
}
