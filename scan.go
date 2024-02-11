package main

import (
        "fmt"
        "net"
        "time"
        "math/rand"
)
// make a random ip
func generateRandomIP() string {
        rand.Seed(time.Now().UnixNano())
        ip := net.IPv4(
                byte(rand.Intn(256)),
                byte(rand.Intn(256)),
                byte(rand.Intn(256)),
                byte(rand.Intn(256)),
        )
        return ip.String()
}
// pinging 
func isPortOpen(ip string, port int, timeout time.Duration) bool {
        address := fmt.Sprintf("%s:%d", ip, port)
        conn, err := net.DialTimeout("tcp", address, timeout)
        if err != nil {
                return false
        }
        defer conn.Close()
        return true
}

func main() {
//loop this
for {
        ipAddress := generateRandomIP()
        // Define the port to check
        port := 25565
        // Set a timeout for the connection attempt
        timeout := 3 * time.Second
        // Check if the specified port is open on the given IP address
        if isPortOpen(ipAddress, port, timeout) {
                fmt.Printf("Port is open on IP %s\n", ipAddress)
        } else {
                fmt.Printf("Port is closed on IP %s\n", ipAddress)
        }
}
}
