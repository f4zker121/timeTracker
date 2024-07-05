package track

type People struct {
	Id             int    `json:"-" db:"id"`
	PassportSerie  int    `json:"passport_serie" binding:"required" db:"passport_serie"`
	PassportNumber int    `json:"passport_number" binding:"required" db:"passport_number"`
	Surname        string `json:"surname" binding:"required" db:"surname"`
	Name           string `json:"name" binding:"required" db:"name"`
	Patronymic     string `json:"patronymic" binding:"required" db:"patronymic"`
	Address        string `json:"address" binding:"required" db:"address"`
}

type Filter struct {
	Id             int    `json:"id"`
	PassportSerie  int    `json:"passport_serie"`
	PassportNumber int    `json:"passport_number"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
	Pelingation    int    `json:"pelingation"`
}
