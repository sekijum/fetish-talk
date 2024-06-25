package model

import (
	"time"
)

type Thread struct {
	Id                int
	BoardId           int
	UserId            int
	Title             string
	Description       *string
	ThumbnailUrl      *string
	IsAutoGenerated   bool
	IsNotifyOnComment bool
	IpAddress         string
	Status            ThreadStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Board             *Board
}

type ThreadStatus int

const (
	ThreadStatusOpen ThreadStatus = iota
	ThreadStatusArchived
	ThreadStatusDeleted
)

func (s ThreadStatus) String() string {
	switch s {
	case ThreadStatusOpen:
		return "Open"
	case ThreadStatusArchived:
		return "Archived"
	case ThreadStatusDeleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}
