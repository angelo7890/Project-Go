package dto

type CreateUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r *CreateUserDTO) Validate() error {
	if r.Name == "" {
		return ErrParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return ErrParamIsRequired("email", "string")
	}
	return nil
}

type ResponseUserDTO struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
