package plugins

type Service struct {
	Target   string
	Protocol string
	Port     int
	Name     string
	Secure   bool
	Version  string
}
