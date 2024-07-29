package gotodo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type items struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []items

func (t *Todos) Add(task string) {
	todo := items{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}
func (t *Todos) Compelete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		errors.New("Inappropirate index")
	}
	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true
	return nil

}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		errors.New("Inappropirate index")
	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil

}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
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

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) countPend() int {
	count := 0
	for _, item := range *t {
		if !item.Done {
      count++
		}
	}

    return count
}

func (t *Todos) Print() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "*"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
			{Align: simpletable.AlignCenter, Text: "Compelete At"},
		},
	}
	var cells [][]*simpletable.Cell
	for idx, item := range *t {
		idx++
		task := blue(item.Task)
    done := blue("no")
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
      done = green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}
	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("The tasks are here, where you at? You have %d Task remaining",t.countPend()))},
	}}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()

}
