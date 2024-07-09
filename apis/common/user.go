package common

type User struct{}

type CreateUserInput struct {
	FullName    string `json:"full_name,omitempty" validate:"required"`
	CompanyName string `json:"company_name,omitempty" validate:"required"`
	CountryID   string `json:"country_id,omitempty" validate:"required"`
	StateID     string `json:"state_id,omitempty" validate:"required"`
	Email       string `json:"email,omitempty" validate:"required"`
	Location    string `json:"location,omitempty" validate:"required"`
	Address     string `json:"address,omitempty" validate:"required"`
	Password    string `json:"password,omitempty" validate:"required,gte=8"`
}

func NewCreateUserInput() *CreateUserInput {
	return &CreateUserInput{}
}

var UserValidationMessages = map[string]string{
	"FullName.required":    "Full Name is required",
	"CompanyName.required": "Company Name is required",
	"CountryID.required":   "Country ID is required",
	"StateID.required":     "State ID is required",
	"Email.required":       "Email is required",
	"Location.required":    "Location is required",
	"Address.required":     "Address is required",
	"Password.required":    "Password is required",
	"Password.gte":         "Password should be at least 8 characters long",
}

type Pagination struct {
	Limit      int         `json:"limit,omitempty" query:"limit"`
	Page       int         `json:"page,omitempty" query:"page"`
	Sort       string      `json:"sort,omitempty" query:"sort"`
	SortField  string      `json:"field,omitempty" query:"sort_field"`
	Offset     int         `json:"-"`
	Search     string      `json:"search,omitempty" query:"search"`
	TotalPages int         `json:"total_pages,omitempty"`
	TotalRows  int64       `json:"total_rows,omitempty"`
	Rows       interface{} `json:"rows,omitempty"`
}
