package result

type Status struct {
	EGovFrameVersion string
	JdkVersion       string
	SpringVersion    string
	BootVersion      string
	// Results of the rules
	VersionRule           bool
	ConfigurationRule     bool
	PresentationLayerRule bool
	ServiceLayerRule      bool
	DataAccessLayerRule   bool
	CompatibilityLevel    int
}

var CheckResult *Status

func init() {
	CheckResult = &Status{}
}

func (s *Status) IsCompatible() bool {
	return s.VersionRule && s.ConfigurationRule && s.PresentationLayerRule && s.ServiceLayerRule && s.DataAccessLayerRule
}
