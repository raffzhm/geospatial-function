package geospatialfunc

import "encoding/json"

func GCHandlerFunc(Mongostring, dbname, colname string) []byte {
	koneksyen := GetConnectionMongo(Mongostring, dbname)
	datageo := GetAllGeoData(koneksyen, colname)

	jsonnya, _ := json.Marshal(datageo)

	return jsonnya
}
