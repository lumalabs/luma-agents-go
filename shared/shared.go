// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/luma-agents-go/internal/apijson"
	"github.com/stainless-sdks/luma-agents-go/packages/param"
	"github.com/stainless-sdks/luma-agents-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Order struct {
	ID       int64     `json:"id"`
	Complete bool      `json:"complete"`
	PetID    int64     `json:"petId"`
	Quantity int64     `json:"quantity"`
	ShipDate time.Time `json:"shipDate" format:"date-time"`
	// Order Status
	//
	// Any of "placed", "approved", "delivered".
	Status OrderStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Complete    respjson.Field
		PetID       respjson.Field
		Quantity    respjson.Field
		ShipDate    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Order) RawJSON() string { return r.JSON.raw }
func (r *Order) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Order to a OrderParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OrderParam.Overrides()
func (r Order) ToParam() OrderParam {
	return param.Override[OrderParam](json.RawMessage(r.RawJSON()))
}

// Order Status
type OrderStatus string

const (
	OrderStatusPlaced    OrderStatus = "placed"
	OrderStatusApproved  OrderStatus = "approved"
	OrderStatusDelivered OrderStatus = "delivered"
)

type OrderParam struct {
	ID       param.Opt[int64]     `json:"id,omitzero"`
	Complete param.Opt[bool]      `json:"complete,omitzero"`
	PetID    param.Opt[int64]     `json:"petId,omitzero"`
	Quantity param.Opt[int64]     `json:"quantity,omitzero"`
	ShipDate param.Opt[time.Time] `json:"shipDate,omitzero" format:"date-time"`
	// Order Status
	//
	// Any of "placed", "approved", "delivered".
	Status OrderStatus `json:"status,omitzero"`
	paramObj
}

func (r OrderParam) MarshalJSON() (data []byte, err error) {
	type shadow OrderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
