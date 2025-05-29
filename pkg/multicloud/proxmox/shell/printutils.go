
package shell

import (
	"yunion.io/x/pkg/util/printutils"
)

func printList(data interface{}, total, offset, limit int, columns []string) {
	printutils.PrintInterfaceList(data, total, offset, limit, columns)
}

func printObject(obj interface{}) {
	printutils.PrintInterfaceObject(obj)
}
