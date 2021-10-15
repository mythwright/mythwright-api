package db

type ItemMetadata struct {
	ItemID   int    `bson:"item_id"`
	Name     string `bson:"name"`
	Picture  string `bson:"picture"`
	ChatLink string `bson:"chat_link"`
	Type     string `bson:"type"`
}
