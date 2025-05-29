package cloudprovider

type SubscriptionCreateInput struct {
	Name                string
	EnrollmentAccountId string
	OfferType           string
}

type SEnrollmentAccount struct {
	// Enrollment Account Id
	Id string `json:"id"`

	// Enrollment Account name
	Name string `json:"name"`

	// Enrollment Account 类型
	Type string `json:"type"`
}
