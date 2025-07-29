package api_test

import (
	"strconv"
	"testing"

	ctfd "github.com/ctfer-io/go-ctfd/api"
	"github.com/ctfer-io/go-ctfdcm/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_F_Run(t *testing.T) {
	// Scenario:
	//
	// This mocks a CTF development and test, first by setting up
	// CTFd, then push a scenario and create a DynamicIaC challenge,
	// finally pops up an instance, and in the end destroy the challenge
	// to check the instance has been deleted.

	// 1a. Get nonce and session to mock a browser first
	nonce, session, err := ctfd.GetNonceAndSession(CTFD_URL)
	require.NoError(t, err)
	admin := ctfd.NewClient(CTFD_URL, nonce, session, "")

	t.Cleanup(func() {
		// Due to relicas, forced to unpause the event elseway the test is not reproducible
		_ = admin.PatchConfigs(&ctfd.PatchConfigsParams{
			Paused: ptr(false),
		})

		_ = admin.Reset(&ctfd.ResetParams{
			Accounts:      ptr("y"),
			Submissions:   ptr("y"),
			Challenges:    ptr("y"),
			Pages:         ptr("y"),
			Notifications: ptr("y"),
		})
	})

	// 1b. Configure the CTF
	err = admin.Setup(&ctfd.SetupParams{
		CTFName:                "CTFer",
		CTFDescription:         "Ephemeral CTFd running for API tests purposes.",
		UserMode:               "users",
		Name:                   "ctfer",
		Email:                  "ctfer-io@protonmail.com",
		Password:               "password", // This is not real, don't bother trying x)
		ChallengeVisibility:    "public",
		AccountVisibility:      "public",
		ScoreVisibility:        "public",
		RegistrationVisibility: "public",
		VerifyEmails:           false,
		TeamSize:               nil,
		CTFLogo:                nil,
		CTFBanner:              nil,
		CTFSmallIcon:           nil,
		CTFTheme:               "core",
		ThemeColor:             "",
		Start:                  "",
		End:                    "",
	})
	require.NoError(t, err)

	// 1c. Create an API Key to avoid session/nonce+cookies dance
	token, err := admin.PostTokens(&ctfd.PostTokensParams{
		Expiration:  "2222-01-01",
		Description: "Example API token.",
	})
	require.NoError(t, err)
	admin.SetAPIKey(*token.Value)

	// 2. Create a DynamicIaC challenge
	ch, err := api.PostChallenges(admin, &api.PostChallengesParams{
		// CTFd
		Name:           "Break The License 1/2",
		Category:       "crypto",
		Description:    "...",
		Attribution:    ptr("pandatix"),
		Function:       ptr("logarithmic"),
		ConnectionInfo: ptr("ssh -l user@crypto1.ctfer.io"),
		MaxAttempts:    ptr(3),
		Initial:        ptr(500),
		Decay:          ptr(17),
		Minimum:        ptr(50),
		State:          "visible",
		Type:           "dynamic_iac",
		// CTFer.io Chall-Manager plugin
		DestroyOnFlag: false, //
		Shared:        false, // keep shared to instanciate
		ManaCost:      1,
		Scenario:      ref,
		Timeout:       ptr(600),
	})
	require.NoError(t, err)

	// 3. Pop up an instance
	ist, err := api.PostInstance(admin, &api.PostInstanceParams{
		ChallengeID: strconv.Itoa(ch.ID),
	})
	require.NoError(t, err)
	assert.NotEmpty(t, ist.ConnectionInfo)

	// 4. Destroy the challenge
	err = admin.DeleteChallenge(ch.ID)
	require.NoError(t, err)

	// 5. Check the instance has been destroyed
	ist, err = api.GetInstance(admin, &api.GetInstanceParams{
		ChallengeID: strconv.Itoa(ch.ID),
	})
	assert.Error(t, err) // challenge does not exist
	assert.Nil(t, ist)   // so does its instance
}

func ptr[T any](t T) *T {
	return &t
}
