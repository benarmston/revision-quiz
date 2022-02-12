package main

func main () {
    slice []string
}

func contains (search string, slice []string) bool {
    for _, value := range slice {
        if value == search {
            return true
        }
    }
    return false
}
