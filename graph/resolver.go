package graph

import (
	"bipper-backend/graph/model"
)

type Resolver struct {
	addedBeepChannel map[string]chan *model.AddedBeep
}

func NewResolver() *Resolver {
	return &Resolver{
		addedBeepChannel: map[string]chan *model.AddedBeep{},
	}
}
