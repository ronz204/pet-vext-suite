package commands

type Scanner struct {
	*Command
}

func NewScanner() *Scanner {
	return &Scanner{
		Command: &Command{
			Use:     "vext scan",
			Example: "vext scan 192.168.1.1",
			Short:   "Scans for vulnerabilities",
			Long:    "The scan command allows you to scan a target for known vulnerabilities using various techniques.",
		},
	}
}

var ScannerCmd = NewScanner()
