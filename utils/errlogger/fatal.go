package errlogger

import "github.com/rs/zerolog/log"

func FatalPanicMessage(msg string) {
	log.Fatal().Msg(msg)
}

func ErrFatalPanic(err error) {
	if err != nil {
		log.Fatal().Err(err)
		panic(err)
	}
}
