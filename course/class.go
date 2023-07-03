package course

import (
	"scheduler/account"
	"scheduler/util"
	"time"
)

// Class is an instance of a section
type Class struct {
	Meta
	SectionID  util.ID            `json:"section"`
	Order      int                `json:"order"`
	Title      string             `json:"title"`
	Time       time.Time          `json:"time"`
	Duration   time.Duration      `json:"duration"`
	Instructor account.Instructor `json:"instructor,omitempty"`
	Remark     string             `json:"remark,omitempty"`
}
