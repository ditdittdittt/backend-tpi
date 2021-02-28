package http

type LoginRequest struct {
	Username	string		`json:"username" binding:"required"`
	Password	string		`json:"password" binding:"required"`
}

type GetUserRequest struct {

}

type CreateTpiAdminRequest struct {
	RoleID		int			`json:"role_id" binding:"required,eq=3"`
	TpiID		int			`json:"tpi_id" binding:"required"`
	Nik			string		`json:"nik" binding:"required"`
	Name		string		`json:"name" binding:"required"`
	Address		string		`json:"address" binding:"required"`
	Username	string		`json:"username" binding:"required"`
}

type CreateTpiOfficerRequest struct {
	RoleID		int			`json:"role_id" binding:"required,eq=4"`
	TpiID		int			`json:"tpi_id" binding:"required"`
	Nik			string		`json:"nik" binding:"required"`
	Name		string		`json:"name" binding:"required"`
	Address		string		`json:"address" binding:"required"`
	Username	string		`json:"username" binding:"required"`
}

type CreateTpiCashierRequest struct {
	RoleID		int			`json:"role_id" binding:"required,eq=5"`
	TpiID		int			`json:"tpi_id" binding:"required"`
	Nik			string		`json:"nik" binding:"required"`
	Name		string		`json:"name" binding:"required"`
	Address		string		`json:"address" binding:"required"`
	Username	string		`json:"username" binding:"required"`
}

type CreateDistrictAdminRequest struct {
	RoleID		int			`json:"role_id" binding:"required,eq=2"`
	DistrictID	int			`json:"district_id" binding:"required"`
	Nik			string		`json:"nik" binding:"required"`
	Name		string		`json:"name" binding:"required"`
	Address		string		`json:"address" binding:"required"`
	Username	string		`json:"username" binding:"required"`
}