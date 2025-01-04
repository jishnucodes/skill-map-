package common



type PostCreationInput struct {
	Title      string `json:"title"` //this is called field tag (the format of object)
	Content    string `json:"content"`
	Author     string `json:"author"`
}

type PostUpdateInput struct {
	Title      string `json:"title"` //this is called field tag (the format of object)
	Content    string `json:"content"`
	Author     string `json:"author"`
}

func NewPostCreationInput() *PostCreationInput {
	return &PostCreationInput{}
}

func NewPostUpdateInput() *PostUpdateInput {
	return &PostUpdateInput{}
}



