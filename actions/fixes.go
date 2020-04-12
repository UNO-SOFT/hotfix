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
	v := models.Vote{CreatedAt: now, UpdatedAt: now}
	if err := c.Bind(&v); err != nil {
		return err
	}
	if verrs, err := db.ValidateAndCreate(&v); err != nil {
		return fmt.Errorf("validate(%+v): %w", v, err)
	} else if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("errors.html"))
	}
	c.Flash().Add("success", fmt.Sprintf("%s: %s", v.Name, v.Vote))
	return c.Redirect(302, "/")
}
