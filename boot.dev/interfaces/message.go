package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	m := msg.getMessage()
	return m, len(m)
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main () {

	birthday := birthdayMessage{birthdayTime: time.Now(), recipientName: "John"}
	report := sendingReport{reportName: "Birthday", numberOfSends: 10}

	birthdayMessage, cost1 := sendMessage(birthday)
	sendingReport, cost2 := sendMessage(report)

	fmt.Println(birthdayMessage, cost1)
	fmt.Println(sendingReport, cost2)


}
