package commercetools_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/trendhim/commercetools-go-sdk-legacy/commercetools"
	"github.com/stretchr/testify/assert"

	"github.com/trendhim/commercetools-go-sdk-legacy/testutil"
)

func TestAPIClientCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	input := &commercetools.APIClientDraft{
		Scope: "manage_project",
		Name:  "api-client-name",
	}

	fmt.Println(output)

	_, err := client.APIClientCreate(context.TODO(), input)
	assert.Nil(t, err)

	expectedBody := `{
		"name": "api-client-name",
		"scope": "manage_project"
	}`

	assert.JSONEq(t, expectedBody, output.JSON)
}

func TestAPIClientDelete(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.APIClientDeleteWithID(context.TODO(), "1234")
	assert.Nil(t, err)

	assert.Equal(t, "/unittest/api-clients/1234", output.URL.Path)
}

func TestAPIClientGetByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("api-client.example.json"), nil, nil)
	defer server.Close()

	timestamp, _ := time.Parse(time.RFC3339, "2018-01-01T00:00:00.001Z")

	input := &commercetools.APIClient{
		ID:         "api-client-id",
		Scope:      "view_products",
		Name:       "api-client-name",
		CreatedAt:  &timestamp,
		LastUsedAt: &commercetools.Date{Year: 2017, Month: 9, Day: 10},
		Secret:     "secret-passphrase",
	}

	result, err := client.APIClientGetWithID(context.TODO(), "1234")
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}

func TestAPIClientQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.QueryInput{
		Limit: 500,
	}
	_, err := client.APIClientQuery(context.TODO(), &queryInput)
	assert.Nil(t, err)

	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
