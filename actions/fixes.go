package actions

import (
	"fmt"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"

	"github.com/UNO-SOFT/hotfix/models"
)

// FixHandler is a default handler to serve up
// the events page.
func FixHandler(c buffalo.Context) error {
	db := c.Value("tx").(*pop.Connection)
	now := time.Now()
	e := models.Event{
		With: c.Param("with"), What: c.Param("what"),
		CreatedAt: now, UpdatedAt: now, When: now,
		Where: c.Param("what")[:1] + "___"}
	if verrs, err := db.ValidateAndCreate(&e); err != nil {
		return fmt.Errorf("validate(%+v): %w", e, err)
	} else if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("index.html"))
	}
	c.Flash().Add("success", fmt.Sprintf("%s: %s", e.With, e.What))
	return c.Redirect(302, "/")
}
