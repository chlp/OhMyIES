package model

import (
	"time"
)

type MsgType string

const (
	MsgTypeLessonGeneral  MsgType = "general"
	MsgTypeAttendance     MsgType = "attendance"
	MsgTypeLessonPlanning MsgType = "lesson_planning"
)

type Msg struct {
	Guid        string
	Title       string
	Description string
	PubDate     time.Time
	Type        MsgType
}
