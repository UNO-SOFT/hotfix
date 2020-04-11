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
	vote, err := models.ParseFixState(c.Param("vote"))
	if err != nil {
		return err
	}
	now := time.Now()
	v := models.Vote{
		Name: c.Param("name"), Vote: vote,
		CreatedAt: now, UpdatedAt: now,
		Author: c.Param("author")}
	if verrs, err := db.ValidateAndCreate(&v); err != nil {
		return fmt.Errorf("validate(%+v): %w", v, err)
	} else if verrs.HasAny() {
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("errors.html"))
	}
	c.Flash().Add("success", fmt.Sprintf("%s: %s", v.Name, vote))
	return c.Redirect(302, "/")
}
