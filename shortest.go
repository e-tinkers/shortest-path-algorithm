package main

import (
	"fmt"
	"time"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Array []string

func (arr Array) hasPropertyOf(str string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func ReadStationsData () map[string][]string {
	var stations map[string][]string

	jsonFile, err := os.Open("./data/stations_sg.json")
	defer jsonFile.Close()
	if err != nil {
	    fmt.Println(err)
	}

	jsonBytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(jsonBytes), &stations)
	    return stations
}

func ShortestPath(graph map[string][]string, start string, end string, path Array) []string {

	if _, exist := graph[start]; !exist {
		return path
	}

	path = append(path, start)
	if start == end {
		return path
	}

	shortest := make([]string, 0)
        for _, node := range graph[start] {
		if !path.hasPropertyOf(node) {
			newPath := ShortestPath(graph, node, end, path)
			if len(newPath) > 0 {
				if (len(shortest) == 0 || (len(newPath) < len(shortest))) {
					shortest = newPath
				}
			}
		}
	}
	return shortest
}

func main() {

	graph := ReadStationsData()
	var s string = "ns1/ew24"
	var e string = "ew2/dt32"
	var path = make([]string, 0 , 50)

	startTime := time.Now()
	shortestPath := ShortestPath(graph, s, e, path)
	elapsed := time.Since(startTime)
	fmt.Println(shortestPath)
	fmt.Println(elapsed)

}
