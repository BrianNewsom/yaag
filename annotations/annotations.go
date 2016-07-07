package annotations

import (
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

func AddDescription(w http.ResponseWriter, desc string) {
	w.Header().Add("Yaag-Annotation-Description", desc)
}

func AddReturnType(w http.ResponseWriter, returnType string) {
	w.Header().Add("Yaag-Annotation-Return-Type", returnType)
}

func Add(w http.ResponseWriter, key, value string) {
	w.Header().Add("Yaag-Annotation-"+key, value)
}
