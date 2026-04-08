// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lumaagents_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/luma-agents-go"
	"github.com/stainless-sdks/luma-agents-go/internal/testutil"
	"github.com/stainless-sdks/luma-agents-go/option"
)

func TestUsage(t *testing.T) {
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
	order, err := client.Store.Orders.New(context.TODO(), lumaagents.StoreOrderNewParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", order.ID)
}
