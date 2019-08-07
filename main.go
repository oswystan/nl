//===============================================================================
//                      Copyright (C) 2019 wystan
//
//        filename: main.go
//     description:
//         created: 2019-08-07 16:17:08
//          author: wystan
//
//===============================================================================

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"
)

var ip string
var interval int
var total uint

func init() {
	flag.StringVar(&ip, "a", "10.33.11.31:7000", "address of server with port")
	flag.IntVar(&interval, "i", 100, "interval in ms between 2 packets")
	flag.UintVar(&total, "c", 1000, "total packets")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func udp_newconn(server string) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", server)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func main() {
	flag.Parse()

	conn, err := udp_newconn(ip)
	if err != nil {
		fmt.Printf("fail to dail udp\n")
		os.Exit(1)
	}

	defer conn.Close()

	data := ""
	var i uint
	for i = 0; i < total; i++ {
		data = fmt.Sprintf("%d\n", i)
		fmt.Printf("%s", data)
		conn.Write([]byte(data))
		time.Sleep(time.Millisecond * time.Duration(interval))
	}

	return
}

//==================================== END ======================================
