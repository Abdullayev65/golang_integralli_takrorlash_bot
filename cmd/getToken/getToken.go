package getToken

import (
	"bufio"
	"os"
	"strings"
)

// TokenFromConsole token hafsizligi uchun uni consoledan qabul qilish
func TokenFromConsole() string {
	reader := bufio.NewReader(os.Stdin)
	//ok
	token, _ := reader.ReadString('\n')
	token = strings.Replace(token, "\n", "", -1)
	return token
}
