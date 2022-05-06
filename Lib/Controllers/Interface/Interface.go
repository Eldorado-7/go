/*
*** @author: Javad Bayzavi
*** @version: 1.0.1
*** @email: javadbayzavi@gmail.com
*** @year: 2021
*/

package Interface

import (
	"net/http"
)

type ControllerInterface interface {
	Run(params map[string]string) http.HandlerFunc
	processMe([]string) string
}
