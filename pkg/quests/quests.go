package quests

import (
	"fmt"
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
	title       string
	description string
	check       bool
}

// AddQuestLine adds a new questLine to the questLog.
func (ql *questLog) AddQuestLine(name string) {
	ql.mutex.Lock()
	defer ql.mutex.Unlock()

	*ql.questLines = append(*ql.questLines, questLine{
		name:   name,
		quests: []quest{},
	})
}

// AddQuest adds a new quest to a specific questLine by name.
func (ql *questLog) AddQuest(lineName string, newQuest quest) error {
	ql.mutex.Lock()
	defer ql.mutex.Unlock()

	for i, line := range *ql.questLines {
		if line.name == lineName {
			(*ql.questLines)[i].quests = append(line.quests, newQuest)
			return nil
		}
	}

	return fmt.Errorf("questLine '%s' not found", lineName)
}

// ListQuests lists all quests in a specific questLine by name.
func (ql *questLog) ListQuests(lineName string) ([]quest, error) {
	ql.mutex.Lock()
	defer ql.mutex.Unlock()

	for _, line := range *ql.questLines {
		if line.name == lineName {
			return line.quests, nil
		}
	}

	return nil, fmt.Errorf("questLine '%s' not found", lineName)
}
