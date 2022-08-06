package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
	"embed"
	"os"

	"github.com/gorilla/mux"
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
		HasNext bool `json:"has_next"`
		Page    int  `json:"page"`
		Results []struct {
			Entry           int       `json:"entry"`
			EntryName       string    `json:"entry_name"`
			JoinedTime      time.Time `json:"joined_time"`
			PlayerFirstName string    `json:"player_first_name"`
			PlayerLastName  string    `json:"player_last_name"`
		} `json:"results"`
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
	PageTitle  string
	Rows       []row
	NewEntries []NewEntries
}

type entry struct {
	ID                       int       `json:"id"`
	JoinedTime               time.Time `json:"joined_time"`
	StartedEvent             int       `json:"started_event"`
	FavouriteTeam            int       `json:"favourite_team"`
	PlayerFirstName          string    `json:"player_first_name"`
	PlayerLastName           string    `json:"player_last_name"`
	PlayerRegionID           int       `json:"player_region_id"`
	PlayerRegionName         string    `json:"player_region_name"`
	PlayerRegionIsoCodeShort string    `json:"player_region_iso_code_short"`
	PlayerRegionIsoCodeLong  string    `json:"player_region_iso_code_long"`
	SummaryOverallPoints     int       `json:"summary_overall_points"`
	SummaryOverallRank       int       `json:"summary_overall_rank"`
	SummaryEventPoints       int       `json:"summary_event_points"`
	SummaryEventRank         int       `json:"summary_event_rank"`
	CurrentEvent             int       `json:"current_event"`
	Leagues                  struct {
		Classic []struct {
			ID             int         `json:"id"`
			Name           string      `json:"name"`
			ShortName      string      `json:"short_name"`
			Created        time.Time   `json:"created"`
			Closed         bool        `json:"closed"`
			Rank           interface{} `json:"rank"`
			MaxEntries     interface{} `json:"max_entries"`
			LeagueType     string      `json:"league_type"`
			Scoring        string      `json:"scoring"`
			AdminEntry     interface{} `json:"admin_entry"`
			StartEvent     int         `json:"start_event"`
			EntryRank      int         `json:"entry_rank"`
			EntryLastRank  int         `json:"entry_last_rank"`
			EntryCanLeave  bool        `json:"entry_can_leave"`
			EntryCanAdmin  bool        `json:"entry_can_admin"`
			EntryCanInvite bool        `json:"entry_can_invite"`
		} `json:"classic"`
		H2H []interface{} `json:"h2h"`
		Cup struct {
			Matches []interface{} `json:"matches"`
			Status  struct {
				QualificationEvent   int         `json:"qualification_event"`
				QualificationNumbers int         `json:"qualification_numbers"`
				QualificationRank    interface{} `json:"qualification_rank"`
				QualificationState   string      `json:"qualification_state"`
			} `json:"status"`
		} `json:"cup"`
	} `json:"leagues"`
	Name                       string `json:"name"`
	Kit                        string `json:"kit"`
	LastDeadlineBank           int    `json:"last_deadline_bank"`
	LastDeadlineValue          int    `json:"last_deadline_value"`
	LastDeadlineTotalTransfers int    `json:"last_deadline_total_transfers"`
}

type row struct {
	Rank      int
	TeamID    int
	TeamName  string
	GWTotal   int
	LiveTotal int
	PrevTotal int
	LastRank  int
	BenchPts  int
	Captain   string
	TotalPlayed	int
}

type NewEntries struct {
	TeamID    int
	TeamName  string
	FirstName string
	LastName  string
}

type livePlayerData struct {
	Elements []struct {
		ID    int `json:"id"`
		Stats struct {
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
			TotalPoints     int    `json:"total_points"`
			InDreamteam     bool   `json:"in_dreamteam"`
		} `json:"stats"`
		Explain []struct {
			Fixture int `json:"fixture"`
			Stats   []struct {
				Identifier string `json:"identifier"`
				Points     int    `json:"points"`
				Value      int    `json:"value"`
			} `json:"stats"`
		} `json:"explain"`
	} `json:"elements"`
}

