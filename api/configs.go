package api

import ctfd "github.com/ctfer-io/go-ctfd/api"

type (
	PatchConfigsParams struct {
		APIURL     string `json:"chall-manager:chall-manager_api_url"`
		ManaTotal  int    `json:"chall-manager:chall-manager_mana_total"`
		APITimeout int    `json:"chall-manager:chall-manager_api_timeout"`
	}
)

// PatchConfigs handle the patch of CTFd-Chall-Manager configs.
func PatchConfigs(client *ctfd.Client, params PatchConfigsParams, opts ...ctfd.Option) error {
	return client.Patch("/configs", params, nil, opts...)
}
