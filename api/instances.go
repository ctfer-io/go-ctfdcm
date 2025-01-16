package api

import (
	ctfd "github.com/ctfer-io/go-ctfd/api"
)

type GetInstanceParams struct {
	ChallengeID string `schema:"challengeId"`
}

func GetInstance(client *ctfd.Client, params *GetInstanceParams, opts ...ctfd.Option) (*Instance, error) {
	ist := &Instance{}
	if err := client.Get("/plugins/ctfd-chall-manager/instance", params, ist, opts...); err != nil {
		return nil, err
	}
	return ist, nil
}

type PostInstanceParams struct {
	ChallengeID string `json:"challengeId"`
}

func PostInstance(client *ctfd.Client, params *PostInstanceParams, opts ...ctfd.Option) (*Instance, error) {
	ist := &Instance{}
	if err := client.Post("/plugins/ctfd-chall-manager/instance", params, ist, opts...); err != nil {
		return nil, err
	}
	return ist, nil
}

type PatchInstanceParams struct {
	ChallengeID string `json:"challengeId"`
}

func PatchInstance(client *ctfd.Client, params *PatchInstanceParams, opts ...ctfd.Option) (*Instance, error) {
	ist := &Instance{}
	if err := client.Patch("/plugins/ctfd-chall-manager/instance", params, ist, opts...); err != nil {
		return nil, err
	}
	return ist, nil
}

type DeleteInstanceParams struct {
	ChallengeID string `json:"challengeId"`
}

func DeleteInstance(client *ctfd.Client, params *DeleteInstanceParams, opts ...ctfd.Option) (*Instance, error) {
	ist := &Instance{}
	if err := client.Delete("/plugins/ctfd-chall-manager/instance", params, ist, opts...); err != nil {
		return nil, err
	}
	return ist, nil
}

type GetAdminInstanceParams struct {
	ChallengeID string `schema:"challengeId"`
	SourceID    string `schema:"soureId"`
}

func GetAdminInstance(client *ctfd.Client, params *GetAdminInstanceParams, opts ...ctfd.Option) (*Instance, error) {
	ist := &Instance{}
	if err := client.Get("/plugins/ctfd-chall-manager/admin/instance", params, ist, opts...); err != nil {
		return nil, err
	}
	return ist, nil
}

type PostAdminInstanceParams struct {
	ChallengeID string `json:"challengeId"`
	SourceID    string `json:"sourceId"`
}

func PostAdminInstance(client *ctfd.Client, params *PostAdminInstanceParams, opts ...ctfd.Option) (*Instance, error) {
	ist := &Instance{}
	if err := client.Post("/plugins/ctfd-chall-manager/admin/instance", params, ist, opts...); err != nil {
		return nil, err
	}
	return ist, nil
}

type PatchAdminInstanceParams struct {
	ChallengeID string `json:"challengeId"`
	SourceID    string `json:"sourceId"`
}

func PatchAdminInstance(client *ctfd.Client, params *PatchAdminInstanceParams, opts ...ctfd.Option) (*Instance, error) {
	ist := &Instance{}
	if err := client.Patch("/plugins/ctfd-chall-manager/admin/instance", params, ist, opts...); err != nil {
		return nil, err
	}
	return ist, nil
}

type DeleteAdminInstanceParams struct {
	ChallengeID string `json:"challengeId"`
	SourceID    string `json:"sourceId"`
}

func DeleteAdminInstance(client *ctfd.Client, params *DeleteAdminInstanceParams, opts ...ctfd.Option) (*Instance, error) {
	ist := &Instance{}
	if err := client.Delete("/plugins/ctfd-chall-manager/admin/instance", params, ist, opts...); err != nil {
		return nil, err
	}
	return ist, nil
}
