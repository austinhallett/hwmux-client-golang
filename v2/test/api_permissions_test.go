/*
HW Mux Reservation System

Testing PermissionsAPIService

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

func Test_hwmux_PermissionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PermissionsAPIService PermissionsGroupsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsDestroy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsDestroy(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsDeviceGroupsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsDeviceGroupsCreate(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsDeviceGroupsDestroy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32
		var nameOrId string

		httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsDeviceGroupsDestroy(context.Background(), id, nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsDeviceGroupsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsDeviceGroupsList(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsDeviceGroupsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32
		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsDeviceGroupsPartialUpdate(context.Background(), id, nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsDeviceGroupsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32
		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsDeviceGroupsUpdate(context.Background(), id, nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsDevicesCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsDevicesCreate(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsDevicesDestroy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32
		var nameOrId string

		httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsDevicesDestroy(context.Background(), id, nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsDevicesList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsDevicesList(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsDevicesPartialUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32
		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsDevicesPartialUpdate(context.Background(), id, nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsDevicesUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32
		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsDevicesUpdate(context.Background(), id, nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsLabelsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsLabelsCreate(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsLabelsDestroy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32
		var nameOrId string

		httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsLabelsDestroy(context.Background(), id, nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsLabelsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsLabelsList(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsLabelsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32
		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsLabelsPartialUpdate(context.Background(), id, nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsLabelsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32
		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsLabelsUpdate(context.Background(), id, nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsPartialUpdate(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsRetrieve", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsRetrieve(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsUpdate(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsUsersCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsUsersCreate(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PermissionsAPIService PermissionsGroupsUsersDestroy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nameOrId string
		var usernameOrId string

		httpRes, err := apiClient.PermissionsAPI.PermissionsGroupsUsersDestroy(context.Background(), nameOrId, usernameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}