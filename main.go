package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type fpl struct {
	Events []struct {
		ID                     int       `json:"id"`
		Name                   string    `json:"name"`
		DeadlineTime           time.Time `json:"deadline_time"`
		AverageEntryScore      int       `json:"average_entry_score"`
		Finished               bool      `json:"finished"`
		DataChecked            bool      `json:"data_checked"`
		HighestScoringEntry    int       `json:"highest_scoring_entry"`
		DeadlineTimeEpoch      int       `json:"deadline_time_epoch"`
		DeadlineTimeGameOffset int       `json:"deadline_time_game_offset"`
		HighestScore           int       `json:"highest_score"`
		IsPrevious             bool      `json:"is_previous"`
		IsCurrent              bool      `json:"is_current"`
		IsNext                 bool      `json:"is_next"`
		ChipPlays              []struct {
			ChipName  string `json:"chip_name"`
			NumPlayed int    `json:"num_played"`
		} `json:"chip_plays"`
		MostSelected      int `json:"most_selected"`
		MostTransferredIn int `json:"most_transferred_in"`
		TopElement        int `json:"top_element"`
		TopElementInfo    struct {
			ID     int `json:"id"`
			Points int `json:"points"`
		} `json:"top_element_info"`
		TransfersMade     int `json:"transfers_made"`
		MostCaptained     int `json:"most_captained"`
		MostViceCaptained int `json:"most_vice_captained"`
	} `json:"events"`
	GameSettings struct {
		LeagueJoinPrivateMax         int           `json:"league_join_private_max"`
		LeagueJoinPublicMax          int           `json:"league_join_public_max"`
		LeagueMaxSizePublicClassic   int           `json:"league_max_size_public_classic"`
		LeagueMaxSizePublicH2H       int           `json:"league_max_size_public_h2h"`
		LeagueMaxSizePrivateH2H      int           `json:"league_max_size_private_h2h"`
		LeagueMaxKoRoundsPrivateH2H  int           `json:"league_max_ko_rounds_private_h2h"`
		LeaguePrefixPublic           string        `json:"league_prefix_public"`
		LeaguePointsH2HWin           int           `json:"league_points_h2h_win"`
		LeaguePointsH2HLose          int           `json:"league_points_h2h_lose"`
		LeaguePointsH2HDraw          int           `json:"league_points_h2h_draw"`
		LeagueKoFirstInsteadOfRandom bool          `json:"league_ko_first_instead_of_random"`
		CupStartEventID              int           `json:"cup_start_event_id"`
		CupStopEventID               int           `json:"cup_stop_event_id"`
		CupQualifyingMethod          string        `json:"cup_qualifying_method"`
		CupType                      string        `json:"cup_type"`
		SquadSquadplay               int           `json:"squad_squadplay"`
		SquadSquadsize               int           `json:"squad_squadsize"`
		SquadTeamLimit               int           `json:"squad_team_limit"`
		SquadTotalSpend              int           `json:"squad_total_spend"`
		UICurrencyMultiplier         int           `json:"ui_currency_multiplier"`
		UIUseSpecialShirts           bool          `json:"ui_use_special_shirts"`
		UISpecialShirtExclusions     []interface{} `json:"ui_special_shirt_exclusions"`
		StatsFormDays                int           `json:"stats_form_days"`
		SysViceCaptainEnabled        bool          `json:"sys_vice_captain_enabled"`
		TransfersCap                 int           `json:"transfers_cap"`
		TransfersSellOnFee           float64       `json:"transfers_sell_on_fee"`
		LeagueH2HTiebreakStats       []string      `json:"league_h2h_tiebreak_stats"`
		Timezone                     string        `json:"timezone"`
	} `json:"game_settings"`
	Phases []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		StartEvent int    `json:"start_event"`
		StopEvent  int    `json:"stop_event"`
	} `json:"phases"`
	Teams []struct {
		Code                int         `json:"code"`
		Draw                int         `json:"draw"`
		Form                interface{} `json:"form"`
		ID                  int         `json:"id"`
		Loss                int         `json:"loss"`
		Name                string      `json:"name"`
		Played              int         `json:"played"`
		Points              int         `json:"points"`
		Position            int         `json:"position"`
		ShortName           string      `json:"short_name"`
		Strength            int         `json:"strength"`
		TeamDivision        interface{} `json:"team_division"`
		Unavailable         bool        `json:"unavailable"`
		Win                 int         `json:"win"`
		StrengthOverallHome int         `json:"strength_overall_home"`
		StrengthOverallAway int         `json:"strength_overall_away"`
		StrengthAttackHome  int         `json:"strength_attack_home"`
		StrengthAttackAway  int         `json:"strength_attack_away"`
		StrengthDefenceHome int         `json:"strength_defence_home"`
		StrengthDefenceAway int         `json:"strength_defence_away"`
		PulseID             int         `json:"pulse_id"`
	} `json:"teams"`
	TotalPlayers int `json:"total_players"`
	Elements     []struct {
		ChanceOfPlayingNextRound         interface{} `json:"chance_of_playing_next_round"`
		ChanceOfPlayingThisRound         interface{} `json:"chance_of_playing_this_round"`
		Code                             int         `json:"code"`
		CostChangeEvent                  int         `json:"cost_change_event"`
		CostChangeEventFall              int         `json:"cost_change_event_fall"`
		CostChangeStart                  int         `json:"cost_change_start"`
		CostChangeStartFall              int         `json:"cost_change_start_fall"`
		DreamteamCount                   int         `json:"dreamteam_count"`
		ElementType                      int         `json:"element_type"`
		EpNext                           string      `json:"ep_next"`
		EpThis                           string      `json:"ep_this"`
		EventPoints                      int         `json:"event_points"`
		FirstName                        string      `json:"first_name"`
		Form                             string      `json:"form"`
		ID                               int         `json:"id"`
		InDreamteam                      bool        `json:"in_dreamteam"`
		News                             string      `json:"news"`
		NewsAdded                        interface{} `json:"news_added"`
		NowCost                          int         `json:"now_cost"`
		Photo                            string      `json:"photo"`
		PointsPerGame                    string      `json:"points_per_game"`
		SecondName                       string      `json:"second_name"`
		SelectedByPercent                string      `json:"selected_by_percent"`
		Special                          bool        `json:"special"`
		SquadNumber                      interface{} `json:"squad_number"`
		Status                           string      `json:"status"`
		Team                             int         `json:"team"`
		TeamCode                         int         `json:"team_code"`
		TotalPoints                      int         `json:"total_points"`
		TransfersIn                      int         `json:"transfers_in"`
		TransfersInEvent                 int         `json:"transfers_in_event"`
		TransfersOut                     int         `json:"transfers_out"`
		TransfersOutEvent                int         `json:"transfers_out_event"`
		ValueForm                        string      `json:"value_form"`
		ValueSeason                      string      `json:"value_season"`
		WebName                          string      `json:"web_name"`
		Minutes                          int         `json:"minutes"`
		GoalsScored                      int         `json:"goals_scored"`
		Assists                          int         `json:"assists"`
		CleanSheets                      int         `json:"clean_sheets"`
		GoalsConceded                    int         `json:"goals_conceded"`
		OwnGoals                         int         `json:"own_goals"`
		PenaltiesSaved                   int         `json:"penalties_saved"`
		PenaltiesMissed                  int         `json:"penalties_missed"`
		YellowCards                      int         `json:"yellow_cards"`
		RedCards                         int         `json:"red_cards"`
		Saves                            int         `json:"saves"`
		Bonus                            int         `json:"bonus"`
		Bps                              int         `json:"bps"`
		Influence                        string      `json:"influence"`
		Creativity                       string      `json:"creativity"`
		Threat                           string      `json:"threat"`
		IctIndex                         string      `json:"ict_index"`
		InfluenceRank                    int         `json:"influence_rank"`
		InfluenceRankType                int         `json:"influence_rank_type"`
		CreativityRank                   int         `json:"creativity_rank"`
		CreativityRankType               int         `json:"creativity_rank_type"`
		ThreatRank                       int         `json:"threat_rank"`
		ThreatRankType                   int         `json:"threat_rank_type"`
		IctIndexRank                     int         `json:"ict_index_rank"`
		IctIndexRankType                 int         `json:"ict_index_rank_type"`
		CornersAndIndirectFreekicksOrder interface{} `json:"corners_and_indirect_freekicks_order"`
		CornersAndIndirectFreekicksText  string      `json:"corners_and_indirect_freekicks_text"`
		DirectFreekicksOrder             interface{} `json:"direct_freekicks_order"`
		DirectFreekicksText              string      `json:"direct_freekicks_text"`
		PenaltiesOrder                   interface{} `json:"penalties_order"`
		PenaltiesText                    string      `json:"penalties_text"`
	} `json:"elements"`
	ElementStats []struct {
		Label string `json:"label"`
		Name  string `json:"name"`
	} `json:"element_stats"`
	ElementTypes []struct {
		ID                 int    `json:"id"`
		PluralName         string `json:"plural_name"`
		PluralNameShort    string `json:"plural_name_short"`
		SingularName       string `json:"singular_name"`
		SingularNameShort  string `json:"singular_name_short"`
		SquadSelect        int    `json:"squad_select"`
		SquadMinPlay       int    `json:"squad_min_play"`
		SquadMaxPlay       int    `json:"squad_max_play"`
		UIShirtSpecific    bool   `json:"ui_shirt_specific"`
		SubPositionsLocked []int  `json:"sub_positions_locked"`
		ElementCount       int    `json:"element_count"`
	} `json:"element_types"`
}

