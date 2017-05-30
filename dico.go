/*
 * @Author: Gaston Siffert
 * @Date: 2017-05-30 17:59:03
 * @Last Modified by: Gaston Siffert
 * @Last Modified time: 2017-05-30 20:48:25
 */
package morse

import "time"

type dico map[byte]*letter

func characters() dico {
	return dico{
		'A': &letter{code: ".-"},
		'B': &letter{code: "-..."},
		'C': &letter{code: "-.-."},
		'D': &letter{code: "-.."},
		'E': &letter{code: "."},
		'F': &letter{code: "..-."},
		'G': &letter{code: "--."},
		'H': &letter{code: "...."},
		'I': &letter{code: ".."},
		'J': &letter{code: ".---"},
		'K': &letter{code: "-.-"},
		'L': &letter{code: ".-.."},
		'M': &letter{code: "--"},
		'N': &letter{code: "-."},
		'O': &letter{code: "---"},
		'P': &letter{code: ".--."},
		'Q': &letter{code: "--.-"},
		'R': &letter{code: ".-."},
		'S': &letter{code: "..."},
		'T': &letter{code: "-"},
		'U': &letter{code: "..-"},
		'V': &letter{code: "...-"},
		'W': &letter{code: ".--"},
		'X': &letter{code: "-..-"},
		'Y': &letter{code: "-.--"},
		'Z': &letter{code: "--.."},

		'0': &letter{code: "-----"},
		'1': &letter{code: ".----"},
		'2': &letter{code: "..---"},
		'3': &letter{code: "...--"},
		'4': &letter{code: "....-"},
		'5': &letter{code: "....."},
		'6': &letter{code: "-...."},
		'7': &letter{code: "--..."},
		'8': &letter{code: "---.."},
		'9': &letter{code: "----."},
	}
}

func (d *dico) setup(unit time.Duration) {
	for _, value := range *d {
		value.setup(unit)
	}
}
