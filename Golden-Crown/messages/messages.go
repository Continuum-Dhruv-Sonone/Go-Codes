package messages

import (
	"strings"

	"github.com/ContinuumLLC/GO/Golden-Crown/cipher"
	"github.com/ContinuumLLC/GO/Golden-Crown/kingdom"
)

const (
	rulerKingdom = "SPACE"
	noAlliance   = "NONE"
)

// ProcessMessages processes each messages and returns error
func ProcessMessages(inputs [][]string) string {
	alliances := make([]string, 1)
	alliances[0] = rulerKingdom
	for _, input := range inputs {

		name := input[0]

		message := strings.Join(input[1:], "")
		emblem := kingdom.GetEmblem(name)

		msg := cipher.Decrypt(len(emblem), message)

		if checkAlliance(emblem, msg) {
			alliances = append(alliances, name)
		}
	}

	if len(alliances) < 4 {
		return noAlliance
	}
	return strings.Join(alliances, " ")
}

func checkAlliance(emblem, message string) bool {
	emblem = strings.ToLower(emblem)
	message = strings.ToLower(message)

	emblemCount := make(map[int]int)
	messageCount := make(map[int]int)

	for _, val := range emblem {
		emblemCount[int(val)]++

		messageCount[int(val)] = strings.Count(message, string(val))
	}

	if compare(emblemCount, messageCount) {
		return true
	}
	return false
}

func compare(emblem, message map[int]int) bool {

	if len(emblem) != len(message) {
		return false
	}

	for key, val := range message {
		if _, ok := emblem[key]; ok {
			if val < emblem[key] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
