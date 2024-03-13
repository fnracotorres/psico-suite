package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

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

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(*addr, nil))
}

// func handleConn(daemonCluster net.Conn) {
// 	defer daemonCluster.Close()

// 	decoder := json.NewDecoder(daemonCluster)
// 	var schema Schema
// 	err := decoder.Decode(&schema)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	switch data := schema.Data.(type) {
// 	case nil:
// 		switch schema.To {
// 		case "daemon cluster":
// 			switch schema.Kind {
// 			case "connection":
// 				handleDaemonClusterConnection(daemonCluster)
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

// func handleDaemonClusterConnection(daemonCluster net.Conn) {
// 	b, err := json.Marshal(Schema{
// 		Data: nil,
// 		To:   "registry",
// 		Kind: "connection",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = daemonCluster.Write(b)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func handleDaemonClusterCPUStat(data map[string]interface{}) {
// 	var cpuStat CPUStat
// 	err := mapstructure.Decode(data, &cpuStat)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, infoStat := range cpuStat.InfoStatList {
// 		fmt.Println("CacheSize:", infoStat.CacheSize)
// 		fmt.Println("CoreID:", infoStat.CoreID)
// 		fmt.Println("CPU:", infoStat.CPU)
// 		fmt.Println("Cores:", infoStat.Cores)
// 		fmt.Println("Family:", infoStat.Family)
// 		fmt.Println("Flags:", infoStat.Flags)
// 		fmt.Println("Microcode:", infoStat.Microcode)
// 		fmt.Println("Model:", infoStat.Model)
// 		fmt.Println("ModelName:", infoStat.ModelName)
// 		fmt.Println("Mhz:", infoStat.Mhz)
// 		fmt.Println("PhysicalID:", infoStat.PhysicalID)
// 		fmt.Println("Stepping:", infoStat.Stepping)
// 		fmt.Println("VendorID:", infoStat.VendorID)
// 	}

// 	for _, timesStat := range cpuStat.TimesStatList {
// 		fmt.Println("CPU:", timesStat.CPU)
// 		fmt.Println("Guest:", timesStat.Guest)
// 		fmt.Println("GuestNice:", timesStat.GuestNice)
// 		fmt.Println("Idle:", timesStat.Idle)
// 		fmt.Println("Iowait:", timesStat.Iowait)
// 		fmt.Println("Irq:", timesStat.Irq)
// 		fmt.Println("Nice:", timesStat.Nice)
// 		fmt.Println("Softirq:", timesStat.Softirq)
// 		fmt.Println("Steal:", timesStat.Steal)
// 		fmt.Println("System:", timesStat.System)
// 		fmt.Println("User:", timesStat.User)
// 	}
// }

// func handleDaemonClusterHostStat(data map[string]interface{}) {
// 	var hostStat HostStat
// 	err := mapstructure.Decode(data, &hostStat)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Hostname:", hostStat.InfoStat.Hostname)
// 	fmt.Println("Uptime:", hostStat.InfoStat.Uptime)
// 	fmt.Println("BootTime:", hostStat.InfoStat.BootTime)
// 	fmt.Println("Procs:", hostStat.InfoStat.Procs)
// 	fmt.Println("OS:", hostStat.InfoStat.OS)
// 	fmt.Println("Platform:", hostStat.InfoStat.Platform)
// 	fmt.Println("PlatformFamily:", hostStat.InfoStat.PlatformFamily)
// 	fmt.Println("PlatformVersion:", hostStat.InfoStat.PlatformVersion)
// 	fmt.Println("KernelVersion:", hostStat.InfoStat.KernelVersion)
// 	fmt.Println("KernelArch:", hostStat.InfoStat.KernelArch)
// 	fmt.Println("VirtualizationSystem:", hostStat.InfoStat.VirtualizationSystem)
// 	fmt.Println("VirtualizationRole:", hostStat.InfoStat.VirtualizationRole)
// 	fmt.Println("HostID:", hostStat.InfoStat.HostID)

// 	for _, temperatureStat := range hostStat.TemperatureStatList {
// 		fmt.Println("SensorKey:", temperatureStat.SensorKey)
// 		fmt.Println("Temperature:", temperatureStat.Temperature)
// 	}

// 	for _, userStat := range hostStat.UserStatList {
// 		fmt.Println("User:", userStat.User)
// 		fmt.Println("Terminal:", userStat.Terminal)
// 		fmt.Println("Host:", userStat.Host)
// 		fmt.Println("Started:", userStat.Started)
// 	}
// }

