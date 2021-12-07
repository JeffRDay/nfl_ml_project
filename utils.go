package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func GetPlayerId(s string) string {
	return strings.Split(s, "/")[1]
}

func ReadCsv(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	slices, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return slices

}

func WriteCsv(filename string, data [][]string) {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()
	// w := bufio.NewWriter(outFile)

	log.Println("writing csv - hold yur breath")
	start := time.Now()
	writer := csv.NewWriter(outFile)
	// err = writer.WriteAll(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	size := len(data) - 1
	for index, v := range data {
		if index%10000 == 0 {
			log.Printf("Writing %v of %v | elapsed time - %v", index, size, time.Since(start))
		}
		writer.Write(v)
		writer.Flush()
	}
	// w.Flush()
}

func Write1DCsv(filename string, data []string, g *sync.WaitGroup) {
	defer g.Done()
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()
	w := bufio.NewWriter(outFile)
	defer w.Flush()

	writer := csv.NewWriter(outFile)
	err = writer.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	defer writer.Flush()

}

func RemoveHeaders(slice [][]string, s int) [][]string {
	return append(slice[:s], slice[s+1:]...)
}

func ConvertNan(csv [][]string) {
	for r, row := range csv {
		for c, cell := range row {
			if cell == "--" {
				csv[r][c] = "NaN"
			}
		}
	}
}
