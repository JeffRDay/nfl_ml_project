package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type Player struct {
	Pid           string
	SeasonTeamMap map[string]string
}

var (
	mapPlayerToTeam    = make(map[string]Player)
	teamsMap           = make(map[string]string)
	opponentsMap       = make(map[string]string)
	mapPlayerBirthYear = make(map[string]string)
	mapPlayerPosition  = make(map[string]string)

	basicStatsCsv       = "data/Basic_Stats.csv"
	basicStatsData      = [][]string{}
	careerReceivingCsv  = "data/Career_Stats_Receiving.csv"
	careerReceivingData = [][]string{}
	careerRushingCsv    = "data/Career_Stats_Rushing.csv"
	careerRushingData   = [][]string{}
	careerPassingCsv    = "data/Career_Stats_Passing.csv"
	careerPassingData   = [][]string{}
	gameRbCsv           = "data/Game_Logs_Runningback.csv"
	gameRbData          = [][]string{}
	gameWrTeCsv         = "data/Game_Logs_Wide_Receiver_and_Tight_End.csv"
	gameWrTeData        = [][]string{}
	gameQbCsv           = "data/Game_Logs_Quarterback.csv"
	gameQbData          = [][]string{}

	// column headers for the final data set
	mergedHeaders = []string{}

	mergedDataMap = map[int]Record{}
	mergedData    = [][]string{}
)

func main() {
	start := time.Now()
	log.Println("starting pre-processing script at " + start.String())

	// LOAD & CLEAN DATA
	LoadData()
	UpdateMaps()
	SetHeaders()
	log.Println(len(gameRbData))
	AddRecords()

	// PRINT STATS
	log.Printf("STATS ::: Features Included ::: count=%v", len(mergedDataMap[1].Data))
	log.Printf("STATS ::: Records Included ::: count=%v", len(mergedDataMap))
	// log.Printf("STATS ::: Records Not Included (Missing Team Data) ::: count=%v", count)
	log.Printf("STATS ::: Oppenents Map ::: len=%v", len(opponentsMap))

	// mergedData
	// count := 0
	mergedData = append(mergedData, mergedHeaders)
	for _, v := range mergedDataMap {
		// if count > 2000 {
		// 	break
		// }
		// count++
		mergedData = append(mergedData, strings.Split(v.String(), ","))
	}

	WriteCsv("test.csv", mergedData)

	elapsed := time.Since(start)
	log.Println("completed execution in " + elapsed.String())
}

func UpdateMaps() {
	UpdateTeamsMaps(careerReceivingData)
	UpdateTeamsMaps(careerRushingData)
	UpdateTeamsMaps(careerPassingData)
	UpdateOppenentsMap(gameRbData)
	UpdateOppenentsMap(gameWrTeData)
	UpdateBasicMaps(basicStatsData)
}

func UpdateTeamsMaps(csv [][]string) {
	for _, row := range csv {
		// Update map of Teams in Data set
		teamsMap[row[4]] = row[3]

		// update map of players -> Player
		pid := strings.Split(row[0], "/")[1]
		value, ok := mapPlayerToTeam[pid]
		if ok {
			value.SeasonTeamMap[row[3]] = row[4]
		} else {
			mapPlayerToTeam[pid] = Player{
				Pid:           pid,
				SeasonTeamMap: make(map[string]string),
			}
			value = mapPlayerToTeam[pid]
			value.SeasonTeamMap[row[3]] = row[4]
		}
	}
}

func UpdateOppenentsMap(csv [][]string) {
	for _, row := range csv {
		// Update map of Teams in Data set
		opponentsMap[row[8]] = row[8]
	}
}

func UpdateBasicMaps(csv [][]string) {
	for _, row := range csv {
		pid := GetPlayerId(row[12])
		birthData := strings.Split(row[2], "/")

		if len(birthData) < 2 {
			continue
		}

		mapPlayerBirthYear[pid] = birthData[2]
		mapPlayerPosition[pid] = row[B_POSITION]
	}
}

func removeSpecChar(s string) string {
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, ".", "")
	return s
}

