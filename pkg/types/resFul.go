package types

// RestFul json data output pattern
type RestFul struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

// Meta header with information about error, pagination and any other information the frontend needs
type Meta struct {
	Success bool     `json:"success"`
	Error   []string `json:"error"`
}

// AddError prepare the error response
func (e *RestFul) AddError(err error) {
	e.Meta.Success = false

	if e.Meta.Error == nil {
		e.Meta.Error = make([]string, 0)
	}

	e.Meta.Error = append(e.Meta.Error, err.Error())

	// returning a blank array makes it easy for iOS
	e.Data = []int{}
}

// Success prepare the exit of success
func (e *RestFul) Success(data any) {
	e.Meta.Success = true

	// returning a blank array makes it easy for iOS
	e.Meta.Error = make([]string, 0)

	e.Data = data
}
