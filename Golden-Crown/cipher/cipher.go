package cipher

import (
	"fmt"
	"strings"
	"unicode"
)

// Decrypt  is used to decrypt the message
func Decrypt(key int, message string) string {

	if key > 26 {
		key = key % 26
	}

	var (
		decryptedMessage strings.Builder
		val              int
	)
	decryptedMessage.Grow(len(message))

	for _, char := range message {
		asciiVal := int(char)

		if unicode.IsUpper(char) {
			val = (asciiVal+26-key-65)%26 + 65
		} else {
			val = (asciiVal+26-key-97)%26 + 97
		}

		fmt.Fprintf(&decryptedMessage, "%c", rune(val))
	}

	str := decryptedMessage.String()

	return str
}