type bonusPoints []struct {
	Code                 int       `json:"code"`
	Event                int       `json:"event"`
	Finished             bool      `json:"finished"`
	FinishedProvisional  bool      `json:"finished_provisional"`
	ID                   int       `json:"id"`
	KickoffTime          time.Time `json:"kickoff_time"`
	Minutes              int       `json:"minutes"`
	ProvisionalStartTime bool      `json:"provisional_start_time"`
	Started              bool      `json:"started"`
	TeamA                int       `json:"team_a"`
	TeamAScore           int       `json:"team_a_score"`
	TeamH                int       `json:"team_h"`
	TeamHScore           int       `json:"team_h_score"`
	Stats                []struct {
		Identifier string `json:"identifier"`
		A          []struct {
			Value   int `json:"value"`
			Element int `json:"element"`
		} `json:"a"`
		H []struct {
			Value   int `json:"value"`
			Element int `json:"element"`
		} `json:"h"`
	} `json:"stats"`
	TeamHDifficulty int `json:"team_h_difficulty"`
	TeamADifficulty int `json:"team_a_difficulty"`
	PulseID         int `json:"pulse_id"`
}

type bonusPointsCalc struct {
	ID    int
	Score int
}

type managerInfo struct {
	ID                       int         `json:"id"`
	JoinedTime               time.Time   `json:"joined_time"`
	StartedEvent             int         `json:"started_event"`
	FavouriteTeam            int         `json:"favourite_team"`
	PlayerFirstName          string      `json:"player_first_name"`
	PlayerLastName           string      `json:"player_last_name"`
	PlayerRegionID           int         `json:"player_region_id"`
	PlayerRegionName         string      `json:"player_region_name"`
	PlayerRegionIsoCodeShort string      `json:"player_region_iso_code_short"`
	PlayerRegionIsoCodeLong  string      `json:"player_region_iso_code_long"`
	SummaryOverallPoints     interface{} `json:"summary_overall_points"`
	SummaryOverallRank       interface{} `json:"summary_overall_rank"`
	SummaryEventPoints       interface{} `json:"summary_event_points"`
	SummaryEventRank         interface{} `json:"summary_event_rank"`
	CurrentEvent             interface{} `json:"current_event"`
	Leagues                  struct {
		Classic []struct {
			ID             int         `json:"id"`
			Name           string      `json:"name"`
			ShortName      string      `json:"short_name"`
			Created        time.Time   `json:"created"`
			Closed         bool        `json:"closed"`
			Rank           interface{} `json:"rank"`
			MaxEntries     interface{} `json:"max_entries"`
			LeagueType     string      `json:"league_type"`
			Scoring        string      `json:"scoring"`
			AdminEntry     interface{} `json:"admin_entry"`
			StartEvent     int         `json:"start_event"`
			EntryCanLeave  bool        `json:"entry_can_leave"`
			EntryCanAdmin  bool        `json:"entry_can_admin"`
			EntryCanInvite bool        `json:"entry_can_invite"`
			HasCup         bool        `json:"has_cup"`
			CupLeague      interface{} `json:"cup_league"`
			CupQualified   interface{} `json:"cup_qualified"`
			EntryRank      int         `json:"entry_rank"`
			EntryLastRank  int         `json:"entry_last_rank"`
		} `json:"classic"`
		H2H []interface{} `json:"h2h"`
		Cup struct {
			Matches []interface{} `json:"matches"`
			Status  struct {
				QualificationEvent   interface{} `json:"qualification_event"`
				QualificationNumbers interface{} `json:"qualification_numbers"`
				QualificationRank    interface{} `json:"qualification_rank"`
				QualificationState   interface{} `json:"qualification_state"`
			} `json:"status"`
			CupLeague interface{} `json:"cup_league"`
		} `json:"cup"`
		CupMatches []interface{} `json:"cup_matches"`
	} `json:"leagues"`
	Name                       string      `json:"name"`
	NameChangeBlocked          bool        `json:"name_change_blocked"`
	Kit                        string      `json:"kit"`
	LastDeadlineBank           interface{} `json:"last_deadline_bank"`
	LastDeadlineValue          interface{} `json:"last_deadline_value"`
	LastDeadlineTotalTransfers int         `json:"last_deadline_total_transfers"`
}

type managerLeagues struct {
	LeagueID   int
	LeagueName string
}

type managerOutputPageData struct {
	ManagerID				 int
	Leagues          []managerLeagues
	ManagerFirstName string
	ManagerLastName  string
	TeamName         string
	PastFinishes     managerPastData
	CurrentGw				 int
}

type managerPastData struct {
	Current []interface{} `json:"current"`
	Past    []struct {
		SeasonName  string `json:"season_name"`
		TotalPoints int    `json:"total_points"`
		Rank        int    `json:"rank"`
	} `json:"past"`
	Chips []interface{} `json:"chips"`
}

