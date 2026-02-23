package todo

type TodoList struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
type UserList struct {
	ID     int
	UserId int
	ListId int
}
type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
type ListsItem struct {
	ID      int
	ListsId int
	ItemId  int
}
