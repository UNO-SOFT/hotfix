package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"

	"github.com/UNO-SOFT/hotfix/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	db := c.Value("tx").(*pop.Connection)
	var events []models.Event
	if err := db.All(&events); err != nil {
		return fmt.Errorf("failed to load events: %w", err)
	}
	c.Set("events", events)
	return c.Render(http.StatusOK, r.HTML("index.html"))
}
