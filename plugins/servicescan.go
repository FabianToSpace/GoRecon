package plugins

type ServiceScan struct {
	Name        string // Name of the ServiceScan
	Description string // Description of the ServiceScan
	Type        string // Type (e.g. TCP or UPD)
	Tags        []string
	Run         run // Execute Function
}
