package watch

import (
	"path"
	"strings"

	"github.com/rs/zerolog/log"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func toGVR(gvr string) schema.GroupVersionResource {
	tokens := strings.Split(gvr, "/")
	if len(tokens) < 3 {
		tokens = append([]string{""}, tokens...)
	}

	return schema.GroupVersionResource{
		Group:    tokens[0],
		Version:  tokens[1],
		Resource: tokens[2],
	}
}

func namespaced(n string) (string, string) {
	ns, po := path.Split(n)

	return strings.Trim(ns, "/"), po
}

// Dump for debug.
func Dump(f *Factory) {
	log.Debug().Msgf("----------- FACTORIES -------------")
	for ns := range f.factories {
		log.Debug().Msgf("  Factory for NS %q", ns)
	}
	log.Debug().Msgf("-----------------------------------")
}

// Debug for debug.
func Debug(f *Factory, ns string, gvr string) {
	log.Debug().Msgf("----------- DEBUG FACTORY (%s) -------------", gvr)
	inf := f.factories[ns].ForResource(toGVR(gvr))
	for i, k := range inf.Informer().GetStore().ListKeys() {
		log.Debug().Msgf("%d -- %s", i, k)
	}
}