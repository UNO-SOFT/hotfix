package grifts

import (
	"github.com/UNO-SOFT/hotfix/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
