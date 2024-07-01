package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

// Flight结构体
type Flight struct {
	FlightNumber string  `json:"flightNumber"`
	Origin       string  `json:"origin"`
	Destination  string  `json:"destination"`
	Departure    string  `json:"departure"`
	Arrival      string  `json:"arrival"`
	PlaneModel   string  `json:"planeModel"`
	Price        float64 `json:"price"`
}

// FlightData结构体
type FlightData struct {
	flights []Flight
}

// 添加航班记录方法
func (fd *FlightData) Add(flight Flight) {
	fd.flights = append(fd.flights, flight)
}

// 按航班号排序方法
func (fd *FlightData) SortFlightNum() {
	sort.Slice(fd.flights, func(i, j int) bool {
		return fd.flights[i].FlightNumber < fd.flights[j].FlightNumber
	})
}

// 按航班号折半查找方法
func (fd *FlightData) SearchBS(flightNumber string) int {
	left, right := 0, len(fd.flights)-1
	for left <= right {
		mid := left + (right-left)/2
		if fd.flights[mid].FlightNumber == flightNumber {
			return mid
		}
		if fd.flights[mid].FlightNumber < flightNumber {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 按指定字段顺序查找方法
func (fd *FlightData) SearchField(field string, value string) []Flight {
	var result []Flight
	for _, flight := range fd.flights {
		switch field {
		case "origin":
			if flight.Origin == value {
				result = append(result, flight)
			}
		case "destination":
			if flight.Destination == value {
				result = append(result, flight)
			}
		case "departure":
			if flight.Departure == value {
				result = append(result, flight)
			}
		case "arrival":
			if flight.Arrival == value {
				result = append(result, flight)
			}
		}
	}
	return result
}

// 格式化输出为JSON的方法
func Format(data any) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("JSON编码错误: %s", err)
		//return fmt.Sprintf("JSON格式化错误: %s", err)
	}
	return string(jsonData)
}

// 主函数
func main() {
	// 初始化FlightData实例
	fd := FlightData{}

	// 添加航班信息
	fd.Add(Flight{"AA101", "Beijing", "Shanghai", "08:00", "10:00", "Boeing 737", 500.0})
	fd.Add(Flight{"AA102", "Shanghai", "Beijing", "09:00", "11:00", "Airbus A320", 550.0})
	fd.Add(Flight{"AA103", "Beijing", "Guangzhou", "07:00", "09:00", "Boeing 737", 600.0})
	fd.Add(Flight{"AA104", "Guangzhou", "Beijing", "06:00", "08:00", "Airbus A330", 650.0})
	fd.Add(Flight{"AA105", "Beijing", "Shenzhen", "05:00", "07:00", "Boeing 737", 700.0})

	// // 添加航班信息(乱序)
	// fd.Add(Flight{"AA101", "Beijing", "Shanghai", "08:00", "10:00", "Boeing 737", 500.0})

	// fd.Add(Flight{"AA105", "Beijing", "Shenzhen", "05:00", "07:00", "Boeing 737", 700.0})

	// fd.Add(Flight{"AA102", "Shanghai", "Beijing", "09:00", "11:00", "Airbus A320", 550.0})
	// fd.Add(Flight{"AA103", "Beijing", "Guangzhou", "07:00", "09:00", "Boeing 737", 600.0})
	// fd.Add(Flight{"AA104", "Guangzhou", "Beijing", "06:00", "08:00", "Airbus A330", 650.0})

	// 按航班号排序
	fd.SortFlightNum()

	// 按航班号折半查找
	flightNumber := "AA103"
	index := fd.SearchBS(flightNumber)
	if index != -1 {
		fmt.Printf("找到航班 %s: %s\n", flightNumber, Format(fd.flights[index]))
	} else {
		fmt.Printf("未找到航班 %s\n", flightNumber)
	}

	// // 按起点站顺序查找
	// originBJ := "Beijing"
	// flightsFromBeijing := fd.SearchField("origin", originBJ)
	// fmt.Printf("从 %s 出发的航班: %s\n", originBJ, Format(flightsFromBeijing))

	// // 按起点站顺序查找
	// originSH := "Shanghai"
	// flightsFromShanghai := fd.SearchField("origin", originSH)
	// fmt.Printf("从 %s 出发的航班: %s\n", originSH, Format(flightsFromShanghai))

	// // 按起点站顺序查找
	// originSH := "Xi'an"
	// flightsFromShanghai := fd.SearchField("origin", originSH)
	// fmt.Printf("从 %s 出发的航班: %s\n", originSH, Format(flightsFromShanghai))

	// 按起点站顺序查找
	originSH := "上海% 。 、 wtf 123457 😊"
	flightsFromShanghai := fd.SearchField("origin", originSH)
	fmt.Printf("从 %s 出发的航班: %s\n", originSH, Format(flightsFromShanghai))

	// // 按终点站顺序查找
	// destination := "Beijing"
	// flightsToBeijing := fd.SearchField("destination", destination)
	// fmt.Printf("到达 %s 的航班: %s\n", destination, Format(flightsToBeijing))
}
