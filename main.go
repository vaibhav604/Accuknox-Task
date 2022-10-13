package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Create map to store entries of log file
	logs := make(map[int][]int)

	//Read log file and handle errors with file
	var logFile=os.Args[1]
	readFile, err := os.Open(logFile)
	if err != nil {
		fmt.Println("Error while opening file. Please try again!!")
		return
	}
	fi,err:=readFile.Stat()
	if err != nil {
		fmt.Println("Error while opening file. Please try again!!")
		return
	}
	fsize :=fi.Size()
	if fsize==0{
		fmt.Printf("Provided file is empty! Please provide another file.")
		return
	}

	defer readFile.Close()
	fileScanner:= bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))

	//Initialization
	var food_id int
	var user_id int

	//function to store log entries into the map and if eater_id is present with
	//same food_id returns error
	for fileScanner.Scan() {
		var entry string = fileScanner.Text()
		result := strings.Split(entry, ",")
		user_id, err = strconv.Atoi(result[0])
		food_id, err = strconv.Atoi(result[1])
		values := logs[user_id]
		if len(values) != 0 {
			for i := range values {
				if values[i] == food_id {
					fmt.Println("Error! Please provide correct logs!")
					return
				}
			}
		}
		logs[user_id] = append(logs[user_id], food_id)
	}
	
	// call function to get top 3 food_ids
	solve(logs)
}

//functio to get top 3 food_ids
func solve(logs map[int][]int) {
	//create map to store number of occurences of food_ids with its food_id
	count := make(map[int]int)

	//get number of occurences of food_id
	for _, p := range logs {
		for i := range p {
			count[p[i]] += 1
		}
	}

	//Sort count map based on count of food_id in descending order
	keys := make([]int, 0, len(count))
	for key := range count {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})

	//Print the top 3 food_ids consumed
	for i := 0; i < 3; i++ {
		fmt.Println(keys[i])
	}
}
