package constant

const (
	SuccessResponseCode = "00"
	ErrorResponseCode   = "XX"
	Success             = "Success"
	Failed              = "Failed"
)

const (
	CreateDistrictAdmin = "create-district-admin"
	CreateTpiAdmin      = "create-tpi-admin"
	CreateTpiOfficer    = "create-tpi-officer"
	CreateTpiCashier    = "create-tpi-cashier"
	GetUser             = "get-user"
	GetByIDUser         = "getbyid-user"
	UpdateUser          = "update-user"

	CreateTpi  = "create-tpi"
	GetByIDTpi = "getbyid-tpi"
	UpdateTpi  = "update-tpi"
	DeleteTpi  = "delete-tpi"

	CreateCaught  = "create-caught"
	InquiryCaught = "inquiry-caught"
	GetByIDCaught = "getbyid-caught"
	UpdateCaught  = "update-caught"
	DeleteCaught  = "delete-caught"

	CreateAuction  = "create-auction"
	InquiryAuction = "inquiry-auction"
	GetByIDAuction = "getbyid-auction"
	UpdateAuction  = "update-auction"
	DeleteAuction  = "delete-auction"

	CreateTransaction  = "create-transaction"
	GetByIDTransaction = "getbyid-transaction"
	UpdateTransaction  = "update-transaction"
	DeleteTransaction  = "delete-transaction"

	CreateFisher  = "create-fisher"
	UpdateFisher  = "update-fisher"
	GetByIDFisher = "getbyid-fisher"
	DeleteFisher  = "delete-fisher"

	CreateBuyer  = "create-buyer"
	UpdateBuyer  = "update-buyer"
	GetByIDBuyer = "getbyid-buyer"
	DeleteBuyer  = "delete-buyer"

	CreateFishingGear  = "create-fishing-gear"
	UpdateFishingGear  = "update-fishing-gear"
	GetByIDFishingGear = "getbyid-fishing-gear"
	DeleteFishingGear  = "delete-fishing-gear"

	CreateFishingArea  = "create-fishing-area"
	UpdateFishingArea  = "update-fishing-area"
	GetByIDFishingArea = "getbyid-fishing-area"
	DeleteFishingArea  = "delete-fishing-area"

	CreateFishType  = "create-fish-type"
	UpdateFishType  = "update-fish-type"
	GetByIDFishType = "getbyid-fish-type"
	DeleteFishType  = "delete-fish-type"

	Pass = "pass"

	ResetPassword = "reset-password"

	PermanentStatus = "Tetap"
	TemporaryStatus = "Pendatang"

	User        = "user"
	Caught      = "caught"
	Auction     = "auction"
	Transaction = "transactions"
	Fisher      = "fisher"
	Buyer       = "buyer"
	Tpi         = "tpi"

	Daily   = "daily"
	Monthly = "monthly"
	Yearly  = "yearly"
	Period  = "period"

	TransactionPdf = "transaction"
	ProductionPdf  = "production"
)
