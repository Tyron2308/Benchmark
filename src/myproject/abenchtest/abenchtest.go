package abenchtest

import (
    "myproject/decodeurtest"
    "sync"
)

type ABenchTest interface {
    Run(cfgTest decodeurtest.DecodeurTest, channel chan string, wg *sync.WaitGroup) bool
}


