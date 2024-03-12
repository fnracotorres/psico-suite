package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/mitchellh/mapstructure"
	"github.com/shirou/gopsutil/disk"
)

type Schema struct {
	Data interface{} `json:"data"`
	From string      `json:"from"`
	Kind string      `json:"kind"`
}

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

	registry, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	for {
		daemon, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(daemon, registry)
	}
}

func handleConn(daemon net.Conn, registry net.Conn) {
	defer daemon.Close()
	decoder := json.NewDecoder(daemon)
	var schema Schema
	err := decoder.Decode(&schema)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch data := schema.Data.(type) {
	case nil:
		switch schema.Kind {
		case "connection":
			switch schema.From {
			case "daemon":
				b, err := json.Marshal(Schema{
					Kind: "connection",
					Data: nil,
				})
				if err != nil {
					log.Printf("Error marshaling disk stats: %v", err)
					panic(err)
				}

				_, err = registry.Write(b)
				if err != nil {
					panic(err)
				}
			case "registry":
				decoder := json.NewDecoder(daemon)
				var schema Schema
				err := decoder.Decode(&schema)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				b, err := json.Marshal(Schema{
					Data: nil,
					From: "daemon cluster",
					Kind: schema.Kind,
				})
				if err != nil {
					log.Printf("Error marshaling disk stats: %v", err)
					panic(err)
				}

				_, err = daemon.Write(b)
				if err != nil {
					panic(err)
				}
			default:
				fmt.Println("Unexpected Schema's From")
			}
		}
	case map[string]interface{}:
		fmt.Println("Data is a single object:", data)
	case []interface{}:
		switch schema.Kind {
		case "disk stat list":
			diskStatList := make([]DiskStat, len(data))
			for i, d := range data {
				var diskStat DiskStat
				err := mapstructure.Decode(d, &diskStat)
				if err != nil {
					panic(err)
				}
				diskStatList[i] = diskStat
			}

			b, err := json.Marshal(Schema{
				Data: diskStatList,
				Kind: "disk stat list",
			})
			if err != nil {
				log.Printf("Error marshaling disk stats: %v", err)
				panic(err)
			}

			_, err = registry.Write(b)
			if err != nil {
				panic(err)
			}
		default:
			fmt.Println("Unexpected Schema's Type")
		}
	default:
		fmt.Println("Unexpected type of Schema's Data")
	}
}
