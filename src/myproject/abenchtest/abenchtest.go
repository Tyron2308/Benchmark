package abenchtest

import (
    "myproject/decodeurtest"
    "os"
    "sync"
)

type ABenchTest interface {
    Run(cfgTest decodeurtest.DecodeurTest,wg *sync.WaitGroup) bool
}


func DeferWrtiing(err error, output []byte, nameFile string) (int, error) {
    if err == nil {
        file, errOs := os.OpenFile(nameFile, os.O_RDWR|os.O_CREATE, 0755)
        if errOs == nil {
            return file.Write(output)
        }
    }
    return -1, nil
}


