package mdbxcmd

import (
	"flag"
)

var flagPrintVersion bool
var flagOpenNoSubDir bool

func init() {
	flag.BoolVar(&flagPrintVersion, "V", false, "Write the library version number to the standard output, and exit.")
	flag.BoolVar(&flagOpenNoSubDir, "n", false, "Open LDMB environment(s) which do not use subdirectories.")
}

//func printVersion(w io.Writer) {
//	fmt.Fprintln(w, mdbx.VersionString())
//}
//
//// PrintVersion writes the API version in a human readable format to
//// os.Stdout.
//func PrintVersion() {
//	if flagPrintVersion {
//		printVersion(os.Stdout)
//		os.Exit(0)
//	}
//}

// OpenFlag returns the bitwise OR'd set of flags specified by options defined
// in the package.  The returned value may be OR'd with additional flags if
// needed.
//func OpenFlag() uint {
//	var flag uint
//	if flagOpenNoSubDir {
//		flag |= mdbx.NoSubdir
//	}
//	return flag
//}
