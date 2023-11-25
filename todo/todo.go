package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Item struct {
	Task      string
	Done      bool
	Created   time.Time
	Completed time.Time
}

type Todos []Item

func (t *Todos) Add(task string) {
	todo := Item{
		Task:      task,
		Created:   time.Now(),
		Completed: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Completed(index int) error {
	ls := *t

	if index < 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].Completed = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t

	if index < 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[:index]...)

	return nil
}

func (t *Todos) LoadFile(filename string) error {
	file, err := os.ReadFile("sample.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) Store(fileName string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
