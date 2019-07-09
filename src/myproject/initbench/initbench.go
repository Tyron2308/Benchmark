package initbench

import (
    "fmt"
    "os/exec"
)

func CreateContainer(pathToMake string, containerToSpin string) bool {
    //output, err := exec.Command("/bin/ls").Output()
    output, err := exec.Command("/usr/bin/make", "-C", pathToMake, containerToSpin).Output()
    if err!=nil {
        fmt.Println(err.Error())
        return false
    }
    fmt.Println(pathToMake, containerToSpin, string(output))
    return true
}
