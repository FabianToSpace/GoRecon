package plugins

type PortScan struct {
	Name        string   // Name of the Portscan
	Description string   // Description of the Portscan
	Type        string   // Type (e.g. TCP or UPD)
	Tags        []string // ['default', 'default-portscan']
	Run         run      // Execute Function
}
