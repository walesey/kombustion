package properties

	import "fmt"

type Application_KinesisFirehoseInput struct {
	
	
	ResourceARN interface{} `yaml:"ResourceARN"`
	RoleARN interface{} `yaml:"RoleARN"`
}

func (resource Application_KinesisFirehoseInput) Validate() []error {
	errs := []error{}
	
	
	if resource.ResourceARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResourceARN'"))
	}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	return errs
}
