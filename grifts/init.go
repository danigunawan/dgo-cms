package grifts

import (
	"github.com/Danigunawan/indiface/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
