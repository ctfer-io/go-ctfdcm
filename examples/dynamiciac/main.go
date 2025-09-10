package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ctfer-io/chall-manager/pkg/scenario"
	ctfd "github.com/ctfer-io/go-ctfd/api"
	"github.com/ctfer-io/go-ctfdcm/api"
)

const (
	url    = "http://localhost:8000"
	apiKey = "ctfd_e6b5f755e9407ce2e5951d00be1c305776406875e5effd3cdeec83b17d464080"
	ref    = "localhost:5000/scenario:v0.1.0"
)

func main() {
	cli := ctfd.NewClient(url, "", "", apiKey)
	opts := []ctfd.Option{
		ctfd.WithContext(context.Background()),
	}

	fmt.Println("[+] Creating scenario file")
	if err := pushScenario(); err != nil {
		log.Fatalf("    Failed to craft and push scenario: %s", err)
	}
	fmt.Printf("    Created scenario file %s\n", ref)

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
		Scenario:      ref,
		Timeout:       ptr(600),
		Additional: map[string]string{
			"image": "pandatix/license-lvl1:latest",
		},
	}, opts...)
	if err != nil {
		log.Fatalf("Creating challenge: %s", err)
	}
	fmt.Printf("ch: %#+v\n", ch)
	fmt.Printf("    Created challenge %d\n", ch.ID)
}

func pushScenario() error {
	ctx := context.Background()
	return scenario.EncodeOCI(ctx, ref, "../examples/dynamiciac/scenario", true, "", "")
}

func ptr[T any](t T) *T {
	return &t
}
