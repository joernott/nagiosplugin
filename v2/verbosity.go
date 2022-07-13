package nagiosplugin

import (
	"errors"
)

// https://nagios-plugins.org/doc/guidelines.html#AEN41
const (
	VERBOSITY_MINIMAL int = iota
	VERBOSITY_SINGLE_LINE
	VERBOSITY_MULTI_LINE
	VERBOSITY_DEBUG
)

// Set the output verbosity to the desired level
func (c *Check) SetVerbosity(Verbosity int) error {
	switch Verbosity {
	case VERBOSITY_MINIMAL, VERBOSITY_SINGLE_LINE:
		c.verbosity = Verbosity
		c.messageSeparator = ", "
	case VERBOSITY_MULTI_LINE, VERBOSITY_DEBUG:
		c.verbosity = Verbosity
		c.messageSeparator = "\n"
	default:
		return errors.New("Illegal verbosity")
	}
	return nil
}

// Set the messages returned when the verbosity is set to minimal
func (c *Check) SetMinimalResult(Code Status, Message string) {
	c.minimalResults[Code] = Message
}
