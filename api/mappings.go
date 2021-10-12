package api

import (
	"context"
)

type ItemMappings struct {
	ItemID     int   `bson:"item_id" json:"item_id,omitempty"`
	ChildrenID []int `bson:"children_id" json:"children_id,omitempty"`
}

type ItemMetadata struct {
	ItemID  int    `bson:"item_id"`
	Name    string `bson:"name"`
	Picture string `bson:"picture"`
}

func (s *Server) SaveItemMappings(ctx context.Context) error {
	return nil
}
