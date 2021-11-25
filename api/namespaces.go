package api

import (
	"os"
	"strings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getNamespaces() []string {
	namespaces := []string{"master"}

	return namespaces
}

func GetNamespacePods() []NamespacePods {
	namespaces := getNamespaces()
	namespacePods := make([]NamespacePods, len(namespaces))

	for i := 0; i < len(namespaces); i++ {
		namespacePods[i] = NamespacePods{
			Namespace: namespaces[i],
			Pods: getPods(namespaces[i]),
		}
	}

	return namespacePods
}

