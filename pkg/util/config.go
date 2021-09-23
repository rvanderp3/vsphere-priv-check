package util

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/rvanderp3/vsphere-priv-check/pkg/types"
	"io/ioutil"
	"sigs.k8s.io/yaml"
)

const (
	installConfig = "install-config.yaml"
)

// LoadConfig Load configuration from install-config.yaml in the working directory
func LoadConfig() (*types.Platform, error) {
	file, err := ioutil.ReadFile(installConfig)
	if err != nil {
		return nil, errors.New("unable to read install-config.yaml")
	}

	var vsphereConfig types.Platform
	var body map[string]interface{}
	err = yaml.Unmarshal(file, &body)
	var platform map[string]interface{}
	if val, ok := body["platform"]; ok {
		platform = val.(map[string]interface{})
	} else {
		return nil, errors.New("'platform' field not found in install-config.yaml")
	}

	var vsphere map[string]interface{}
	if val, ok := platform["vsphere"]; ok {
		vsphere = val.(map[string]interface{})
	} else {
		return nil, errors.New("'vsphere' field not found in 'platform' install-config.yaml")
	}
	err = mapstructure.Decode(vsphere, &vsphereConfig)
	if err != nil {
		return nil, err
	}
	return &vsphereConfig, nil
}
