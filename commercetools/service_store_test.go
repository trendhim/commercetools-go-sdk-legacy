package commercetools_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/trendhim/commercetools-go-sdk-legacy/commercetools"
	"github.com/stretchr/testify/assert"

	"github.com/trendhim/commercetools-go-sdk-legacy/testutil"
)

func TestStoreCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("store.json"), &output, nil)
	defer server.Close()

	input := &commercetools.StoreDraft{
		Key: "test123",
		Name: &commercetools.LocalizedString{
			"en": "test123",
		},
	}

	store, err := client.StoreCreate(context.TODO(), input)
	fmt.Printf("%#v", store)
	assert.Nil(t, err)

	assert.Equal(t, "test123", store.Key)
	assert.Equal(t, "test123", (*store.Name)["en"])
}

func TestStoreUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *commercetools.StoreUpdateWithIDInput
		requestBody string
	}{
		{
			desc: "Set locale",
			input: &commercetools.StoreUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StoreUpdateAction{
					&commercetools.StoreSetNameAction{
						Name: &commercetools.LocalizedString{
							"en": "foo",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setName",
						"name": {
							"en": "foo"
						}
					}
				]
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}

			client, server := testutil.MockClient(t, "{}", &output, nil)
			defer server.Close()

			_, err := client.StoreUpdateWithID(context.TODO(), tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/stores/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestStoreUpdateByKey(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *commercetools.StoreUpdateWithKeyInput
		requestBody string
	}{
		{
			desc: "Set locale",
			input: &commercetools.StoreUpdateWithKeyInput{
				Key:     "test123",
				Version: 2,
				Actions: []commercetools.StoreUpdateAction{
					&commercetools.StoreSetNameAction{
						Name: &commercetools.LocalizedString{
							"en": "foo",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setName",
						"name": {
							"en": "foo"
						}
					}
				]
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}

			client, server := testutil.MockClient(t, "{}", &output, nil)
			defer server.Close()

			_, err := client.StoreUpdateWithKey(context.TODO(), tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/stores/key=test123", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestStoreDelete(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.StoreDeleteWithID(context.TODO(), "ba8d47e5-6591-4ca2-af4c-d547f062bf35", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/stores/ba8d47e5-6591-4ca2-af4c-d547f062bf35", output.URL.Path)
}

func TestStoreDeleteByKey(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.StoreDeleteWithKey(context.TODO(), "test123", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/stores/key=test123", output.URL.Path)
}

func TestStoreGetByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("store.json"), nil, nil)
	defer server.Close()

	store, err := client.StoreGetWithID(context.TODO(), "ba8d47e5-6591-4ca2-af4c-d547f062bf35")

	assert.Nil(t, err)
	assert.Equal(t, "ba8d47e5-6591-4ca2-af4c-d547f062bf35", store.ID)
}

func TestStoreGetByKey(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("store.json"), nil, nil)
	defer server.Close()

	store, err := client.StoreGetWithKey(context.TODO(), "test123")

	assert.Nil(t, err)
	assert.Equal(t, "ba8d47e5-6591-4ca2-af4c-d547f062bf35", store.ID)
}

func TestStoreQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.QueryInput{
		Limit: 500,
	}
	_, err := client.StoreQuery(context.TODO(), &queryInput)

	assert.Nil(t, err)
	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
