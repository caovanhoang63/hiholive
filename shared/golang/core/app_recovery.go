package core

import log "github.com/sirupsen/logrus"

// AppRecover is an intelligent function to recover from panic
func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