func SetHeaders() {
	teams := []string{}
	for k, _ := range teamsMap {
		teams = append(teams, removeSpecChar(k))
	}

	opponents := []string{}
	for k, _ := range opponentsMap {
		opponents = append(opponents, k)
	}

	mergedHeaders = append(mergedHeaders, "player_id")
	mergedHeaders = append(mergedHeaders, "age_under_20")
	mergedHeaders = append(mergedHeaders, "age_20")
	mergedHeaders = append(mergedHeaders, "age_21")
	mergedHeaders = append(mergedHeaders, "age_22")
	mergedHeaders = append(mergedHeaders, "age_23")
	mergedHeaders = append(mergedHeaders, "age_24")
	mergedHeaders = append(mergedHeaders, "age_25")
	mergedHeaders = append(mergedHeaders, "age_26")
	mergedHeaders = append(mergedHeaders, "age_27")
	mergedHeaders = append(mergedHeaders, "age_28")
	mergedHeaders = append(mergedHeaders, "age_29")
	mergedHeaders = append(mergedHeaders, "age_30")
	mergedHeaders = append(mergedHeaders, "age_31")
	mergedHeaders = append(mergedHeaders, "age_32")
	mergedHeaders = append(mergedHeaders, "age_33")
	mergedHeaders = append(mergedHeaders, "age_34")
	mergedHeaders = append(mergedHeaders, "age_35")
	mergedHeaders = append(mergedHeaders, "age_36")
	mergedHeaders = append(mergedHeaders, "age_37")
	mergedHeaders = append(mergedHeaders, "age_38")
	mergedHeaders = append(mergedHeaders, "age_39")
	mergedHeaders = append(mergedHeaders, "age_40")
	mergedHeaders = append(mergedHeaders, "age_41")
	mergedHeaders = append(mergedHeaders, "age_42")
	mergedHeaders = append(mergedHeaders, "age_43")
	mergedHeaders = append(mergedHeaders, "age_44")
	mergedHeaders = append(mergedHeaders, "age_45")
	mergedHeaders = append(mergedHeaders, "age_over_45")
	mergedHeaders = append(mergedHeaders, teams...)
	mergedHeaders = append(mergedHeaders, opponents...)
	mergedHeaders = append(mergedHeaders, "game_month_08")
	mergedHeaders = append(mergedHeaders, "game_month_09")
	mergedHeaders = append(mergedHeaders, "game_month_10")
	mergedHeaders = append(mergedHeaders, "game_month_11")
	mergedHeaders = append(mergedHeaders, "game_month_12")
	mergedHeaders = append(mergedHeaders, "game_month_01")
	mergedHeaders = append(mergedHeaders, "game_month_02")
	mergedHeaders = append(mergedHeaders, "outcome_win")
	mergedHeaders = append(mergedHeaders, "outcome_tie")
	mergedHeaders = append(mergedHeaders, "outcome_loss")
	mergedHeaders = append(mergedHeaders, "points_scored")
	mergedHeaders = append(mergedHeaders, "rb")
	mergedHeaders = append(mergedHeaders, "fb")
	mergedHeaders = append(mergedHeaders, "wr")
	mergedHeaders = append(mergedHeaders, "te")
	mergedHeaders = append(mergedHeaders, "qb")
	mergedHeaders = append(mergedHeaders, "nfl_year")
	mergedHeaders = append(mergedHeaders, "pre_season")
	mergedHeaders = append(mergedHeaders, "regular_season")
	mergedHeaders = append(mergedHeaders, "post_season")
	mergedHeaders = append(mergedHeaders, "pro_bowl")
	mergedHeaders = append(mergedHeaders, "week")
	mergedHeaders = append(mergedHeaders, "home")
	mergedHeaders = append(mergedHeaders, "away")
	mergedHeaders = append(mergedHeaders, "receptions")
	mergedHeaders = append(mergedHeaders, "receiving_yards")
	mergedHeaders = append(mergedHeaders, "yards_per_reception")
	mergedHeaders = append(mergedHeaders, "longest_reception")
	mergedHeaders = append(mergedHeaders, "receiving_tds")
	mergedHeaders = append(mergedHeaders, "rushing_attempts")
	mergedHeaders = append(mergedHeaders, "rushing_yards")
	mergedHeaders = append(mergedHeaders, "yards_per_carry")
	mergedHeaders = append(mergedHeaders, "longest_rushing_run")
	mergedHeaders = append(mergedHeaders, "rushing_tds")
	mergedHeaders = append(mergedHeaders, "fumbles")
	mergedHeaders = append(mergedHeaders, "fumbles_lost")
	mergedHeaders = append(mergedHeaders, "passes_completed")
	mergedHeaders = append(mergedHeaders, "passes_attempted")
	mergedHeaders = append(mergedHeaders, "completion_percentage")
	mergedHeaders = append(mergedHeaders, "passing_yards")
	mergedHeaders = append(mergedHeaders, "passing_yards_per_attempt")
	mergedHeaders = append(mergedHeaders, "passing_tds")
	mergedHeaders = append(mergedHeaders, "interceptions")
	mergedHeaders = append(mergedHeaders, "sacks")
	mergedHeaders = append(mergedHeaders, "sacks_yards_lost")
	mergedHeaders = append(mergedHeaders, "passer_rating")
	mergedHeaders = append(mergedHeaders, "fs_rush_recieve_yards")
	mergedHeaders = append(mergedHeaders, "fs_passing_yards")
	mergedHeaders = append(mergedHeaders, "fs_rush_recieve_td")
	mergedHeaders = append(mergedHeaders, "fs_passing_td")
	mergedHeaders = append(mergedHeaders, "fs_interceptions")
	mergedHeaders = append(mergedHeaders, "fs_fumble")
	mergedHeaders = append(mergedHeaders, "fs_total_penalty_points")
	mergedHeaders = append(mergedHeaders, "fs_total_points_gained")
	mergedHeaders = append(mergedHeaders, "fs_total")
}