const fplURL string = "https://fantasy.premierleague.com/api/bootstrap-static/"

var fplData fpl

// var rows []row

var threeBp []int
var twoBp []int
var oneBp []int

// var wg sync.WaitGroup

var currentGw int

const (
	templatesDir = "templates/"
	extension    = "/*.html"
)

var (
	//go:embed templates/*
	files     embed.FS
	templates map[string]*template.Template
)

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

	for _, element := range fplData.Events {
		if element.IsCurrent == true {
			currentGw = element.ID
			fmt.Println("Current GW: ", currentGw)
		}
	}

	r := mux.NewRouter()

	// var wg sync.WaitGroup

	r.HandleFunc("/", handler)

	tmpl := template.Must(template.ParseFS(files,templatesDir+"league.html"))
	r.HandleFunc("/league/{league}", func(w http.ResponseWriter, r *http.Request) {
		// wg.Add(1)
		vars := mux.Vars(r)
		i, _ := strconv.Atoi(vars["league"])
		// rows = nil
		getBonusPoints()
		rows := getLeague(i, 1)
		newEntries := getNewLeagueEntries(i, 1)
		// go func() {
		// 	getLeague(i)
		// 	wg.Done()
		// }()

		// wg.Wait()

		data := OutputPageData{
			PageTitle:  "FPL",
			Rows:       rows,
			NewEntries: newEntries,
		}
		tmpl.Execute(w, data)
	})

	tmplManager := template.Must(template.ParseFS(files, templatesDir+"manager.html"))
	r.HandleFunc("/manager/{manager}", func(w http.ResponseWriter, r *http.Request) {
		// wg.Add(1)
		vars := mux.Vars(r)
		i, _ := strconv.Atoi(vars["manager"])
		// rows = nil
		managerInfo := getManagerInfo(i)
		tmplManager.Execute(w, managerInfo)
	})

	r.HandleFunc("/league", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "league_index.html")
		p, _ := ioutil.ReadFile(templatesDir+"league_index.html")
		w.Write(p)
	})

	r.HandleFunc("/manager", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "manager_index.html")
		p, _ := ioutil.ReadFile(templatesDir+"league_index.html")
		w.Write(p)
	})
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
					port = "8080"
					log.Printf("defaulting to port %s", port)
	}
	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
					log.Fatal(err)
	}
	// http.ListenAndServeTLS(":443", "localhost.crt", "localhost.key", r)
	// http.ListenAndServe(":80", r)
}

func getPicks(id, week int) ([]int, int) {
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

	var responseObject picks

	json.Unmarshal(body, &responseObject)

	var players []int
	var captain int

	for _, element := range responseObject.Picks {
		if element.Position <= 11 {
			if element.IsCaptain {
				captain = element.Element
				continue
			}
			players = append(players, element.Element)
		}
		// fmt.Println(getPlayerName(element.Element))
		// fmt.Println(element.IsCaptain)
		// getPlayer(element.Element)
	}
	return players, captain
}

func getCaptain(id, week int) string {
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

	var responseObject picks

	json.Unmarshal(body, &responseObject)

	for _, element := range responseObject.Picks {
		if element.IsCaptain {
			fmt.Println(element.Element)
			return getPlayerName(element.Element)
		}
	}
	return "N/A"
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

	// for index, element := range responseObject.History {
	// 	fmt.Print("GW: ", index+1, " Pts: ", element.TotalPoints, " - ")
	// }
	// fmt.Println(" ")
}

func getLiveScore(ids []int, week int) int {
	client := &http.Client{}

	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/event/%v/live/", currentGw)

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

	var responseObject livePlayerData

	json.Unmarshal(body, &responseObject)

	var liveTotal int

	for _, element := range responseObject.Elements {
		if contains(ids, element.ID) {
			liveTotal = liveTotal + element.Stats.TotalPoints - element.Stats.Bonus
			if contains(threeBp, element.ID) {
				liveTotal = liveTotal + 3
				fmt.Println("ADDING THREE for: ", element.ID)
				fmt.Println("Based off this chart: ", threeBp)
			}
			if contains(twoBp, element.ID) {
				liveTotal = liveTotal + 2
			}
			if contains(oneBp, element.ID) {
				liveTotal = liveTotal + 1
			}
		}
	}
	return liveTotal
}

func contains(s []int, num int) bool {
	for _, v := range s {
		if v == num {
			return true
		}
	}
	return false
}

