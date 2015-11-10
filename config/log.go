//
package config

import (
	"fmt"
	"github.com/jcelliott/lumber"
)

var (
	Console *lumber.ConsoleLogger
	Log     *lumber.FileLogger
)

// init
func init() {

	// create a console logger
	Console = lumber.NewConsoleLogger(lumber.INFO)

	// create a file logger
	if Log, err = lumber.NewTruncateLogger(Root + "/nanobox.log"); err != nil {
		Fatal("[config/log] lumber.NewAppendLogger() failed", err.Error())
	}
}

// Debug
func Debug(msg string, debug bool) {
	if debug {
		fmt.Printf(msg)
	}
}

// Info
func Info(msg string) {
	Log.Info(msg)
}

// Error
func Error(msg, err string) {
	fmt.Printf("%s (See ~/.nanobox/nanobox.log for details)\n", msg)
	Log.Error(err)
}

// Fatal
func Fatal(msg, err string) {
	fmt.Println("A fatal error occurred (See ~/.nanobox/nanobox.log for details). Exiting...")
	Log.Fatal(fmt.Sprintf("%s - %s", msg, err))
	Log.Close()
	Exit(1)
}
