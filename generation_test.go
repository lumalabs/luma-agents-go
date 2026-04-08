// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lumaagents_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/luma-agents-go"
	"github.com/stainless-sdks/luma-agents-go/internal/testutil"
	"github.com/stainless-sdks/luma-agents-go/option"
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
		ImageRef: lumaagents.F([]lumaagents.GenerationNewParamsImageRef{{
			Data:      lumaagents.F("data"),
			MediaType: lumaagents.F("media_type"),
			URL:       lumaagents.F("url"),
		}}),
		Model:        lumaagents.F("model"),
		OutputFormat: lumaagents.F(lumaagents.GenerationNewParamsOutputFormatPng),
		Source: lumaagents.F(lumaagents.GenerationNewParamsSource{
			Data:      lumaagents.F("data"),
			MediaType: lumaagents.F("media_type"),
			URL:       lumaagents.F("url"),
		}),
		Style:     lumaagents.F(lumaagents.GenerationNewParamsStyleAuto),
		Type:      lumaagents.F(lumaagents.GenerationNewParamsTypeImage),
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
