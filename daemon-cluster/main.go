package main

import (
	"encoding/json"
	"net"

	"github.com/shirou/gopsutil/disk"
)

type DiskStat struct {
	IOCountersStats map[string]disk.IOCountersStat `json:"ioCountersStats"`
	PartitionStat   disk.PartitionStat             `json:"partitionStat"`
	UsageStat       disk.UsageStat                 `json:"usageStat"`
}

func main() {
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	cc, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(c, cc)
	}
}

func handleConn(c net.Conn, cc net.Conn) {
	defer c.Close()

	var diskStats []DiskStat
	decoder := json.NewDecoder(c)
	if err := decoder.Decode(&diskStats); err != nil {
		panic(err)
	}

	b, err := json.Marshal(diskStats)
	if err != nil {
		panic(err)
	}

	_, err = cc.Write(b)
	if err != nil {
		panic(err)
	}
}
