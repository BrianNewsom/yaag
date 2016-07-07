package annotations

import (
	"github.com/briannewsom/yaag/yaag"
	"github.com/briannewsom/yaag/yaag/models"
	"net/http"
	"strings"
)

func New(rawHeader, value string) models.YaagAnnotationHeader {
	header := strings.TrimPrefix(rawHeader, "Yaag-Annotation-")
	return models.YaagAnnotationHeader{
		RawHeader: rawHeader,
		Header:    header,
		Value:     value,
	}
}

func Add(w http.ResponseWriter, key, value string) {
	if yaag.IsOn() {
		w.Header().Add("Yaag-Annotation-"+key, value)
	}
}

func AddDescription(w http.ResponseWriter, desc string) {
	Add(w, "Description", desc)
}

func AddReturnType(w http.ResponseWriter, returnType string) {
	Add(w, "Return-Type", returnType)
}
