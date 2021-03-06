package properties

	import "fmt"

type DeploymentGroup_S3Location struct {
	
	
	
	
	
	Bucket interface{} `yaml:"Bucket"`
	BundleType interface{} `yaml:"BundleType,omitempty"`
	ETag interface{} `yaml:"ETag,omitempty"`
	Key interface{} `yaml:"Key"`
	Version interface{} `yaml:"Version,omitempty"`
}

func (resource DeploymentGroup_S3Location) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.Bucket == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Bucket'"))
	}
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	return errs
}
