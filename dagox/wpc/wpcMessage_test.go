package wpc

import (
	"testing"
)

var wpcMessageContent = "{\"id_Ext\":\"28148190\",\"id\":\"00179\",\"entryContainerID\":\"6602\",\"movement\":\"MODIFIED\",\"coreAttribs\":[{\"id\":\"/Supermercado - Marcas LTS\",\"type\":\"SPEC\",\"children\":[{\"id\":\"CodigoMarca\",\"type\":\"STRING\",\"value\":\"00179\"},{\"id\":\"Descripcion\",\"type\":\"STRING\",\"value\":\"N.P.U.\"},{\"id\":\"MetadataFechaModificacion\",\"type\":\"DATE\",\"value\":\"20021217_000000\"},{\"id\":\"ExternalCode\",\"type\":\"STRING\",\"value\":\"28148190\"}]}]}";

func TestUnmarshal(t *testing.T) {
	expectedObjectId := "28148190"

	wpcMessage := Unmarshal([]byte(wpcMessageContent))
	if wpcMessage.ObjectId != expectedObjectId {
		t.Errorf("Unmarshall = %q, want %q", wpcMessage.ObjectId, expectedObjectId)
	}
}
