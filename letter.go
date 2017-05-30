/*
 * @Author: Gaston Siffert
 * @Date: 2017-05-30 20:26:04
 * @Last Modified by: Gaston Siffert
 * @Last Modified time: 2017-05-30 20:48:25
 */
package morse

import (
	"log"
	"time"
)

var codeConverter = map[byte]time.Duration{
	'.': time.Duration(1),
	'-': time.Duration(3),
}

type letter struct {
	code  string
	morse []time.Duration
}

func (l *letter) setup(unit time.Duration) {
	l.morse = make([]time.Duration, len(l.code))

	for i := 0; i < len(l.code); i++ {
		if value, ok := codeConverter[l.code[i]]; ok {
			l.morse[i] = value * unit
		} else {
			log.Fatalf("Invalid character in code: '%s'\n", l.code)
		}
	}
}
