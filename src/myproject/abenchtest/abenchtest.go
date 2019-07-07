package abenchtest

import (
    "myproject/decodeurtest"
)

type ABenchTest interface {
    RunConcreteTest(cfgTest decodeurtest.DecodeurTest) bool
}


