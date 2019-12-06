package kustomizer

// RFC6902 Json Patches

type PatchOp string

const (
	OpAdd     PatchOp = "add"
	OpReplace PatchOp = "replace"
	OpRemove  PatchOp = "remove"
)

type Patches []Patch

type Patch struct {
	Op    PatchOp     `json:"op" yaml:"op"`
	Path  string      `json:"path" yaml:"path"`
	Value interface{} `json:"value" yaml:"value"`
}
