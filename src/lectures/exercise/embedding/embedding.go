//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) average() int {
	var sum = 0
	for _, b := range b.amount {
		sum += int(b)
	}

	return sum / len(b.amount)
}

type CpuTemp struct {
	temp []Celcius
}

func (c *CpuTemp) average() float64 {
	var sum = 0.0
	for _, b := range c.temp {
		sum += float64(b)
	}

	return sum / float64(len(c.temp))
}

type MemoryUsage struct {
	amount []Bytes
}

func (m *MemoryUsage) average() int {
	var sum = 0
	for _, a := range m.amount {
		sum += int(a)
	}

	return sum / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func (d *Dashboard) printAverageUsage() {
	fmt.Println()
	fmt.Println("Average Bandwidth usage:", d.BandwidthUsage.average())
	fmt.Println("Average CPU temperautre:", d.CpuTemp.average())
	fmt.Println("Average Memory usage:", d.MemoryUsage.average())
	fmt.Println()
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}
	dashboard.printAverageUsage()
}
