package entity

type CredentialSet struct {
	IamTicket IamTicket
	Credentials[] Credentials
	ProfileId string
	ProfileType string
	ProviderId string
	CredentialSetId string
}

type Credentials struct {
	AuthenticationFieldId string
	AuthenticationFieldType string
	AuthenticationFieldText string
	AuthenticationFieldValue string
	AuthenticationSourceId string
	Encrypted bool
}