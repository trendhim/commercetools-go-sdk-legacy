// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"
)

// ShippingRateTierType is an enum type
type ShippingRateTierType string

// Enum values for ShippingRateTierType
const (
	ShippingRateTierTypeCartValue          ShippingRateTierType = "CartValue"
	ShippingRateTierTypeCartClassification ShippingRateTierType = "CartClassification"
	ShippingRateTierTypeCartScore          ShippingRateTierType = "CartScore"
)

// ShippingMethodUpdateAction uses action as discriminator attribute
type ShippingMethodUpdateAction interface{}

func mapDiscriminatorShippingMethodUpdateAction(input interface{}) (ShippingMethodUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "addShippingRate":
		new := ShippingMethodAddShippingRateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addZone":
		new := ShippingMethodAddZoneAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeIsDefault":
		new := ShippingMethodChangeIsDefaultAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := ShippingMethodChangeNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTaxCategory":
		new := ShippingMethodChangeTaxCategoryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeShippingRate":
		new := ShippingMethodRemoveShippingRateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeZone":
		new := ShippingMethodRemoveZoneAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := ShippingMethodSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := ShippingMethodSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := ShippingMethodSetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := ShippingMethodSetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocalizedName":
		new := ShippingMethodSetLocalizedNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocalizedDescription":
		new := ShippingMethodSetLocalizedDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setPredicate":
		new := ShippingMethodSetPredicateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// ShippingRatePriceTier uses type as discriminator attribute
type ShippingRatePriceTier interface{}

func mapDiscriminatorShippingRatePriceTier(input interface{}) (ShippingRatePriceTier, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "CartClassification":
		new := CartClassificationTier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CartScore":
		new := CartScoreTier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CartValue":
		new := CartValueTier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// CartClassificationTier implements the interface ShippingRatePriceTier
type CartClassificationTier struct {
	Value      string `json:"value"`
	Price      *Money `json:"price"`
	IsMatching bool   `json:"isMatching"`
}

// MarshalJSON override to set the discriminator value
func (obj CartClassificationTier) MarshalJSON() ([]byte, error) {
	type Alias CartClassificationTier
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartClassification", Alias: (*Alias)(&obj)})
}

// CartScoreTier implements the interface ShippingRatePriceTier
type CartScoreTier struct {
	Score         float64        `json:"score"`
	PriceFunction *PriceFunction `json:"priceFunction,omitempty"`
	Price         *Money         `json:"price,omitempty"`
	IsMatching    bool           `json:"isMatching"`
}

// MarshalJSON override to set the discriminator value
func (obj CartScoreTier) MarshalJSON() ([]byte, error) {
	type Alias CartScoreTier
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartScore", Alias: (*Alias)(&obj)})
}

// CartValueTier implements the interface ShippingRatePriceTier
type CartValueTier struct {
	Price             *Money `json:"price"`
	MinimumCentAmount int    `json:"minimumCentAmount"`
	IsMatching        bool   `json:"isMatching"`
}

// MarshalJSON override to set the discriminator value
func (obj CartValueTier) MarshalJSON() ([]byte, error) {
	type Alias CartValueTier
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartValue", Alias: (*Alias)(&obj)})
}

// PriceFunction is a standalone struct
type PriceFunction struct {
	Function     string       `json:"function"`
	CurrencyCode CurrencyCode `json:"currencyCode"`
}

// ShippingMethod is of type BaseResource
type ShippingMethod struct {
	ZoneRates            []ZoneRate            `json:"zoneRates"`
	Version              int                   `json:"version"`
	TaxCategory          *TaxCategoryReference `json:"taxCategory"`
	Predicate            string                `json:"predicate,omitempty"`
	Name                 string                `json:"name"`
	LocalizedName        *LocalizedString      `json:"localizedName,omitempty"`
	LocalizedDescription *LocalizedString      `json:"localizedDescription,omitempty"`
	LastModifiedBy       *LastModifiedBy       `json:"lastModifiedBy,omitempty"`
	LastModifiedAt       time.Time             `json:"lastModifiedAt"`
	Key                  string                `json:"key,omitempty"`
	IsDefault            bool                  `json:"isDefault"`
	ID                   string                `json:"id"`
	Description          string                `json:"description,omitempty"`
	Custom               *CustomFields         `json:"custom,omitempty"`
	CreatedBy            *CreatedBy            `json:"createdBy,omitempty"`
	CreatedAt            time.Time             `json:"createdAt"`
}

// ShippingMethodAddShippingRateAction implements the interface ShippingMethodUpdateAction
type ShippingMethodAddShippingRateAction struct {
	Zone         *ZoneResourceIdentifier `json:"zone"`
	ShippingRate *ShippingRateDraft      `json:"shippingRate"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodAddShippingRateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodAddShippingRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingRate", Alias: (*Alias)(&obj)})
}

// ShippingMethodAddZoneAction implements the interface ShippingMethodUpdateAction
type ShippingMethodAddZoneAction struct {
	Zone *ZoneResourceIdentifier `json:"zone"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodAddZoneAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodAddZoneAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addZone", Alias: (*Alias)(&obj)})
}

// ShippingMethodChangeIsDefaultAction implements the interface ShippingMethodUpdateAction
type ShippingMethodChangeIsDefaultAction struct {
	IsDefault bool `json:"isDefault"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodChangeIsDefaultAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeIsDefaultAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsDefault", Alias: (*Alias)(&obj)})
}

// ShippingMethodChangeNameAction implements the interface ShippingMethodUpdateAction
type ShippingMethodChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ShippingMethodChangeTaxCategoryAction implements the interface ShippingMethodUpdateAction
type ShippingMethodChangeTaxCategoryAction struct {
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodChangeTaxCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeTaxCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxCategory", Alias: (*Alias)(&obj)})
}

// ShippingMethodDraft is a standalone struct
type ShippingMethodDraft struct {
	ZoneRates            []ZoneRateDraft                `json:"zoneRates"`
	TaxCategory          *TaxCategoryResourceIdentifier `json:"taxCategory"`
	Predicate            string                         `json:"predicate,omitempty"`
	Name                 string                         `json:"name"`
	LocalizedName        *LocalizedString               `json:"localizedName,omitempty"`
	LocalizedDescription *LocalizedString               `json:"localizedDescription,omitempty"`
	Key                  string                         `json:"key,omitempty"`
	IsDefault            bool                           `json:"isDefault"`
	Description          string                         `json:"description,omitempty"`
	Custom               *CustomFieldsDraft             `json:"custom,omitempty"`
}

// ShippingMethodPagedQueryResponse is a standalone struct
type ShippingMethodPagedQueryResponse struct {
	Total   int              `json:"total,omitempty"`
	Results []ShippingMethod `json:"results"`
	Offset  int              `json:"offset,omitempty"`
	Limit   int              `json:"limit,omitempty"`
	Count   int              `json:"count"`
}

// ShippingMethodReference implements the interface Reference
type ShippingMethodReference struct {
	ID  string          `json:"id"`
	Obj *ShippingMethod `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodReference) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "shipping-method", Alias: (*Alias)(&obj)})
}

