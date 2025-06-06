/*
 * Copyright 2021 kloeckner.i GmbH
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"os"

	"github.com/db-operator/db-operator/v2/pkg/utils/kci"
	yaml "gopkg.in/yaml.v2"
)

// LoadConfig reads config file for db-operator from defined path and parse
func LoadConfig() (*Config, error) {
	path := kci.StringNotEmpty(os.Getenv("CONFIG_PATH"), "/srv/config/config.yaml")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := &Config{}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