type picks struct {
	ActiveChip    interface{}   `json:"active_chip"`
	AutomaticSubs []interface{} `json:"automatic_subs"`
	EntryHistory  struct {
		Event              int `json:"event"`
		Points             int `json:"points"`
		TotalPoints        int `json:"total_points"`
		Rank               int `json:"rank"`
		RankSort           int `json:"rank_sort"`
		OverallRank        int `json:"overall_rank"`
		Bank               int `json:"bank"`
		Value              int `json:"value"`
		EventTransfers     int `json:"event_transfers"`
		EventTransfersCost int `json:"event_transfers_cost"`
		PointsOnBench      int `json:"points_on_bench"`
	} `json:"entry_history"`
	Picks []struct {
		Element       int  `json:"element"`
		Position      int  `json:"position"`
		Multiplier    int  `json:"multiplier"`
		IsCaptain     bool `json:"is_captain"`
		IsViceCaptain bool `json:"is_vice_captain"`
	} `json:"picks"`
}

type player struct {
	Fixtures []struct {
		ID                   int         `json:"id"`
		Code                 int         `json:"code"`
		TeamH                int         `json:"team_h"`
		TeamHScore           interface{} `json:"team_h_score"`
		TeamA                int         `json:"team_a"`
		TeamAScore           interface{} `json:"team_a_score"`
		Event                int         `json:"event"`
		Finished             bool        `json:"finished"`
		Minutes              int         `json:"minutes"`
		ProvisionalStartTime bool        `json:"provisional_start_time"`
		KickoffTime          time.Time   `json:"kickoff_time"`
		EventName            string      `json:"event_name"`
		IsHome               bool        `json:"is_home"`
		Difficulty           int         `json:"difficulty"`
	} `json:"fixtures"`
	History []struct {
		Element          int       `json:"element"`
		Fixture          int       `json:"fixture"`
		OpponentTeam     int       `json:"opponent_team"`
		TotalPoints      int       `json:"total_points"`
		WasHome          bool      `json:"was_home"`
		KickoffTime      time.Time `json:"kickoff_time"`
		TeamHScore       int       `json:"team_h_score"`
		TeamAScore       int       `json:"team_a_score"`
		Round            int       `json:"round"`
		Minutes          int       `json:"minutes"`
		GoalsScored      int       `json:"goals_scored"`
		Assists          int       `json:"assists"`
		CleanSheets      int       `json:"clean_sheets"`
		GoalsConceded    int       `json:"goals_conceded"`
		OwnGoals         int       `json:"own_goals"`
		PenaltiesSaved   int       `json:"penalties_saved"`
		PenaltiesMissed  int       `json:"penalties_missed"`
		YellowCards      int       `json:"yellow_cards"`
		RedCards         int       `json:"red_cards"`
		Saves            int       `json:"saves"`
		Bonus            int       `json:"bonus"`
		Bps              int       `json:"bps"`
		Influence        string    `json:"influence"`
		Creativity       string    `json:"creativity"`
		Threat           string    `json:"threat"`
		IctIndex         string    `json:"ict_index"`
		Value            int       `json:"value"`
		TransfersBalance int       `json:"transfers_balance"`
		Selected         int       `json:"selected"`
		TransfersIn      int       `json:"transfers_in"`
		TransfersOut     int       `json:"transfers_out"`
	} `json:"history"`
	HistoryPast []struct {
		SeasonName      string `json:"season_name"`
		ElementCode     int    `json:"element_code"`
		StartCost       int    `json:"start_cost"`
		EndCost         int    `json:"end_cost"`
		TotalPoints     int    `json:"total_points"`
		Minutes         int    `json:"minutes"`
		GoalsScored     int    `json:"goals_scored"`
		Assists         int    `json:"assists"`
		CleanSheets     int    `json:"clean_sheets"`
		GoalsConceded   int    `json:"goals_conceded"`
		OwnGoals        int    `json:"own_goals"`
		PenaltiesSaved  int    `json:"penalties_saved"`
		PenaltiesMissed int    `json:"penalties_missed"`
		YellowCards     int    `json:"yellow_cards"`
		RedCards        int    `json:"red_cards"`
		Saves           int    `json:"saves"`
		Bonus           int    `json:"bonus"`
		Bps             int    `json:"bps"`
		Influence       string `json:"influence"`
		Creativity      string `json:"creativity"`
		Threat          string `json:"threat"`
		IctIndex        string `json:"ict_index"`
	} `json:"history_past"`
}