// func handleDaemonClusterLoadStat(data map[string]interface{}) {
// 	var loadStat LoadStat
// 	err := mapstructure.Decode(data, &loadStat)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Load1:", loadStat.AvgStat.Load1)
// 	fmt.Println("Load5:", loadStat.AvgStat.Load5)
// 	fmt.Println("Load15:", loadStat.AvgStat.Load15)

// 	fmt.Println("ProcsTotal:", loadStat.MiscStat.ProcsTotal)
// 	fmt.Println("ProcsCreated:", loadStat.MiscStat.ProcsCreated)
// 	fmt.Println("ProcsRunning:", loadStat.MiscStat.ProcsRunning)
// 	fmt.Println("ProcsBlocked:", loadStat.MiscStat.ProcsBlocked)
// 	fmt.Println("Ctxt:", loadStat.MiscStat.Ctxt)
// }

// func handleDaemonClusterMemStat(data map[string]interface{}) {
// 	var memStat MemStat
// 	err := mapstructure.Decode(data, &memStat)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, swapDevice := range memStat.SwapDeviceList {
// 		fmt.Println("FreeBytes:", swapDevice.FreeBytes)
// 		fmt.Println("Name:", swapDevice.Name)
// 		fmt.Println("UsedBytes:", swapDevice.UsedBytes)
// 	}

// 	fmt.Println("Free:", memStat.SwapMemoryStat.Free)
// 	fmt.Println("PgFault:", memStat.SwapMemoryStat.PgFault)
// 	fmt.Println("PgIn:", memStat.SwapMemoryStat.PgIn)
// 	fmt.Println("PgMajFault:", memStat.SwapMemoryStat.PgMajFault)
// 	fmt.Println("PgOut:", memStat.SwapMemoryStat.PgOut)
// 	fmt.Println("Sin:", memStat.SwapMemoryStat.Sin)
// 	fmt.Println("Sout:", memStat.SwapMemoryStat.Sout)
// 	fmt.Println("Total:", memStat.SwapMemoryStat.Total)
// 	fmt.Println("Used:", memStat.SwapMemoryStat.Used)
// 	fmt.Println("UsedPercent:", memStat.SwapMemoryStat.UsedPercent)

// 	fmt.Println("ActiveAnon:", memStat.VirtualMemoryExStat.ActiveAnon)
// 	fmt.Println("ActiveFile:", memStat.VirtualMemoryExStat.ActiveFile)
// 	fmt.Println("InactiveAnon:", memStat.VirtualMemoryExStat.InactiveAnon)
// 	fmt.Println("InactiveFile:", memStat.VirtualMemoryExStat.InactiveFile)
// 	fmt.Println("Unevictable:", memStat.VirtualMemoryExStat.Unevictable)

