package track

type Task struct {
	Id          int    `json:"-" db:"id"`
	UserId      int    `json:"user_id" db:"user_id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	EndTime     string `json:"end_time" db:"end_time"`
	StartTime   string `json:"start_time" db:"start_time"`
	DiffTime    string `json:"diff_time" db:"diff_time"`
	Done        bool   `json:"done"`
}
