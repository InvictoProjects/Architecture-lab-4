package benchmark

import (
	"fmt"
	"github.com/InvictoProjects/Architecture-lab-4/tools"
	"math/rand"
	"strings"
	"testing"

	"github.com/InvictoProjects/Architecture-lab-4/commands"
)

var cntResult commands.Command
var argumentLen = 20
var maxRuneByte = 300

func BenchmarkParse_CommandAmount(b *testing.B) {
	const baseAmount = 5

	for i := 0; i < 20; i++ {
		l := baseAmount * 1 << i

		var commandsSlice []string
		for j := 0; j < l; j++ {
			commandsSlice = append(commandsSlice, generateCommand())
		}

		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			for k := 0; k < len(commandsSlice); k++ {
				cntResult = tools.Parse(commandsSlice[k])
			}
		})
	}
}

func generateCommand() string {
	if param := rand.Intn(2); param%2 == 0 {
		return "print " + generateArgument()
	} else {
		return "cat " + generateArgument() + " " + generateArgument()
	}
}

func generateArgument() string {
	res := strings.Builder{}
	for i := 0; i < argumentLen+1; i++ {
		res.WriteRune(rune(rand.Intn(maxRuneByte)))
	}
	return res.String()
}
