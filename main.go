package main

import(
	"fmt"
	"os"
	"strconv"
	"time"
	"math"
	"strings"
)

const filledIndicator = "█"
const emptyIndicator = "░"

const capacityLevelFile = "/sys/class/power_supply/BATT/capacity"

const red = "\033[1;31m"
const green = "\033[1;32m"
const white = "\033[1;37m"
const reset = "\033[00m"

const columns = 20

func main(){
	var levelColor string
	filled, empty := createBar()
	for{
		level := getChargeLevel()
		switch{
			case level <= 20:
				levelColor = red
				break;
			default:
				levelColor = green
		}
		fmt.Printf("\033[2J \033[H \033[?25l")
		fmt.Printf("%s%s%s%d%%%s\r", levelColor, filled, empty, level, reset)
		time.Sleep(10 * time.Second)
	}
}

func createBar()(string, string){
	filledColumns := math.Round(float64(columns) * float64(getChargeLevel()) / 100)
	emptyColumns := float64(columns) - filledColumns

	filledColumnsStr := strings.Repeat(filledIndicator, int(filledColumns))
	emptyColumnsStr := strings.Repeat(emptyIndicator, int(emptyColumns))
	
	return filledColumnsStr, emptyColumnsStr
}

func getChargeLevel()(int){
	data, readErr := os.ReadFile(capacityLevelFile)
	if readErr != nil{
		fmt.Printf("%s%s%s", red, readErr, reset)
		os.Exit(1)
	}
	
	strData := string(data)
	level, _ := strconv.Atoi(strData[0:len(strData)-1])

	return level
}