func getPlayerName(id int) string {
	for _, element := range fplData.Elements {
		if element.ID == id {
			// fmt.Println(element.ID, element.FirstName, element.SecondName, element.PointsPerGame, element.Team)
			// fullName := element.FirstName + " " + element.SecondName
			return element.WebName
		}
	}
	return ("")
}

func getBonusPoints() {
	client := &http.Client{}

	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/fixtures/?event=%v", currentGw)

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

	var responseObject bonusPoints

	json.Unmarshal(body, &responseObject)

	// fmt.Println(responseObject[0].Stats[9].A)

	threeBp = nil
	twoBp = nil
	oneBp = nil
	for _, element := range responseObject {
		for _, match := range element.Stats {
			if match.Identifier == "bps" {
				var bps []bonusPointsCalc
				// fmt.Println(match.H[0].Element)
				// h := map[int]int{match.H[0].Element: match.H[0].Value, match.H[1].Element: match.H[1].Value, match.H[2].Element: match.H[2].Value, match.A[0].Element: match.A[0].Value, match.A[1].Element: match.A[1].Value, match.A[2].Element: match.A[2].Value}
				// a := map[int]int{match.A[0].Element: match.A[0].Value, match.A[1].Element: match.A[1].Value, match.A[2].Element: match.A[2].Value}
				// fmt.Println(h)
				for i := 0; i < 3; i++ {
					playerH := bonusPointsCalc{match.H[i].Element, match.H[i].Value}
					playerA := bonusPointsCalc{match.A[i].Element, match.A[i].Value}
					bps = append(bps, playerH)
					bps = append(bps, playerA)
				}
				sort.Slice(bps, func(i, j int) bool {
					return bps[i].Score > bps[j].Score
				})
				bps = bps[:3]
				if bps[0].Score == bps[1].Score {
					threeBp = append(threeBp, bps[0].ID)
					threeBp = append(threeBp, bps[1].ID)
				} else {
					threeBp = append(threeBp, bps[0].ID)
					twoBp = append(twoBp, bps[1].ID)
				}
				if bps[1].Score == bps[2].Score {
					twoBp = append(twoBp, bps[1].ID)
					twoBp = append(twoBp, bps[2].ID)
				} else {
					oneBp = append(oneBp, bps[2].ID)
				}
			}
		}
	}
	// fmt.Printf("%+v\n", bps)
	// bps = bps[:3]
	// fmt.Println(bps)
	fmt.Println("3 Bonus Points: ", threeBp)
	fmt.Println("2 Bonus Points: ", twoBp)
	fmt.Println("1 Bonus Points: ", oneBp)
}

func getLiveTotal(id int) int {
	client := &http.Client{}

	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/entry/%v/", id)

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

	var responseObject entry

	json.Unmarshal(body, &responseObject)

	return (responseObject.SummaryOverallPoints)
}

func getBenchPts(id, week int) int {
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

	var responseObject picks

	json.Unmarshal(body, &responseObject)

	return (responseObject.EntryHistory.PointsOnBench)

}

func getPrevTotal(id, week int) int {
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

	var responseObject picks

	json.Unmarshal(body, &responseObject)

	return (responseObject.EntryHistory.TotalPoints)

}

func getLeague(id, offset int) []row {
	client := &http.Client{}

	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/leagues-classic/%v/standings/", id)

	if offset > 1 {
		fmt.Println(("RUNNNG OFFSET API BIT"))
		apiURL = fmt.Sprintf("https://fantasy.premierleague.com/api/leagues-classic/%v/standings?page_standings=%v", id, offset)
	}

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
	var rows []row

	json.Unmarshal(body, &responseObject)

	// wg.Add(len(responseObject.Standings.Results))
	for _, element := range responseObject.Standings.Results {
		fmt.Println(element.EntryName)
		fmt.Println("Team ID: ", element.Entry)
		fmt.Println("Event Total: ", element.EventTotal)
		fmt.Println("Total: ", element.Total)
		// go func() {
		// 	getPicks(element.Entry, currentGw)
		// }()
		fmt.Println("---------")
		fmt.Println("---------")
		benchPts := getBenchPts(element.Entry, currentGw)
		prevTotal := getPrevTotal(element.Entry, currentGw-1)
		picks, captainPick := getPicks(element.Entry, currentGw)
		eventTotal := getLiveScore(picks, currentGw) + (getLiveScore([]int{captainPick}, currentGw) * 2)
		liveTotal := eventTotal + prevTotal
		captain := getCaptain(element.Entry, currentGw)
		picks = append(picks, captainPick)
		totalPlayed := hasPlayed(picks)
		result := row{element.RankSort, element.Entry, element.EntryName, eventTotal, liveTotal, prevTotal, element.LastRank, benchPts, captain, totalPlayed}
		rows = append(rows, result)
	}
	if responseObject.Standings.HasNext == true {
		if offset < 5 {
			fmt.Println(("RUNNING OFFSET BIT"))
			offset = offset + 1
			offsetResult := getLeague(id, offset)
			fmt.Println("OFFSET RESULT: ", offsetResult)

			rows = append(rows, offsetResult...)
			fmt.Println("Combined: ", rows)
		}
	}
	// wg.Wait()
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].LiveTotal > rows[j].LiveTotal
	})
	for i := range rows {
		rows[i].Rank = i + 1
	}
	fmt.Println(rows)
	return rows
}

