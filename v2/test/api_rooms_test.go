/*
HW Mux Reservation System

Testing RoomsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package hwmux

import (
	"context"
	"testing"

	openapiclient "github.com/Silabs-UTF/hwmux-client-golang/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_hwmux_RoomsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoomsAPIService RoomsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RoomsAPI.RoomsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomsAPIService RoomsRetrieve", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.RoomsAPI.RoomsRetrieve(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}