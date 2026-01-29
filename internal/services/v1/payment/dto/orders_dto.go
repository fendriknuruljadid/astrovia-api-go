package dto

type DuitkuItemDetail struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}

type DuitkuCustomerDetail struct {
	FirstName       string        `json:"firstName"`
	LastName        string        `json:"lastName"`
	Email           string        `json:"email"`
	PhoneNumber     string        `json:"phoneNumber"`
	BillingAddress  DuitkuAddress `json:"billingAddress"`
	ShippingAddress DuitkuAddress `json:"shippingAddress"`
}

type DuitkuAddress struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  string `json:"postalCode"`
	Phone       string `json:"phone"`
	CountryCode string `json:"countryCode"`
}

type DuitkuInquiryRequest struct {
	MerchantCode     string               `json:"merchantCode"`
	PaymentAmount    int64                `json:"paymentAmount"`
	PaymentMethod    string               `json:"paymentMethod"`
	MerchantOrderID  string               `json:"merchantOrderId"`
	ProductDetails   string               `json:"productDetails"`
	AdditionalParam  string               `json:"additionalParam,omitempty"`
	MerchantUserInfo string               `json:"merchantUserInfo,omitempty"`
	CustomerVaName   string               `json:"customerVaName"`
	Email            string               `json:"email"`
	PhoneNumber      string               `json:"phoneNumber,omitempty"`
	ItemDetails      []DuitkuItemDetail   `json:"itemDetails"`
	CustomerDetail   DuitkuCustomerDetail `json:"customerDetail"`
	CallbackURL      string               `json:"callbackUrl"`
	ReturnURL        string               `json:"returnUrl"`
	Signature        string               `json:"signature"`
	ExpiryPeriod     int                  `json:"expiryPeriod"`
}

type DuitkuInquiryResponse struct {
	MerchantCode  string `json:"merchantCode"`
	Reference     string `json:"reference"`
	PaymentURL    string `json:"paymentUrl"`
	VANumber      string `json:"vaNumber"`
	Amount        string `json:"amount"`
	QRString      string `json:"qrString"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

type CreateOrderRequest struct {
	PricingId     string `json:"pricing_id" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
}
