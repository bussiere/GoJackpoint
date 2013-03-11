package main

import (
        "encoding/base64"
        "fmt"
        "log"
)

func main() {
        str := base64.StdEncoding.EncodeToString([]byte("Hello, playground"))
        fmt.Println(str)

        data, err := base64.StdEncoding.DecodeString(str)
        if err != nil {
                log.Fatal("error:", err)
        }

        fmt.Printf("%q\n", data)
}