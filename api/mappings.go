package api

type ItemMappings struct {
	ItemID     int   `bson:"item_id"`
	ChildrenID []int `bson:"children_id"`
}

type ItemMetadata struct {
	ItemID  int    `bson:"item_id"`
	Name    string `bson:"name"`
	Picture string `bson:"picture"`
}
