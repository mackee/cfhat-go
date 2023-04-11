package main

import "time"

func (c ChatMessage) CreatedAtStr() string {
	return time.Unix(c.CreatedAt, 0).String()
}
