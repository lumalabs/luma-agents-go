// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lumaagents

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/luma-agents-go/internal/apijson"
	shimjson "github.com/stainless-sdks/luma-agents-go/internal/encoding/json"
	"github.com/stainless-sdks/luma-agents-go/internal/requestconfig"
	"github.com/stainless-sdks/luma-agents-go/option"
	"github.com/stainless-sdks/luma-agents-go/shared"
)

// Access to Petstore orders
//
// StoreOrderService contains methods and other services that help with interacting
// with the luma-agents API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreOrderService] method instead.
type StoreOrderService struct {
	options []option.RequestOption
}

// NewStoreOrderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStoreOrderService(opts ...option.RequestOption) (r StoreOrderService) {
	r = StoreOrderService{}
	r.options = opts
	return
}

// Place a new order in the store
func (r *StoreOrderService) New(ctx context.Context, body StoreOrderNewParams, opts ...option.RequestOption) (res *shared.Order, err error) {
	opts = slices.Concat(r.options, opts)
	path := "store/order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// For valid response try integer IDs with value <= 5 or > 10. Other values will
// generate exceptions.
func (r *StoreOrderService) Get(ctx context.Context, orderID int64, opts ...option.RequestOption) (res *shared.Order, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("store/order/%v", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// For valid response try integer IDs with value < 1000. Anything above 1000 or
// nonintegers will generate API errors
func (r *StoreOrderService) Delete(ctx context.Context, orderID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("store/order/%v", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type StoreOrderNewParams struct {
	Order shared.OrderParam
	paramObj
}

func (r StoreOrderNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Order)
}
func (r *StoreOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
