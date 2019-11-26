package main

import (
	"github.com/securenative/packman/pkg"
	"time"
)

type ReplyModel struct {
	Flags     map[string]string
	Timestamp time.Time
}

func main() {
	flags := packman.ReadFlags()

	reply := ReplyModel{
		Flags:     flags,
		Timestamp: time.Now(),
	}

	packman.WriteReply(reply)
}
