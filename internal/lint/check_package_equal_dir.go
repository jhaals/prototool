// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package lint

import (
	"path/filepath"
	"strings"

	"github.com/emicklei/proto"
	"github.com/uber/prototool/internal/text"
)

var packagesEqualDirLinter = NewLinter(
	"PACKAGE_EQUAL_DIR",
	"Verifies that package name matches file path.",
	checkPackageEqualDir,
)

func checkPackageEqualDir(add func(*text.Failure), dirPath string, descriptors []*proto.Proto) error {
	return runVisitor(&packageEqualDirVisitor{baseAddVisitor: newBaseAddVisitor(add)}, descriptors)
}

type packageEqualDirVisitor struct {
	baseAddVisitor

	filename string
}

func (v *packageEqualDirVisitor) OnStart(descriptor *proto.Proto) error {
	v.filename = descriptor.Filename
	return nil
}

func (v *packageEqualDirVisitor) VisitPackage(pkg *proto.Package) {
	packagePath := strings.Replace(pkg.Name, ".", "/", -1)
	if !strings.HasSuffix(filepath.Dir(v.filename), packagePath) {
		v.AddFailuref(pkg.Position, "package name does not match file path. Expect %s to be stored in %s", pkg.Name, packagePath)
	}
}
