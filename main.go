package main

import (
	"time"
)
func main() {
	todos := todos{
		{
			title:     "Learn Go",
			completed: false,
			createdAt: time.Now(),
		},
		{
			title:     "Write blog post",
			completed: true,
			createdAt: time.Now().Add(-48 * time.Hour),
			completedAt: func() *time.Time {
				t := time.Now().Add(-24 * time.Hour)
				return &t
			}(),
		},
	}

	todos.printManual()
}
