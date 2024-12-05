package payload

// UserAddRequest is a struct for request
type UserAddRequest struct {
	UserName    *string `json:"username" validate:"required"`
	Email       *string `json:"email" validate:"required,email"`
	Password    *string `json:"password" validate:"required,e164"`
	PhoneNumber *string `json:"phonenumber" validate:"required"`
	RoleCode    *string `json:"rolecode" validate:"required,oneof=user manager admin superadmin"`
}

func (r *UserAddRequest) StructName() string {
	return "UserAddRequest"
}

func (payl *UserAddRequest) Validate() error {
	if err := Validate(payl); err != nil {
		return err
	}

	return nil
}

// UserLoginRequest is a struct for request
type UserLoginRequest struct {
	Email       *string `json:"email,omitempty" validate:"omitempty,email"`
	Phone       *string `json:"phone,omitempty" validate:"omitempty,e164"`
	Password    *string `json:"password" validate:"required"`
}

func (r *UserLoginRequest) StructName() string {
	return "UserLoginRequest"
}

func (payl *UserLoginRequest) Validate() error {
	if err := Validate(payl); err != nil {
		return err
	}

	return nil
}

// UserRefreshTokenRequest is a struct for request
type UserRefreshTokenRequest struct {
	RefreshToken *string `json:"refresh_token" validate:"required"`
}

func (r *UserRefreshTokenRequest) StructName() string {
	return "UserRefreshTokenRequest"
}

func (payl *UserRefreshTokenRequest) Validate() error {
	if err := Validate(payl); err != nil {
		return err
	}

	return nil
}
