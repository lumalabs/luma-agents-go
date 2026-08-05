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

// Submit an image or video generation job. Returns immediately with an opaque job
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

// Per-signal manual conditioning controls for video edits
type AdvancedControlsParam struct {
	// Depth / scene-geometry conditioning control
	Depth param.Field[DepthControlParam] `json:"depth"`
	// Face-identity conditioning control
	Face param.Field[FaceControlParam] `json:"face"`
	// Surface-normals conditioning control
	Normals param.Field[NormalsControlParam] `json:"normals"`
	// Pose / skeleton conditioning control
	Pose param.Field[PoseControlParam] `json:"pose"`
	// Motion-trajectory conditioning control
	Trajectory param.Field[TrajectoryControlParam] `json:"trajectory"`
}

func (r AdvancedControlsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Depth / scene-geometry conditioning control
type DepthControlParam struct {
	// Depth-map blur amount from 0 to 1. Higher values allow more geometric freedom.
	Blur param.Field[float64] `json:"blur"`
	// Enable or disable depth conditioning. Omit to use the model default.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r DepthControlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Face-identity conditioning control
type FaceControlParam struct {
	// Enable or disable face conditioning. Omit to use the model default.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r FaceControlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	GenerationTypeImage        GenerationType = "image"
	GenerationTypeImageEdit    GenerationType = "image_edit"
	GenerationTypeVideo        GenerationType = "video"
	GenerationTypeVideoEdit    GenerationType = "video_edit"
	GenerationTypeVideoReframe GenerationType = "video_reframe"
	GenerationTypeLayering     GenerationType = "layering"
)

func (r GenerationType) IsKnown() bool {
	switch r {
	case GenerationTypeImage, GenerationTypeImageEdit, GenerationTypeVideo, GenerationTypeVideoEdit, GenerationTypeVideoReframe, GenerationTypeLayering:
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
	// Media type (e.g. image, video)
	Type string `json:"type" api:"required"`
	// Presigned URL (1hr expiry)
	URL string `json:"url" api:"required" format:"uri"`
	// Per-layer semantics for a type=layering output
	Layer GenerationOutputLayer `json:"layer" api:"nullable"`
	JSON  generationOutputJSON  `json:"-"`
}

// generationOutputJSON contains the JSON metadata for the struct
// [GenerationOutput]
type generationOutputJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	Layer       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GenerationOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r generationOutputJSON) RawJSON() string {
	return r.raw
}

// Per-layer semantics for a type=layering output
type GenerationOutputLayer struct {
	// Edge treatment of the layer's transparency — soft (hair/fur/glass), hard (solid
	// edges), or none (the opaque background)
	AlphaHint string `json:"alpha_hint" api:"required"`
	// Complete-element caption for the layer's content
	Description string `json:"description" api:"required"`
	// Layer position, front-to-back; the last layer is the background
	Index int64 `json:"index" api:"required"`
	// Short (1-2 word) layer name
	Label string                    `json:"label" api:"required"`
	JSON  generationOutputLayerJSON `json:"-"`
}

// generationOutputLayerJSON contains the JSON metadata for the struct
// [GenerationOutputLayer]
type generationOutputLayerJSON struct {
	AlphaHint   apijson.Field
	Description apijson.Field
	Index       apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GenerationOutputLayer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r generationOutputLayerJSON) RawJSON() string {
	return r.raw
}

// Media reference for guided generation. Provide exactly one of url, inline base64
// data, generation_id, or file_id. URL/data references accept image media at image
// positions; video_edit and video_reframe sources also accept source.url or
// source.data when source.media_type is a video/\* MIME. generation_id chains
// image_edit off a prior image output, video_edit/video_reframe off a prior video
// output, and video.start_frame/end_frame for extension. file_id references a file
// previously uploaded via POST /files — see the Files API.
type ImageRefParam struct {
	// Base64-encoded image or video data
	Data param.Field[string] `json:"data"`
	// UUID of a file previously uploaded via POST /files. Skips URL fetch / base64
	// decode and reuses the file's pre-moderated backing artifact. The referenced file
	// must be owned by the same client and in state=ready. See the Files API for the
	// upload flow.
	FileID param.Field[string] `json:"file_id" format:"uuid"`
	// UUID of a prior generation owned by the same caller. Used on source for
	// image_edit, video_edit, and video_reframe chaining and on video.start_frame /
	// video.end_frame for video extension.
	GenerationID param.Field[string] `json:"generation_id" format:"uuid"`
	// MIME type (for example, image/jpeg or video/mp4). Required with data. Required
	// with source.url on video_edit/video_reframe so the route can dispatch video
	// ingest before fetching bytes; optional for image URLs.
	MediaType param.Field[string] `json:"media_type"`
	// Publicly accessible image URL, or a video URL when used as source for
	// video_edit/video_reframe with media_type=video/\*.
	URL param.Field[string] `json:"url"`
}

func (r ImageRefParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Model identifier. `uni-1` is the default image tier; `uni-1-max` produces
// higher-quality output than `uni-1` at a higher per-image price. `ray-3.2` is the
// public video model for text-to-video, image-to-video, and video-to-video
// editing.
type Model string

const (
	ModelUni1    Model = "uni-1"
	ModelUni1Max Model = "uni-1-max"
	ModelRay3_2  Model = "ray-3.2"
)

func (r Model) IsKnown() bool {
	switch r {
	case ModelUni1, ModelUni1Max, ModelRay3_2:
		return true
	}
	return false
}

// Surface-normals conditioning control
type NormalsControlParam struct {
	// Surface-normals augmentation from 0 to 1. Higher values allow more
	// reinterpretation of surface geometry.
	Augmentation param.Field[float64] `json:"augmentation"`
	// Enable or disable normals conditioning. Omit to use the model default.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r NormalsControlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Pose / skeleton conditioning control
type PoseControlParam struct {
	// Enable or disable pose conditioning. Omit to use the model default.
	Enabled param.Field[bool] `json:"enabled"`
	// Pose-conditioning strength
	Strength param.Field[PoseControlStrength] `json:"strength"`
}

func (r PoseControlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Pose-conditioning strength
type PoseControlStrength string

const (
	PoseControlStrengthPrecise PoseControlStrength = "precise"
	PoseControlStrengthCoarse  PoseControlStrength = "coarse"
)

func (r PoseControlStrength) IsKnown() bool {
	switch r {
	case PoseControlStrengthPrecise, PoseControlStrengthCoarse:
		return true
	}
	return false
}

// Normalized source rectangle inside the output canvas for video_reframe. Omit to
// let the model choose the default centered-fit crop.
type SourcePositionParam struct {
	// Source rectangle height, as a fraction of canvas height. Up to 2.0 so the source
	// can bleed off-canvas.
	HNorm param.Field[float64] `json:"h_norm" api:"required"`
	// Source rectangle width, as a fraction of canvas width. Up to 2.0 so the source
	// can bleed off-canvas.
	WNorm param.Field[float64] `json:"w_norm" api:"required"`
	// Left edge of the source rectangle, as a fraction of canvas width. May be
	// negative when the source extends off-canvas.
	XNorm param.Field[float64] `json:"x_norm" api:"required"`
	// Top edge of the source rectangle, as a fraction of canvas height. May be
	// negative when the source extends off-canvas.
	YNorm param.Field[float64] `json:"y_norm" api:"required"`
}

func (r SourcePositionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Motion-trajectory conditioning control
type TrajectoryControlParam struct {
	// Enable or disable trajectory conditioning. Omit to use the model default.
	Enabled param.Field[bool] `json:"enabled"`
	// Point-trajectory sparsity from 0 to 1. Higher values use fewer motion anchors.
	Sparsity param.Field[float64] `json:"sparsity"`
}

func (r TrajectoryControlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Video duration
type VideoDuration string

const (
	VideoDuration5s  VideoDuration = "5s"
	VideoDuration10s VideoDuration = "10s"
)

func (r VideoDuration) IsKnown() bool {
	switch r {
	case VideoDuration5s, VideoDuration10s:
		return true
	}
	return false
}

// Ray 3.2 video-to-video edit controls. Only valid under `video.edit` when `type`
// is `video_edit`. The source video must be 18 seconds or shorter; output duration
// matches the source.
type VideoEditOptionsParam struct {
	// When true, the model derives the control schedule from the source video. When
	// omitted, supplying strength or controls implies manual mode.
	AutoControls param.Field[bool] `json:"auto_controls"`
	// Per-signal manual conditioning controls for video edits
	Controls param.Field[AdvancedControlsParam] `json:"controls"`
	// Parallel list of non-negative, unique frame positions in the source video's
	// frame grid where each keyframes[i] is anchored. Must match keyframes in length.
	KeyframeIndexes param.Field[[]int64] `json:"keyframe_indexes"`
	// Multi-anchor guide-frame images at arbitrary source-frame positions (parallel
	// with keyframe_indexes). Up to 64 anchors. Mutually exclusive with
	// video.start_frame (the single-anchor case). Each entry takes the same ImageRef
	// shape as source / image_ref[].
	Keyframes param.Field[[]ImageRefParam] `json:"keyframes"`
	// How much a video edit preserves or reimagines the source
	Strength param.Field[VideoEditStrength] `json:"strength"`
}

func (r VideoEditOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How much a video edit preserves or reimagines the source
type VideoEditStrength string

const (
	VideoEditStrengthAdhere1    VideoEditStrength = "adhere_1"
	VideoEditStrengthAdhere2    VideoEditStrength = "adhere_2"
	VideoEditStrengthAdhere3    VideoEditStrength = "adhere_3"
	VideoEditStrengthFlex1      VideoEditStrength = "flex_1"
	VideoEditStrengthFlex2      VideoEditStrength = "flex_2"
	VideoEditStrengthFlex3      VideoEditStrength = "flex_3"
	VideoEditStrengthReimagine1 VideoEditStrength = "reimagine_1"
	VideoEditStrengthReimagine2 VideoEditStrength = "reimagine_2"
	VideoEditStrengthReimagine3 VideoEditStrength = "reimagine_3"
)

func (r VideoEditStrength) IsKnown() bool {
	switch r {
	case VideoEditStrengthAdhere1, VideoEditStrengthAdhere2, VideoEditStrengthAdhere3, VideoEditStrengthFlex1, VideoEditStrengthFlex2, VideoEditStrengthFlex3, VideoEditStrengthReimagine1, VideoEditStrengthReimagine2, VideoEditStrengthReimagine3:
		return true
	}
	return false
}

// Ray 3.2 video request options. Common output settings live at the top level for
// `type=video`, `type=video_edit`, and `type=video_reframe`; video-to-video
// conditioning lives under `edit`.
type VideoOptionsParam struct {
	// Video duration
	Duration param.Field[VideoDuration] `json:"duration"`
	// Ray 3.2 video-to-video edit controls. Only valid under `video.edit` when `type`
	// is `video_edit`. The source video must be 18 seconds or shorter; output duration
	// matches the source.
	Edit param.Field[VideoEditOptionsParam] `json:"edit"`
	// Media reference for guided generation. Provide exactly one of url, inline base64
	// data, generation_id, or file_id. URL/data references accept image media at image
	// positions; video_edit and video_reframe sources also accept source.url or
	// source.data when source.media_type is a video/\* MIME. generation_id chains
	// image_edit off a prior image output, video_edit/video_reframe off a prior video
	// output, and video.start_frame/end_frame for extension. file_id references a file
	// previously uploaded via POST /files — see the Files API.
	EndFrame param.Field[ImageRefParam] `json:"end_frame"`
	// Export EXR alongside the MP4. Requires hdr=true.
	ExrExport param.Field[bool] `json:"exr_export"`
	// Media reference for guided generation. Provide exactly one of url, inline base64
	// data, generation_id, or file_id. URL/data references accept image media at image
	// positions; video_edit and video_reframe sources also accept source.url or
	// source.data when source.media_type is a video/\* MIME. generation_id chains
	// image_edit off a prior image output, video_edit/video_reframe off a prior video
	// output, and video.start_frame/end_frame for extension. file_id references a file
	// previously uploaded via POST /files — see the Files API.
	GuideFrame param.Field[ImageRefParam] `json:"guide_frame"`
	// Generate HDR video. Requires HDR access. Not supported for video_reframe.
	Hdr param.Field[bool] `json:"hdr"`
	// Parallel list of non-negative, unique output-frame positions where each
	// keyframes[i] is anchored, in the duration x 24fps grid (5s -> 0..120, 10s ->
	// 0..240). Must match keyframes in length.
	KeyframeIndexes param.Field[[]int64] `json:"keyframe_indexes"`
	// Image-to-video guide frames (type=video only), each pinned to an output-frame
	// position via the parallel keyframe_indexes. 1-64 anchors: a single anchor is a
	// valid start-pinned i2v (an alternate to start_frame), and any count up to 64
	// places guide frames at arbitrary positions. Unlike start_frame/end_frame (the
	// legacy 2-frame surface), this supports arbitrary positions, 10s durations, and
	// HDR. Mutually exclusive with start_frame / end_frame / loop. Only supported on
	// model ray-3.2. For video-to-video keyframes use video.edit.keyframes on
	// type=video_edit instead.
	Keyframes param.Field[[]ImageRefParam] `json:"keyframes"`
	// Generate a seamlessly looping video. Only valid for type=video; not supported
	// with duration=10s or hdr=true.
	Loop param.Field[bool] `json:"loop"`
	// Ray 3.2 video output resolution. 360p is the draft tier (fast, low-cost
	// previews), accepted on type=video, video_edit, and video_reframe; on type=video
	// it is SDR-only (not valid with hdr=true). 1080p is public for video generation;
	// video_reframe 1080p is still rolling out and may return a coming-soon validation
	// error until enabled for the caller.
	Resolution param.Field[VideoResolution] `json:"resolution"`
	// Normalized source rectangle inside the output canvas for video_reframe. Omit to
	// let the model choose the default centered-fit crop.
	SourcePosition param.Field[SourcePositionParam] `json:"source_position"`
	// Media reference for guided generation. Provide exactly one of url, inline base64
	// data, generation_id, or file_id. URL/data references accept image media at image
	// positions; video_edit and video_reframe sources also accept source.url or
	// source.data when source.media_type is a video/\* MIME. generation_id chains
	// image_edit off a prior image output, video_edit/video_reframe off a prior video
	// output, and video.start_frame/end_frame for extension. file_id references a file
	// previously uploaded via POST /files — see the Files API.
	StartFrame param.Field[ImageRefParam] `json:"start_frame"`
}

func (r VideoOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Ray 3.2 video output resolution. 360p is the draft tier (fast, low-cost
// previews), accepted on type=video, video_edit, and video_reframe; on type=video
// it is SDR-only (not valid with hdr=true). 1080p is public for video generation;
// video_reframe 1080p is still rolling out and may return a coming-soon validation
// error until enabled for the caller.
type VideoResolution string

const (
	VideoResolution360p  VideoResolution = "360p"
	VideoResolution540p  VideoResolution = "540p"
	VideoResolution720p  VideoResolution = "720p"
	VideoResolution1080p VideoResolution = "1080p"
)

func (r VideoResolution) IsKnown() bool {
	switch r {
	case VideoResolution360p, VideoResolution540p, VideoResolution720p, VideoResolution1080p:
		return true
	}
	return false
}

type GenerationNewParams struct {
	// Text prompt
	Prompt param.Field[string] `json:"prompt" api:"required"`
	// Output aspect ratio. Valid values depend on the selected model and generation
	// type; the server validates the final model-specific set.
	AspectRatio param.Field[GenerationNewParamsAspectRatio] `json:"aspect_ratio"`
	// Reference images for style/content guidance. Up to 9 for type 'image', up to 8
	// for type 'image_edit'.
	ImageRef param.Field[[]ImageRefParam] `json:"image_ref"`
	// Layer-extraction options for type=layering (model uni-1). The image to decompose
	// rides body.source; body.prompt optionally guides how to split it (max 500
	// characters). The server plans the layers automatically before generating.
	Layering param.Field[GenerationNewParamsLayering] `json:"layering"`
	// Model identifier. `uni-1` is the default image tier; `uni-1-max` produces
	// higher-quality output than `uni-1` at a higher per-image price. `ray-3.2` is the
	// public video model for text-to-video, image-to-video, and video-to-video
	// editing.
	Model param.Field[Model] `json:"model"`
	// Output image format
	OutputFormat param.Field[GenerationNewParamsOutputFormat] `json:"output_format"`
	// Media reference for guided generation. Provide exactly one of url, inline base64
	// data, generation_id, or file_id. URL/data references accept image media at image
	// positions; video_edit and video_reframe sources also accept source.url or
	// source.data when source.media_type is a video/\* MIME. generation_id chains
	// image_edit off a prior image output, video_edit/video_reframe off a prior video
	// output, and video.start_frame/end_frame for extension. file_id references a file
	// previously uploaded via POST /files — see the Files API.
	Source param.Field[ImageRefParam] `json:"source"`
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
	// Ray 3.2 video request options. Common output settings live at the top level for
	// `type=video`, `type=video_edit`, and `type=video_reframe`; video-to-video
	// conditioning lives under `edit`.
	Video param.Field[VideoOptionsParam] `json:"video"`
	// Enable web search grounding — the agent can search the web and download
	// reference images before generating.
	WebSearch param.Field[bool] `json:"web_search"`
}

func (r GenerationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Output aspect ratio. Valid values depend on the selected model and generation
// type; the server validates the final model-specific set.
type GenerationNewParamsAspectRatio string

const (
	GenerationNewParamsAspectRatio3_1  GenerationNewParamsAspectRatio = "3:1"
	GenerationNewParamsAspectRatio2_1  GenerationNewParamsAspectRatio = "2:1"
	GenerationNewParamsAspectRatio21_9 GenerationNewParamsAspectRatio = "21:9"
	GenerationNewParamsAspectRatio16_9 GenerationNewParamsAspectRatio = "16:9"
	GenerationNewParamsAspectRatio4_3  GenerationNewParamsAspectRatio = "4:3"
	GenerationNewParamsAspectRatio3_2  GenerationNewParamsAspectRatio = "3:2"
	GenerationNewParamsAspectRatio1_1  GenerationNewParamsAspectRatio = "1:1"
	GenerationNewParamsAspectRatio3_4  GenerationNewParamsAspectRatio = "3:4"
	GenerationNewParamsAspectRatio2_3  GenerationNewParamsAspectRatio = "2:3"
	GenerationNewParamsAspectRatio9_16 GenerationNewParamsAspectRatio = "9:16"
	GenerationNewParamsAspectRatio1_2  GenerationNewParamsAspectRatio = "1:2"
	GenerationNewParamsAspectRatio1_3  GenerationNewParamsAspectRatio = "1:3"
)

func (r GenerationNewParamsAspectRatio) IsKnown() bool {
	switch r {
	case GenerationNewParamsAspectRatio3_1, GenerationNewParamsAspectRatio2_1, GenerationNewParamsAspectRatio21_9, GenerationNewParamsAspectRatio16_9, GenerationNewParamsAspectRatio4_3, GenerationNewParamsAspectRatio3_2, GenerationNewParamsAspectRatio1_1, GenerationNewParamsAspectRatio3_4, GenerationNewParamsAspectRatio2_3, GenerationNewParamsAspectRatio9_16, GenerationNewParamsAspectRatio1_2, GenerationNewParamsAspectRatio1_3:
		return true
	}
	return false
}

// Layer-extraction options for type=layering (model uni-1). The image to decompose
// rides body.source; body.prompt optionally guides how to split it (max 500
// characters). The server plans the layers automatically before generating.
type GenerationNewParamsLayering struct {
	// Output resolution for every extracted layer. 1k is faster and lower cost; 2k
	// re-renders each layer at higher quality (priced higher, per layer).
	Resolution param.Field[GenerationNewParamsLayeringResolution] `json:"resolution"`
}

func (r GenerationNewParamsLayering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Output resolution for every extracted layer. 1k is faster and lower cost; 2k
// re-renders each layer at higher quality (priced higher, per layer).
type GenerationNewParamsLayeringResolution string

const (
	GenerationNewParamsLayeringResolution1k GenerationNewParamsLayeringResolution = "1k"
	GenerationNewParamsLayeringResolution2k GenerationNewParamsLayeringResolution = "2k"
)

func (r GenerationNewParamsLayeringResolution) IsKnown() bool {
	switch r {
	case GenerationNewParamsLayeringResolution1k, GenerationNewParamsLayeringResolution2k:
		return true
	}
	return false
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
	GenerationNewParamsTypeImage        GenerationNewParamsType = "image"
	GenerationNewParamsTypeImageEdit    GenerationNewParamsType = "image_edit"
	GenerationNewParamsTypeVideo        GenerationNewParamsType = "video"
	GenerationNewParamsTypeVideoEdit    GenerationNewParamsType = "video_edit"
	GenerationNewParamsTypeVideoReframe GenerationNewParamsType = "video_reframe"
	GenerationNewParamsTypeLayering     GenerationNewParamsType = "layering"
)

func (r GenerationNewParamsType) IsKnown() bool {
	switch r {
	case GenerationNewParamsTypeImage, GenerationNewParamsTypeImageEdit, GenerationNewParamsTypeVideo, GenerationNewParamsTypeVideoEdit, GenerationNewParamsTypeVideoReframe, GenerationNewParamsTypeLayering:
		return true
	}
	return false
}
