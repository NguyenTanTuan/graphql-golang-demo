package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nguyentantuan/graphql-golang-demo/graph/generated"
	"github.com/nguyentantuan/graphql-golang-demo/graph/model"
	"math/rand"
)

func (r *mutationResolver) CreateNote(ctx context.Context, input model.NewNote) (*model.Note, error) {
	note := &model.Note{
		Text:        input.Text,
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Description: input.Description,
	}
	r.notes = append(r.notes, note)
	return note, nil
}

func (r *queryResolver) Notes(ctx context.Context) ([]*model.Note, error) {
	return r.notes, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
