package main

import (
	"log"
	"strconv"
	"strings"
)

type Record struct {
	Data map[string]string
}

func (r *Record) NewRecord(keys []string) {
	r.Data = make(map[string]string)

	for _, key := range keys {
		r.Data[key] = "0"
	}

}

func (r Record) Set(key, value string) {

	if value == "" || value == "--" || value == "NaN" || len(value) < 1 {
		r.Data[key] = "0"
		return
	}

	if strings.ContainsAny(value, "T") {
		value = strings.Replace(value, "T", "", -1)
	}

	r.Data[key] = value
}

func (r Record) SetAge(age int) {

	if age < 20 {
		r.Data["age_under_20"] = "1"
		return
	}

	if age > 45 {
		r.Data["age_over_45"] = "1"
		return
	}

	formattedAge := "age_" + strconv.Itoa(age)
	r.Data[formattedAge] = "1"

}

func (r Record) SetTeam(pid, year string) {
	team := removeSpecChar(mapPlayerToTeam[pid].SeasonTeamMap[year])
	_, ok := r.Data[team]
	if !ok {
		log.Fatalf("team not found")
	}
	r.Data[team] = "1"
}

func (r Record) SetGameMonth(date string) {
	month := strings.Split(date, "/")[0]
	formattedMonth := "game_month_" + strings.TrimSpace(month)
	r.Data[formattedMonth] = "1"

	if month == "08" ||
		month == "09" ||
		month == "10" ||
		month == "11" ||
		month == "12" ||
		month == "01" ||
		month == "02" {

	} else {
		log.Fatalf("month - |%s|", month)
	}

	if len(strings.Split(date, "/")) != 2 {
		log.Fatalf("failed set game month with %s", date)
	}
}

func (r Record) SetOutcome(outcome string) {
	switch outcome {
	case "L":
		r.Data["outcome_loss"] = "1"
	case "W":
		r.Data["outcome_win"] = "1"
	case "T":
		r.Data["outcome_tie"] = "1"
	default:
		log.Fatalf("Unknown outcome - %s", outcome)
	}
}

func (r Record) SetPointsScore(data string) {
	score := strings.Split(data, "to")[0]
	r.Data["points_scored"] = strings.TrimSpace(score)
}

func (r Record) SetPosition(data string) {
	switch data {
	case "QB":
		r.Data["qb"] = "1"
	case "WR":
		r.Data["wr"] = "1"
	case "TE":
		r.Data["te"] = "1"
	case "RB":
		r.Data["rb"] = "1"
	case "FB":
		r.Data["fb"] = "1"
	}
}

func (r Record) SetSeason(data string) {
	switch data {
	case "Preseason":
		r.Data["pre_season"] = "1"
	case "Regular Season":
		r.Data["regular_season"] = "1"
	case "Postseason":
		r.Data["post_season"] = "1"
	case "Pro Bowl":
		r.Data["pro_bowl"] = "1"
	default:
		log.Fatalf("unknown season - %s", data)
	}

}

func (r Record) SetHomeAway(data string) {
	switch data {
	case "Home":
		r.Data["home"] = "1"
	case "Away":
		r.Data["away"] = "1"
	default:
		log.Fatalf("unknown home/away - %s", data)
	}
}

func (r Record) SetFantasyScores() {

	rushingTd := asInt("rushing_tds", r)
	receivingTd := asInt("receiving_tds", r)
	passingTd := asInt("passing_tds", r)
	rushingYards := asInt("rushing_yards", r)
	receivingYards := asInt("receiving_yards", r)
	passingYards := asInt("passing_yards", r)
	interceptions := asInt("interceptions", r)
	fumblesLost := asInt("fumbles_lost", r)

	r.Data["fs_rush_recieve_yards"] = strconv.Itoa((receivingYards / 10) + (rushingYards / 10))
	r.Data["fs_passing_yards"] = strconv.Itoa(passingYards / 25)
	r.Data["fs_rush_recieve_td"] = strconv.Itoa((rushingTd * 6) + (receivingTd * 6))
	r.Data["fs_passing_td"] = strconv.Itoa(passingTd * 4)
	r.Data["fs_interceptions"] = strconv.Itoa(interceptions * -2)
	r.Data["fs_fumble"] = strconv.Itoa(fumblesLost * -2)

	r.Data["fs_total_penalty_points"] = strconv.Itoa(
		(fumblesLost * -2) + (interceptions * -2),
	)

	r.Data["fs_total"] = strconv.Itoa(
		(receivingYards / 10) + (rushingYards / 10) + (passingYards / 25) +
			(rushingTd * 6) + (receivingTd * 6) + (passingTd * 4) +
			(interceptions * -2) + (fumblesLost * -2),
	)

}

