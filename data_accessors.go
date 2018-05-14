package pagination

import (
	"github.com/gobuffalo/pop"
	"context"
	"errors"
	"time"
)

// PageTokenDataAccessor provides access to the Page Token database
type PageTokenDataAccessor interface {
	Create(context.Context, *PageToken) error
	FindByToken(context.Context, string) (*PageToken, error)
}

// PageTokenDB implements PageTokenDataAccessor and wraps a db connection
type PageTokenDB struct {
	tx *pop.Connection
}

// NewPageTokenDB returns a PageTokenDB with the provided db connection
func NewPageTokenDB(tx *pop.Connection) *PageTokenDB {
	return &PageTokenDB{tx}
}

// Create creates a new PageToken in the database, returning validation or save errors
func (ptdb *PageTokenDB) Create(ctx context.Context, pageToken *PageToken) error {
	vErrs, err := ptdb.tx.ValidateAndCreate(pageToken)

	if err != nil {
		return err
	}

	if vErrs.HasAny() {
		return vErrs
	}

	return nil
}

// FindByToken returns the PageToken corresponding to the provided token
func (ptdb *PageTokenDB) FindByToken(ctx context.Context, token string) (*PageToken, error) {
	tokens := []*PageToken{}
	//err := ptdb.tx.Scope(NotDeleted()).Where("token = ?", token).All(&tokens)
	err := ptdb.tx.Where("token = ?", token).All(&tokens)
	if err != nil {
		//if IsPopRecordNotFound(err) {
		//	return nil, ErrRecordNotFound
		//}
		return nil, err
	}
	if len(tokens) > 1 {
		return nil, errors.New("multiple tokens found")
	}

	return tokens[0], nil
}

func GetCursor(ctx context.Context, db PageTokenDataAccessor, token string) (*time.Time, error) {
	var toke *PageToken
	var err error
	if token != "" {
		toke, err = db.FindByToken(ctx, token)
		if err != nil {
			return nil, err
		}
	}

	var cursor *time.Time
	if toke == nil {
		cursor = nil
	} else {
		cursor = &toke.PageCursor
	}
	return cursor, nil
}