type league struct {
	League struct {
		ID          int         `json:"id"`
		Name        string      `json:"name"`
		Created     time.Time   `json:"created"`
		Closed      bool        `json:"closed"`
		MaxEntries  interface{} `json:"max_entries"`
		LeagueType  string      `json:"league_type"`
		Scoring     string      `json:"scoring"`
		AdminEntry  int         `json:"admin_entry"`
		StartEvent  int         `json:"start_event"`
		CodePrivacy string      `json:"code_privacy"`
		Rank        interface{} `json:"rank"`
	} `json:"league"`
	NewEntries struct {
		HasNext bool          `json:"has_next"`
		Page    int           `json:"page"`
		Results []interface{} `json:"results"`
	} `json:"new_entries"`
	Standings struct {
		HasNext bool `json:"has_next"`
		Page    int  `json:"page"`
		Results []struct {
			ID         int    `json:"id"`
			EventTotal int    `json:"event_total"`
			PlayerName string `json:"player_name"`
			Rank       int    `json:"rank"`
			LastRank   int    `json:"last_rank"`
			RankSort   int    `json:"rank_sort"`
			Total      int    `json:"total"`
			Entry      int    `json:"entry"`
			EntryName  string `json:"entry_name"`
		} `json:"results"`
	} `json:"standings"`
}

