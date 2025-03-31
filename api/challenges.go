package api

import (
	"fmt"

	ctfd "github.com/ctfer-io/go-ctfd/api"
)

// GetChallenge return the challenge corresponding to the given id,
// in the CTFd-Chall-Manager plugin datamodel.
func GetChallenge(client *ctfd.Client, id string, opts ...ctfd.Option) (*Challenge, error) {
	ch := &Challenge{}
	if err := client.Get(fmt.Sprintf("/challenges/%s", id), nil, ch, opts...); err != nil {
		return nil, err
	}
	return ch, nil
}

type (
	PostChallengesParams struct {
		// CTFd Dynamic arguments

		Name           string             `json:"name"`
		Category       string             `json:"category"`
		Description    string             `json:"description"`
		Attribution    *string            `json:"attribution,omitempty"`
		Function       *string            `json:"function,omitempty"`
		ConnectionInfo *string            `json:"connection_info,omitempty"`
		Initial        *int               `json:"initial,omitempty"`
		Decay          *int               `json:"decay,omitempty"`
		Minimum        *int               `json:"minimum,omitempty"`
		MaxAttempts    *int               `json:"max_attempts,omitempty"`
		NextID         *int               `json:"next_id,omitempty"`
		Requirements   *ctfd.Requirements `json:"requirements,omitempty"`
		State          string             `json:"state"`
		Type           string             `json:"type"`

		// CTFer.io DynamicIaC arguments

		DestroyOnFlag bool              `json:"destroy_on_flag"`
		Shared        bool              `json:"shared"`
		ManaCost      int               `json:"mana_cost"`
		ScenarioID    int               `json:"scenario_id"`
		Timeout       *int              `json:"timeout,omitempty"`
		Until         *string           `json:"until,omitempty"`
		Additional    map[string]string `json:"additional,omitempty"`
	}
)

// PostChallenges creates a CTFd-Chall-Manager plugin challenge.
func PostChallenges(client *ctfd.Client, params *PostChallengesParams, opts ...ctfd.Option) (*Challenge, error) {
	ch := &Challenge{}
	if err := client.Post("/challenges", params, ch, opts...); err != nil {
		return nil, err
	}
	return ch, nil
}

type (
	PatchChallengeParams struct {
		// CTFd Dynamic arguments

		Name           string  `json:"name"`
		Category       string  `json:"category"`
		Description    string  `json:"description"`
		Attribution    *string `json:"attribution,omitempty"`
		Function       *string `json:"function,omitempty"`
		ConnectionInfo *string `json:"connection_info,omitempty"`
		Value          *int    `json:"value,omitempty"`
		Initial        *int    `json:"initial,omitempty"`
		Decay          *int    `json:"decay,omitempty"`
		Minimum        *int    `json:"minimum,omitempty"`
		MaxAttempts    *int    `json:"max_attempts,omitempty"`
		NextID         *int    `json:"next_id,omitempty"`
		// Requirements can update the challenge's behavior and prerequisites i.e.
		// the other challenges the team/user must have solved before.
		// WARNING: it won't return those in the response body, so updating this
		// field requires you to do it manually through *Client.GetChallengeRequirements
		Requirements *ctfd.Requirements `json:"requirements,omitempty"`
		State        string             `json:"state"`

		// CTFer.io DynamicIaC arguments

		DestroyOnFlag bool    `json:"destroy_on_flag"`
		Shared        bool    `json:"shared"`
		ManaCost      int     `json:"mana_cost"`
		ScenarioID    int     `json:"scenario_id"`
		Timeout       *int    `json:"timeout,omitempty"`
		Until         *string `json:"until,omitempty"`
	}
)

// PatchChallenges updates a challenge configuration.
func PatchChallenges(client *ctfd.Client, id string, params *PatchChallengeParams, opts ...ctfd.Option) (*Challenge, error) {
	ch := &Challenge{}
	if err := client.Patch(fmt.Sprintf("/challenges/%s", id), params, ch, opts...); err != nil {
		return nil, err
	}
	return ch, nil
}
