package api

import (
	ctfd "github.com/ctfer-io/go-ctfd/api"
)

type (
	Challenge struct {
		// CTFd Dynamic arguments

		ID             int                `json:"id"`
		Name           string             `json:"name"`
		Description    string             `json:"description"`
		Attribution    *string            `json:"attribution,omitempty"`
		ConnectionInfo *string            `json:"connection_info,omitempty"`
		MaxAttempts    *int               `json:"max_attempts,omitempty"`
		Function       *string            `json:"function,omitempty"`
		Value          int                `json:"value"`
		Initial        *int               `json:"initial,omitempty"`
		Decay          *int               `json:"decay,omitempty"`
		Minimum        *int               `json:"minimum,omitempty"`
		Logic          string             `json:"logic"`
		Category       string             `json:"category"`
		Type           string             `json:"type"`
		TypeData       *ctfd.Type         `json:"type_data,omitempty"`
		State          string             `json:"state"`
		NextID         *int               `json:"next_id"`
		Requirements   *ctfd.Requirements `json:"requirements"` // List of challenge IDs to complete before
		Solves         int                `json:"solves"`
		SolvedByMe     bool               `json:"solved_by_me"`

		// CTFer.io DynamicIaC arguments

		DestroyOnFlag bool              `json:"destroy_on_flag"`
		Shared        bool              `json:"shared"`
		ManaCost      int               `json:"mana_cost"`
		Scenario      string            `json:"scenario"`
		Until         *string           `json:"until,omitempty"`
		Timeout       *int              `json:"timeout,omitempty"`
		Additional    map[string]string `json:"additional,omitempty"`
		Min           int               `json:"min"`
		Max           int               `json:"max"`
	}

	Instance struct {
		ConnectionInfo string  `json:"connectionInfo"`
		Until          *string `json:"until,omitempty"`
		Since          string  `json:"since"`
	}

	Mana struct {
		Used  int `json:"mana_used"`
		Total int `json:"mana_total"`
	}
)
