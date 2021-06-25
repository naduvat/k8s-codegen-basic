package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/runtime/schema"
)

// Define your schema name and the version
var SchemeGroupVersion = schema.GroupVersion {
    Group: "k8s.crd.io",
    Version: "v1alpha1",
}

var (
    SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
    localSchemeBuilder = &SchemeBuilder
    AddToScheme        = SchemeBuilder.AddToScheme
)

func init() {
    // We only register manually written functions here. The registration of the
    // generated functions takes place in the generated files. The separation
    // makes the code compile even when the generated files are missing.
    localSchemeBuilder.Register(addKnownTypes)
}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
    return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
    return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to the given scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
    scheme.AddKnownTypes(SchemeGroupVersion,
                         &PolicyDefinition{},
                         &PolicyDefinitionList{},
    )
    metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
    return nil
}
