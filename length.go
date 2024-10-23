package passwordvalidator

const (
	seqNums      = "0123456789"
	seqKeyboard0 = "qwertyuiop"
	seqKeyboard1 = "asdfghjkl"
	seqKeyboard2 = "zxcvbnm"
	seqAlphabet  = "abcdefghijklmnopqrstuvwxyz"
)

func getLength(password string) int {
	password = removeRepeatingChars(password)
	password = removeSequence(password, seqNums)
	password = removeSequence(password, seqKeyboard2)
	password = removeSequence(password, seqKeyboard1)
	password = removeSequence(password, seqKeyboard2)
	password = removeSequence(password, seqAlphabet)
	password = removeSequence(password, getReversedString(seqNums))
	password = removeSequence(password, getReversedString(seqKeyboard0))
	password = removeSequence(password, getReversedString(seqKeyboard1))
	password = removeSequence(password, getReversedString(seqKeyboard2))
	password = removeSequence(password, getReversedString(seqAlphabet))
	return len(password)
}

func removeRepeatingChars(s string) string {
	var prevPrev rune
	var prev rune

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		rune := runes[i]
		if rune == prev && rune == prevPrev {
			runes = deleteRuneAt(runes, i)
			i--
		}
		prevPrev = prev
		prev = rune
	}

	return string(runes)
}

func removeSequence(s, seq string) string {
	seqRunes := []rune(seq)
	runes := []rune(s)
	mathces := 0

	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(seqRunes); j++ {
			if i >= len(runes) {
				break
			}

			r := runes[i]
			r2 := seqRunes[j]

			if r != r2 {
				mathces = 0
				continue
			}

			mathces++

			if mathces > 2 {
				runes = deleteRuneAt(runes, i)
			} else {
				i++
			}
		}
	}
	return string(runes)
}

func getReversedString(s string) string {
	n := 0
	runes := make([]rune, len(s))

	for _, rune := range s {
		runes[n] = rune
		n++
	}

	runes = runes[0:n]

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func deleteRuneAt(runes []rune, i int) []rune {
	if i >= len(runes) ||
		i < 0 {
		return runes
	}
	copy(runes[i:], runes[i+1:])
	runes[len(runes)-1] = 0
	runes = runes[:len(runes)-1]
	return runes
}
