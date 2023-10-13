package geospatialfunc

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/petapedia/peda"
)

func init() {
	functions.HTTP("GeoJson", petaPedia)
}

func petaPedia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://raffzhm.github.io")
	fmt.Fprintf(w, peda.GCFHandler("MONGO_URI", "pointmap", "pointmap"))
}
