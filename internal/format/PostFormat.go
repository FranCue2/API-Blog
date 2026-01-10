package format


type PostFormat struct {
	Title       string `json:"title" bson:"title"`
	Content     string `json:"content" bson:"content"`
	Author      string `json:"author" bson:"author"`
	PublishedAt string `json:"published_at" bson:"published_at"`
}