// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/trendhim/commercetools-go-sdk-legacy/commercetools"
	"github.com/trendhim/commercetools-go-sdk-legacy/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedMessageGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "174adf2f-783f-4ce5-a2d5-ee7d3ee7caf4",
	  "version": 1,
	  "sequenceNumber": 1,
	  "resource": {
	    "typeId": "category",
	    "id": "45251684-d693-409d-9864-f93f486adb95"
	  },
	  "resourceVersion": 1,
	  "type": "CategoryCreated",
	  "category": {
	    "id": "45251684-d693-409d-9864-f93f486adb95",
	    "version": 1,
	    "name": {
	      "en": "myCategory"
	    },
	    "slug": {
	      "en": "my-category"
	    },
	    "ancestors": [],
	    "orderHint": "0.000014556311799451762128695",
	    "createdAt": "2016-02-16T13:59:39.945Z",
	    "lastModifiedAt": "2016-02-16T13:59:39.945Z",
	    "assets": [],
	    "lastMessageSequenceNumber": 1
	  },
	  "createdAt": "2016-02-16T13:59:39.945Z",
	  "lastModifiedAt": "2016-02-16T13:59:39.945Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	message, err := client.MessageGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, message)

}

func TestGeneratedMessageQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "id": "7d165eba-6032-4614-8e6a-fb78d2a02139",
	      "version": 1,
	      "sequenceNumber": 1,
	      "resource": {
	        "typeId": "category",
	        "id": "c979f0f7-2e6a-4e8b-92eb-59e6e840e797"
	      },
	      "resourceVersion": 1,
	      "type": "CategoryCreated",
	      "category": {
	        "id": "c979f0f7-2e6a-4e8b-92eb-59e6e840e797",
	        "version": 1,
	        "name": {
	          "en": "test-Category"
	        },
	        "slug": {
	          "en": "test-category"
	        },
	        "ancestors": [],
	        "orderHint": "0.00001455631179379695951525",
	        "createdAt": "2016-02-16T13:59:39.379Z",
	        "lastModifiedAt": "2016-02-16T13:59:39.379Z",
	        "assets": [],
	        "lastMessageSequenceNumber": 1
	      },
	      "createdAt": "2016-02-16T13:59:39.379Z",
	      "lastModifiedAt": "2016-02-16T13:59:39.379Z"
	    },
	    {
	      "id": "174adf2f-783f-4ce5-a2d5-ee7d3ee7caf4",
	      "version": 1,
	      "sequenceNumber": 1,
	      "resource": {
	        "typeId": "category",
	        "id": "45251684-d693-409d-9864-f93f486adb95"
	      },
	      "resourceVersion": 1,
	      "type": "CategoryCreated",
	      "category": {
	        "id": "45251684-d693-409d-9864-f93f486adb95",
	        "version": 1,
	        "name": {
	          "en": "myCategory"
	        },
	        "slug": {
	          "en": "my-category"
	        },
	        "ancestors": [],
	        "orderHint": "0.000014556311799451762128695",
	        "createdAt": "2016-02-16T13:59:39.945Z",
	        "lastModifiedAt": "2016-02-16T13:59:39.945Z",
	        "assets": [],
	        "lastMessageSequenceNumber": 1
	      },
	      "createdAt": "2016-02-16T13:59:39.945Z",
	      "lastModifiedAt": "2016-02-16T13:59:39.945Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.MessageQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
