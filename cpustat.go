package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type clocks struct {
	lsClock  []int
	maxClock int `require: min 0`
	avgClock int `require: min 0`
}

func main() {
	stat := clocks{
		maxClock: 0,
		avgClock: 0,
	}
	for {
		stat.lsClock = []int{}
		cpu(&stat)
		cpuPrint(&stat)
		time.Sleep(time.Millisecond * 500)
	}
}

// Given a file and a string to look for, find all lines
// in the file containing the search string as a substring
// Returns a pointer to the array of strings and an error field
func badGrep(file string, exp string) (*[]string, error) {
	var x []string
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, exp) {
			x = append(x, line)
		}
	}
	return &x, nil
}

// Process data obtained through badGrep
func calcMHz(arr *[]string, stat *clocks) {
	s := 0
	for _, v := range *arr {
		curr := strings.Split(v, ": ")[1]
		curr = strings.Split(curr, ".")[0]
		i, err := strconv.Atoi(curr)
		if err != nil {
			log.Fatal(err)
		}
		if i >= stat.maxClock {
			stat.maxClock = i
		}
		s += i
		stat.lsClock = append(stat.lsClock, i)
	}
	stat.avgClock = s / len(stat.lsClock)
}

// Read Clock information from /proc/cpuinfo
// and have calcMHZ() parse it.
func cpu(stat *clocks) {
	var (
		file string = "/proc/cpuinfo"
		exp  string = "cpu MHz"
	)
	arr, err := badGrep(file, exp)
	if err != nil {
		panic("something went wrong when reading /proc/cpuinfo")
	}
	calcMHz(arr, stat)
}

// clear terminal (linux only)
func callClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Print info to stdout
func cpuPrint(stat *clocks) {
	callClear()
	println("Max Clock: ", stat.maxClock)
	println("Avg Clock: ", stat.avgClock)
	println("\nClocks:")
	for _, v := range stat.lsClock {
		println(v)
	}
}
