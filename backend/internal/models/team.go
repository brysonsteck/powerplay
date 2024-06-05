package models

type Team struct {
	DbModel
	CorrelationId uint   `json:"correlation_id"`
	Name          string `json:"name"`
	LogoId        uint   `json:"logo_id"`
	Color         string `json:"color"`
	LeagueID      uint   `json:"league_id"`
	League        League `json:"league"`
	Roster        Roster `json:"roster"`
	RosterID      uint   `json:"roster_id"`

	Wins   int `json:"wins"`
	Losses int `json:"losses"`
}
