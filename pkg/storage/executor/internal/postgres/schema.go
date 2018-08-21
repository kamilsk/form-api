package postgres

import (
	"context"
	"database/sql"
	"encoding/xml"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/query"
)

// NewSchemaContext TODO
func NewSchemaContext(ctx context.Context, conn *sql.Conn) schemaScope {
	return schemaScope{ctx, conn}
}

type schemaScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Create TODO
func (scope schemaScope) Create(token *query.Token, data query.CreateSchema) (query.Schema, error) {
	entity := query.Schema{AccountID: token.User.AccountID, Title: data.Title, Definition: data.Definition}
	encoded, encodeErr := xml.Marshal(entity.Definition)
	if encodeErr != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, encodeErr,
			"user %q of account %q tried to marshal definition `%#v` of a schema into XML",
			token.UserID, token.User.AccountID, entity.Definition)
	}
	q := `INSERT INTO "schema" ("id", "account_id", "title", "definition") VALUES ($1, $2, $3, $4)
	      RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, entity.AccountID, entity.Title, encoded)
	if err := row.Scan(&entity.ID, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to create a schema %q",
			token.UserID, token.User.AccountID, entity.Title)
	}
	return entity, nil
}

// Read TODO
func (scope schemaScope) Read(token *query.Token, data query.ReadSchema) (query.Schema, error) {
	entity, encoded := query.Schema{ID: data.ID, AccountID: token.User.AccountID}, []byte(nil)
	q := `SELECT "title", "definition", "created_at", "updated_at", "deleted_at" FROM "schema"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if err := row.Scan(&entity.Title, &encoded, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the schema %q", token.UserID, token.User.AccountID, entity.ID)
	}
	if err := xml.Unmarshal(encoded, &entity.Definition); err != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, err,
			"trying to unmarshal XML `%s` of the schema definition %q",
			encoded, entity.ID)
	}
	return entity, nil
}

// ReadByID TODO
func (scope schemaScope) ReadByID(id domain.ID) (query.Schema, error) {
	entity, encoded := query.Schema{ID: id}, []byte(nil)
	q := `SELECT "title", "definition", "created_at", "updated_at" FROM "schema"
	       WHERE "id" = $1 AND "deleted_at" IS NULL`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID)
	if err := row.Scan(&entity.Title, &encoded, &entity.CreatedAt, &entity.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return entity, errors.NotFound(errors.SchemaNotFoundMessage, err, "the schema %q not found", entity.ID)
		}
		return entity, errors.Database(errors.ServerErrorMessage, err, "trying to populate the schema %q", entity.ID)
	}
	if err := xml.Unmarshal(encoded, &entity.Definition); err != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, err,
			"trying to unmarshal XML `%s` of the schema definition %q",
			encoded, entity.ID)
	}
	return entity, nil
}

// Update TODO
func (scope schemaScope) Update(token *query.Token, data query.UpdateSchema) (query.Schema, error) {
	entity, readErr := scope.Read(token, query.ReadSchema{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	if data.Title != "" {
		entity.Title = data.Title
	}
	if !data.Definition.IsEmpty() {
		entity.Definition = data.Definition
	}
	encoded, encodeErr := xml.Marshal(entity.Definition)
	if encodeErr != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, encodeErr,
			"user %q of account %q tried to marshal definition `%#v` of the schema %q into XML",
			token.UserID, token.User.AccountID, entity.Definition, entity.ID)
	}
	q := `UPDATE "schema" SET "title" = $1, "definition" = $2
	       WHERE "id" = $3 AND "account_id" = $4
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.Title, encoded, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.UpdatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to update the schema %q",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Delete TODO
func (scope schemaScope) Delete(token *query.Token, data query.DeleteSchema) (query.Schema, error) {
	entity, readErr := scope.Read(token, query.ReadSchema{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	if data.Permanently {
		q := `DELETE FROM "schema" WHERE "id" = $1 AND "account_id" = $2 RETURNING now()`
		row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
		if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
			return entity, errors.Database(errors.ServerErrorMessage, scanErr,
				"user %q of account %q tried to delete the schema %q permanently",
				token.UserID, token.User.AccountID, entity.ID)
		}
		return entity, nil
	}
	q := `UPDATE "schema" SET "deleted_at" = now()
	       WHERE "id" = $1 AND "account_id" = $2
	   RETURNING "deleted_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to delete the schema %q safely",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}
