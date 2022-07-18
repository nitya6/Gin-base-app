package models
type Category struct{
	ID int64 `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
type Tag struct{
	ID int64 `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
type Pet struct {
	ID int64 `json:"id" bson:"id"`
	Category Category `json:"category" bson:"category"`
	Name string `json:"name" bson:"name"`
	Tags []Tag `json:"tags" bson:"tags"`
	PhotoURLs []string `json:"photoUrls" bson:"photoUrls"`
	Status    string   `json:"status" bson:"status"`
}
