package api

import ctfd "github.com/ctfer-io/go-ctfd/api"

func GetMana(client *ctfd.Client, opts ...ctfd.Option) (*Mana, error) {
	m := &Mana{}
	if err := client.Get("/plugins/ctfd-chall-manager/mana", nil, m, opts...); err != nil {
		return nil, err
	}
	return m, nil
}
