package actions

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"

	"github.com/UNO-SOFT/hotfix/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("join", func(slice []string, sep string) string { return strings.Join(slice, sep) })
	db := c.Value("tx").(*pop.Connection)
	fixes, err := models.GetFixes(db)
	if err != nil {
		return fmt.Errorf("failed to load fixes: %w", err)
	}
	c.Set("fixes", fixes)
	return c.Render(http.StatusOK, r.HTML("index.html"))
}
