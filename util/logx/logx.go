package logx

import "log"

func MustNil(err error) {
	if err == nil {
		return
	}

	log.Fatal(err)
}