type OutputPageData struct {
	PageTitle string
	Rows      []row
}

type row struct {
	TeamName string
	GWTotal  int
	Total    int
}

const fplURL string = "https://fantasy.premierleague.com/api/bootstrap-static/"

var fplData fpl

var rows []row

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fplURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "PostmanRuntime/7.18.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// log.Println(string(body))

	json.Unmarshal(body, &fplData)

	// fmt.Println(responseObject.Teams[0].Name)

	// for index, element := range fplData.Teams {
	// 	fmt.Println(index, element.ShortName, element.Name)
	// }

	// for _, element := range fplData.Elements {
	// 	fmt.Println(element.ID, element.SecondName, element.PointsPerGame, element.Team)
	// }
	// getPicks(575369, 6)
	// getPlayerName(390)
	getLeague(113899)

	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := OutputPageData{
			PageTitle: "FPL",
			Rows:      rows,
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}

func getPicks(id, week int) {
	client := &http.Client{}

	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/entry/%v/event/%v/picks/", id, week)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "PostmanRuntime/7.18.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// log.Println(string(body))

	var responseObject picks

	json.Unmarshal(body, &responseObject)

	// fmt.Println(responseObject.Teams[0].Name)

	// fmt.Println(responseObject.EntryHistory, responseObject)

	for _, element := range responseObject.Picks {
		// fmt.Println(index, element)
		fmt.Println(getPlayerName(element.Element))
		getPlayer(element.Element)
	}
}

func getPlayer(id int) {
	client := &http.Client{}

	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/element-summary/%v/", id)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "PostmanRuntime/7.18.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// log.Println(string(body))

	var responseObject player

	json.Unmarshal(body, &responseObject)

	// fmt.Println(responseObject.Teams[0].Name)

	// fmt.Println(responseObject.History)

	for index, element := range responseObject.History {
		fmt.Print("GW: ", index+1, " Pts: ", element.TotalPoints, " - ")
	}
	fmt.Println(" ")
}

func getPlayerName(id int) string {
	for _, element := range fplData.Elements {
		if element.ID == id {
			// fmt.Println(element.ID, element.FirstName, element.SecondName, element.PointsPerGame, element.Team)
			fullName := element.FirstName + " " + element.SecondName
			return fullName
		}
	}
	return ("")
}

func getLeague(id int) []row {
	client := &http.Client{}

	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/leagues-classic/%v/standings/", id)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "PostmanRuntime/7.18.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseObject league

	json.Unmarshal(body, &responseObject)

	for _, element := range responseObject.Standings.Results {
		fmt.Println(element.EntryName)
		fmt.Println("Team ID: ", element.Entry)
		fmt.Println("Event Total: ", element.EventTotal)
		fmt.Println("Total: ", element.Total)
		getPicks(element.Entry, 6)
		fmt.Println("---------")
		fmt.Println("---------")
		result := row{element.EntryName, element.EventTotal, element.Total}
		rows = append(rows, result)
	}
	fmt.Println(rows)
	return rows
}
