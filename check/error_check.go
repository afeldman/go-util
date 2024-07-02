// Description: check your data for errors
package check

import log "github.com/sirupsen/logrus"

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
