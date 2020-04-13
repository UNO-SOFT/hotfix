package actions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"

	"github.com/UNO-SOFT/hotfix/models"
)

var AllowedPubkeys = []string{
	NaCLPublicPrefix + `rmWa1cGJb38fTh/JFAVMP1H5G2f2jIk1qKG0kxxryEU=`,
}

var identitiesMu sync.Mutex
var identities map[[32]byte]struct{}

func ParseAllowedPubkeys() error {
	identitiesMu.Lock()
	if identities == nil {
		identities = make(map[[32]byte]struct{}, len(AllowedPubkeys))
	} else {
		for k := range identities {
			delete(identities, k)
		}
	}
	err := ParsePubkeys(identities, AllowedPubkeys)
	identitiesMu.Unlock()
	return err
}

func ParsePubkeys(m map[[32]byte]struct{}, keys []string) error {
	var k [32]byte
	var key []byte
	for _, s := range keys {
		var err error
		key, err = ParseKey(key[:0], []byte(s))
		if err != nil {
			return err
		}
		copy(k[:], key[:32])
		m[k] = struct{}{}
	}
	return nil
}

func PutEventHandler(c buffalo.Context) error {
	signer := c.Param("NaCL-Signer")
	pubKey, err := ParseKey(nil, []byte(signer))
	if err != nil {
		return err
	}
	var pk [32]byte
	copy(pk[:], pubKey[:32])

	identitiesMu.Lock()
	_, ok := identities[pk]
	identitiesMu.Unlock()
	if !ok {
		return fmt.Errorf("unknown sender % x", pk)
	}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}
	if body, ok = Open(body[:0], body, pubKey); !ok {
		return fmt.Errorf("signature mismatch")
	}
	db := c.Value("tx").(*pop.Connection)
	_ = db
	return nil
}

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
