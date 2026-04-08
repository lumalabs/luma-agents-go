// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lumaagents

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/lumalabs/luma-agents-go/internal/apijson"
	"github.com/lumalabs/luma-agents-go/internal/param"
	"github.com/lumalabs/luma-agents-go/internal/requestconfig"
	"github.com/lumalabs/luma-agents-go/option"
)

// GenerationService contains methods and other services that help with interacting
// with the luma API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGenerationService] method instead.
type GenerationService struct {
	Options []option.RequestOption
}

// NewGenerationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGenerationService(opts ...option.RequestOption) (r *GenerationService) {
	r = &GenerationService{}
	r.Options = opts
	return
}

// Submit an image generation or edit job. Returns immediately with an opaque job
// ID to poll via GET /generations/{id}.
func (r *GenerationService) New(ctx context.Context, body GenerationNewParams, opts ...option.RequestOption) (res *Generation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "generations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Poll for generation status and output. On completion, the response includes
// presigned URLs to download the generated images.
func (r *GenerationService) Get(ctx context.Context, generationID string, opts ...option.RequestOption) (res *Generation, err error) {
	opts = slices.Concat(r.Options, opts)
	if generationID == "" {
		err = errors.New("missing required generation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("generations/%s", generationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Generation status and output
type Generation struct {
	// Generation identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Creation timestamp
	CreatedAt string `json:"created_at" api:"required"`
	// Model used
	Model string `json:"model" api:"required"`
	// Current state of the generation
	State GenerationState `json:"state" api:"required"`
	// The kind of generation to perform
	Type GenerationType `json:"type" api:"required"`
	// Error description (populated on failure)
	FailureReason string `json:"failure_reason" api:"nullable"`
	// Generated outputs (populated on completion)
	Output []GenerationOutput `json:"output"`
	JSON   generationJSON     `json:"-"`
}

// generationJSON contains the JSON metadata for the struct [Generation]
type generationJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Model         apijson.Field
	State         apijson.Field
	Type          apijson.Field
	FailureReason apijson.Field
	Output        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Generation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r generationJSON) RawJSON() string {
	return r.raw
}

// Current state of the generation
type GenerationState string

const (
	GenerationStateQueued     GenerationState = "queued"
	GenerationStateProcessing GenerationState = "processing"
	GenerationStateCompleted  GenerationState = "completed"
	GenerationStateFailed     GenerationState = "failed"
)

func (r GenerationState) IsKnown() bool {
	switch r {
	case GenerationStateQueued, GenerationStateProcessing, GenerationStateCompleted, GenerationStateFailed:
		return true
	}
	return false
}

// The kind of generation to perform
type GenerationType string

const (
	GenerationTypeImage     GenerationType = "image"
	GenerationTypeImageEdit GenerationType = "image_edit"
)

func (r GenerationType) IsKnown() bool {
	switch r {
	case GenerationTypeImage, GenerationTypeImageEdit:
		return true
	}
	return false
}

// A single generated output
type GenerationOutput struct {
	// Media type (e.g. image)
	Type string `json:"type" api:"required"`
	// Presigned URL (1hr expiry)
	URL  string               `json:"url" api:"required" format:"uri"`
	JSON generationOutputJSON `json:"-"`
}

// generationOutputJSON contains the JSON metadata for the struct
// [GenerationOutput]
type generationOutputJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GenerationOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r generationOutputJSON) RawJSON() string {
	return r.raw
}

type GenerationNewParams struct {
	// Text prompt
	Prompt param.Field[string] `json:"prompt" api:"required"`
	// Output aspect ratio
	AspectRatio param.Field[GenerationNewParamsAspectRatio] `json:"aspect_ratio"`
	// Up to 8 reference images for style/content guidance
	ImageRef param.Field[[]GenerationNewParamsImageRef] `json:"image_ref"`
	// Model to use
	Model param.Field[string] `json:"model"`
	// Output image format
	OutputFormat param.Field[GenerationNewParamsOutputFormat] `json:"output_format"`
	// Reference image for guided generation. Provide either url or inline base64 data
	// (not both).
	Source param.Field[GenerationNewParamsSource] `json:"source"`
	// Style preset (auto, manga)
	Style param.Field[GenerationNewParamsStyle] `json:"style"`
	// The kind of generation to perform
	Type param.Field[GenerationNewParamsType] `json:"type"`
	// Enable web search grounding — the agent can search the web and download
	// reference images before generating.
	WebSearch param.Field[bool] `json:"web_search"`
}

func (r GenerationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Output aspect ratio
type GenerationNewParamsAspectRatio string

const (
	GenerationNewParamsAspectRatio3_1  GenerationNewParamsAspectRatio = "3:1"
	GenerationNewParamsAspectRatio2_1  GenerationNewParamsAspectRatio = "2:1"
	GenerationNewParamsAspectRatio16_9 GenerationNewParamsAspectRatio = "16:9"
	GenerationNewParamsAspectRatio3_2  GenerationNewParamsAspectRatio = "3:2"
	GenerationNewParamsAspectRatio1_1  GenerationNewParamsAspectRatio = "1:1"
	GenerationNewParamsAspectRatio2_3  GenerationNewParamsAspectRatio = "2:3"
	GenerationNewParamsAspectRatio9_16 GenerationNewParamsAspectRatio = "9:16"
	GenerationNewParamsAspectRatio1_2  GenerationNewParamsAspectRatio = "1:2"
	GenerationNewParamsAspectRatio1_3  GenerationNewParamsAspectRatio = "1:3"
)

func (r GenerationNewParamsAspectRatio) IsKnown() bool {
	switch r {
	case GenerationNewParamsAspectRatio3_1, GenerationNewParamsAspectRatio2_1, GenerationNewParamsAspectRatio16_9, GenerationNewParamsAspectRatio3_2, GenerationNewParamsAspectRatio1_1, GenerationNewParamsAspectRatio2_3, GenerationNewParamsAspectRatio9_16, GenerationNewParamsAspectRatio1_2, GenerationNewParamsAspectRatio1_3:
		return true
	}
	return false
}

// Reference image for guided generation. Provide either url or inline base64 data
// (not both).
type GenerationNewParamsImageRef struct {
	// Base64-encoded image data
	Data param.Field[string] `json:"data"`
	// MIME type (e.g. image/jpeg). Required with data.
	MediaType param.Field[string] `json:"media_type"`
	// Publicly accessible image URL
	URL param.Field[string] `json:"url"`
}

func (r GenerationNewParamsImageRef) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Output image format
type GenerationNewParamsOutputFormat string

const (
	GenerationNewParamsOutputFormatPng  GenerationNewParamsOutputFormat = "png"
	GenerationNewParamsOutputFormatJpeg GenerationNewParamsOutputFormat = "jpeg"
)

func (r GenerationNewParamsOutputFormat) IsKnown() bool {
	switch r {
	case GenerationNewParamsOutputFormatPng, GenerationNewParamsOutputFormatJpeg:
		return true
	}
	return false
}

// Reference image for guided generation. Provide either url or inline base64 data
// (not both).
type GenerationNewParamsSource struct {
	// Base64-encoded image data
	Data param.Field[string] `json:"data"`
	// MIME type (e.g. image/jpeg). Required with data.
	MediaType param.Field[string] `json:"media_type"`
	// Publicly accessible image URL
	URL param.Field[string] `json:"url"`
}

func (r GenerationNewParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Style preset (auto, manga)
type GenerationNewParamsStyle string

const (
	GenerationNewParamsStyleAuto  GenerationNewParamsStyle = "auto"
	GenerationNewParamsStyleManga GenerationNewParamsStyle = "manga"
)

func (r GenerationNewParamsStyle) IsKnown() bool {
	switch r {
	case GenerationNewParamsStyleAuto, GenerationNewParamsStyleManga:
		return true
	}
	return false
}

// The kind of generation to perform
type GenerationNewParamsType string

const (
	GenerationNewParamsTypeImage     GenerationNewParamsType = "image"
	GenerationNewParamsTypeImageEdit GenerationNewParamsType = "image_edit"
)

func (r GenerationNewParamsType) IsKnown() bool {
	switch r {
	case GenerationNewParamsTypeImage, GenerationNewParamsTypeImageEdit:
		return true
	}
	return false
}
