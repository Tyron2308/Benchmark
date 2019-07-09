package abenchtest

import (
    "myproject/decodeurtest"
)

type ABenchTest interface {
    Run(cfgTest decodeurtest.DecodeurTest) bool
}


