package imageutil

import (
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"golang.org/x/xerrors"
)

// GetImageHash retrieves image hash from remote registry
func GetImageHash(path string) (string, error) {
	desc, err := GetImageDescriptor(path)

	if err != nil {
		return "", xerrors.Errorf("failed to get image descriptor: %w", err)
	}

	return desc.Digest.Hex, nil
}

// GetImageDescriptor retrieves image descriptor
func GetImageDescriptor(path string) (*remote.Descriptor, error) {
	opts := []remote.Option{}

	authOpt := remote.WithAuthFromKeychain(authn.DefaultKeychain)
	opts = append(opts, authOpt)

	ref, err := name.ParseReference(path)

	if err != nil {
		return nil, xerrors.Errorf("failed to parse reference: %w", err)
	}

	desc, err := remote.Get(ref, opts...)

	if err != nil {
		return nil, xerrors.Errorf("failed to get descriptor of images: %w", err)
	}

	return desc, nil
}
