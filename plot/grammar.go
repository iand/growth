package plot

import (
	"bytes"
	"fmt"
)

type Grammar struct {
	Seed  string
	Rules map[rune]string
}

func (g *Grammar) Iterate(n int, complexityLimit int) (string, error) {
	s := g.Seed
	for i := 0; i < n; i++ {
		s = g.Expand(s)
		if len(s) > complexityLimit {
			return "", fmt.Errorf("ruleset is too complex (%d steps), try lowering the number of iterations", len(s))
		}
	}
	return s, nil
}

func (g *Grammar) Expand(s string) string {
	var buf bytes.Buffer
	for _, c := range s {
		if r, ok := g.Rules[c]; ok {
			buf.WriteRune('<')
			buf.WriteString(r)
			buf.WriteRune('>')
		} else {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}
