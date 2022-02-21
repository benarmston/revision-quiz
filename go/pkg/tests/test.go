package tests

import (
    "fmt"
    "github.com/crus4d3/revision-quiz/go/pkg/utils"
)

func Test() bool {
    fmt.Println("Running tests")
    test1, msg1 := test1()
    if test1 {
        fmt.Println("Test 1 passed")
    } else {
        fmt.Printf("Test 1 failed: %s\n", msg1)
    }
    fmt.Println("Tests finished")
    if test1 == false {
        return false
    } else {
        return true
    }
}

func test1() (bool, string) {
    slice1 := []string{"a"}
    slice2 := []string{"b"}
    slice3 := []string{"a", "b"}
    result1T := utils.Contains("a", slice1)
    result2F := utils.Contains("a", slice2)
    result3T := utils.Contains("b", slice3)
    result4F := utils.Contains("c", slice3)
    if result1T == false {
        return false, "check1 failed"
    } else if result2F == true {
        return false, "check2 failed"
    } else if result3T == false {
        return false, "check3 failed"
    } else if result4F == true {
        return false, "check4 failed"
    } else {
        return true, ""
    }
}
