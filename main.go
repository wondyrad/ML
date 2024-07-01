package main

import (
    "encoding/binary"
    "fmt"
    "math/rand"
    "net"
    "net/http"
    "os"
    "time"
)

func main() {
    url := "https://petstore.wcfdemo.com"

    fmt.Println("------------------------------------------------------------------------")
    fmt.Printf("Sending API GET requests to %s/\n", url)
    fmt.Println("------------------------------------------------------------------------")

    for i := 1; i < 100; i++ {
        ipAddress := generateRandomIP()
        fmt.Printf("GET : %s - HTTP status = ", url)
        statusCode := sendGetRequest(url+"/api/pet/findByStatus?status=available", ipAddress)
        fmt.Println(statusCode)

        ipAddress = generateRandomIP()
        fmt.Printf("GET : %s - HTTP status = ", url)
        statusCode = sendGetRequest(url+"/api/pet/v3/findByStatus?status=pending", ipAddress)
        fmt.Println(statusCode)

        ipAddress = generateRandomIP()
        fmt.Printf("GET : %s - HTTP status = ", url)
        statusCode = sendGetRequest(url+"/api/pet/findByStatus?status=sold", ipAddress)
        fmt.Println(statusCode)
    }

    fmt.Println("-------------------------------------------------------------------------------------------")
    fmt.Printf("FortiWeb ML-API trained with GET method on %s/\n", url)
    fmt.Println("-------------------------------------------------------------------------------------------")
}

func generateRandomIP() string {
    rand.Seed(time.Now().UnixNano())
    ip := make(net.IP, 4)
    binary.BigEndian.PutUint32(ip, rand.Uint32())
    return ip.String()
}

func sendGetRequest(url, ipAddress string) int {
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        os.Exit(1)
    }

    req.Header.Set("X-Forwarded-For", ipAddress)
    req.Header.Set("User-Agent", "ML-Requester")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept-Encoding", "identity") // Disable compression

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    return resp.StatusCode
}

