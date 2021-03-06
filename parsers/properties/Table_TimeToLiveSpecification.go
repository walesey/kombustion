package properties

	import "fmt"

type Table_TimeToLiveSpecification struct {
	
	
	AttributeName interface{} `yaml:"AttributeName"`
	Enabled interface{} `yaml:"Enabled"`
}

func (resource Table_TimeToLiveSpecification) Validate() []error {
	errs := []error{}
	
	
	if resource.AttributeName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AttributeName'"))
	}
	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	return errs
}
