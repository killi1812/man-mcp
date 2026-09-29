// Package app preforms basic app functions like setup, loading config and global definitions
package app

import (
	"fmt"
)

// Setup will preform app setup or panic of it fails
// Can only be called once
func Setup() {
	// Logger setup
	{
		var err error

		if Build == BuildDev {
			err = devLoggerSetup()
			if err != nil {
				fmt.Printf("err: %v\n", err)
				panic("faled to setup logger")
			}
		} else {
			err = prodLoggerSetup()
			if err != nil {
				fmt.Printf("err: %v\n", err)
				panic("faled to setup logger")
			}
		}
	}
}
