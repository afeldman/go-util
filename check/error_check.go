// Description: check your data for errors
package check

import log "github.com/sirupsen/logrus"

// CheckError: check for errors
// err: error
// return: bool
func CheckError(err error) bool {
	if err != nil {
		log.Error(err.Error())
		return false
	}
	return true
}

// CheckErrorFatal: check for errors
// err: error
func CheckErrorFatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
