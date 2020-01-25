package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

var usage = make(map[string]int64)

func bytesInUse(username string) int64 {
	return usage[username]
}

// Email sender configuration.
// Note: never put passwords in source code.
const sender = "notifications@example"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

// Quota exceeded message template.
const template = `Warning: you are using %d bytes of storage, %d%% of your quota.`

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000 * 1000 * 1000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return // OK
	}

	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}