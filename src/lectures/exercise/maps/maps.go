//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(servers map[string]int) {
	fmt.Println("==============================================================")
	fmt.Printf("There are %d servers\n", len(servers))

	stats := make(map[int]int)

	for name, status := range servers {
		fmt.Printf("%s is ", name)
		switch status {
		case Online:
			fmt.Printf("online.\n")
			stats[Online] += 1
		case Offline:
			fmt.Printf("offline.\n")
			stats[Offline] += 1
		case Maintenance:
			fmt.Printf("on maintenance.\n")
			stats[Maintenance] += 1
		case Retired:
			fmt.Printf("retired.\n")
			stats[Retired] += 1
		default:
			fmt.Printf("unknown.\n")
			panic("unhandled server status")
		}
	}

	fmt.Println("\nSummary:")
	fmt.Println("\tservers online:", stats[Online])
	fmt.Println("\tservers offline:", stats[Offline])
	fmt.Println("\tservers under maintenance:", stats[Maintenance])
	fmt.Println("\tservers retired:", stats[Retired])
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatusMap := make(map[string]int)

	for _, v := range servers {
		serverStatusMap[v] = 0
	}
	printServerStatus(serverStatusMap)

	serverStatusMap["darkstar"] = Retired
	serverStatusMap["aiur"] = Offline
	printServerStatus(serverStatusMap)

	for k := range serverStatusMap {
		serverStatusMap[k] = Maintenance
	}
	printServerStatus(serverStatusMap)

	delete(serverStatusMap, "w359")
	fmt.Println("\nremoved w359 server")
	printServerStatus(serverStatusMap)
}
