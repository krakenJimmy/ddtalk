package main

import (
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("0 0 10 * * *", func() {
		word := GetWord()
		SendLinkMsg(word)
	})

	c.AddFunc("* * * * * *", func() {
		SendMsg("主人, 您该点外卖了")
	})

	c.Start()
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
