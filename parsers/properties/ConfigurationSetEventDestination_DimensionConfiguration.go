package properties

	import "fmt"

type ConfigurationSetEventDestination_DimensionConfiguration struct {
	
	
	
	DefaultDimensionValue interface{} `yaml:"DefaultDimensionValue"`
	DimensionName interface{} `yaml:"DimensionName"`
	DimensionValueSource interface{} `yaml:"DimensionValueSource"`
}

func (resource ConfigurationSetEventDestination_DimensionConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	if resource.DefaultDimensionValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DefaultDimensionValue'"))
	}
	if resource.DimensionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DimensionName'"))
	}
	if resource.DimensionValueSource == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DimensionValueSource'"))
	}
	return errs
}
