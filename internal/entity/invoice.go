package entity

type AccountType int32

const (
	AccountTypeFactory = iota
	AccountTypeShop
)

type Invoice struct {
	Id          int         `json:"id"`
	LastName    string      `json:"lastName"`
	FirstName   string      `json:"firstName"`
	MiddleName  string      `json:"middleName"`
	OrgName     string      `json:"orgName"`
	AccountType AccountType `json:"accountType"`
}
