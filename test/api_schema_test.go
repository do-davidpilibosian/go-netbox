/*
NetBox REST API

Testing SchemaApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package netbox

import (
	"context"
	"testing"

	openapiclient "github.com/netbox-community/go-netbox/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_netbox_SchemaApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SchemaApiService SchemaRetrieve", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SchemaApi.SchemaRetrieve(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
