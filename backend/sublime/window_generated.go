// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

// This file was generated as part of a build step and shouldn't be manually modified
package sublime

import (
	"fmt"
	"github.com/limetext/gopy/lib"
	"github.com/limetext/lime/backend"
	"github.com/quarnster/util/text"
)

var (
	_ = backend.View{}
	_ = text.Region{}
	_ = fmt.Errorf
)

var _windowClass = py.Class{
	Name:    "sublime.Window",
	Pointer: (*Window)(nil),
}

type Window struct {
	py.BaseObject
	data *backend.Window
}

func (o *Window) PyInit(args *py.Tuple, kwds *py.Dict) error {
	return fmt.Errorf("Can't initialize type Window")
}
func (o *Window) Py_active_view() (py.Object, error) {
	ret0 := o.data.ActiveView()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *Window) Py_id() (py.Object, error) {
	ret0 := o.data.Id()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *Window) Py_new_file() (py.Object, error) {
	ret0 := o.data.NewFile()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *Window) Py_settings() (py.Object, error) {
	ret0 := o.data.Settings()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *Window) Py_views() (py.Object, error) {
	ret0 := o.data.Views()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}
