package properties

	import "fmt"

type ConfigurationTemplate_ConfigurationOptionSetting struct {
	
	
	
	
	Namespace interface{} `yaml:"Namespace"`
	OptionName interface{} `yaml:"OptionName"`
	ResourceName interface{} `yaml:"ResourceName,omitempty"`
	Value interface{} `yaml:"Value,omitempty"`
}

func (resource ConfigurationTemplate_ConfigurationOptionSetting) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.Namespace == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Namespace'"))
	}
	if resource.OptionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'OptionName'"))
	}
	return errs
}
