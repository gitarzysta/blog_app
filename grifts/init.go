package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gitarzysta/blog_app/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