// ShippingMethodRemoveShippingRateAction implements the interface ShippingMethodUpdateAction
type ShippingMethodRemoveShippingRateAction struct {
	Zone         *ZoneResourceIdentifier `json:"zone"`
	ShippingRate *ShippingRateDraft      `json:"shippingRate"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodRemoveShippingRateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodRemoveShippingRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingRate", Alias: (*Alias)(&obj)})
}

// ShippingMethodRemoveZoneAction implements the interface ShippingMethodUpdateAction
type ShippingMethodRemoveZoneAction struct {
	Zone *ZoneResourceIdentifier `json:"zone"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodRemoveZoneAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodRemoveZoneAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeZone", Alias: (*Alias)(&obj)})
}

// ShippingMethodResourceIdentifier implements the interface ResourceIdentifier
type ShippingMethodResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "shipping-method", Alias: (*Alias)(&obj)})
}

// ShippingMethodSetCustomFieldAction implements the interface ShippingMethodUpdateAction
type ShippingMethodSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// ShippingMethodSetCustomTypeAction implements the interface ShippingMethodUpdateAction
type ShippingMethodSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// ShippingMethodSetDescriptionAction implements the interface ShippingMethodUpdateAction
type ShippingMethodSetDescriptionAction struct {
	Description string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// ShippingMethodSetKeyAction implements the interface ShippingMethodUpdateAction
type ShippingMethodSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ShippingMethodSetLocalizedNameAction implements the interface ShippingMethodUpdateAction
type ShippingMethodSetLocalizedNameAction struct {
	LocalizedName *LocalizedString `json:"localizedName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetLocalizedNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetLocalizedNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocalizedName", Alias: (*Alias)(&obj)})
}

// ShippingMethodSetLocalizedDescriptionAction implements the interface ShippingMethodUpdateAction
type ShippingMethodSetLocalizedDescriptionAction struct {
	LocalizedDescription *LocalizedString `json:"localizedDescription,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetLocalizedDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetLocalizedDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocalizedDescription", Alias: (*Alias)(&obj)})
}

// ShippingMethodSetPredicateAction implements the interface ShippingMethodUpdateAction
type ShippingMethodSetPredicateAction struct {
	Predicate string `json:"predicate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPredicate", Alias: (*Alias)(&obj)})
}

// ShippingMethodUpdate is a standalone struct
type ShippingMethodUpdate struct {
	Version int                          `json:"version"`
	Actions []ShippingMethodUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingMethodUpdate) UnmarshalJSON(data []byte) error {
	type Alias ShippingMethodUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorShippingMethodUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// ShippingRate is a standalone struct
type ShippingRate struct {
	Tiers      []ShippingRatePriceTier `json:"tiers"`
	Price      TypedMoney              `json:"price"`
	IsMatching bool                    `json:"isMatching"`
	FreeAbove  TypedMoney              `json:"freeAbove,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingRate) UnmarshalJSON(data []byte) error {
	type Alias ShippingRate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.FreeAbove != nil {
		var err error
		obj.FreeAbove, err = mapDiscriminatorTypedMoney(obj.FreeAbove)
		if err != nil {
			return err
		}
	}
	if obj.Price != nil {
		var err error
		obj.Price, err = mapDiscriminatorTypedMoney(obj.Price)
		if err != nil {
			return err
		}
	}
	for i := range obj.Tiers {
		var err error
		obj.Tiers[i], err = mapDiscriminatorShippingRatePriceTier(obj.Tiers[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// ShippingRateDraft is a standalone struct
type ShippingRateDraft struct {
	Tiers     []ShippingRatePriceTier `json:"tiers,omitempty"`
	Price     *Money                  `json:"price"`
	FreeAbove *Money                  `json:"freeAbove,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingRateDraft) UnmarshalJSON(data []byte) error {
	type Alias ShippingRateDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Tiers {
		var err error
		obj.Tiers[i], err = mapDiscriminatorShippingRatePriceTier(obj.Tiers[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// ZoneRate is a standalone struct
type ZoneRate struct {
	Zone          *ZoneReference `json:"zone"`
	ShippingRates []ShippingRate `json:"shippingRates"`
}

// ZoneRateDraft is a standalone struct
type ZoneRateDraft struct {
	Zone          *ZoneResourceIdentifier `json:"zone"`
	ShippingRates []ShippingRateDraft     `json:"shippingRates"`
}
