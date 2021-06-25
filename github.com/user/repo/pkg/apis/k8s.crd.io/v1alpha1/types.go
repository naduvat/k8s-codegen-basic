package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PolicyDefinition struct {
    metav1.TypeMeta             `json:",inline"`
    // +optional
    metav1.ObjectMeta           `json:"metadata"`
    Name string                 `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PolicyDefinitionList struct {
    metav1.TypeMeta             `json:",inline"`
    metav1.ListMeta             `json:"metadata"`
    Items []PolicyDefinition    `json:"items"`
}
