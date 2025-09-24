package sportsblaze

type DailyBoxScores struct {
	League  League `json:"league"`
	Games   []Game `json:"games"`
	Updated string `json:"updated"`
}

type League struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Sport string `json:"sport"`
}

type Game struct {
	Season     Season      `json:"season"`
	ID         string      `json:"id"`
	Teams      Teams       `json:"teams"`
	Date       string      `json:"date"`
	Status     string      `json:"status"`
	Venue      Venue       `json:"venue"`
	Officials  []Official  `json:"officials"`
	Broadcasts []Broadcast `json:"broadcasts"`
	Scores     Scores      `json:"scores"`
	Stats      TeamStats   `json:"stats"`
	Rosters    Rosters     `json:"rosters"`
}

type Season struct {
	Year int    `json:"year"`
	Type string `json:"type"`
}

type Teams struct {
	Away Team `json:"away"`
	Home Team `json:"home"`
}

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Venue struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Official struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Type   string `json:"type"`
}

type Broadcast struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Country string `json:"country"`
	Market  string `json:"market"`
}

type Scores struct {
	Periods []Period `json:"periods"`
	Total   Total    `json:"total"`
}

type Period struct {
	Number int    `json:"number"`
	Type   string `json:"type"`
	Away   Score  `json:"away"`
	Home   Score  `json:"home"`
}

type Score struct {
	Points int `json:"points"`
}

type Total struct {
	Away Score `json:"away"`
	Home Score `json:"home"`
}

type TeamStats struct {
	Away Stats `json:"away"`
	Home Stats `json:"home"`
}

type Stats struct {
	Assists                     int     `json:"assists"`
	AssistsTurnoverRatio        float64 `json:"assists_turnover_ratio"`
	BenchPoints                 int     `json:"bench_points"`
	BiggestLead                 int     `json:"biggest_lead"`
	BiggestLeadScore            string  `json:"biggest_lead_score"`
	BiggestScoringRun           int     `json:"biggest_scoring_run"`
	BiggestScoringRunScore      string  `json:"biggest_scoring_run_score"`
	Blocks                      int     `json:"blocks"`
	BlocksReceived              int     `json:"blocks_received"`
	FastBreakPointsAttempts     int     `json:"fast_break_points_attempts"`
	FastBreakPointsMade         int     `json:"fast_break_points_made"`
	FastBreakPointsPct          float64 `json:"fast_break_points_pct"`
	FieldGoalsAttempts          int     `json:"field_goals_attempts"`
	FieldGoalsEffectiveAdjusted float64 `json:"field_goals_effective_adjusted"`
	FieldGoalsMade              int     `json:"field_goals_made"`
	FieldGoalsPct               float64 `json:"field_goals_pct"`
	FoulsDrawn                  int     `json:"fouls_drawn"`
	FoulsOffensive              int     `json:"fouls_offensive"`
	FoulsPersonal               int     `json:"fouls_personal"`
	FoulsTeam                   int     `json:"fouls_team"`
	FoulsTeamTechnical          int     `json:"fouls_team_technical"`
	FoulsTechnical              int     `json:"fouls_technical"`
	FreeThrowsAttempts          int     `json:"free_throws_attempts"`
	FreeThrowsMade              int     `json:"free_throws_made"`
	FreeThrowsPct               float64 `json:"free_throws_pct"`
	LeadChanges                 int     `json:"lead_changes"`
	Minutes                     int     `json:"minutes"`
	Points                      int     `json:"points"`
	PointsAgainst               int     `json:"points_against"`
	PointsFastBreak             int     `json:"points_fast_break"`
	PointsFromTurnovers         int     `json:"points_from_turnovers"`
	PointsInThePaint            int     `json:"points_in_the_paint"`
	PointsInThePaintAttempts    int     `json:"points_in_the_paint_attempts"`
	PointsInThePaintMade        int     `json:"points_in_the_paint_made"`
	PointsInThePaintPct         float64 `json:"points_in_the_paint_pct"`
	PointsSecondChance          int     `json:"points_second_chance"`
	ReboundsDefensive           int     `json:"rebounds_defensive"`
	ReboundsOffensive           int     `json:"rebounds_offensive"`
	ReboundsPersonal            int     `json:"rebounds_personal"`
	ReboundsTeam                int     `json:"rebounds_team"`
	ReboundsTeamDefensive       int     `json:"rebounds_team_defensive"`
	ReboundsTeamOffensive       int     `json:"rebounds_team_offensive"`
	Rebounds                    int     `json:"rebounds"`
	SecondChancePointsAttempts  int     `json:"second_chance_points_attempts"`
	SecondChancePointsMade      int     `json:"second_chance_points_made"`
	SecondChancePointsPct       float64 `json:"second_chance_points_pct"`
	Steals                      int     `json:"steals"`
	ThreePointersAttempts       int     `json:"three_pointers_attempts"`
	ThreePointersMade           int     `json:"three_pointers_made"`
	ThreePointersPct            float64 `json:"three_pointers_pct"`
	TimeLeading                 string  `json:"time_leading"`
	TimeOnCourt                 string  `json:"time_on_court"`
	TimesTied                   int     `json:"times_tied"`
	TrueShootingAttempts        float64 `json:"true_shooting_attempts"`
	TrueShootingPct             float64 `json:"true_shooting_pct"`
	TurnoversPersonal           int     `json:"turnovers_personal"`
	TurnoversTeam               int     `json:"turnovers_team"`
	Turnovers                   int     `json:"turnovers"`
	TwoPointersAttempts         int     `json:"two_pointers_attempts"`
	TwoPointersMade             int     `json:"two_pointers_made"`
	TwoPointersPct              float64 `json:"two_pointers_pct"`
}

