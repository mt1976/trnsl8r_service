package text

import (
	"fmt"

	"github.com/mt1976/frantic-core/timing"
	"github.com/mt1976/trnsl8r_service/app/dao/textStore"
)

func GetLocalised(signature, localeFilter string) (*textStore.TextStore, error) {
	watch := timing.Start("Localise", "Get", signature)
	// Log the start of the retrieval operation
	// logger.Info.Printf("GET: [%v] Id=[%v]", strings.ToUpper(tableName), id)

	// Log the ID of the dest object being retrieved
	// logger.Info.Printf("GET: %v Object: %v", tableName, fmt.Sprintf("%+v", id))

	// Initialize an empty d object
	u := textStore.New()

	// Retrieve the dest object from the database based on the given IDs
	u, err := textStore.GetBy(textStore.Fields.Signature, signature)
	if err != nil {
		// Log and panic if there is an error reading the dest object
		// logger.Info.Printf("Reading %v: [%v] %v ", tableName, id, err.Error())
		return nil, fmt.Errorf("reading %v Id=[%v] %v ", "Localise", signature, err.Error())
		//	panic(err)
	}

	// Log the retrieved dest object
	// u.Spew()

	// Log the completion of the retrieval operation
	// logger.Info.Printf("GET: [%v] Id=[%v] RealName=[%v] ", strings.ToUpper(tableName), u.ID, u.RealName)

	watch.Stop(1)

	return u, nil
}
