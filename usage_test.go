// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lumaagents_test

import (
	"context"
	"os"
	"testing"

	"github.com/lumalabs/luma-agents-go"
	"github.com/lumalabs/luma-agents-go/internal/testutil"
	"github.com/lumalabs/luma-agents-go/option"
)

func TestUsage(t *testing.T) {
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
	generation, err := client.Generations.New(context.TODO(), lumaagents.GenerationNewParams{
		Prompt:      lumaagents.F("A glass of iced coffee on a marble countertop, morning light streaming through a window"),
		AspectRatio: lumaagents.F(lumaagents.GenerationNewParamsAspectRatio16_9),
		Model:       lumaagents.F("uni-1"),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", generation.ID)
}
