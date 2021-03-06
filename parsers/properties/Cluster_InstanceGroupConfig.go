package properties

	import "fmt"

type Cluster_InstanceGroupConfig struct {
	
	
	
	
	
	
	
	
	BidPrice interface{} `yaml:"BidPrice,omitempty"`
	InstanceCount interface{} `yaml:"InstanceCount"`
	InstanceType interface{} `yaml:"InstanceType"`
	Market interface{} `yaml:"Market,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	Configurations interface{} `yaml:"Configurations,omitempty"`
	EbsConfiguration *Cluster_EbsConfiguration `yaml:"EbsConfiguration,omitempty"`
	AutoScalingPolicy *Cluster_AutoScalingPolicy `yaml:"AutoScalingPolicy,omitempty"`
}

func (resource Cluster_InstanceGroupConfig) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	if resource.InstanceCount == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceCount'"))
	}
	if resource.InstanceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceType'"))
	}
	return errs
}
