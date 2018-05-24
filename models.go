package pagination

import (
	"time"
	"github.com/gobuffalo/uuid"
	"encoding/json"
	"github.com/gobuffalo/pop"
	"crypto/rand"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"encoding/base64"
)

type PageToken struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	Token      string     `json:"token" db:"token"`
	PageCursor time.Time  `json:"page_cursor" db:"page_cursor"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
}

// String is not required by pop and may be deleted
func (p PageToken) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// NewPageToken returns a pointer to a PageToken containing the provided cursor
func NewPageToken(cursor time.Time) *PageToken {
	return &PageToken{
		Token:      createTokenString(),
		PageCursor: cursor,
	}
}

func createTokenString() string {
	b := make([]byte, 8)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

// PageTokens is not required by pop and may be deleted
type PageTokens []PageToken

// String is not required by pop and may be deleted
func (p PageTokens) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *PageToken) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Token, Name: "Token"},
		&validators.TimeIsPresent{Field: p.PageCursor, Name: "PageCursor"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *PageToken) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *PageToken) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

