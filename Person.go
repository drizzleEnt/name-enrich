package nameenrich

type Person struct {
	Id          int
	Surname     string `json:"surname" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string
}

type Country struct {
	Country []struct {
		CountryId   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}
