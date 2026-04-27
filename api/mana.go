package api

import ctfd "github.com/ctfer-io/go-ctfd/api"

func GetMana(client *ctfd.Client, opts ...ctfd.Option) (*Mana, *ctfd.MetaResponse, error) {
	m := &Mana{}
	meta, err := client.Get("/plugins/ctfd-chall-manager/mana", nil, m, opts...)
	if err != nil {
		return nil, meta, err
	}
	return m, meta, nil
}
