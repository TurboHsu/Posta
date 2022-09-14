package util

import "log"

func LogFatal(m string) {
	log.Println("[FATAL] ", m)
	panic("")
}

func LogInfo(m string) {
	log.Println("[I] ", m)
}

func LogError(err error) {
	log.Panicln("[E] ", err.Error())
}
