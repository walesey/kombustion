package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2Host struct {
	Type       string                      `yaml:"Type"`
	Properties EC2HostProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2HostProperties struct {
	AutoPlacement interface{} `yaml:"AutoPlacement,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone"`
	InstanceType interface{} `yaml:"InstanceType"`
}

func NewEC2Host(properties EC2HostProperties, deps ...interface{}) EC2Host {
	return EC2Host{
		Type:       "AWS::EC2::Host",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2Host(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2Host
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2Host - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2Host) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2HostProperties) Validate() []error {
	errs := []error{}
	if resource.AvailabilityZone == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AvailabilityZone'"))
	}
	if resource.InstanceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceType'"))
	}
	return errs
}
