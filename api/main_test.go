package api_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ctfer-io/chall-manager/pkg/scenario"
)

var (
	CTFD_URL = ""
	REGISTRY = ""

	ref = "scenario:v0.1.0"
)

func TestMain(m *testing.M) {
	u, ok := os.LookupEnv("CTFD_URL")
	if !ok {
		fmt.Println("Environment variable CTFD_URL is not set, please indicate the domain name/IP address to reach out the cluster.")
		os.Exit(1)
	}
	CTFD_URL = u

	r, ok := os.LookupEnv("REGISTRY")
	if !ok {
		fmt.Println("Environment variable REGISTRY is not set, please indicate the domain name/IP address to reach out the registry.")
		os.Exit(1)
	}
	REGISTRY = r
	ref = fmt.Sprintf("%s/%s", REGISTRY, ref)

	if err := pushScenario(); err != nil {
		fmt.Printf("Pushing scenario %s: %s", ref, err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func pushScenario() error {
	ctx := context.Background()
	return scenario.EncodeOCI(ctx, ref, "../examples/dynamiciac/scenario", true, nil, nil)
}
