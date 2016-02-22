package main

import (
    "fmt"
    "time"
    "os"
    "log"
    "bufio"
    "github.com/JeremyOne/subsets"
)

func main() {

    if len(os.Args) < 2{
        fmt.Println("You must provide two arguments: fileA and fileB")
        return
    }

    start := time.Now()

    //read words
    aWords := readFileWords(os.Args[1])
    bWords := readFileWords(os.Args[2])

    fmt.Printf("Set A contains: %[1]d words, set B contains: %[2]d words, loading took: %[3]dms\n", len(aWords), len(bWords), getMillsToNow(start))

    fmt.Println("Running hash map subset tests...")

    start = time.Now()
    result := subsets.ArrayIsSubset_HashMap(aWords, bWords, true, true)
    fmt.Printf("- Do all the words in '%[1]s' appear in '%[2]s'? Function returned: %[3]t\n", os.Args[1], os.Args[2], result)
    fmt.Printf("- Actual search took: %[2]dms\n", result, getMillsToNow(start))

    start = time.Now()
    subsets.ArrayIsSubset_HashMap(aWords, bWords, false, false)
    fmt.Printf("- Worst case (full search) took: %dms\n", getMillsToNow(start))


    fmt.Println("Running rune tree subset tests...")

    start = time.Now()
    result = subsets.ArrayIsSubset_RuneTree(aWords, bWords, true)
    fmt.Printf("- Actual search took: %[2]dms\n", result, getMillsToNow(start))

    start = time.Now()
    subsets.ArrayIsSubset_RuneTree(aWords, bWords, false)
    fmt.Printf("- Worst case (full search) took: %dms\n", getMillsToNow(start))

    
    fmt.Println("Running brute-force subset tests...")

    start = time.Now()
    result = subsets.ArrayIsSubset_BruteForce(aWords, bWords, true)
    fmt.Printf("- Returned %[1]t, actual search took: %[2]dms\n", result, getMillsToNow(start))


    start = time.Now()
    subsets.ArrayIsSubset_BruteForce(aWords, bWords, false)
    fmt.Printf("- Worst case (full search) took: %dms\n", getMillsToNow(start))

}

func getMillsToNow(startTime time.Time) int64 {
    return (time.Now().UnixNano() - startTime.UnixNano()) / int64(time.Millisecond)
}

func readFileWords(path string) []string{

    file, err := os.Open(path)
    if err != nil {
        //if we can't read the file, we can't continue
        log.Fatal(err)
    }

    //close the file when we are done
    defer file.Close()

    scanner := bufio.NewScanner(file)
    return subsets.ReadScannerWords(scanner)
}