// 	fmt.Println("Active:", memStat.VirtualMemoryStat.Active)
// 	fmt.Println("Available:", memStat.VirtualMemoryStat.Available)
// 	fmt.Println("Buffers:", memStat.VirtualMemoryStat.Buffers)
// 	fmt.Println("Cached:", memStat.VirtualMemoryStat.Cached)
// 	fmt.Println("CommitLimit:", memStat.VirtualMemoryStat.CommitLimit)
// 	fmt.Println("CommittedAS:", memStat.VirtualMemoryStat.CommittedAS)
// 	fmt.Println("Dirty:", memStat.VirtualMemoryStat.Dirty)
// 	fmt.Println("Free:", memStat.VirtualMemoryStat.Free)
// 	fmt.Println("HighFree:", memStat.VirtualMemoryStat.HighFree)
// 	fmt.Println("HighTotal:", memStat.VirtualMemoryStat.HighTotal)
// 	fmt.Println("HugePageSize:", memStat.VirtualMemoryStat.HugePageSize)
// 	fmt.Println("HugePagesFree:", memStat.VirtualMemoryStat.HugePagesFree)
// 	fmt.Println("HugePagesTotal:", memStat.VirtualMemoryStat.HugePagesTotal)
// 	fmt.Println("Inactive:", memStat.VirtualMemoryStat.Inactive)
// 	fmt.Println("Laundry:", memStat.VirtualMemoryStat.Laundry)
// 	fmt.Println("LowFree:", memStat.VirtualMemoryStat.LowFree)
// 	fmt.Println("LowTotal:", memStat.VirtualMemoryStat.LowTotal)
// 	fmt.Println("Mapped:", memStat.VirtualMemoryStat.Mapped)
// 	fmt.Println("PageTables:", memStat.VirtualMemoryStat.PageTables)
// 	fmt.Println("Shared:", memStat.VirtualMemoryStat.Shared)
// 	fmt.Println("SReclaimable:", memStat.VirtualMemoryStat.SReclaimable)
// 	fmt.Println("Slab:", memStat.VirtualMemoryStat.Slab)
// 	fmt.Println("SwapCached:", memStat.VirtualMemoryStat.SwapCached)
// 	fmt.Println("SwapFree:", memStat.VirtualMemoryStat.SwapFree)
// 	fmt.Println("SwapTotal:", memStat.VirtualMemoryStat.SwapTotal)
// 	fmt.Println("Total:", memStat.VirtualMemoryStat.Total)
// 	fmt.Println("Used:", memStat.VirtualMemoryStat.Used)
// 	fmt.Println("UsedPercent:", memStat.VirtualMemoryStat.UsedPercent)
// 	fmt.Println("VMallocChunk:", memStat.VirtualMemoryStat.VMallocChunk)
// 	fmt.Println("VMallocTotal:", memStat.VirtualMemoryStat.VMallocTotal)
// 	fmt.Println("VMallocUsed:", memStat.VirtualMemoryStat.VMallocUsed)
// 	fmt.Println("Wired:", memStat.VirtualMemoryStat.Wired)
// 	fmt.Println("Writeback:", memStat.VirtualMemoryStat.Writeback)
// 	fmt.Println("WritebackTmp:", memStat.VirtualMemoryStat.WritebackTmp)
// }

// func handleDaemonClusterNetStat(data map[string]interface{}) {
// 	var netStat NetStat
// 	err := mapstructure.Decode(data, &netStat)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, connectionStat := range netStat.ConnectionStatList {
// 		fmt.Println("Family:", connectionStat.Family)
// 		fmt.Println("Fd:", connectionStat.Fd)
// 		fmt.Println("Laddr:", connectionStat.Laddr)
// 		fmt.Println("Pid:", connectionStat.Pid)
// 		fmt.Println("Raddr:", connectionStat.Raddr)
// 		fmt.Println("Status:", connectionStat.Status)
// 		fmt.Println("Type:", connectionStat.Type)
// 		fmt.Println("Uids:", connectionStat.Uids)
// 	}

// 	for _, conntrackStat := range netStat.ConntrackStatListAggregation {
// 		fmt.Println("Delete:", conntrackStat.Delete)
// 		fmt.Println("DeleteList:", conntrackStat.DeleteList)
// 		fmt.Println("Drop:", conntrackStat.Drop)
// 		fmt.Println("EarlyDrop:", conntrackStat.EarlyDrop)
// 		fmt.Println("Entries:", conntrackStat.Entries)
// 		fmt.Println("ExpectCreate:", conntrackStat.ExpectCreate)
// 		fmt.Println("ExpectDelete:", conntrackStat.ExpectDelete)
// 		fmt.Println("ExpectNew:", conntrackStat.ExpectNew)
// 		fmt.Println("Found:", conntrackStat.Found)
// 		fmt.Println("Ignore:", conntrackStat.Ignore)
// 		fmt.Println("IcmpError:", conntrackStat.IcmpError)
// 		fmt.Println("Insert:", conntrackStat.Insert)
// 		fmt.Println("InsertFailed:", conntrackStat.InsertFailed)
// 		fmt.Println("Invalid:", conntrackStat.Invalid)
// 		fmt.Println("New:", conntrackStat.New)
// 		fmt.Println("SearchRestart:", conntrackStat.SearchRestart)
// 		fmt.Println("Searched:", conntrackStat.Searched)
// 	}

