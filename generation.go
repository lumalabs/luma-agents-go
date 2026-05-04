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
	Model Model `json:"model" api:"required"`
	// Current state of the generation
	State GenerationState `json:"state" api:"required"`
	// The kind of generation to perform
	Type GenerationType `json:"type" api:"required"`
	// Machine-readable failure code for programmatic handling
	FailureCode GenerationFailureCode `json:"failure_code" api:"nullable"`
	// Human-readable failure description
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
	FailureCode   apijson.Field
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

// Machine-readable failure code for programmatic handling
type GenerationFailureCode string

const (
	GenerationFailureCodeContentModerated  GenerationFailureCode = "content_moderated"
	GenerationFailureCodeGenerationFailed  GenerationFailureCode = "generation_failed"
	GenerationFailureCodeBudgetExhausted   GenerationFailureCode = "budget_exhausted"
	GenerationFailureCodeOutputNotFound    GenerationFailureCode = "output_not_found"
	GenerationFailureCodeImageTooLarge     GenerationFailureCode = "image_too_large"
	GenerationFailureCodeUnsupportedFormat GenerationFailureCode = "unsupported_format"
	GenerationFailureCodeCorruptInput      GenerationFailureCode = "corrupt_input"
	GenerationFailureCodeInvalidRequest    GenerationFailureCode = "invalid_request"
	GenerationFailureCodeRateLimited       GenerationFailureCode = "rate_limited"
)

func (r GenerationFailureCode) IsKnown() bool {
	switch r {
	case GenerationFailureCodeContentModerated, GenerationFailureCodeGenerationFailed, GenerationFailureCodeBudgetExhausted, GenerationFailureCodeOutputNotFound, GenerationFailureCodeImageTooLarge, GenerationFailureCodeUnsupportedFormat, GenerationFailureCodeCorruptInput, GenerationFailureCodeInvalidRequest, GenerationFailureCodeRateLimited:
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

// Model identifier. `uni-1` is the default tier; `uni-1-max` produces
// higher-quality output than `uni-1` at a higher per-image price. Both models are
// available to all accounts — see Pricing for per-image rates.
type Model string

const (
	ModelUni1    Model = "uni-1"
	ModelUni1Max Model = "uni-1-max"
)

func (r Model) IsKnown() bool {
	switch r {
	case ModelUni1, ModelUni1Max:
		return true
	}
	return false
}

type GenerationNewParams struct {
	// Text prompt
	Prompt param.Field[string] `json:"prompt" api:"required"`
	// Output aspect ratio
	AspectRatio param.Field[GenerationNewParamsAspectRatio] `json:"aspect_ratio"`
	// Reference images for style/content guidance. Up to 9 for type 'image', up to 8
	// for type 'image_edit'.
	ImageRef param.Field[[]GenerationNewParamsImageRef] `json:"image_ref"`
	// Model identifier. `uni-1` is the default tier; `uni-1-max` produces
	// higher-quality output than `uni-1` at a higher per-image price. Both models are
	// available to all accounts — see Pricing for per-image rates.
	Model param.Field[Model] `json:"model"`
	// Output image format
	OutputFormat param.Field[GenerationNewParamsOutputFormat] `json:"output_format"`
	// Reference image for guided generation. Provide either url or inline base64 data
	// (not both).
	Source param.Field[GenerationNewParamsSource] `json:"source"`
	// Style preset (auto, manga)
	Style param.Field[GenerationNewParamsStyle] `json:"style"`
	// The kind of generation to perform
	Type param.Field[GenerationNewParamsType] `json:"type"`
	// Your end-user's stable opaque identifier (no PII). Forwarded to upstream model
	// providers as their per-user tagging field so trust & safety violations can be
	// attributed to a specific end-user rather than the whole API account. Also used
	// for per-end-user usage breakdowns in /v1/usage. Strongly recommended for partner
	// integrations.
	UserID param.Field[string] `json:"user_id"`
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
