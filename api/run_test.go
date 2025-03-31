package api_test

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"os/exec"
	"strconv"
	"testing"

	ctfd "github.com/ctfer-io/go-ctfd/api"
	"github.com/ctfer-io/go-ctfdcm/api"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_F_Run(t *testing.T) {
	// Scenario:
	//
	// This mocks a CTF development and test, first by setting up
	// CTFd, then push a scenario and create a DynamicIaC challenge,
	// finally pops up an instance, and in the end destroy the challenge
	// to check the instance has been deleted.

	assert := assert.New(t)

	// 1a. Get nonce and session to mock a browser first
	nonce, session, err := ctfd.GetNonceAndSession(CTFD_URL)
	if !assert.NoError(err) {
		return
	}
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
	if !assert.NoError(err) {
		return
	}

	// 1c. Create an API Key to avoid session/nonce+cookies dance
	token, err := admin.PostTokens(&ctfd.PostTokensParams{
		Expiration:  "2222-01-01",
		Description: "Example API token.",
	})
	if !assert.NoError(err) {
		return
	}
	admin.SetAPIKey(*token.Value)

	// 2. Add the scenario
	scn, err := scenario()
	if !assert.NoError(err) {
		return
	}
	f, err := admin.PostFiles(&ctfd.PostFilesParams{
		Files: []*ctfd.InputFile{
			{
				Name:    "scenario.zip",
				Content: scn,
			},
		},
	})
	if !assert.NoError(err) {
		return
	}

	// 3. Create a DynamicIaC challenge
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
		ScenarioID:    f[0].ID,
		Timeout:       ptr(600),
	})
	if !assert.NoError(err) {
		return
	}

	// 4. Pop up an instance
	ist, err := api.PostInstance(admin, &api.PostInstanceParams{
		ChallengeID: strconv.Itoa(ch.ID),
	})
	if !assert.NoError(err) {
		return
	}
	assert.NotEmpty(ist.ConnectionInfo)

	// 5. Destroy the challenge
	err = admin.DeleteChallenge(ch.ID)
	if !assert.NoError(err) {
		return
	}

	// 6. Check the instance has been destroyed
	ist, err = api.GetInstance(admin, &api.GetInstanceParams{
		ChallengeID: strconv.Itoa(ch.ID),
	})
	assert.Error(err) // challenge does not exist
	assert.Nil(ist)   // so does its instance
}

func scenario() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	archive := zip.NewWriter(buf)

	// Add Pulumi.yaml file
	mp := map[string]any{
		"name": "scenario",
		"runtime": map[string]any{
			"name": "go",
			"options": map[string]any{
				"binary": "./main",
			},
		},
		"description": "An example scenario.",
	}
	b, err := yaml.Marshal(mp)
	if err != nil {
		return nil, err
	}
	w, err := archive.Create("Pulumi.yaml")
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(w, bytes.NewBuffer(b)); err != nil {
		return nil, err
	}

	// Add binary file
	fs, err := compile()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = fs.Close()
	}()

	fst, err := fs.Stat()
	if err != nil {
		return nil, err
	}
	header, err := zip.FileInfoHeader(fst)
	if err != nil {
		return nil, err
	}
	header.Name = "main"

	// Create archive
	f, err := archive.CreateHeader(header)
	if err != nil {
		return nil, err
	}

	// Copy the file's contents into the archive.
	_, err = io.Copy(f, fs)
	if err != nil {
		return nil, err
	}

	// Complete zip creation
	if err := archive.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func compile() (*os.File, error) {
	cmd := exec.Command("go", "build", "-o", "main", "../examples/dynamiciac/scenario/main.go")
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0")
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	defer func() {
		cmd := exec.Command("rm", "main")
		_ = cmd.Run()
	}()
	return os.Open("main")
}

func ptr[T any](t T) *T {
	return &t
}
