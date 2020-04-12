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
	with := c.Param("with")
	if with != "" {
		qry = qry.Where("f_with = ?", with)
	}
	var events []models.Event
	if err := qry.All(&events); err != nil {
		return fmt.Errorf("failed to load events: %w", err)
	}
	c.Set("events", events)
	var votes []models.Vote
	qry = db.Order("created_at DESC")
	if with != "" {
		qry = qry.Where("name = ?", with)
	}
	if err := qry.All(&votes); err != nil {
		return fmt.Errorf("failed to load votes: %w", err)
	}
	c.Set("votes", votes)
	return c.Render(http.StatusOK, r.HTML("events.html"))
}
