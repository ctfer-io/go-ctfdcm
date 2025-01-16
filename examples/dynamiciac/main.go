package main

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	ctfd "github.com/ctfer-io/go-ctfd/api"
	"github.com/ctfer-io/go-ctfdcm/api"
	"gopkg.in/yaml.v3"
)

const (
	url    = "http://localhost:8000"
	apiKey = "ctfd_e6b5f755e9407ce2e5951d00be1c305776406875e5effd3cdeec83b17d464080"
)

func main() {
	cli := ctfd.NewClient(url, "", "", apiKey)
	opts := []ctfd.Option{
		ctfd.WithContext(context.Background()),
	}

	fmt.Println("[+] Creating scenario file")
	scn, err := scenario()
	if err != nil {
		log.Fatalf("    Failed to craft scenario: %s", err)
	}
	f, err := cli.PostFiles(&ctfd.PostFilesParams{
		Files: []*ctfd.InputFile{
			{
				Name:    "scenario.zip",
				Content: scn,
			},
		},
	}, opts...)
	if err != nil {
		log.Fatalf("    Failed to create scenario file: %s", err)
	}
	fmt.Printf("    Created scenario file %d\n", f[0].ID)

	fmt.Println("[+] Creating dynamic_iac challenge")
	ch, err := api.PostChallenges(cli, &api.PostChallengesParams{
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
		DestroyOnFlag: false,
		Shared:        true,
		ManaCost:      1,
		ScenarioID:    f[0].ID,
		Timeout:       ptr(600),
	}, opts...)
	if err != nil {
		log.Fatalf("Creating challenge: %s", err)
	}
	fmt.Printf("ch: %#+v\n", ch)
	fmt.Printf("    Created challenge %d\n", ch.ID)
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
	bmain, err := compile()
	if err != nil {
		return nil, err
	}
	w, err = archive.Create("main")
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(w, bytes.NewBuffer(bmain)); err != nil {
		return nil, err
	}

	// Complete zip creation
	if err := archive.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func compile() ([]byte, error) {
	cmd := exec.Command("go", "build", "-o", "scenario/main", "scenario/main.go")
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0")
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	defer func() {
		cmd := exec.Command("rm", "scenario/main")
		_ = cmd.Run()
	}()
	return os.ReadFile("scenario/main")
}

func ptr[T any](t T) *T {
	return &t
}
