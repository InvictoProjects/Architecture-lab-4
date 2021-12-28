package benchmark

import (
	"fmt"
	"strings"
	"testing"

	"github.com/InvictoProjects/Architecture-lab-4/commands"
	"github.com/InvictoProjects/Architecture-lab-4/tools"
)

var cntRes commands.Command

func BenchmarkParse(b *testing.B) {
	const baseLen = 5

	for i := 0; i < 25; i++ {
		l := baseLen * 1 << i

		commandLine := strings.Join([]string{"cat", strings.Repeat("5", l), strings.Repeat("5", l)}, " ")

		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
			cntRes = tools.Parse(commandLine)
		})
	}
}
