package defalt

import (
	"fmt"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"This is HOME")
}
