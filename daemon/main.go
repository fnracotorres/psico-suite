package main

import (
	"flag"
	"fmt"
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
	From string      `json:"from"`
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

type Message struct {
	Data   interface{} `json:"data"`
	Sender string      `json:"sender"`
	Type   string      `json:"type"`
}

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}

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
			var msg Message
			err := c.ReadJSON(&msg)
			if err != nil {
				log.Println("read json:", err)
			}

			switch data := msg.Data.(type) {
			case nil:
				switch msg.Sender {
				case "daemon cluster":
					fmt.Println(data)
					switch msg.Type {
					case "disk stat list":
						sendDiskStatList(c)
					default:
						fmt.Println("Unexpected msg.Type")
					}
				default:
					fmt.Println("Unexpected msg.Sender")
				}
			case map[string]interface{}:
			case []interface{}:
			default:
				fmt.Println("Unexpected msg.Data.(type)")
			}

			switch msg.Sender {
			case "daemon":
				switch msg.Type {

				}
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteJSON(t.String())
			if err != nil {
				log.Println("write:", err)
			}
		case <-interrupt:
			log.Println("interrupt")

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

// func sendCPUStat(daemonCluster net.Conn) {
// 	infoStatList, err := cpu.Info()
// 	if err != nil {
// 		panic(err)
// 	}

// 	timesStatList, err := cpu.Times(true)
// 	if err != nil {
// 		panic(err)
// 	}

// 	b, err := json.Marshal(Response{
// 		Type: "cpu stat",
// 		Data: CPUStat{
// 			InfoStatList:  infoStatList,
// 			TimesStatList: timesStatList,
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = daemonCluster.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func sendDiskStatList(c *websocket.Conn) {
	diskStatList := []DiskStat{}

	partitionStats, err := disk.Partitions(true)
	if err != nil {
		panic(err)
	}

	for _, partitionStat := range partitionStats {
		usageStat, err := disk.Usage(partitionStat.Mountpoint)
		if err != nil {
			panic(err)
		}

		ioCountersStats, err := disk.IOCounters(partitionStat.Device)
		if err != nil {
			panic(err)
		}

		diskStat := DiskStat{
			IOCountersStatList: ioCountersStats,
			PartitionStat:      partitionStat,
			UsageStat:          *usageStat,
		}

		diskStatList = append(diskStatList, diskStat)
	}

	msg := Message{Data: diskStatList, Sender: "daemon", Type: "dist stat list"}
	err = c.WriteJSON(msg)
	if err != nil {
		log.Println(err)
	}
}

// func sendHostStat(daemonCluster net.Conn) {
// 	infoStat, err := host.Info()
// 	if err != nil {
// 		panic(err)
// 	}

// 	temperatureStatList, err := host.SensorsTemperatures()
// 	if err != nil {
// 		panic(err)
// 	}

// 	userStatList, err := host.Users()
// 	if err != nil {
// 		panic(err)
// 	}

// 	b, err := json.Marshal(Response{
// 		Type: "host stat",
// 		Data: HostStat{
// 			InfoStat:            *infoStat,
// 			TemperatureStatList: temperatureStatList,
// 			UserStatList:        userStatList,
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = daemonCluster.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func sendLoadStat(daemonCluster net.Conn) {
// 	avgStat, err := load.Avg()
// 	if err != nil {
// 		panic(err)
// 	}

// 	miscStat, err := load.Misc()
// 	if err != nil {
// 		panic(err)
// 	}

// 	b, err := json.Marshal(Response{
// 		Type: "load stat",
// 		Data: LoadStat{
// 			AvgStat:  *avgStat,
// 			MiscStat: *miscStat,
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = daemonCluster.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func sendMemStat(daemonCluster net.Conn) {
// 	swapDevices, err := mem.SwapDevices()
// 	if err != nil {
// 		panic(err)
// 	}

// 	swapMemoryStat, err := mem.SwapMemory()
// 	if err != nil {
// 		panic(err)
// 	}

// 	virtualMemoryExStat, err := mem.VirtualMemoryEx()
// 	if err != nil {
// 		panic(err)
// 	}

// 	virtualMemoryStat, err := mem.VirtualMemory()
// 	if err != nil {
// 		panic(err)
// 	}

// 	b, err := json.Marshal(Response{
// 		Type: "disk stat list",
// 		Data: MemStat{
// 			SwapDeviceList:      swapDevices,
// 			SwapMemoryStat:      *swapMemoryStat,
// 			VirtualMemoryExStat: *virtualMemoryExStat,
// 			VirtualMemoryStat:   *virtualMemoryStat,
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = daemonCluster.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func sendNetStat(daemonCluster net.Conn) {
// 	// .ConnectionsMax
// 	// .ConnectionsPidMax
// 	// .ConnectionsPidMaxWithoutUids
// 	// .ConnectionsPidWithoutUids
// 	// "all"
// 	// "inet"
// 	// "inet4"
// 	// "inet6"
// 	// "tcp"
// 	// "tcp6"
// 	// "tcp4"
// 	// "udp"
// 	// "udp6"
// 	// "udp4"
// 	// "unix"
// 	connectionStatList, err := gopsutilnet.Connections("all")
// 	if err != nil {
// 		panic(err)
// 	}

// 	conntrackStatListPerCPU, err := gopsutilnet.ConntrackStats(true)
// 	if err != nil {
// 		// panic(err) panic: open /proc/net/stat/nf_conntrack: no such file or directory
// 	}

// 	conntrackStatAggregation, err := gopsutilnet.ConntrackStats(false)
// 	if err != nil {
// 		// panic(err) panic: open /proc/net/stat/nf_conntrack: no such file or directory
// 	}

// 	filterStatList, err := gopsutilnet.FilterCounters()
// 	if err != nil {
// 		// panic(err) panic: open /proc/sys/net/netfilter/nf_conntrack_count: no such file or directory
// 	}

// 	ioCountersStatListPerNIC, err := gopsutilnet.IOCounters(true)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, filterStat := range filterStatList {
// 		fmt.Println("ConnTrackCount:", filterStat.ConnTrackCount)
// 		fmt.Println("ConnTrackMax:", filterStat.ConnTrackMax)
// 	}

// 	ioCountersStatListAggregation, err := gopsutilnet.IOCounters(true)
// 	if err != nil {
// 		panic(err)
// 	}

// 	interfaceStatList, err := gopsutilnet.Interfaces()
// 	if err != nil {
// 		panic(err)
// 	}

// 	//ip,icmp,icmpmsg,tcp,udp,udplite
// 	protocols := []string{"tcp"}
// 	protoCountersStatList, err := gopsutilnet.ProtoCounters(protocols)
// 	if err != nil {
// 		panic(err)
// 	}

// 	b, err := json.Marshal(Response{
// 		Type: "net stat",
// 		Data: NetStat{
// 			ConnectionStatList:            connectionStatList,
// 			ConntrackStatListPerCPU:       conntrackStatListPerCPU,
// 			ConntrackStatListAggregation:  conntrackStatAggregation,
// 			FilterStatList:                filterStatList,
// 			IOCountersStatListPerNIC:      ioCountersStatListPerNIC,
// 			IOCountersStatListAggregation: ioCountersStatListAggregation,
// 			InterfaceStatList:             interfaceStatList,
// 			ProtoCountersStatList:         protoCountersStatList,
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = daemonCluster.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func sendProcessStat(daemonCluster net.Conn) {
// 	var processStatList []ProcessStat

// 	pids, err := process.Pids()
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, pid := range pids {
// 		proc, err := process.NewProcess(pid)
// 		if err != nil {
// 			fmt.Printf("Error creating process object for PID %d: %s\n", pid, err)
// 			panic(err)
// 		}

// 		// ioCountersStat, err := proc.IOCounters()
// 		// if err != nil {
// 		// 	fmt.Printf("Error retrieving IO counters for PID %d: %s\n", pid, err)
// 		// 	// panic(err) panic: open /proc/1/io: permission denied
// 		// }

// 		memoryInfoExStat, err := proc.MemoryInfoEx()
// 		if err != nil {
// 			fmt.Printf("Error retrieving extended memory info for PID %d: %s\n", pid, err)
// 			panic(err)
// 		}

// 		memoryInfoStat, err := proc.MemoryInfo()
// 		if err != nil {
// 			fmt.Printf("Error retrieving memory info for PID %d: %s\n", pid, err)
// 			panic(err)
// 		}

// 		// memoryMapsStatList, err := proc.MemoryMaps(true) // also false
// 		// if err != nil {
// 		// 	fmt.Printf("Error retrieving memory maps for PID %d: %s\n", pid, err)
// 		// 	// panic(err) panic: open /proc/1/smaps: permission denied
// 		// }

// 		numCtxSwitchesStat, err := proc.NumCtxSwitches()
// 		if err != nil {
// 			fmt.Printf("Error retrieving context switches for PID %d: %s\n", pid, err)
// 			panic(err)
// 		}

// 		// openFilesStatList, err := proc.OpenFiles()
// 		// if err != nil {
// 		// 	fmt.Printf("Error retrieving open files for PID %d: %s\n", pid, err)
// 		// 	// panic(err) panic: open /proc/1/smaps: permission denied
// 		// }

// 		// rlimitStatList, err := proc.Rlimit()
// 		// if err != nil {
// 		// 	fmt.Printf("Error retrieving resource limits for PID %d: %s\n", pid, err)
// 		// 	panic(err)
// 		// } Not working, TOOD: Fork, fix, and pull request

// 		// fmt.Println(ioCountersStat.ReadBytes)
// 		// fmt.Println(ioCountersStat.ReadCount)
// 		// fmt.Println(ioCountersStat.WriteBytes)
// 		// fmt.Println(ioCountersStat.WriteCount)
// 		// if panic above then: Error retrieving IO counters for PID 1: open /proc/1/io: permission denied

// 		// for _, memoryMapsStat := range *memoryMapsStatList {
// 		// 	fmt.Println(memoryMapsStat.Anonymous)
// 		// 	fmt.Println(memoryMapsStat.Path)
// 		// 	fmt.Println(memoryMapsStat.PrivateClean)
// 		// 	fmt.Println(memoryMapsStat.PrivateDirty)
// 		// 	fmt.Println(memoryMapsStat.Pss)
// 		// 	fmt.Println(memoryMapsStat.Referenced)
// 		// 	fmt.Println(memoryMapsStat.Rss)
// 		// 	fmt.Println(memoryMapsStat.SharedClean)
// 		// 	fmt.Println(memoryMapsStat.SharedDirty)
// 		// 	fmt.Println(memoryMapsStat.Size)
// 		// 	fmt.Println(memoryMapsStat.Swap)
// 		// }
// 		// Error retrieving memory maps for PID 1: open /proc/1/smaps: permission denied

// 		// for _, openFilesStat := range openFilesStatList {
// 		// 	fmt.Println(openFilesStat.Fd)
// 		// 	fmt.Println(openFilesStat.Path)
// 		// }
// 		// Error retrieving open files for PID 1: open /proc/1/fd: permission denied

// 		// for _, rlimitStat := range rlimitStatList {
// 		// 	fmt.Println(rlimitStat.Hard)
// 		// 	fmt.Println(rlimitStat.Resource)
// 		// 	fmt.Println(rlimitStat.Soft)
// 		// 	fmt.Println(rlimitStat.Used)
// 		// } // Not working

// 		processStatList = append(processStatList, ProcessStat{
// 			MemoryInfoExStat:   *memoryInfoExStat,
// 			MemoryInfoStat:     *memoryInfoStat,
// 			NumCtxSwitchesStat: *numCtxSwitchesStat,
// 		})
// 	}

// 	b, err := json.Marshal(Response{
// 		Type: "process stat list",
// 		Data: processStatList,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = daemonCluster.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func powerOnRemoteDesk(macAddr string) error {
// 	// Parse MAC address
// 	mac, err := net.ParseMAC(macAddr)
// 	if err != nil {
// 		return fmt.Errorf("error parsing MAC address: %v", err)
// 	}

// 	// Create magic packet
// 	magicPacket := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
// 	for i := 0; i < 16; i++ {
// 		magicPacket = append(magicPacket, mac...)
// 	}

// 	// Resolve broadcast address
// 	broadcastAddr := net.IPv4(255, 255, 255, 255)

// 	// Setup UDP connection
// 	udpAddr := &net.UDPAddr{
// 		IP:   broadcastAddr,
// 		Port: 9,
// 	}
// 	conn, err := net.DialUDP("udp", nil, udpAddr)
// 	if err != nil {
// 		return fmt.Errorf("error setting up UDP connection: %v", err)
// 	}
// 	defer conn.Close()

// 	// Send magic packet
// 	if _, err := conn.Write(magicPacket); err != nil {
// 		return fmt.Errorf("error sending magic packet: %v", err)
// 	}

// 	return nil
// }

// func powerOffLocalDesk() error {
// 	var cmd *exec.Cmd
// 	switch os := runtime.GOOS; os {
// 	case "windows":
// 		cmd = exec.Command("shutdown", "/s", "/t", "0")
// 	case "linux":
// 		cmd = exec.Command("shutdown", "-h", "now")
// 	case "darwin":
// 		cmd = exec.Command("sudo", "shutdown", "-h", "now")
// 	default:
// 		return fmt.Errorf("unsupported operating system: %s", os)
// 	}

// 	if err := cmd.Run(); err != nil {
// 		return fmt.Errorf("error shutting down computer: %v", err)
// 	}
// 	return nil
// }
