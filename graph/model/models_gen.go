// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Courses     []*Course `json:"courses"`
}

type Chapter struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Course   *Course   `json:"course"`
	Category *Category `json:"category"`
}

type Course struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Category    *Category  `json:"category"`
	Chapters    []*Chapter `json:"chapters"`
}

type NewCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewChapter struct {
	Name     string `json:"name"`
	CourseID string `json:"courseId"`
}

type NewCourse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  string `json:"categoryId"`
}
