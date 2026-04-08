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

func TestStoreListInventory(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lumaagents.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.ListInventory(context.TODO())
	if err != nil {
		var apierr *lumaagents.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
