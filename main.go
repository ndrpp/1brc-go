package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Measurement struct {
	min   float64
	max   float64
	count float64
	sum   float64
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
            fmt.Println("Duplicate entry!", station)
		} else {
			hashmap[station] = &Measurement{
				min:   value,
				max:   value,
				count: 1,
				sum:   value,
			}
            
            fmt.Println("obj: ",station,"-", hashmap[station])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	printOutput(hashmap)
}

func parseLine(line string) (string, float64) {
	name := strings.Split(line, ";")[0]
	value, err := strconv.ParseFloat(strings.Split(line, ";")[1], 8)
	if err != nil {
		fmt.Println(err)
	}
	return name, value
}

func printOutput(hashmap map[string]*Measurement) {
	for key, val := range hashmap {
		fmt.Printf("%s:%f/%f/%f\n", key, val.min, val.sum/val.count, val.max)
	}
}
