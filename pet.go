// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lumaagents

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/luma-agents-go/internal/apiform"
	"github.com/stainless-sdks/luma-agents-go/internal/apijson"
	"github.com/stainless-sdks/luma-agents-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/luma-agents-go/internal/encoding/json"
	"github.com/stainless-sdks/luma-agents-go/internal/requestconfig"
	"github.com/stainless-sdks/luma-agents-go/option"
	"github.com/stainless-sdks/luma-agents-go/packages/param"
	"github.com/stainless-sdks/luma-agents-go/packages/respjson"
)

// Everything about your Pets
//
// PetService contains methods and other services that help with interacting with
// the luma-agents API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPetService] method instead.
type PetService struct {
	options []option.RequestOption
}

// NewPetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPetService(opts ...option.RequestOption) (r PetService) {
	r = PetService{}
	r.options = opts
	return
}

// Add a new pet to the store
func (r *PetService) New(ctx context.Context, body PetNewParams, opts ...option.RequestOption) (res *Pet, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single pet
func (r *PetService) Get(ctx context.Context, petID int64, opts ...option.RequestOption) (res *Pet, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing pet by Id
func (r *PetService) Update(ctx context.Context, body PetUpdateParams, opts ...option.RequestOption) (res *Pet, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// delete a pet
func (r *PetService) Delete(ctx context.Context, petID int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Multiple status values can be provided with comma separated strings
func (r *PetService) FindByStatus(ctx context.Context, query PetFindByStatusParams, opts ...option.RequestOption) (res *[]Pet, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet/findByStatus"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3
// for testing.
func (r *PetService) FindByTags(ctx context.Context, query PetFindByTagsParams, opts ...option.RequestOption) (res *[]Pet, err error) {
	opts = slices.Concat(r.options, opts)
	path := "pet/findByTags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates a pet in the store with form data
func (r *PetService) UpdateByID(ctx context.Context, petID int64, body PetUpdateByIDParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("pet/%v", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// uploads an image
func (r *PetService) UploadImage(ctx context.Context, petID int64, image io.Reader, params PetUploadImageParams, opts ...option.RequestOption) (res *PetUploadImageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", image)}, opts...)
	path := fmt.Sprintf("pet/%v/uploadImage", petID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Category) RawJSON() string { return r.JSON.raw }
func (r *Category) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Category to a CategoryParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CategoryParam.Overrides()
func (r Category) ToParam() CategoryParam {
	return param.Override[CategoryParam](json.RawMessage(r.RawJSON()))
}

type CategoryParam struct {
	ID   param.Opt[int64]  `json:"id,omitzero"`
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r CategoryParam) MarshalJSON() (data []byte, err error) {
	type shadow CategoryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CategoryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Pet struct {
	Name      string   `json:"name" api:"required"`
	PhotoURLs []string `json:"photoUrls" api:"required"`
	ID        int64    `json:"id"`
	Category  Category `json:"category"`
	// pet status in the store
	//
	// Any of "available", "pending", "sold".
	Status PetStatus `json:"status"`
	Tags   []Tag     `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		PhotoURLs   respjson.Field
		ID          respjson.Field
		Category    respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pet) RawJSON() string { return r.JSON.raw }
func (r *Pet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Pet to a PetParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PetParam.Overrides()
func (r Pet) ToParam() PetParam {
	return param.Override[PetParam](json.RawMessage(r.RawJSON()))
}

// pet status in the store
type PetStatus string

const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)

// The properties Name, PhotoURLs are required.
type PetParam struct {
	Name      string           `json:"name" api:"required"`
	PhotoURLs []string         `json:"photoUrls,omitzero" api:"required"`
	ID        param.Opt[int64] `json:"id,omitzero"`
	Category  CategoryParam    `json:"category,omitzero"`
	// pet status in the store
	//
	// Any of "available", "pending", "sold".
	Status PetStatus  `json:"status,omitzero"`
	Tags   []TagParam `json:"tags,omitzero"`
	paramObj
}

func (r PetParam) MarshalJSON() (data []byte, err error) {
	type shadow PetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PetParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tag) RawJSON() string { return r.JSON.raw }
func (r *Tag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Tag to a TagParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TagParam.Overrides()
func (r Tag) ToParam() TagParam {
	return param.Override[TagParam](json.RawMessage(r.RawJSON()))
}

type TagParam struct {
	ID   param.Opt[int64]  `json:"id,omitzero"`
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r TagParam) MarshalJSON() (data []byte, err error) {
	type shadow TagParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetUploadImageResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PetUploadImageResponse) RawJSON() string { return r.JSON.raw }
func (r *PetUploadImageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetNewParams struct {
	Pet PetParam
	paramObj
}

func (r PetNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Pet)
}
func (r *PetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetUpdateParams struct {
	Pet PetParam
	paramObj
}

func (r PetUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Pet)
}
func (r *PetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PetFindByStatusParams struct {
	// Status values that need to be considered for filter
	//
	// Any of "available", "pending", "sold".
	Status PetFindByStatusParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PetFindByStatusParams]'s query parameters as `url.Values`.
func (r PetFindByStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Status values that need to be considered for filter
type PetFindByStatusParamsStatus string

const (
	PetFindByStatusParamsStatusAvailable PetFindByStatusParamsStatus = "available"
	PetFindByStatusParamsStatusPending   PetFindByStatusParamsStatus = "pending"
	PetFindByStatusParamsStatusSold      PetFindByStatusParamsStatus = "sold"
)

type PetFindByTagsParams struct {
	// Tags to filter by
	Tags []string `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PetFindByTagsParams]'s query parameters as `url.Values`.
func (r PetFindByTagsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetUpdateByIDParams struct {
	// Name of pet that needs to be updated
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Status of pet that needs to be updated
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PetUpdateByIDParams]'s query parameters as `url.Values`.
func (r PetUpdateByIDParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PetUploadImageParams struct {
	// Additional Metadata
	AdditionalMetadata param.Opt[string] `query:"additionalMetadata,omitzero" json:"-"`
	paramObj
}

func (r PetUploadImageParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [PetUploadImageParams]'s query parameters as `url.Values`.
func (r PetUploadImageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
