/*
 * @Author: Gaston Siffert
 * @Date: 2017-05-30 17:57:52
 * @Last Modified by: Gaston Siffert
 * @Last Modified time: 2017-05-30 21:39:19
 */
package morse

import (
	"strings"
	"time"
)

// Morse is a data structure defining the Morse algorithm to translate
// a string in ISO basic Latin alphabet into morse code
type Morse struct {
	dico dico
	unit time.Duration
}

// New create and initialize the Morse structure
func New(unit time.Duration) *Morse {
	m := new(Morse)
	m.unit = unit
	m.dico = characters()
	m.dico.setup(unit)
	return m
}

const (
	spaceInLetter      = time.Duration(1)
	spaceBetweenLetter = time.Duration(3)
	spaceBetweenWords  = time.Duration(7)
)

func (l *letter) print(unit time.Duration, in func(), out func()) {

	for i, duration := range l.morse {
		in()
		time.Sleep(duration)
		out()

		if i+1 < len(l.morse) {
			time.Sleep(spaceInLetter * unit)
		}
	}
}

// Translate will call back the in function when entering a morse phase
// and will call the out function when leaving the phase.
// It will handle internally the duration each phase must run.
func (m Morse) Translate(data string, in func(), out func()) {

	upperCase := strings.ToUpper(data)
	for i := 0; i < len(upperCase); i++ {
		if upperCase[i] == ' ' {
			time.Sleep(spaceBetweenWords * m.unit)
		} else if letter, ok := m.dico[upperCase[i]]; ok {
			letter.print(m.unit, in, out)
		}

		if i+1 < len(upperCase) && upperCase[i+1] != ' ' {
			time.Sleep(spaceBetweenLetter * m.unit)
		}
	}
}
