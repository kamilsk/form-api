package postgres_test

import (
	"context"
	"testing"

	"go.octolab.org/ecosystem/forma/internal/storage/executor"
	. "go.octolab.org/ecosystem/forma/internal/storage/executor/internal/postgres"
)

func TestSchemaEditor(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.SchemaEditor = NewSchemaContext(ctx, nil)
	})
	t.Run("read", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.SchemaEditor = NewSchemaContext(ctx, nil)
	})
	t.Run("update", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.SchemaEditor = NewSchemaContext(ctx, nil)
	})
	t.Run("delete", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.SchemaEditor = NewSchemaContext(ctx, nil)
	})
}

func TestSchemaReader(t *testing.T) {
	t.Run("read by ID", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.SchemaReader = NewSchemaContext(ctx, nil)
	})
}
