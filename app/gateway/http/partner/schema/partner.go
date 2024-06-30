package schema

// CreatePartnerRequest represents the request to create a new partner.
type CreatePartnerRequest struct {
	PartnerName  string `json:"partner_name"`
	Phone        string `json:"phone"`
	CelPhone     string `json:"cel_phone"`
	Email        string `json:"email"`
	ZipCode      string `json:"zip_code"`
	Address      string `json:"address"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Reference    string `json:"reference"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	UF           string `json:"uf"`
}
