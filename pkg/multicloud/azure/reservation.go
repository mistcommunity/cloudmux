
package azure

import "time"

type SReservationOrder struct {
	Id         string
	Name       string
	Type       string
	Properties struct {
		OriginalQuantity int
		RequestDateTime  time.Time
		BillingPlan      string
		Reservations     []struct {
			Id string
		}
		Term              string
		DisplayName       string
		CreatedDateTime   time.Time
		BenefitStartTime  time.Time
		ProvisioningState string
		BillingProfileId  string
		BillingAccountId  string
		ExpiryDate        time.Time
		ExpiryDateTime    time.Time
	}
}

func (client *SAzureClient) ListReservationOrders() ([]SReservationOrder, error) {
	result := []SReservationOrder{}
	resp, err := client.list_v2("/providers/Microsoft.Capacity/reservationOrders", "2022-11-01", nil)
	if err != nil {
		return nil, err
	}
	err = resp.Unmarshal(&result, "value")
	if err != nil {
		return nil, err
	}
	return result, nil
}

type SReservation struct {
	Id       string
	Name     string
	Type     string
	Location string
	Etag     int
	Sku      struct {
		Name string
	}
	Properties struct {
		ReservedResourceType         string
		UserFriendlyRenewState       string
		SkuDescription               string
		Renew                        bool
		Archived                     bool
		Quantity                     int
		AppliedScopeType             string
		DisplayName                  string
		ProvisioningState            string
		Term                         string
		DisplayProvisioningState     string
		UserFriendlyAppliedScopeType string
		ExpiryDateTime               time.Time
		PurchaseDateTime             time.Time
		BenefitStartTime             time.Time
		LastUpdatedDateTime          time.Time
		ExpiryDate                   string
		PurchaseDate                 string
		EffectiveDateTime            time.Time
		InstanceFlexibility          string
		Utilization                  struct {
			Trend      string
			Aggregates []struct {
				Grain     float64
				GrainUnit string
				Value     float64
				ValueUnit string
			}
		}
		BillingPlan    string
		BillingScopeId string
	}
}

func (client *SAzureClient) ListReservations() ([]SReservation, error) {
	result := []SReservation{}
	resp, err := client.list_v2("/providers/Microsoft.Capacity/reservations", "2022-11-01", nil)
	if err != nil {
		return nil, err
	}
	err = resp.Unmarshal(&result, "value")
	if err != nil {
		return nil, err
	}
	return result, nil
}
