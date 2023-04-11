package main

import "time"

func (c ChatMessage) DefaultInsertHook(q chatMessageInsertSQL) (chatMessageInsertSQL, error) {
	now := time.Now()
	return q.ValueCreatedAt(now.Unix()), nil
}