// 	for _, conntrackStat := range netStat.ConntrackStatListPerCPU {
// 		fmt.Println("Delete:", conntrackStat.Delete)
// 		fmt.Println("DeleteList:", conntrackStat.DeleteList)
// 		fmt.Println("Drop:", conntrackStat.Drop)
// 		fmt.Println("EarlyDrop:", conntrackStat.EarlyDrop)
// 		fmt.Println("Entries:", conntrackStat.Entries)
// 		fmt.Println("ExpectCreate:", conntrackStat.ExpectCreate)
// 		fmt.Println("ExpectDelete:", conntrackStat.ExpectDelete)
// 		fmt.Println("ExpectNew:", conntrackStat.ExpectNew)
// 		fmt.Println("Found:", conntrackStat.Found)
// 		fmt.Println("Ignore:", conntrackStat.Ignore)
// 		fmt.Println("IcmpError:", conntrackStat.IcmpError)
// 		fmt.Println("Insert:", conntrackStat.Insert)
// 		fmt.Println("InsertFailed:", conntrackStat.InsertFailed)
// 		fmt.Println("Invalid:", conntrackStat.Invalid)
// 		fmt.Println("New:", conntrackStat.New)
// 		fmt.Println("SearchRestart:", conntrackStat.SearchRestart)
// 		fmt.Println("Searched:", conntrackStat.Searched)
// 	}

// 	for _, filterStat := range netStat.FilterStatList {
// 		fmt.Println("ConnTrackCount:", filterStat.ConnTrackCount)
// 		fmt.Println("ConnTrackMax:", filterStat.ConnTrackMax)
// 	}

// 	for _, ioCountersStat := range netStat.IOCountersStatListAggregation {
// 		fmt.Println("BytesRecv:", ioCountersStat.BytesRecv)
// 		fmt.Println("BytesSent:", ioCountersStat.BytesSent)
// 		fmt.Println("Dropin:", ioCountersStat.Dropin)
// 		fmt.Println("Dropout:", ioCountersStat.Dropout)
// 		fmt.Println("Errin:", ioCountersStat.Errin)
// 		fmt.Println("Errout:", ioCountersStat.Errout)
// 		fmt.Println("Fifoin:", ioCountersStat.Fifoin)
// 		fmt.Println("Fifoout:", ioCountersStat.Fifoout)
// 		fmt.Println("Name:", ioCountersStat.Name)
// 		fmt.Println("PacketsRecv:", ioCountersStat.PacketsRecv)
// 		fmt.Println("PacketsSent:", ioCountersStat.PacketsSent)
// 	}

// 	for _, ioCountersStat := range netStat.IOCountersStatListPerNIC {
// 		fmt.Println("BytesRecv:", ioCountersStat.BytesRecv)
// 		fmt.Println("BytesSent:", ioCountersStat.BytesSent)
// 		fmt.Println("Dropin:", ioCountersStat.Dropin)
// 		fmt.Println("Dropout:", ioCountersStat.Dropout)
// 		fmt.Println("Errin:", ioCountersStat.Errin)
// 		fmt.Println("Errout:", ioCountersStat.Errout)
// 		fmt.Println("Fifoin:", ioCountersStat.Fifoin)
// 		fmt.Println("Fifoout:", ioCountersStat.Fifoout)
// 		fmt.Println("Name:", ioCountersStat.Name)
// 		fmt.Println("PacketsRecv:", ioCountersStat.PacketsRecv)
// 		fmt.Println("PacketsSent:", ioCountersStat.PacketsSent)
// 	}

// 	for _, interfaceStat := range netStat.InterfaceStatList {
// 		fmt.Println("Addrs:", interfaceStat.Addrs)
// 		fmt.Println("Flags:", interfaceStat.Flags)
// 		fmt.Println("HardwareAddr:", interfaceStat.HardwareAddr)
// 		fmt.Println("Index:", interfaceStat.Index)
// 		fmt.Println("MTU:", interfaceStat.MTU)
// 		fmt.Println("Name:", interfaceStat.Name)
// 	}

// 	for _, protoCountersStat := range netStat.ProtoCountersStatList {
// 		fmt.Println("Protocol:", protoCountersStat.Protocol)
// 		fmt.Println("Stats:", protoCountersStat.Stats)
// 	}
// }

// func handleRegistryCPUStat(data map[string]interface{}) {}

// func handleRegistryHostStat(data map[string]interface{}) {}

// func handleRegistryLoadStat(data map[string]interface{}) {}

// func handleRegistryMemStat(data map[string]interface{}) {}

// func handleRegistryNetStat(data map[string]interface{}) {}

// func handleDaemonClusterDiskStatList(data []interface{}) {
// 	diskStatList := make([]DiskStat, len(data))
// 	for i, d := range data {
// 		var diskStat DiskStat
// 		err := mapstructure.Decode(d, &diskStat)
// 		if err != nil {
// 			panic(err)
// 		}
// 		diskStatList[i] = diskStat
// 	}