func getNewLeagueEntries(id, offset int) []NewEntries {
	client := &http.Client{}

	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/leagues-classic/%v/standings/", id)
	if offset > 1 {
		fmt.Println(("RUNNNG OFFSET API BIT"))
		apiURL = fmt.Sprintf("https://fantasy.premierleague.com/api/leagues-classic/%v/standings?page_new_entries=%v", id, offset)
	}

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
	var newEntries []NewEntries

	json.Unmarshal(body, &responseObject)

	fmt.Println(responseObject.NewEntries.HasNext)

	for _, element := range responseObject.NewEntries.Results {
		// fmt.Println(element.EntryName)
		// fmt.Println("Team ID: ", element.Entry)
		// fmt.Println("Player First Name: ", element.PlayerFirstName)
		// fmt.Println("Player Second Name: ", element.PlayerLastName)
		// fmt.Println("Team Name: ", element.EntryName)
		// fmt.Println("---------")
		// fmt.Println("---------")
		result := NewEntries{element.Entry, element.EntryName, element.PlayerFirstName, element.PlayerLastName}
		newEntries = append(newEntries, result)
	}
	if responseObject.NewEntries.HasNext == true {
		if offset < 5 {
			fmt.Println(("RUNNING OFFSET BIT"))
			offset = offset + 1
			offsetResult := getNewLeagueEntries(id, offset)
			fmt.Println("OFFSET RESULT: ", offsetResult)

			newEntries = append(newEntries, offsetResult...)
			fmt.Println("Combined: ", newEntries)
		}
	}
	fmt.Println("Array and Len:")
	fmt.Println(newEntries)
	fmt.Println(len(newEntries))
	return newEntries
}

func getManagerInfo(id int) managerOutputPageData {
	client := &http.Client{}

	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/entry/%v/", id)

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
	var responseObject managerInfo
	var managerLeaguess []managerLeagues

	json.Unmarshal(body, &responseObject)

	for _, element := range responseObject.Leagues.Classic {
		fmt.Println("League ID: ", element.ID)
		fmt.Println("League Name: ", element.Name)
		fmt.Println("---------")
		fmt.Println("---------")
		result := managerLeagues{element.ID, element.Name}
		managerLeaguess = append(managerLeaguess, result)
	}

	managerPast := getManagerPast(id)

	managerOutput := managerOutputPageData{id, managerLeaguess, responseObject.PlayerFirstName, responseObject.PlayerLastName, responseObject.Name, managerPast, currentGw}

	return managerOutput
}

func getManagerPast(id int) managerPastData {
	client := &http.Client{}

	apiURL := fmt.Sprintf("https://fantasy.premierleague.com/api/entry/%v/history/", id)

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
	var responseObject managerPastData

	json.Unmarshal(body, &responseObject)
	return responseObject
}

func hasPlayed(ids []int) int {
	client := &http.Client{}
	var playersPlayed int
	for _, id := range ids {
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

		var responseObject player

		json.Unmarshal(body, &responseObject)

		for _, id := range responseObject.History {
			// if time.Parse(time.RFC3339Nano, id.KickoffTime).Before(time.Now()) {
			if id.KickoffTime.Before(time.Now()) {

				// if id.Minutes > 0 {
					playersPlayed++
				}
			// }
		}
	}
	return playersPlayed
}

func handler(w http.ResponseWriter, r *http.Request) {
        name := os.Getenv("NAME")
        if name == "" {
                name = "World"
        }
        fmt.Fprintf(w, "Hello %s!\n", name)
}