type Rosters struct {
	Away []Player `json:"away"`
	Home []Player `json:"home"`
}

type Player struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Position string      `json:"position,omitempty"`
	Number   string      `json:"number"`
	Played   bool        `json:"played"`
	Started  bool        `json:"started,omitempty"`
	Stats    PlayerStats `json:"stats"`
}

type PlayerStats struct {
	Assists               int     `json:"assists"`
	Blocks                int     `json:"blocks"`
	BlocksReceived        int     `json:"blocks_received"`
	FieldGoalsAttempts    int     `json:"field_goals_attempts"`
	FieldGoalsMade        int     `json:"field_goals_made"`
	FieldGoalsPct         float64 `json:"field_goals_pct"`
	FoulsDrawn            int     `json:"fouls_drawn"`
	FoulsOffensive        int     `json:"fouls_offensive"`
	FoulsPersonal         int     `json:"fouls_personal"`
	FoulsTechnical        int     `json:"fouls_technical"`
	FreeThrowsAttempts    int     `json:"free_throws_attempts"`
	FreeThrowsMade        int     `json:"free_throws_made"`
	FreeThrowsPct         float64 `json:"free_throws_pct"`
	Minus                 int     `json:"minus"`
	Minutes               int     `json:"minutes"`
	Plus                  int     `json:"plus"`
	PlusMinus             int     `json:"plus_minus"`
	Points                int     `json:"points"`
	PointsFastBreak       int     `json:"points_fast_break"`
	PointsInThePaint      int     `json:"points_in_the_paint"`
	PointsSecondChance    int     `json:"points_second_chance"`
	ReboundsDefensive     int     `json:"rebounds_defensive"`
	ReboundsOffensive     int     `json:"rebounds_offensive"`
	Rebounds              int     `json:"rebounds"`
	Steals                int     `json:"steals"`
	ThreePointersAttempts int     `json:"three_pointers_attempts"`
	ThreePointersMade     int     `json:"three_pointers_made"`
	ThreePointersPct      float64 `json:"three_pointers_pct"`
	TimeOnCourt           string  `json:"time_on_court"`
	TurnoversPersonal     int     `json:"turnovers_personal"`
	TwoPointersAttempts   int     `json:"two_pointers_attempts"`
	TwoPointersMade       int     `json:"two_pointers_made"`
	TwoPointersPct        float64 `json:"two_pointers_pct"`
}
