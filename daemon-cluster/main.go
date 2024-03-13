package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	gopsutilnet "github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

type Schema struct {
	Data interface{} `json:"data"`
	To   string      `json:"from"`
	Kind string      `json:"kind"`
}

type CPUStat struct {
	InfoStatList  []cpu.InfoStat  `json:"infoStatList"`
	TimesStatList []cpu.TimesStat `json:"timeStatList"`
}

type DiskStat struct {
	IOCountersStatList map[string]disk.IOCountersStat `json:"ioCountersStatList"`
	PartitionStat      disk.PartitionStat             `json:"partitionStat"`
	UsageStat          disk.UsageStat                 `json:"usageStat"`
}

type HostStat struct {
	InfoStat            host.InfoStat          `json:"infoStat"`
	TemperatureStatList []host.TemperatureStat `json:"temperatureStatList"`
	UserStatList        []host.UserStat        `json:"userStatList"`
}

type LoadStat struct {
	AvgStat  load.AvgStat
	MiscStat load.MiscStat
}

type MemStat struct {
	SwapDeviceList      []*mem.SwapDevice       `json:"swapDeviceList"`
	SwapMemoryStat      mem.SwapMemoryStat      `json:"swapMemoryStat"`
	VirtualMemoryExStat mem.VirtualMemoryExStat `json:"virtualMemoryExStat"`
	VirtualMemoryStat   mem.VirtualMemoryStat   `json:"virtualMemoryStat"`
}

type NetStat struct {
	ConnectionStatList            []gopsutilnet.ConnectionStat    `json:"connectionStatList"`
	ConntrackStatListAggregation  []gopsutilnet.ConntrackStat     `json:"conntrackStatListAggregation"`
	ConntrackStatListPerCPU       []gopsutilnet.ConntrackStat     `json:"conntrackStatListPerCPU"`
	FilterStatList                []gopsutilnet.FilterStat        `json:"filterStatList"`
	IOCountersStatListAggregation []gopsutilnet.IOCountersStat    `json:"ioCountersStatListAggregation"`
	IOCountersStatListPerNIC      []gopsutilnet.IOCountersStat    `json:"ioCountersStatListPerNIC"`
	InterfaceStatList             []gopsutilnet.InterfaceStat     `json:"swapDeviceList"`
	ProtoCountersStatList         []gopsutilnet.ProtoCountersStat `json:"protoCountersStatList"`
}

// type RlimitStat struct {
// 	Resource uint64   `json:"resource"`
// 	Soft     *big.Int `json:"soft"`
// 	Hard     *big.Int `json:"hard"`
// 	Used     *big.Int `json:"used"`
// }

type ProcessStat struct {
	// IOCountersStat     process.IOCountersStat
	MemoryInfoExStat   process.MemoryInfoExStat   `json:"memoryInfoExStat"`
	MemoryInfoStat     process.MemoryInfoStat     `json:"memoryInfoStat"`
	NumCtxSwitchesStat process.NumCtxSwitchesStat `json:"numCtxSwitchesStat"`
	// RlimitStat         []RlimitStat // Modified // Does not work
	// OpenFilesStat  []process.OpenFilesStat
	// MemoryMapsStat []process.MemoryMapsStat
}

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, b, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			var msg []byte
			if err := json.Unmarshal(b, &msg); err != nil {
				panic(err)
			}

			log.Println("Message: ", string(msg))
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			timestamp, err := json.Marshal([]byte(t.String()))
			if err != nil {
				panic(err)
			}

			err = c.WriteMessage(websocket.TextMessage, timestamp)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

// func handleConn(daemon net.Conn, registry net.Conn) {
// 	defer daemon.Close()

// 	decoder := json.NewDecoder(daemon)
// 	var schema Schema
// 	err := decoder.Decode(&schema)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	switch data := schema.Data.(type) {
// 	case nil:
// 		switch schema.To {
// 		case "daemon":
// 			switch schema.Kind {
// 			case "connection":
// 				handleDaemonConnection(registry)
// 			default:
// 				fmt.Println("Unexpected Schema.Kind")
// 			}
// 		case "registry":
// 			switch schema.Kind {
// 			case "connection":
// 				handleRegistryConnection(daemon)
// 			default:
// 				fmt.Println("Unexpected Schema.Kind")
// 			}
// 		default:
// 			fmt.Println("Unexpected Schema.To")
// 		}
// 	case map[string]interface{}:
// 		switch schema.To {
// 		case "daemon cluster":
// 			switch schema.Kind {
// 			case "cpu stat":
// 				handleDaemonClusterCPUStat(data)
// 			case "host stat":
// 				handleDaemonClusterHostStat(data)
// 			case "load stat":
// 				handleDaemonClusterLoadStat(data)
// 			case "mem stat":
// 				handleDaemonClusterMemStat(data)
// 			case "net stat":
// 				handleDaemonClusterNetStat(data)
// 			default:
// 				fmt.Println("Unexpected Schema.Kind")
// 			}
// 		case "registry":
// 			switch schema.Kind {
// 			case "cpu stat":
// 				handleRegistryCPUStat(data)
// 			case "host stat":
// 				handleRegistryHostStat(data)
// 			case "load stat":
// 				handleRegistryLoadStat(data)
// 			case "mem stat":
// 				handleRegistryMemStat(data)
// 			case "net stat":
// 				handleRegistryNetStat(data)
// 			default:
// 				fmt.Println("Unexpected Schema.Kind")
// 			}
// 		default:
// 			fmt.Println("Unexpected Schema.To")
// 		}
// 	case []interface{}:
// 		switch schema.To {
// 		case "daemon cluster":
// 			switch schema.Kind {
// 			case "disk stat list":
// 				handleDaemonClusterDiskStatList(data)
// 			case "process stat list":
// 				handleDaemonClusterProcessStatList(data)
// 			default:
// 				fmt.Println("Unexpected Schema.Kind")
// 			}
// 		case "registry":
// 			switch schema.Kind {
// 			case "disk stat list":
// 				handleRegistryDiskStatList(data)
// 			case "process stat list":
// 				handleRegistryProcessStatList(data)
// 			default:
// 				fmt.Println("Unexpected Schema.Kind")
// 			}
// 		default:
// 			fmt.Println("Unexpected Schema.To")
// 		}
// 	default:
// 		fmt.Println("Unexpected Schema.Data.(type)")
// 	}
// }

// func handleDaemonConnection(registry net.Conn) {
// 	b, err := json.Marshal(Schema{
// 		Data: nil,
// 		To:   "registry",
// 		Kind: "connection",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = registry.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func handleRegistryConnection(daemon net.Conn) {
// 	b, err := json.Marshal(Schema{
// 		Data: nil,
// 		To:   "daemon cluster",
// 		Kind: "connection",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = daemon.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }
