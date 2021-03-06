package storage

import (
	"context"

	"go.octolab.org/ecosystem/forma/internal/domain"
	"go.octolab.org/ecosystem/forma/internal/storage/query"
	"go.octolab.org/ecosystem/forma/internal/storage/types"
)

// Schema returns the form schema by provided ID.
func (storage *Storage) Schema(ctx context.Context, id domain.ID) (domain.Schema, error) {
	var schema domain.Schema

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return schema, connErr
	}
	defer func() { _ = closer() }()

	entity, readErr := storage.exec.SchemaReader(ctx, conn).ReadByID(id)
	if readErr != nil {
		return schema, readErr
	}

	// TODO issue#logic duplicated
	{
		ptr := &entity.Definition
		ptr.ID = entity.ID.String()
		ptr.Title = entity.Title
		for i, input := range ptr.Inputs {
			ptr.Inputs[i].ID = ptr.ID + "_" + input.Name
		}
	}

	return entity.Definition, nil
}

// Template returns the form template by provided ID.
func (storage *Storage) Template(ctx context.Context, id domain.ID) (domain.Template, error) {
	var template domain.Template

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return template, connErr
	}
	defer func() { _ = closer() }()

	entity, readErr := storage.exec.TemplateReader(ctx, conn).ReadByID(id)
	if readErr != nil {
		return template, readErr
	}
	return entity.Definition, nil
}

// StoreInput stores an user input data.
func (storage *Storage) StoreInput(ctx context.Context, schemaID domain.ID, verified domain.InputData) (*types.Input, error) {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return nil, connErr
	}
	defer func() { _ = closer() }()

	entity, writeErr := storage.exec.InputWriter(ctx, conn).Write(query.WriteInput{SchemaID: schemaID, VerifiedData: verified})
	if writeErr != nil {
		return nil, writeErr
	}
	return &entity, nil
}

// LogInput stores an input event.
func (storage *Storage) LogInput(ctx context.Context, event domain.InputEvent) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer func() { _ = closer() }()

	// TODO issue#51
	_, writeErr := storage.exec.LogWriter(ctx, conn).Write(query.WriteLog{InputEvent: event})

	return writeErr
}