// 	for _, diskStat := range diskStatList {
// 		fmt.Println("Device:", diskStat.PartitionStat.Device)
// 		fmt.Println("Fstype:", diskStat.PartitionStat.Fstype)
// 		fmt.Println("Mountpoint:", diskStat.PartitionStat.Mountpoint)
// 		fmt.Println("Opts:", diskStat.PartitionStat.Opts)

// 		for _, ioCountersStat := range diskStat.IOCountersStatList {
// 			fmt.Println("IoTime:", ioCountersStat.IoTime)
// 			fmt.Println("IopsInProgress:", ioCountersStat.IopsInProgress)
// 			fmt.Println("Label:", ioCountersStat.Label)
// 			fmt.Println("MergedReadCount:", ioCountersStat.MergedReadCount)
// 			fmt.Println("MergedWriteCount:", ioCountersStat.MergedWriteCount)
// 			fmt.Println("Name:", ioCountersStat.Name)
// 			fmt.Println("ReadBytes:", ioCountersStat.ReadBytes)
// 			fmt.Println("ReadCount:", ioCountersStat.ReadCount)
// 			fmt.Println("ReadTime:", ioCountersStat.ReadTime)
// 			fmt.Println("SerialNumber:", ioCountersStat.SerialNumber)
// 			fmt.Println("WeightedIO:", ioCountersStat.WeightedIO)
// 			fmt.Println("WriteBytes:", ioCountersStat.WriteBytes)
// 			fmt.Println("WriteCount:", ioCountersStat.WriteCount)
// 			fmt.Println("WriteTime:", ioCountersStat.WriteTime)
// 		}

// 		fmt.Println("Free:", diskStat.UsageStat.Free)
// 		fmt.Println("Fstype:", diskStat.UsageStat.Fstype)
// 		fmt.Println("InodesFree:", diskStat.UsageStat.InodesFree)
// 		fmt.Println("InodesTotal:", diskStat.UsageStat.InodesTotal)
// 		fmt.Println("InodesUsed:", diskStat.UsageStat.InodesUsed)
// 		fmt.Println("InodesUsedPercent:", diskStat.UsageStat.InodesUsedPercent)
// 		fmt.Println("Path:", diskStat.UsageStat.Path)
// 		fmt.Println("Total:", diskStat.UsageStat.Total)
// 		fmt.Println("Used:", diskStat.UsageStat.Used)
// 		fmt.Println("UsedPercent:", diskStat.UsageStat.UsedPercent)
// 	}
// }

// func handleDaemonClusterProcessStatList(data []interface{}) {
// 	processStatList := make([]ProcessStat, len(data))
// 	for i, d := range data {
// 		var processStat ProcessStat
// 		err := mapstructure.Decode(d, &processStat)
// 		if err != nil {
// 			panic(err)
// 		}
// 		processStatList[i] = processStat
// 	}

// 	for _, processStat := range processStatList {
// 		fmt.Println("Data:", processStat.MemoryInfoExStat.Data)
// 		fmt.Println("Dirty:", processStat.MemoryInfoExStat.Dirty)
// 		fmt.Println("Lib:", processStat.MemoryInfoExStat.Lib)
// 		fmt.Println("RSS:", processStat.MemoryInfoExStat.RSS)
// 		fmt.Println("Shared:", processStat.MemoryInfoExStat.Shared)
// 		fmt.Println("Text:", processStat.MemoryInfoExStat.Text)
// 		fmt.Println("VMS:", processStat.MemoryInfoExStat.VMS)

// 		fmt.Println("Data:", processStat.MemoryInfoStat.Data)
// 		fmt.Println("HWM:", processStat.MemoryInfoStat.HWM)
// 		fmt.Println("Locked:", processStat.MemoryInfoStat.Locked)
// 		fmt.Println("RSS:", processStat.MemoryInfoStat.RSS)
// 		fmt.Println("Stack:", processStat.MemoryInfoStat.Stack)
// 		fmt.Println("Swap:", processStat.MemoryInfoStat.Swap)
// 		fmt.Println("VMS:", processStat.MemoryInfoStat.VMS)

// 		fmt.Println("Involuntary:", processStat.NumCtxSwitchesStat.Involuntary)
// 		fmt.Println("Voluntary:", processStat.NumCtxSwitchesStat.Voluntary)
// 	}
// }

// func handleRegistryDiskStatList(data []interface{}) {}

// func handleRegistryProcessStatList(data []interface{}) {}
