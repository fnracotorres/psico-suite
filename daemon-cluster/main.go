package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
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

type Message struct {
	Type      string `json:"type"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Define a welcome message for the server
var welcomeMessage = Message{
	Type:      "info",
	Sender:    "Server",
	Recipient: "All",
	Content:   "Welcome to the chat!",
}

func connect(url string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket server: %w", err)
	}
	return conn, nil
}

func main() {
	// Connect to WebSocket server
	conn, err := connect("ws://localhost:8000/ws")
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Start a goroutine to handle incoming messages from the server
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			var msg Message
			if err := conn.ReadJSON(&msg); err != nil {
				log.Println("read:", err)
				break
			}
			fmt.Printf("%s: %s\n", msg.Sender, msg.Content)
		}
	}()

	// Close connection gracefully
	err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close:", err)
	}
	wg.Wait()

	// Create a new Gorilla WebSocket router
	router := http.NewServeMux()
	router.HandleFunc("/ws", handleConnections)

	// Create a HTTP server with graceful shutdown
	server := &http.Server{
		Addr:    ":8001",
		Handler: router,
	}

	// Start the server in a separate goroutine
	go func() {
		log.Println("Server started on :8000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	// Start handling messages
	go handleMessages()

	// Handle graceful shutdown
	gracefulShutdown(server)
}

// handleConnections upgrades initial HTTP requests to WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a WebSocket
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	// Register client
	clients[ws] = true

	// Send welcome message to client
	err = ws.WriteJSON(welcomeMessage)
	if err != nil {
		log.Printf("error sending welcome message: %v", err)
		return
	}

	// Read incoming messages from client
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error reading message: %v", err)
			delete(clients, ws)
			break
		}
		handleMessage(msg)
	}
}

// handleMessage processes different types of messages
func handleMessage(msg Message) {
	switch msg.Type {
	case "text":
		broadcast <- msg
	case "command":
		// Handle command messages
		fmt.Printf("Received command from %s: %s\n", msg.Sender, msg.Content)
	default:
		fmt.Printf("Unknown message type: %s\n", msg.Type)
	}
}

// handleMessages broadcasts messages to all connected clients
func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error writing message: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// gracefulShutdown handles server shutdown gracefully
func gracefulShutdown(server *http.Server) {
	// Create a channel to listen for interrupt and terminate signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Wait for the interrupt signal
	<-stop

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Shutdown: %v", err)
	} else {
		log.Println("Server gracefully stopped")
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
