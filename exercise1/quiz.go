package main

import (
  "encoding/csv"
  "fmt"
  "os"
  "log"
  "strings"
  "flag"
  "time"
)

func main()  {
    filename := flag.String("csv", "problems.csv", "Specify the name of the csv file; must be of question, answer format")
    limit := flag.String("limit", "5000ms", "Specify the amount of time that you want in milliseconds (ms); Ex: 5000ms = 5s")
    flag.Parse()
    file, error := os.Open(*filename)
    if error != nil {
      log.Fatal(error)
    }
    reader := csv.NewReader(file)
    reader.Comma = ','

    records, error := reader.ReadAll()
    if error != nil {
      log.Fatal(error)
    }

    duration, error := time.ParseDuration(*limit)

    if error != nil{
      log.Fatal(error)
    }

    score := 0
    timer := time.AfterFunc(duration, func(){
      fmt.Printf("\nFinal Score: %d", score)
      os.Exit(0)
    })

    defer timer.Stop()

    for index, value := range records {
        fmt.Printf("Problem #%d %s = ", index+1, value[0])
        var answer string
        fmt.Scanf("%s\n", &answer)
        if(value[1] == strings.TrimRight(answer, "\n")){
          score = score + 10
        }
    }

    fmt.Printf("Final Score: %d", score)
}
