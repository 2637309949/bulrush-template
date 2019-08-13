// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadYaml defined load yaml file
func LoadYaml(path string, v interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
