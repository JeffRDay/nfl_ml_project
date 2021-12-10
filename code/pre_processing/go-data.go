package main

import "sync"

func LoadData() {
	var w sync.WaitGroup
	w.Add(7)
	go func() {
		basicStatsData = ReadCsv(basicStatsCsv)
		basicStatsData = RemoveHeaders(basicStatsData, 0)
		ConvertNan(basicStatsData)
		w.Done()
	}()
	go func() {
		careerReceivingData = ReadCsv(careerReceivingCsv)
		careerReceivingData = RemoveHeaders(careerReceivingData, 0)
		ConvertNan(careerReceivingData)
		w.Done()
	}()
	go func() {
		careerRushingData = ReadCsv(careerRushingCsv)
		careerRushingData = RemoveHeaders(careerRushingData, 0)
		ConvertNan(careerRushingData)
		w.Done()
	}()
	go func() {
		careerPassingData = ReadCsv(careerPassingCsv)
		careerPassingData = RemoveHeaders(careerPassingData, 0)
		ConvertNan(careerPassingData)
		w.Done()
	}()
	go func() {
		gameRbData = ReadCsv(gameRbCsv)
		gameRbData = RemoveHeaders(gameRbData, 0)
		ConvertNan(gameRbData)
		w.Done()
	}()
	go func() {
		gameWrTeData = ReadCsv(gameWrTeCsv)
		gameWrTeData = RemoveHeaders(gameWrTeData, 0)
		ConvertNan(gameWrTeData)
		w.Done()
	}()
	go func() {
		gameQbData = ReadCsv(gameQbCsv)
		gameQbData = RemoveHeaders(gameQbData, 0)
		ConvertNan(gameQbData)
		w.Done()
	}()
	w.Wait()

}
