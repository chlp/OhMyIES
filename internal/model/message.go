package model

import (
	"time"
)

type MsgType string

const (
	MsgTypeAttendance     MsgType = "attendance"
	MsgTypeLessonPlanning MsgType = "lesson_planning"
	MsgTypeLessonGeneral  MsgType = "general"
)

type Msg struct {
	Guid        string
	Title       string
	Description string
	PubDate     time.Time
	Type        MsgType
}
