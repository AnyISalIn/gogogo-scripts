package main

import (
	"fmt"
	"net"
	"os"
)

func checkICMP(ip *string) {
	_, err := net.Dial("ip4:icmp", *ip)
	if err != nil {
		fmt.Printf("%v check fail\n", *ip)
	} else {
		fmt.Printf("%v check icmp success\n", *ip)
	}
}
func check(ips *[]string) {
	for _, ip := range *ips {
		checkICMP(&ip)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage ./live_check IP [...]")
	}
	ips := os.Args[1:]
	check(&ips)
}
