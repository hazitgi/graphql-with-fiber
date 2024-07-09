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
	SortField  string      `json:"sort_field,omitempty" query:"sort_field"`
	Offset     int         `json:"-"`
	Search     string      `json:"search,omitempty" query:"search"`
	TotalPages int         `json:"total_pages,omitempty"`
	TotalRows  int64       `json:"total_rows,omitempty"`
	Rows       interface{} `json:"rows,omitempty"`
}

func (pagination *Pagination) Setup() *Pagination {
	if pagination.Page < 1 {
		pagination.Page = 1
	}
	if pagination.Limit < 1 {
		pagination.Limit = 10
	}
	if pagination.Sort == "" {
		pagination.Sort = "asc" // Default to ascending order if sort is not provided
	}
	if pagination.Sort != "desc" && pagination.Sort != "asc" {
		pagination.Sort = "asc"
	}
	pagination.Offset = (pagination.Page - 1) * pagination.Limit
	if pagination.SortField == "" {
		pagination.SortField = "id" // Default to id if field is not provided
	}
	return &Pagination{}
}
