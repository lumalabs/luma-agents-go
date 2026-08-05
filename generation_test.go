// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lumaagents_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/lumalabs/luma-agents-go"
	"github.com/lumalabs/luma-agents-go/internal/testutil"
	"github.com/lumalabs/luma-agents-go/option"
)

func TestGenerationNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lumaagents.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	_, err := client.Generations.New(context.TODO(), lumaagents.GenerationNewParams{
		Prompt:      lumaagents.F("A glass of iced coffee on a marble countertop, morning light streaming through a window"),
		AspectRatio: lumaagents.F(lumaagents.GenerationNewParamsAspectRatio3_1),
		ImageRef: lumaagents.F([]lumaagents.ImageRefParam{{
			Data:         lumaagents.F("data"),
			FileID:       lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			GenerationID: lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			MediaType:    lumaagents.F("media_type"),
			URL:          lumaagents.F("url"),
		}}),
		Layering: lumaagents.F(lumaagents.GenerationNewParamsLayering{
			Resolution: lumaagents.F(lumaagents.GenerationNewParamsLayeringResolution1k),
		}),
		Model:        lumaagents.F(lumaagents.ModelUni1),
		OutputFormat: lumaagents.F(lumaagents.GenerationNewParamsOutputFormatPng),
		Source: lumaagents.F(lumaagents.ImageRefParam{
			Data:         lumaagents.F("data"),
			FileID:       lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			GenerationID: lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			MediaType:    lumaagents.F("media_type"),
			URL:          lumaagents.F("url"),
		}),
		Style:  lumaagents.F(lumaagents.GenerationNewParamsStyleAuto),
		Type:   lumaagents.F(lumaagents.GenerationNewParamsTypeImage),
		UserID: lumaagents.F("user_id"),
		Video: lumaagents.F(lumaagents.VideoOptionsParam{
			Duration: lumaagents.F(lumaagents.VideoDuration5s),
			Edit: lumaagents.F(lumaagents.VideoEditOptionsParam{
				AutoControls: lumaagents.F(true),
				Controls: lumaagents.F(lumaagents.AdvancedControlsParam{
					Depth: lumaagents.F(lumaagents.DepthControlParam{
						Blur:    lumaagents.F(0.000000),
						Enabled: lumaagents.F(true),
					}),
					Face: lumaagents.F(lumaagents.FaceControlParam{
						Enabled: lumaagents.F(true),
					}),
					Normals: lumaagents.F(lumaagents.NormalsControlParam{
						Augmentation: lumaagents.F(0.000000),
						Enabled:      lumaagents.F(true),
					}),
					Pose: lumaagents.F(lumaagents.PoseControlParam{
						Enabled:  lumaagents.F(true),
						Strength: lumaagents.F(lumaagents.PoseControlStrengthPrecise),
					}),
					Trajectory: lumaagents.F(lumaagents.TrajectoryControlParam{
						Enabled:  lumaagents.F(true),
						Sparsity: lumaagents.F(0.000000),
					}),
				}),
				KeyframeIndexes: lumaagents.F([]int64{int64(0)}),
				Keyframes: lumaagents.F([]lumaagents.ImageRefParam{{
					Data:         lumaagents.F("data"),
					FileID:       lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					GenerationID: lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					MediaType:    lumaagents.F("media_type"),
					URL:          lumaagents.F("url"),
				}}),
				Strength: lumaagents.F(lumaagents.VideoEditStrengthAdhere1),
			}),
			EndFrame: lumaagents.F(lumaagents.ImageRefParam{
				Data:         lumaagents.F("data"),
				FileID:       lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				GenerationID: lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				MediaType:    lumaagents.F("media_type"),
				URL:          lumaagents.F("url"),
			}),
			ExrExport: lumaagents.F(true),
			GuideFrame: lumaagents.F(lumaagents.ImageRefParam{
				Data:         lumaagents.F("data"),
				FileID:       lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				GenerationID: lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				MediaType:    lumaagents.F("media_type"),
				URL:          lumaagents.F("url"),
			}),
			Hdr:             lumaagents.F(true),
			KeyframeIndexes: lumaagents.F([]int64{int64(0)}),
			Keyframes: lumaagents.F([]lumaagents.ImageRefParam{{
				Data:         lumaagents.F("data"),
				FileID:       lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				GenerationID: lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				MediaType:    lumaagents.F("media_type"),
				URL:          lumaagents.F("url"),
			}}),
			Loop:       lumaagents.F(true),
			Resolution: lumaagents.F(lumaagents.VideoResolution360p),
			SourcePosition: lumaagents.F(lumaagents.SourcePositionParam{
				HNorm: lumaagents.F(1.000000),
				WNorm: lumaagents.F(1.000000),
				XNorm: lumaagents.F(-2.000000),
				YNorm: lumaagents.F(-2.000000),
			}),
			StartFrame: lumaagents.F(lumaagents.ImageRefParam{
				Data:         lumaagents.F("data"),
				FileID:       lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				GenerationID: lumaagents.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				MediaType:    lumaagents.F("media_type"),
				URL:          lumaagents.F("url"),
			}),
		}),
		WebSearch: lumaagents.F(true),
	})
	if err != nil {
		var apierr *lumaagents.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGenerationGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lumaagents.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	_, err := client.Generations.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lumaagents.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
