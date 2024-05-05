package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Measurement struct {
	min   int64
	max   int64
	count int64
	sum   int64
}

func main() {
	hashmap := make(map[string]*Measurement)

	file, err := os.Open("measurements.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		station, value := parseLine(scanner.Text())

		if hashmap[station] != nil {
			obj := hashmap[station]
			obj.count += 1
			obj.sum += value
			if value > obj.max {
				obj.max = value
			} else if value < obj.min {
				obj.min = value
			}
			hashmap[station] = obj
		} else {
			hashmap[station] = &Measurement{
				min:   value,
				max:   value,
				count: 1,
				sum:   value,
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	printOutput(hashmap)
}

func parseLine(line string) (string, int64) {
	arr := strings.Split(line, ";")
	name := arr[0]
	temp := arr[1]
	value, _ := strconv.ParseInt((temp[:len(temp)-2] + temp[len(temp)-1:]), 10, 64)
	return name, value
}

func printOutput(hashmap map[string]*Measurement) {
	for key, val := range hashmap {
		fmt.Printf("%s:%.1f/%.1f/%.1f\n", key, float64(val.min)/10.0, float64(val.sum)/float64(val.count)/10.0, float64(val.max)/10.0)
	}
}
