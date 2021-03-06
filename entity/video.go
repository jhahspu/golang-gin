package entity

// Person as author for video
type Person struct {
	Name  string `json:"name" binding:"required"`
	Age   int8   `json:"age" binding:"gte=1,lte=130"`
	Email string `json:"email" binding:"required,email"`
}

// Video struct
type Video struct {
	Title       string `json:"title" binding:"min=2,max=10" validate:"is-cool"`
	Description string `json:"description" binding:"max=100"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
