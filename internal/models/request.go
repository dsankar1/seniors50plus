package models

const (
	RequestStatusPending  = "pending"
	RequestStatusAccepted = "accepted"
	RequestStatusDenied   = "denied"
)

type OthersRequest struct {
	Id              uint   `json:"id"`
	Email           string `json:"email"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	ProfileImageUrl string `json:"profileImageUrl"`
	RequestedOn     string `json:"requestedOn"`
	Status          string `json:"status"`
}

type YourRequest struct {
	Id               uint   `json:"id"`
	OfferId          uint   `json:"offerId"`
	PropertyImageUrl string `json:"propertyImageUrl"`
	State            string `json:"State"`
	City             string `json:"city"`
	Zip              string `json:"zip"`
	RequestedOn      string `json:"requestedOn"`
	Status           string `json:"status"`
}