func asInt(element string, r Record) int {

	i, err := strconv.Atoi(r.Data[element])
	if err != nil {
		log.Fatalf("failed to convert element value to int ::: element-|%s|, pid-|%s|, year-%s, e-%s", element, r.Data["player_id"], r.Data["nfl_year"], err)
	}

	return i
}

func (r Record) String() string {
	return r.Data["player_id"] + "," + r.Data["age_under_20"] + "," + r.Data["age_20"] + "," + r.Data["age_21"] + "," + r.Data["age_22"] + "," + r.Data["age_23"] + "," + r.Data["age_24"] + "," + r.Data["age_25"] + "," + r.Data["age_26"] + "," + r.Data["age_27"] + "," + r.Data["age_28"] + "," + r.Data["age_29"] + "," + r.Data["age_30"] + "," + r.Data["age_31"] + "," + r.Data["age_32"] + "," + r.Data["age_33"] + "," + r.Data["age_34"] + "," + r.Data["age_35"] + "," + r.Data["age_36"] + "," + r.Data["age_37"] + "," + r.Data["age_38"] + "," + r.Data["age_39"] + "," + r.Data["age_40"] + "," + r.Data["age_41"] + "," + r.Data["age_42"] + "," + r.Data["age_43"] + "," + r.Data["age_44"] + "," + r.Data["age_45"] + "," + r.Data["age_over_45"] + "," + r.Data["Tampa_Bay_Buccaneers"] + "," + r.Data["Los_Angeles_Rams"] + "," + r.Data["New_Orleans_Saints"] + "," + r.Data["Green_Bay_Packers"] + "," + r.Data["San_Diego_Chargers"] + "," + r.Data["Los_Angeles_Raiders"] + "," + r.Data["Buffalo_Bisons"] + "," + r.Data["Chicago_Bears"] + "," + r.Data["Tennessee_Titans"] + "," + r.Data["St_Louis_Cardinals"] + "," + r.Data["Dallas_Cowboys"] + "," + r.Data["Card-Pitt_Combine"] + "," + r.Data["Washington_Redskins"] + "," + r.Data["Jacksonville_Jaguars"] + "," + r.Data["St_Louis_Rams"] + "," + r.Data["Boston_Bulldogs"] + "," + r.Data["Philadelphia_Eagles"] + "," + r.Data["San_Francisco_49ers"] + "," + r.Data["Chicago_Cardinals"] + "," + r.Data["Boston_Patriots"] + "," + r.Data["Phoenix_Cardinals"] + "," + r.Data["Staten_Island_Stapletons"] + "," + r.Data["Phil-Pitt_Combine"] + "," + r.Data["Cleveland_Indians"] + "," + r.Data["Duluth_Eskimos"] + "," + r.Data["Detroit_Wolverines"] + "," + r.Data["Cleveland_Bulldogs"] + "," + r.Data["New_York_Giants"] + "," + r.Data["Los_Angeles_Chargers"] + "," + r.Data["New_York_Yanks"] + "," + r.Data["Minnesota_Vikings"] + "," + r.Data["Houston_Oilers"] + "," + r.Data["Boston_Yanks"] + "," + r.Data["Brooklyn_Tigers"] + "," + r.Data["Chicago_Rockets"] + "," + r.Data["Indianapolis_Colts"] + "," + r.Data["Detroit_Lions"] + "," + r.Data["Miami_Dolphins"] + "," + r.Data["Pittsburgh_Pirates"] + "," + r.Data["Cleveland_Browns"] + "," + r.Data["Atlanta_Falcons"] + "," + r.Data["Boston_Redskins"] + "," + r.Data["Los_Angeles_Buccaneers"] + "," + r.Data["Frankford_Yellow_Jackets"] + "," + r.Data["Pottsville_Maroons"] + "," + r.Data["Tennessee_Oilers"] + "," + r.Data["Oakland_Raiders"] + "," + r.Data["Newark_Tornadoes"] + "," + r.Data["Milwaukee_Badgers"] + "," + r.Data["Rock_Island_Independents"] + "," + r.Data["Carolina_Panthers"] + "," + r.Data["Cincinnati_Bengals"] + "," + r.Data["Seattle_Seahawks"] + "," + r.Data["Pittsburgh_Steelers"] + "," + r.Data["New_York_Yankees"] + "," + r.Data["Cincinnati_Reds"] + "," + r.Data["Baltimore_Ravens"] + "," + r.Data["Los_Angeles_Dons"] + "," + r.Data["Chicago_Hornets"] + "," + r.Data["Buffalo_Bills"] + "," + r.Data["Kansas_City_Chiefs"] + "," + r.Data["New_York_Jets"] + "," + r.Data["New_York_Titans"] + "," + r.Data["Houston_Texans"] + "," + r.Data["Dallas_Texans"] + "," + r.Data["Portsmouth_Spartans"] + "," + r.Data["St_Louis_Gunners"] + "," + r.Data["Brooklyn_Dodgers"] + "," + r.Data["Boston_Braves"] + "," + r.Data["Cleveland_Rams"] + "," + r.Data["Minneapolis_Red_Jackets"] + "," + r.Data["Miami_Seahawks"] + "," + r.Data["Denver_Broncos"] + "," + r.Data["New_England_Patriots"] + "," + r.Data["New_York_Bulldogs"] + "," + r.Data["Providence_Steam_Roller"] + "," + r.Data["Arizona_Cardinals"] + "," + r.Data["Baltimore_Colts"] + "," + r.Data["ATL"] + "," + r.Data["MIA"] + "," + r.Data["RAM"] + "," + r.Data["SAN"] + "," + r.Data["SEA"] + "," + r.Data["DAL"] + "," + r.Data["SF"] + "," + r.Data["GB"] + "," + r.Data["STL"] + "," + r.Data["JAX"] + "," + r.Data["NYG"] + "," + r.Data["SD"] + "," + r.Data["RAI"] + "," + r.Data["CIN"] + "," + r.Data["NPR"] + "," + r.Data["TEN"] + "," + r.Data["PIT"] + "," + r.Data["HOU"] + "," + r.Data["NYJ"] + "," + r.Data["CLE"] + "," + r.Data["IND"] + "," + r.Data["BOS"] + "," + r.Data["DEN"] + "," + r.Data["OAK"] + "," + r.Data["RIC"] + "," + r.Data["TB"] + "," + r.Data["LA"] + "," + r.Data["APR"] + "," + r.Data["BAL"] + "," + r.Data["PHI"] + "," + r.Data["WAS"] + "," + r.Data["BUF"] + "," + r.Data["ARI"] + "," + r.Data["CHI"] + "," + r.Data["NO"] + "," + r.Data["NE"] + "," + r.Data["PHO"] + "," + r.Data["DET"] + "," + r.Data["MIN"] + "," + r.Data["JAC"] + "," + r.Data["KC"] + "," + r.Data["CRT"] + "," + r.Data["CAR"] + "," + r.Data["game_month_08"] + "," + r.Data["game_month_09"] + "," + r.Data["game_month_10"] + "," + r.Data["game_month_11"] + "," + r.Data["game_month_12"] + "," + r.Data["game_month_01"] + "," + r.Data["game_month_02"] + "," + r.Data["outcome_win"] + "," + r.Data["outcome_tie"] + "," + r.Data["outcome_loss"] + "," + r.Data["points_scored"] + "," + r.Data["rb"] + "," + r.Data["fb"] + "," + r.Data["wr"] + "," + r.Data["te"] + "," + r.Data["qb"] + "," + r.Data["nfl_year"] + "," + r.Data["pre_season"] + "," + r.Data["regular_season"] + "," + r.Data["post_season"] + "," + r.Data["pro_bowl"] + "," + r.Data["week"] + "," + r.Data["home"] + "," + r.Data["away"] + "," + r.Data["receptions"] + "," + r.Data["receiving_yards"] + "," + r.Data["yards_per_reception"] + "," + r.Data["longest_reception"] + "," + r.Data["receiving_tds"] + "," + r.Data["rushing_attempts"] + "," + r.Data["rushing_yards"] + "," + r.Data["yards_per_carry"] + "," + r.Data["longest_rushing_run"] + "," + r.Data["rushing_tds"] + "," + r.Data["fumbles"] + "," + r.Data["fumbles_lost"] + "," + r.Data["passes_completed"] + "," + r.Data["passes_attempted"] + "," + r.Data["completion_percentage"] + "," + r.Data["passing_yards"] + "," + r.Data["passing_yards_per_attempt"] + "," + r.Data["passing_tds"] + "," + r.Data["interceptions"] + "," + r.Data["sacks"] + "," + r.Data["sacks_yards_lost"] + "," + r.Data["passer_rating"] + "," + r.Data["fs_rush_recieve_yards"] + "," + r.Data["fs_passing_yards"] + "," + r.Data["fs_rush_recieve_td"] + "," + r.Data["fs_passing_td"] + "," + r.Data["fs_interceptions"] + "," + r.Data["fs_fumble"] + "," + r.Data["fs_total_penalty_points"] + "," + r.Data["fs_total_points_gained"] + "," + r.Data["fs_total"]
}