func AddRecords() {
	addRbRecords("RB Data", gameRbData)
	addWrRecords("WR & TE Data", gameWrTeData)
	addQbRecords("QB Data", gameQbData)
}

func addRbRecords(name string, csv [][]string) {
	mergedDataMapSize := len(mergedDataMap) + 1
	size := len(csv)
	// check for team
	droppedTeam := 0
	// check for birth year
	droppedBirthYear := 0
	// check for game year
	droppedGameYear := 0
	// check for position
	droppedPosition := 0

	for index, row := range csv {

		if index%10000 == 0 {
			log.Printf("%s ::: %v of %v", name, index, size)
		}

		pid := GetPlayerId(row[RB_PLAYER_ID])
		var r Record
		r.NewRecord(mergedHeaders)

		gameYear, err := strconv.Atoi(row[RB_YEAR])
		if err != nil {
			droppedGameYear++
			continue
		}

		t := removeSpecChar(mapPlayerToTeam[pid].SeasonTeamMap[row[RB_YEAR]])
		_, ok := r.Data[t]
		if !ok {
			droppedTeam++
			continue
		}

		birthYearString, ok := mapPlayerBirthYear[pid]
		if !ok {
			droppedBirthYear++
			continue
		}

		birthYear, err := strconv.Atoi(birthYearString)
		if err != nil {
			droppedBirthYear++
			continue
		}

		p, ok := mapPlayerPosition[pid]
		if !ok {
			droppedPosition++
		}

		if p == "" {
			droppedPosition++
		}

		r.Set("player_id", pid)
		r.SetPosition(row[RB_POSITION])
		r.Set("nfl_year", row[RB_YEAR])
		r.SetSeason(row[RB_SEASON])
		r.Set("week", row[RB_WEEK])
		r.SetGameMonth(row[RB_GAME_DATE])
		r.SetHomeAway(row[RB_HOME_OR_AWAY])
		r.Set(row[RB_OPPONENT], "1")
		r.SetOutcome(row[RB_OUTCOME])
		r.SetPointsScore(row[RB_SCORE])
		r.Set("rushing_attempts", row[RB_RUSHING_ATTEMPTS])
		r.Set("rushing_yards", row[RB_RUSHING_YARDS])
		r.Set("yards_per_carry", row[RB_YARDS_PER_CARRY])
		r.Set("longest_rushing_run", row[RB_LONGEST_RUSHING_RUN])
		r.Set("rushing_tds", row[RB_RUSHING_TDS])
		r.Set("receptions", row[RB_RECEPTIONS])
		r.Set("receiving_yards", row[RB_RECEIVING_YARDS])
		r.Set("yards_per_reception", row[RB_YARDS_PER_RECEPTION])
		r.Set("longest_reception", row[RB_LONGEST_RECEPTION])
		r.Set("receiving_tds", row[RB_RECEIVING_TDS])
		r.Set("fumbles", row[RB_FUMBLES])
		r.Set("fumbles_lost", row[RB_FUMBLES_LOST])

		r.SetAge(gameYear - birthYear)
		r.SetTeam(pid, row[RB_YEAR])

		r.SetFantasyScores()

		mergedDataMap[index+mergedDataMapSize] = r

	}

	log.Printf("%s ::: %v of %v", name, size, size)
	log.Printf("%s ::: Drop %v Record(s) due to %v", name, droppedTeam, "(NO TEAM DATA)")
	log.Printf("%s ::: Drop %v Record(s) due to %v", name, droppedBirthYear, "(NO BIRTH DATA)")
	log.Printf("%s ::: Drop %v Record(s) due to %v", name, droppedGameYear, "(NO GAME DATE DATA)")
	log.Printf("%s ::: WARN %v Record(s) DO NOT HAVE POSITION DATA", name, droppedPosition)
}

