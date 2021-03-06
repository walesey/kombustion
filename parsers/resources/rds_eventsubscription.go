package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RDSEventSubscription struct {
	Type       string                      `yaml:"Type"`
	Properties RDSEventSubscriptionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RDSEventSubscriptionProperties struct {
	Enabled interface{} `yaml:"Enabled,omitempty"`
	SnsTopicArn interface{} `yaml:"SnsTopicArn"`
	SourceType interface{} `yaml:"SourceType,omitempty"`
	EventCategories interface{} `yaml:"EventCategories,omitempty"`
	SourceIds interface{} `yaml:"SourceIds,omitempty"`
}

func NewRDSEventSubscription(properties RDSEventSubscriptionProperties, deps ...interface{}) RDSEventSubscription {
	return RDSEventSubscription{
		Type:       "AWS::RDS::EventSubscription",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRDSEventSubscription(name string, data string) (cf types.ValueMap, err error) {
	var resource RDSEventSubscription
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RDSEventSubscription - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RDSEventSubscription) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RDSEventSubscriptionProperties) Validate() []error {
	errs := []error{}
	if resource.SnsTopicArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SnsTopicArn'"))
	}
	return errs
}
