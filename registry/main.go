package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/shirou/gopsutil/disk"
)

type DiskStat struct {
	IOCountersStats map[string]disk.IOCountersStat `json:"ioCountersStats"`
	PartitionStat   disk.PartitionStat             `json:"partitionStat"`
	UsageStat       disk.UsageStat                 `json:"usageStat"`
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	var diskStats []DiskStat
	decoder := json.NewDecoder(c)
	if err := decoder.Decode(&diskStats); err != nil {
		panic(err)
	}

	for _, diskStat := range diskStats {
		fmt.Println("Device:", diskStat.PartitionStat.Device)
		fmt.Println("Fstype:", diskStat.PartitionStat.Fstype)
		fmt.Println("Mountpoint:", diskStat.PartitionStat.Mountpoint)
		fmt.Println("Opts:", diskStat.PartitionStat.Opts)

		for _, ioCountersStat := range diskStat.IOCountersStats {
			fmt.Println("IoTime:", ioCountersStat.IoTime)
			fmt.Println("IopsInProgress:", ioCountersStat.IopsInProgress)
			fmt.Println("Label:", ioCountersStat.Label)
			fmt.Println("MergedReadCount:", ioCountersStat.MergedReadCount)
			fmt.Println("MergedWriteCount:", ioCountersStat.MergedWriteCount)
			fmt.Println("Name:", ioCountersStat.Name)
			fmt.Println("ReadBytes:", ioCountersStat.ReadBytes)
			fmt.Println("ReadCount:", ioCountersStat.ReadCount)
			fmt.Println("ReadTime:", ioCountersStat.ReadTime)
			fmt.Println("SerialNumber:", ioCountersStat.SerialNumber)
			fmt.Println("WeightedIO:", ioCountersStat.WeightedIO)
			fmt.Println("WriteBytes:", ioCountersStat.WriteBytes)
			fmt.Println("WriteCount:", ioCountersStat.WriteCount)
			fmt.Println("WriteTime:", ioCountersStat.WriteTime)
		}

		fmt.Println("Free:", diskStat.UsageStat.Free)
		fmt.Println("Fstype:", diskStat.UsageStat.Fstype)
		fmt.Println("InodesFree:", diskStat.UsageStat.InodesFree)
		fmt.Println("InodesTotal:", diskStat.UsageStat.InodesTotal)
		fmt.Println("InodesUsed:", diskStat.UsageStat.InodesUsed)
		fmt.Println("InodesUsedPercent:", diskStat.UsageStat.InodesUsedPercent)
		fmt.Println("Path:", diskStat.UsageStat.Path)
		fmt.Println("Total:", diskStat.UsageStat.Total)
		fmt.Println("Used:", diskStat.UsageStat.Used)
		fmt.Println("UsedPercent:", diskStat.UsageStat.UsedPercent)

		fmt.Println()
		fmt.Println("-----")
		fmt.Println()
	}
}

// CPU

// fmt.Println("CPU:", infoStat.CPU)
// fmt.Println("VendorID:", infoStat.VendorID)
// fmt.Println("Family:", infoStat.Family)
// fmt.Println("Model:", infoStat.Model)
// fmt.Println("Stepping:", infoStat.Stepping)
// fmt.Println("PhysicalID:", infoStat.PhysicalID)
// fmt.Println("CoreID:", infoStat.CoreID)
// fmt.Println("Cores:", infoStat.Cores)
// fmt.Println("ModelName:", infoStat.ModelName)
// fmt.Println("Mhz:", infoStat.Mhz)
// fmt.Println("CacheSize:", infoStat.CacheSize)
// fmt.Println("Flags:", infoStat.Flags)
// fmt.Println("Microcode:", infoStat.Microcode)

// fmt.Println("CPU:", timesStat.CPU)
// fmt.Println("User:", timesStat.User)
// fmt.Println("System:", timesStat.System)
// fmt.Println("Idle:", timesStat.Idle)
// fmt.Println("Nice:", timesStat.Nice)
// fmt.Println("Iowait:", timesStat.Iowait)
// fmt.Println("Irq:", timesStat.Irq)
// fmt.Println("Softirq:", timesStat.Softirq)
// fmt.Println("Steal:", timesStat.Steal)
// fmt.Println("Guest:", timesStat.Guest)
// fmt.Println("GuestNice:", timesStat.GuestNice)

//

// Host

// fmt.Println("Hostname:", infoStat.Hostname)
// fmt.Println("Uptime:", infoStat.Uptime)
// fmt.Println("BootTime:", infoStat.BootTime)
// fmt.Println("Procs:", infoStat.Procs)
// fmt.Println("OS:", infoStat.OS)
// fmt.Println("Platform:", infoStat.Platform)
// fmt.Println("PlatformFamily:", infoStat.PlatformFamily)
// fmt.Println("PlatformVersion:", infoStat.PlatformVersion)
// fmt.Println("KernelVersion:", infoStat.KernelVersion)
// fmt.Println("KernelArch:", infoStat.KernelArch)
// fmt.Println("VirtualizationSystem:", infoStat.VirtualizationSystem)
// fmt.Println("VirtualizationRole:", infoStat.VirtualizationRole)
// fmt.Println("HostID:", infoStat.HostID)

// fmt.Println("SensorKey:", temperatureStat.SensorKey)
// fmt.Println("Temperature:", temperatureStat.Temperature)

// fmt.Println("User:", userStat.User)
// fmt.Println("Terminal:", userStat.Terminal)
// fmt.Println("Host:", userStat.Host)
// fmt.Println("Started:", userStat.Started)
