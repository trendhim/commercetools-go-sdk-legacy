package commercetools_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/trendhim/commercetools-go-sdk-legacy/commercetools"
	"github.com/stretchr/testify/assert"

	"github.com/trendhim/commercetools-go-sdk-legacy/testutil"
)

func TestProductProjectionSearch(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.ProductProjectionSearchInput{
		Text:   map[string]string{"nl": "foobar"},
		Filter: []string{"category.id:foo", "category.id:bar"},
	}
	_, err := client.ProductProjectionSearch(context.TODO(), &queryInput)

	assert.Nil(t, err)
	assert.Equal(t, url.Values{
		"text.nl": {"foobar"},
		"filter":  {"category.id:foo", "category.id:bar"},
	}, output.URL.Query())
}
