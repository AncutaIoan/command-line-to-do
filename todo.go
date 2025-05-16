package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type todo struct {
	title       string
	completed   bool
	createdAt   time.Time
	completedAt *time.Time
}

type todos []todo

func (t *todos) add(title string) {
	newTodo := todo{
		title:     title,
		completed: false,
		createdAt: time.Now(),
	}
	*t = append(*t, newTodo)
}

func (t *todos) validateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		return fmt.Errorf("index %d out of range", index)
	}
	return nil
}

func (t *todos) delete(index int) error {
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*t = append((*t)[:index], (*t)[index+1:]...)
	return nil
}

func (t *todos) toggle(index int) error {
	if err := t.validateIndex(index); err != nil {
		return err
	}
	item := &(*t)[index]
	if !item.completed {
		now := time.Now()
		item.completedAt = &now
	} else {
		item.completedAt = nil
	}
	item.completed = !item.completed
	return nil
}

func (t *todos) printManual() {
	// calculate max width for each column
	maxIndexWidth := len("#")
	maxTitleWidth := len("title")
	maxCompletedWidth := len("completed")
	maxCreatedAtWidth := len("created at")
	maxCompletedAtWidth := len("completed at")

	rows := make([][]string, len(*t))
	for i, todo := range *t {
		completed := "❌"
		completedAt := ""
		if todo.completed {
			completed = "✅"
			if todo.completedAt != nil {
				completedAt = todo.completedAt.Format(time.RFC1123)
			}
		}

		indexStr := strconv.Itoa(i)
		createdAtStr := todo.createdAt.Format(time.RFC1123)

		rows[i] = []string{indexStr, todo.title, completed, createdAtStr, completedAt}

		if len(indexStr) > maxIndexWidth {
			maxIndexWidth = len(indexStr)
		}
		if len(todo.title) > maxTitleWidth {
			maxTitleWidth = len(todo.title)
		}
		if len(completed) > maxCompletedWidth {
			maxCompletedWidth = len(completed)
		}
		if len(createdAtStr) > maxCreatedAtWidth {
			maxCreatedAtWidth = len(createdAtStr)
		}
		if len(completedAt) > maxCompletedAtWidth {
			maxCompletedAtWidth = len(completedAt)
		}
	}

	headerFmt := fmt.Sprintf("%%-%ds | %%-%ds | %%-%ds | %%-%ds | %%-%ds\n",
		maxIndexWidth, maxTitleWidth, maxCompletedWidth, maxCreatedAtWidth, maxCompletedAtWidth)
	rowFmt := headerFmt

	fmt.Printf(headerFmt, "#", "title", "completed", "created at", "completed at")

	totalWidth := maxIndexWidth + maxTitleWidth + maxCompletedWidth + maxCreatedAtWidth + maxCompletedAtWidth + 12 // for separators
	fmt.Println(strings.Repeat("-", totalWidth))

	for _, row := range rows {
		fmt.Printf(rowFmt, row[0], row[1], row[2], row[3], row[4])
	}
}

