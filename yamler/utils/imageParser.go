package utils

import (
	"regexp"
	"strings"

	"golang.org/x/xerrors"
)

// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0
// https://github.com/kubernetes-sigs/kustomize/blob/02f9b98b5ac562e07c28041835b86503fda0bfdb/kustomize/internal/commands/edit/set/setimage.go

var (
	ErrImageInvalidArgs = xerrors.New(`invalid format of image, use one of the following options:
- <image>=<newimage>:<newtag>
- <image>=<newimage>@<newtag>
- <image>=<newimage>
- <image>:<newtag>
- <image>@<digest>`)
	pattern = regexp.MustCompile("^(.*):([a-zA-Z0-9._-]*)$")
)

type Overwrite struct {
	Name   string
	Digest string
	Tag    string
}

// ParseOverwrite parses the overwrite parameters
// from the given arg into a struct
func ParseOverwrite(arg string, overwriteImage bool) (*Overwrite, error) {
	// match <image>@<digest>
	if d := strings.Split(arg, "@"); len(d) > 1 {
		return &Overwrite{
			Name:   d[0],
			Digest: d[1],
		}, nil
	}

	// match <image>:<tag>
	if t := pattern.FindStringSubmatch(arg); len(t) == 3 {
		return &Overwrite{
			Name: t[1],
			Tag:  t[2],
		}, nil
	}

	// match <image>
	if len(arg) > 0 && overwriteImage {
		return &Overwrite{
			Name: arg,
		}, nil
	}
	return &Overwrite{}, ErrImageInvalidArgs
}
