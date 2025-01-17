//go:build integration
// +build integration

package legalentity

import (
	"context"
	openapiclient "github.com/adyen/adyen-go-api-library/v6/src/legalentity"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test_Integration_LegalEntity_HostedOnboardingApiService(t *testing.T) {

	godotenv.Load("./../../../.env")

	var (
		username = os.Getenv("ADYEN_LEM_USERNAME")
		password = os.Getenv("ADYEN_LEM_PASSWORD")
		env      = openapiclient.TestEnv
	)

	configuration, err := openapiclient.NewConfiguration(username, password, env)
	require.Nil(t, err, "Error creating Config object")
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HostedOnboardingApiService ListHostedOnboardingPageThemes", func(t *testing.T) {

		resp, httpRes, err := apiClient.HostedOnboardingApi.ListHostedOnboardingPageThemes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HostedOnboardingApiService GetOnboardingLinkTheme", func(t *testing.T) {

		themeId := os.Getenv("LEM_THEME_ID")

		resp, httpRes, err := apiClient.HostedOnboardingApi.GetOnboardingLinkTheme(context.Background(), themeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
