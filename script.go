package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"sync"
	"time"
)

var (
	careerOlineCsv      = "data/Career_Stats_Offensive_Line.csv"
	careerOlineData     = [][]string{}
	careerReceivingCsv  = "data/Career_Stats_Receiving.csv"
	careerReceivingData = [][]string{}
	careerDlineCsv      = "data/Game_Logs_Defensive_Lineman.csv"
	careerDlineData     = [][]string{}
	careerRushingCsv    = "data/Career_Stats_Rushing.csv"
	careerRushingData   = [][]string{}
	gameDlineCsv        = "data/Game_Logs_Defensive_Lineman.csv"
	gameDlineData       = [][]string{}
	gameOlineCsv        = "data/Game_Logs_Offensive_Line.csv"
	gameOlineData       = [][]string{}
	gameRbCsv           = "data/Game_Logs_Runningback.csv"
	gameRbData          = [][]string{}
	gameWrTeCsv         = "data/Game_Logs_Wide_Receiver_and_Tight_End.csv"
	gameWrTeData        = [][]string{}

	// column headers for the final data set
	mergedHeaders = []string{
		"player_id",
		"player_name",
		"position",
		"nfl_year",
		"season",
		"week",
		"game_date",
		"home_or_away",
		"opponent",
		"outcome",
		"score",
		"games_played",
		"games_started",
		"receptions",
		"receiving_yards",
		"yards_per_reception",
		"longest_reception",
		"receiving_tds",
		"rushing_attempts",
		"rushing_yards",
		"yards_per_carry",
		"longest_rushing_run",
		"rushing_tds",
		"fumbles",
		"fumbles_lost",
	}

	mergedData = [][]string{}
)

func main() {
	start := time.Now()
	log.Println("starting pre-processing script at " + start.String())

	var w sync.WaitGroup
	w.Add(8)
	go func() {
		careerOlineData = ReadCsv(careerOlineCsv)
		ConvertNan(careerOlineData)
		w.Done()
	}()

	go func() {
		careerReceivingData = ReadCsv(careerReceivingCsv)
		ConvertNan(careerReceivingData)
		w.Done()
	}()

	go func() {
		careerDlineData = ReadCsv(careerDlineCsv)
		ConvertNan(careerDlineData)
		w.Done()
	}()

	go func() {
		careerRushingData = ReadCsv(careerRushingCsv)
		ConvertNan(careerRushingData)
		w.Done()
	}()

	go func() {
		gameDlineData = ReadCsv(gameDlineCsv)
		ConvertNan(gameDlineData)
		w.Done()
	}()

	go func() {
		gameOlineData = ReadCsv(gameOlineCsv)
		ConvertNan(gameOlineData)
		mergedHeaders = append(mergedHeaders, gameOlineData[1][0])
		w.Done()
	}()
	go func() {
		gameRbData = ReadCsv(gameRbCsv)
		ConvertNan(gameRbData)
		gameRbData = removeHeaders(gameRbData, 0)
		mergedData = append(mergedData, gameRbData...)
		w.Done()
	}()
	go func() {
		gameWrTeData = ReadCsv(gameWrTeCsv)
		ConvertNan(gameWrTeData)
		gameWrTeData = removeHeaders(gameWrTeData, 0)
		mergedData = append(mergedData, gameWrTeData...)
		w.Done()
	}()
	w.Wait()
	log.Println(gameRbData[0])

	log.Println("all csvs are in dude - writing merged data csv")
	var formattedData [][]string
	formattedData = append(formattedData, mergedHeaders)
	formattedData = append(formattedData, mergedData...)

	WriteCsv("merged-data.csv", formattedData)

	elapsed := time.Since(start)
	log.Println("completed execution in " + elapsed.String())

}

func ReadCsv(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func WriteCsv(filename string, data [][]string) {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()
	w := bufio.NewWriter(outFile)

	writer := csv.NewWriter(outFile)
	err = writer.WriteAll(data)
	if err != nil {
		log.Fatal(err)
	}

	w.Flush()
}

func removeHeaders(slice [][]string, s int) [][]string {
	return append(slice[:s], slice[s+1:]...)
}

func ConvertNan(csv [][]string) {
	for _, row := range csv {
		for _, cell := range row {
			if cell == "--" {
				cell = "NaN"
			}
		}
	}
}
