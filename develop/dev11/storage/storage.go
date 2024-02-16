package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	Day   = "day"
	Week  = "week"
	Month = "month"
)

type Event struct {
	TaskId int       `json:"task_id"`
	UserId int       `json:"user_id"`
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
}

type Storage struct {
	Events []Event `json:"events"`
	DBPath string
}

func InitStorage(dbPath string) (*Storage, error) {
	data, err := os.ReadFile(dbPath)
	if err != nil {
		return nil, err
	}

	var store Storage
	err = json.Unmarshal(data, &store)
	if err != nil {
		return nil, err
	}
	store.DBPath = dbPath

	return &store, nil
}

func (s *Storage) CreateEvent(evn Event) error {
	if evn.TaskId == 0 || evn.UserId == 0 || evn.Name == "" {
		return fmt.Errorf("user must have all filled fields")
	}

	for _, v := range s.Events {
		if v.UserId == evn.UserId && v.TaskId == evn.TaskId {
			return fmt.Errorf("task already exists")
		}
	}
	s.Events = append(s.Events, evn)
	return nil
}

func (s *Storage) UpdateEvent(evn Event) error {
	for k, v := range s.Events {
		if v.TaskId == evn.TaskId && v.UserId == evn.UserId {
			s.Events[k] = evn
			return nil
		}
	}
	return fmt.Errorf("task doesn't exist")
}

func (s *Storage) GetEvents(userId int, startDate time.Time, period string) []Event {
	outEvents := make([]Event, 0)
	var endDate time.Time
	switch period {
	case Day:
		endDate = startDate.AddDate(0, 0, 1)
	case Week:
		endDate = startDate.AddDate(0, 0, 7)
	case Month:
		endDate = startDate.AddDate(0, 1, 0)
	default:
		fmt.Println("unsupported period")
		return nil
	}

	for _, v := range s.Events {
		if v.UserId == userId && (startDate.Before(v.Date) && endDate.After(v.Date)) {
			outEvents = append(outEvents, v)
		}
	}
	return outEvents
}

func (s *Storage) DeleteEvent(userId int, taskId int) error {
	for k, v := range s.Events {
		if v.UserId == userId && v.TaskId == taskId {
			s.Events[k] = s.Events[len(s.Events)-1]
			s.Events[len(s.Events)-1] = Event{}
			s.Events = s.Events[:len(s.Events)-1]
			return nil
		}
	}
	return fmt.Errorf("task doesn't exist")
}

func (s *Storage) SaveToFile() error {
	data, err := json.Marshal(s.Events)
	if err != nil {
		return fmt.Errorf("marshaling error: %v", err)
	}

	resStr := "{\"events\":" + string(data) + "}"

	err = os.WriteFile(s.DBPath, []byte(resStr), os.ModePerm)
	if err != nil {
		return fmt.Errorf("save error: %v", err)
	}
	return nil
}
