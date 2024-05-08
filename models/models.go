package models

import "time"

// Player struct

type Player struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"tag" gorm:"type:varchar(15)"`
	Rating    Rating    `json:"current_rating" gorm:"embedded;embeddedPrefix:rating_"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

// Rating struct to capture nested JSON data
type Rating struct {
	CurrentRating float32 `json:"rating" gorm:"column:rating_rating"`
	VsProtoss     float32 `json:"tot_vp" gorm:"column:rating_vs_protoss"`
	VsTerran      float32 `json:"tot_vt" gorm:"column:rating_vs_terran"`
	VsZerg        float32 `json:"tot_vz" gorm:"column:rating_vs_zerg"`
}

// Map struct
type GameMap struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Name         string    `json:"name" gorm:"type:varchar(30)"`
	Height       uint8     `json:"height" gorm:"type:uint"`
	Width        uint8     `json:"width" gorm:"type:uint"`
	RushDistance uint8     `json:"rush_distance" gorm:"type:uint"`
	TvZ          float32   `json:"tvz" gorm:"type:float"`
	ZvP          float32   `json:"zvp" gorm:"type:float"`
	PvT          float32   `json:"pvt" gorm:"type:float"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:datetime"`
}

// Match struct
type Match struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Player1   Player    `json:"player1_id" gorm:"type:uint"`
	Player2   Player    `json:"player2_id" gorm:"type:uint"`
	GameMap   GameMap   `json:"map_id" gorm:"type:uint"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

// Result struct
type Result struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Match     Match     `json:"match_id" gorm:"type:int"`
	Winner    Player    `json:"winner" gorm:"foreignKey:WinnerID"`
	Loser     Player    `json:"loser" gorm:"foreignKey:LoserID"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

type Bracket struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Players   []Player  `json:"players" gorm:"type:json"`
	Timestamp string    `json:"timestamp" gorm:"type:datetime"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

// TODO: ModelWeights struct to be incorporated later
type ModelWeights struct {
}