func addWrRecords(name string, csv [][]string) {
	mergedDataMapSize := len(mergedDataMap) + 1
	size := len(csv)
	// check for team
	droppedTeam := 0
	// check for birth year
	droppedBirthYear := 0
	// check for game year
	droppedGameYear := 0
	// check for position
	droppedPosition := 0

	for index, row := range csv {

		if index%10000 == 0 {
			log.Printf("%s ::: %v of %v", name, index, size)
		}

		pid := GetPlayerId(row[PLAYER_ID])
		var r Record
		r.NewRecord(mergedHeaders)

		gameYear, err := strconv.Atoi(row[YEAR])
		if err != nil {
			droppedGameYear++
			continue
		}

		t := removeSpecChar(mapPlayerToTeam[pid].SeasonTeamMap[row[YEAR]])
		_, ok := r.Data[t]
		if !ok {
			droppedTeam++
			continue
		}

		birthYearString, ok := mapPlayerBirthYear[pid]
		if !ok {
			droppedBirthYear++
			continue
		}

		birthYear, err := strconv.Atoi(birthYearString)
		if err != nil {
			droppedBirthYear++
			continue
		}

		p, ok := mapPlayerPosition[pid]
		if !ok {
			droppedPosition++
		}

		if p == "" {
			droppedPosition++
		}

		r.Set("player_id", pid)
		r.Set("nfl_year", row[YEAR])
		r.SetAge(gameYear - birthYear)
		r.SetTeam(pid, row[YEAR])
		r.Set(row[OPPONENT], "1")
		r.SetGameMonth(row[GAME_DATE])
		r.SetOutcome(row[OUTCOME])
		r.SetPosition(row[POSITION])
		r.SetSeason(row[SEASON])
		r.Set("week", row[WEEK])
		r.SetHomeAway(row[HOME_OR_AWAY])
		r.SetPointsScore(row[SCORE])
		r.Set("receptions", row[RECEPTIONS])
		r.Set("receiving_yards", row[RECEIVING_YARDS])
		r.Set("yards_per_reception", row[YARDS_PER_RECEPTION])
		r.Set("longest_reception", row[LONGEST_RECEPTION])
		r.Set("receiving_tds", row[RECEIVING_TDS])
		r.Set("rushing_attempts", row[RUSHING_ATTEMPTS])
		r.Set("rushing_yards", row[RUSHING_YARDS])
		r.Set("yards_per_carry", row[YARDS_PER_CARRY])
		r.Set("longest_rushing_run", row[LONGEST_RUSHING_RUN])
		r.Set("rushing_tds", row[RUSHING_TDS])
		r.Set("fumbles", row[FUMBLES])
		r.Set("fumbles_lost", row[FUMBLES_LOST])
		r.SetFantasyScores()

		mergedDataMap[index+mergedDataMapSize] = r

	}

	log.Printf("%s ::: %v of %v", name, size, size)
	log.Printf("%s ::: Drop %v Record(s) due to %v", name, droppedTeam, "(NO TEAM DATA)")
	log.Printf("%s ::: Drop %v Record(s) due to %v", name, droppedBirthYear, "(NO BIRTH DATA)")
	log.Printf("%s ::: Drop %v Record(s) due to %v", name, droppedGameYear, "(NO GAME DATE DATA)")
	log.Printf("%s ::: WARN %v Record(s) DO NOT HAVE POSITION DATA", name, droppedPosition)
}

