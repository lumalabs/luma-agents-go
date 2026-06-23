// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lumaagents

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/lumalabs/luma-agents-go/internal/apijson"
	"github.com/lumalabs/luma-agents-go/internal/apiquery"
	"github.com/lumalabs/luma-agents-go/internal/param"
	"github.com/lumalabs/luma-agents-go/internal/requestconfig"
	"github.com/lumalabs/luma-agents-go/option"
)

// FileService contains methods and other services that help with interacting with
// the luma API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// Upload a file to your namespace, then reference it from a generation via
// ImageRef.file_id (as source, image_ref[], video.start_frame, keyframes, and so
// on). Two upload modes share this endpoint, selected by Content-Type:
//
//   - multipart/form-data — send the bytes inline in the `file` part. Best for small
//     files (subject to an inline size cap; larger files must use the presigned
//     flow). The returned file is already `pending` ingest.
//
//   - application/json — request a presigned upload. The response `upload` envelope
//     tells you where to PUT the bytes; afterward call POST
//     /files/{file_id}/complete to start ingest. Use this for larger files.
func (r *FileService) New(ctx context.Context, body FileNewParams, opts ...option.RequestOption) (res *CreateFileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List the files in your namespace, newest first. Keyset-paginated: when has_more
// is true, pass next_cursor back as cursor.
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *FileList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Soft-delete a file. It can no longer be referenced from new generations. Returns
// 204 with no body.
func (r *FileService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return err
	}
	path := fmt.Sprintf("files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Finalize a presigned upload after you have PUT the bytes to the upload URL.
// Kicks off ingest/moderation and returns the file, which transitions to `ready`
// (or `failed`) asynchronously — poll GET /files/{file_id} to observe the terminal
// state.
func (r *FileService) Complete(ctx context.Context, fileID string, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/%s/complete", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Retrieve metadata for a single file in your namespace.
func (r *FileService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Result of POST /files. In the multipart (inline) flow `upload` is null and the
// file is already `pending` ingest. In the presigned (JSON) flow `upload` carries
// the PUT envelope and the file stays `pending` until you call POST
// /files/{file_id}/complete. Top-level `id` and `state` are conveniences that
// mirror `file.id` and `file.state`; the full record is always under `file`.
type CreateFileResponse struct {
	// File identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// A file in the caller's namespace.
	File File `json:"file" api:"required"`
	// Lifecycle state of an uploaded file. `pending` until bytes are received and the
	// ingest pipeline runs; `ready` once it can be referenced from a generation;
	// `failed` if ingest/moderation rejected it; `deleted` after a soft-delete.
	State FileState `json:"state" api:"required"`
	// Where to PUT the file bytes for a presigned (JSON) upload. Issue an HTTP PUT of
	// the raw bytes to `url` with the given `headers`, then call POST
	// /files/{file_id}/complete.
	Upload PresignedUpload        `json:"upload" api:"nullable"`
	JSON   createFileResponseJSON `json:"-"`
}

// createFileResponseJSON contains the JSON metadata for the struct
// [CreateFileResponse]
type createFileResponseJSON struct {
	ID          apijson.Field
	File        apijson.Field
	State       apijson.Field
	Upload      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateFileResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createFileResponseJSON) RawJSON() string {
	return r.raw
}

// A file in the caller's namespace.
type File struct {
	// File identifier, referenced as ImageRef.file_id.
	ID string `json:"id" api:"required" format:"uuid"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// MIME type of the stored bytes (for example, image/jpeg).
	MimeType string `json:"mime_type" api:"required"`
	// How the file is intended to be used in a generation. `input` is the primary
	// subject (e.g. the source image for an edit); `reference` is style/content
	// guidance.
	Purpose FilePurpose `json:"purpose" api:"required"`
	// Size of the stored object in bytes.
	SizeBytes int64 `json:"size_bytes" api:"required"`
	// Lifecycle state of an uploaded file. `pending` until bytes are received and the
	// ingest pipeline runs; `ready` once it can be referenced from a generation;
	// `failed` if ingest/moderation rejected it; `deleted` after a soft-delete.
	State FileState `json:"state" api:"required"`
	// Soft-delete timestamp, if the file was deleted.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// TTL set at upload, if any. After this time Luma may automatically delete the
	// file and reclaim its bytes — you don't need to call DELETE yourself.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Human-readable reason when state is failed.
	FailureReason string `json:"failure_reason" api:"nullable"`
	// Original filename supplied at upload, if any.
	Filename string `json:"filename" api:"nullable"`
	// The opaque end-user tag supplied at upload, echoed back unchanged.
	// Abuse-attribution only; not an access-control primitive.
	UserID string   `json:"user_id" api:"nullable"`
	JSON   fileJSON `json:"-"`
}

// fileJSON contains the JSON metadata for the struct [File]
type fileJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	MimeType      apijson.Field
	Purpose       apijson.Field
	SizeBytes     apijson.Field
	State         apijson.Field
	DeletedAt     apijson.Field
	ExpiresAt     apijson.Field
	FailureReason apijson.Field
	Filename      apijson.Field
	UserID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *File) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileJSON) RawJSON() string {
	return r.raw
}

// Keyset-paginated page of files, newest first. When has_more is true, pass
// next_cursor back as the cursor query parameter to fetch the next page.
// next_cursor is opaque.
type FileList struct {
	// Files in this page.
	Data []File `json:"data" api:"required"`
	// Whether more files exist beyond this page.
	HasMore bool `json:"has_more" api:"required"`
	// Opaque cursor for the next page, when has_more is true.
	NextCursor string       `json:"next_cursor" api:"nullable"`
	JSON       fileListJSON `json:"-"`
}

// fileListJSON contains the JSON metadata for the struct [FileList]
type fileListJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileListJSON) RawJSON() string {
	return r.raw
}

// How the file is intended to be used in a generation. `input` is the primary
// subject (e.g. the source image for an edit); `reference` is style/content
// guidance.
type FilePurpose string

const (
	FilePurposeInput     FilePurpose = "input"
	FilePurposeReference FilePurpose = "reference"
)

func (r FilePurpose) IsKnown() bool {
	switch r {
	case FilePurposeInput, FilePurposeReference:
		return true
	}
	return false
}

// Lifecycle state of an uploaded file. `pending` until bytes are received and the
// ingest pipeline runs; `ready` once it can be referenced from a generation;
// `failed` if ingest/moderation rejected it; `deleted` after a soft-delete.
type FileState string

const (
	FileStatePending FileState = "pending"
	FileStateReady   FileState = "ready"
	FileStateFailed  FileState = "failed"
	FileStateDeleted FileState = "deleted"
)

func (r FileState) IsKnown() bool {
	switch r {
	case FileStatePending, FileStateReady, FileStateFailed, FileStateDeleted:
		return true
	}
	return false
}

// Where to PUT the file bytes for a presigned (JSON) upload. Issue an HTTP PUT of
// the raw bytes to `url` with the given `headers`, then call POST
// /files/{file_id}/complete.
type PresignedUpload struct {
	// When the presigned URL expires.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// HTTP method to use for the upload — always PUT.
	Method string `json:"method" api:"required"`
	// Presigned S3 URL to PUT the bytes to.
	URL string `json:"url" api:"required" format:"uri"`
	// Headers that must be sent with the PUT request.
	Headers map[string]string   `json:"headers"`
	JSON    presignedUploadJSON `json:"-"`
}

// presignedUploadJSON contains the JSON metadata for the struct [PresignedUpload]
type presignedUploadJSON struct {
	ExpiresAt   apijson.Field
	Method      apijson.Field
	URL         apijson.Field
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PresignedUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r presignedUploadJSON) RawJSON() string {
	return r.raw
}

type FileNewParams struct {
	// MIME type of the bytes you will upload.
	MimeType param.Field[string] `json:"mime_type" api:"required"`
	// Exact size in bytes of the object you will PUT. Up to 5 GiB (the S3 single-PUT
	// ceiling).
	SizeBytes param.Field[int64] `json:"size_bytes" api:"required"`
	// Optional TTL. After this time Luma may automatically delete the file and reclaim
	// its bytes.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// Optional original filename to record.
	Filename param.Field[string] `json:"filename"`
	// How the file is intended to be used in a generation. `input` is the primary
	// subject (e.g. the source image for an edit); `reference` is style/content
	// guidance.
	Purpose param.Field[FilePurpose] `json:"purpose"`
	// Optional opaque end-user tag for abuse attribution. Mirrors the user_id field on
	// POST /generations.
	UserID param.Field[string] `json:"user_id"`
}

func (r FileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FileListParams struct {
	// Opaque pagination cursor from a prior response's next_cursor.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum files to return (1–100). Defaults to 25.
	Limit param.Field[int64] `query:"limit"`
	// Filter to files with this purpose.
	Purpose param.Field[FilePurpose] `query:"purpose"`
	// Filter to files in this state.
	State param.Field[FileState] `query:"state"`
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
