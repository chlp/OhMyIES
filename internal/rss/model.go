package rss

import (
	"ohmyies/internal/model"
	"strings"
)

type feed struct {
	Channel channel `xml:"channel"`
}

type channel struct {
	Items []item `xml:"item"`
}

type item struct {
	Guid        string `xml:"guid"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func getTypeByTitle(title string) model.MsgType {
	if strings.HasPrefix(title, "Attendance/Absence:") {
		return model.MsgTypeAttendance
	}
	if strings.HasPrefix(title, "Lesson planning:") {
		return model.MsgTypeLessonPlanning
	}
	return model.MsgTypeLessonGeneral
}
