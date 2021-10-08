package api

import "context"

func (s *Server) LoadItems() {

}

type DroppedItem struct {
	ItemID          int     `bson:"item_id" json:"item_id,omitempty"`
	QuantityDropped int     `bson:"quantity_dropped" json:"quantity_dropped,omitempty"`
	SourceItemID    int     `bson:"source_item_id" json:"source_item_id,omitempty"`
	PriceForSample  float64 `bson:"price_for_sample" json:"price_for_sample,omitempty"`
}

func (s *Server) PostInputs(ctx context.Context) {

}
