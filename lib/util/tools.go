package util

import "github.com/duxphp/duxgo-ui/lib/form"

func modelToSelectOptions(data []map[string]any, key string, name string) []form.SelectOptions {
	options := []form.SelectOptions{}
	for _, datum := range data {
		options = append(options, form.SelectOptions{
			Key:  datum[key],
			Name: datum[name],
		})
	}
	return options
}

func modelToRadioOptions(data []map[string]any, key string, name string) []form.RadioOptions {
	options := []form.RadioOptions{}
	for _, datum := range data {
		options = append(options, form.RadioOptions{
			Key:  datum[key],
			Name: datum[name],
		})
	}
	return options
}
