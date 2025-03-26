package models

type StatusReturn struct {
	Status string `json:"status"`
}

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostReturn struct {
	Posts []Post `json:"posts"`
}

type GreetArgs struct {
	Message string `json:"message"`
}

type GreetReturn struct {
	Message string `json:"message"`
}

type AskArgs struct {
	Question string `json:"question"`
}

type AskReturn struct {
	Answer string `json:"answer"`
}

type ErrorReturn struct {
	Error string `json:"error"`
}
