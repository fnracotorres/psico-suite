package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

type CPUStat struct {
	InfoStat  cpu.InfoStat  `json:"infoStat"`
	TimesStat cpu.TimesStat `json:"timeStat"`
}

type DiskStat struct {
	IOCountersStats map[string]disk.IOCountersStat `json:"ioCountersStats"`
	PartitionStat   disk.PartitionStat             `json:"partitionStat"`
	UsageStat       disk.UsageStat                 `json:"usageStat"`
}

type HostStat struct {
	InfoStat        host.InfoStat        `json:"infoStat"`
	TemperatureStat host.TemperatureStat `json:"temperatureStat"`
	UserStat        host.UserStat        `json:"userStat"`
}

func main() {
	c, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		panic(err)
	}

	writeCPU() // Not writing yet
	writeDiskStat(c)
	writeHostStat() // Not writing yet
}

func writeCPU() {
	infoStats, err := cpu.Info()
	if err != nil {
		panic(err)
	}

	for _, infoStat := range infoStats {
		fmt.Println("CPU:", infoStat.CPU)
		fmt.Println("VendorID:", infoStat.VendorID)
		fmt.Println("Family:", infoStat.Family)
		fmt.Println("Model:", infoStat.Model)
		fmt.Println("Stepping:", infoStat.Stepping)
		fmt.Println("PhysicalID:", infoStat.PhysicalID)
		fmt.Println("CoreID:", infoStat.CoreID)
		fmt.Println("Cores:", infoStat.Cores)
		fmt.Println("ModelName:", infoStat.ModelName)
		fmt.Println("Mhz:", infoStat.Mhz)
		fmt.Println("CacheSize:", infoStat.CacheSize)
		fmt.Println("Flags:", infoStat.Flags)
		fmt.Println("Microcode:", infoStat.Microcode)
	}

	timesStats, err := cpu.Times(true)
	if err != nil {
		panic(err)
	}

	for _, timesStat := range timesStats {
		fmt.Println("CPU:", timesStat.CPU)
		fmt.Println("User:", timesStat.User)
		fmt.Println("System:", timesStat.System)
		fmt.Println("Idle:", timesStat.Idle)
		fmt.Println("Nice:", timesStat.Nice)
		fmt.Println("Iowait:", timesStat.Iowait)
		fmt.Println("Irq:", timesStat.Irq)
		fmt.Println("Softirq:", timesStat.Softirq)
		fmt.Println("Steal:", timesStat.Steal)
		fmt.Println("Guest:", timesStat.Guest)
		fmt.Println("GuestNice:", timesStat.GuestNice)
	}
}

func writeDiskStat(c net.Conn) {
	diskStats := []DiskStat{}

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
			IOCountersStats: ioCountersStats,
			PartitionStat:   partitionStat,
			UsageStat:       *usageStat,
		}

		diskStats = append(diskStats, diskStat)
	}

	b, err := json.Marshal(diskStats)
	if err != nil {
		panic(err)
	}

	_, err = c.Write(b)
	if err != nil {
		panic(err)
	}
}

func writeHostStat() {
	infoStat, err := host.Info()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hostname:", infoStat.Hostname)
	fmt.Println("Uptime:", infoStat.Uptime)
	fmt.Println("BootTime:", infoStat.BootTime)
	fmt.Println("Procs:", infoStat.Procs)
	fmt.Println("OS:", infoStat.OS)
	fmt.Println("Platform:", infoStat.Platform)
	fmt.Println("PlatformFamily:", infoStat.PlatformFamily)
	fmt.Println("PlatformVersion:", infoStat.PlatformVersion)
	fmt.Println("KernelVersion:", infoStat.KernelVersion)
	fmt.Println("KernelArch:", infoStat.KernelArch)
	fmt.Println("VirtualizationSystem:", infoStat.VirtualizationSystem)
	fmt.Println("VirtualizationRole:", infoStat.VirtualizationRole)
	fmt.Println("HostID:", infoStat.HostID)

	temperatureStats, err := host.SensorsTemperatures()
	if err != nil {
		panic(err)
	}

	for _, temperatureStat := range temperatureStats {
		fmt.Println("SensorKey:", temperatureStat.SensorKey)
		fmt.Println("Temperature:", temperatureStat.Temperature)
	}

	userStats, err := host.Users()
	if err != nil {
		panic(err)
	}

	for _, userStat := range userStats {
		fmt.Println("User:", userStat.User)
		fmt.Println("Terminal:", userStat.Terminal)
		fmt.Println("Host:", userStat.Host)
		fmt.Println("Started:", userStat.Started)
	}
}