func addQbRecords(name string, csv [][]string) {
	mergedDataMapSize := len(mergedDataMap) + 1
	size := len(csv)
	// check for team
	droppedTeam := 0
	// check for birth year
	droppedBirthYear := 0
	// check for game year
	droppedGameYear := 0
	// check for position
	droppedPosition := 0

	for index, row := range csv {

		if index%10000 == 0 {
			log.Printf("%s ::: %v of %v", name, index, size)
		}

		pid := GetPlayerId(row[QB_PLAYER_ID])
		var r Record
		r.NewRecord(mergedHeaders)

		gameYear, err := strconv.Atoi(row[QB_YEAR])
		if err != nil {
			droppedGameYear++
			continue
		}

		t := removeSpecChar(mapPlayerToTeam[pid].SeasonTeamMap[row[QB_YEAR]])
		_, ok := r.Data[t]
		if !ok {
			droppedTeam++
			continue
		}

		birthYearString, ok := mapPlayerBirthYear[pid]
		if !ok {
			droppedBirthYear++
			continue
		}

		birthYear, err := strconv.Atoi(birthYearString)
		if err != nil {
			droppedBirthYear++
			continue
		}

		p, ok := mapPlayerPosition[pid]
		if !ok {
			droppedPosition++
		}

		if p == "" {
			droppedPosition++
		}

		r.Set("player_id", pid)
		r.Set("qb", "1")
		r.Set("nfl_year", row[QB_YEAR])
		r.SetSeason(row[QB_SEASON])
		r.Set("week", row[QB_WEEK])
		r.SetGameMonth(row[QB_GAME_DATE])
		r.SetHomeAway(row[QB_HOME_OR_AWAY])
		r.Set(row[QB_OPPONENT], "1")
		r.SetOutcome(row[QB_OUTCOME])
		r.SetPointsScore(row[QB_SCORE])
		r.Set("passes_completed", row[QB_PASSES_COMPLETED])
		r.Set("passes_attempted", row[QB_PASSES_ATTEMPTED])
		r.Set("completion_percentage", row[QB_COMPLETION_PERCENTAGE])
		r.Set("passing_yards", row[QB_PASSING_YARDS])
		r.Set("passing_yards_per_attempt", row[QB_PASSING_YARDS_PER_ATTEMPT])
		r.Set("passing_tds", row[QB_TD_PASSES])
		r.Set("interceptions", row[QB_INTS])
		r.Set("sacks", row[QB_SACKS])
		r.Set("sacks_yards_lost", row[QB_SACKED_YARDS_LOST])
		r.Set("passer_rating", row[QB_PASSER_RATING])
		r.Set("rushing_attempts", row[QB_RUSHING_ATTEMPTS])
		r.Set("rushing_yards", row[QB_RUSHING_YARDS])
		r.Set("yards_per_carry", row[QB_YARDS_PER_CARRY])
		r.Set("rushing_tds", row[QB_RUSHING_TDS])
		r.Set("fumbles", row[QB_FUMBLES])
		r.Set("fumbles_lost", row[QB_FUMBLES_LOST])
		r.SetFantasyScores()

		r.SetAge(gameYear - birthYear)
		r.SetTeam(pid, row[QB_YEAR])

		mergedDataMap[index+mergedDataMapSize] = r

	}

	log.Printf("%s ::: %v of %v", name, size, size)
	log.Printf("%s ::: Drop %v Record(s) due to %v", name, droppedTeam, "(NO TEAM DATA)")
	log.Printf("%s ::: Drop %v Record(s) due to %v", name, droppedBirthYear, "(NO BIRTH DATA)")
	log.Printf("%s ::: Drop %v Record(s) due to %v", name, droppedGameYear, "(NO GAME DATE DATA)")
	log.Printf("%s ::: WARN %v Record(s) DO NOT HAVE POSITION DATA", name, droppedPosition)
}
