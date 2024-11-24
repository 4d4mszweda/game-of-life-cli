package quests

import (
	"sync"
)

var Log *questLog

type questLog struct {
	questLines *[]questLine
	mutex      sync.Mutex
}

type questLine struct {
	name   string
	quests []quest
}

type quest struct {
	title string
	desc  string
	check bool
}
