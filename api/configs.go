package api

import ctfd "github.com/ctfer-io/go-ctfd/api"

type (
	PatchSettingsParams struct {
		APIURL    string `json:"chall-manager:chall-manager_api_url"`
		ManaTotal int    `json:"chall-manager:chall-manager_mana_total"`
	}
)

// PatchSettings handle the patch of CTFd-Chall-Manager settings.
func PatchSettings(client *ctfd.Client, params PatchSettingsParams, opts ...ctfd.Option) error {
	return client.Patch("/configs", params, nil, opts...)
}
