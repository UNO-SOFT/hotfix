package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"

	"github.com/UNO-SOFT/hotfix/models"
)

// EventsHandler is a default handler to serve up
// the events page.
func EventsHandler(c buffalo.Context) error {
	db := c.Value("tx").(*pop.Connection)
	qry := db.Order("f_when DESC")
	if with := c.Param("with"); with != "" {
		qry = qry.Where("f_with = ?", c.Param("with"))
	}
	var events []models.Event
	if err := qry.All(&events); err != nil {
		return fmt.Errorf("failed to load events: %w", err)
	}
	c.Set("events", events)
	return c.Render(http.StatusOK, r.HTML("events.html"))
}
