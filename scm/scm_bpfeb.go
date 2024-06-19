// Code generated by bpf2go; DO NOT EDIT.
//go:build mips || mips64 || ppc64 || s390x

package scm

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadScm returns the embedded CollectionSpec for scm.
func loadScm() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ScmBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load scm: %w", err)
	}

	return spec, err
}

// loadScmObjects loads scm and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*scmObjects
//	*scmPrograms
//	*scmMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadScmObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadScm()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// scmSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type scmSpecs struct {
	scmProgramSpecs
	scmMapSpecs
}

// scmSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type scmProgramSpecs struct {
	BpfProg *ebpf.ProgramSpec `ebpf:"bpf_prog"`
}

// scmMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type scmMapSpecs struct {
	SysCalls *ebpf.MapSpec `ebpf:"sys_calls"`
}

// scmObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadScmObjects or ebpf.CollectionSpec.LoadAndAssign.
type scmObjects struct {
	scmPrograms
	scmMaps
}

func (o *scmObjects) Close() error {
	return _ScmClose(
		&o.scmPrograms,
		&o.scmMaps,
	)
}

// scmMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadScmObjects or ebpf.CollectionSpec.LoadAndAssign.
type scmMaps struct {
	SysCalls *ebpf.Map `ebpf:"sys_calls"`
}

func (m *scmMaps) Close() error {
	return _ScmClose(
		m.SysCalls,
	)
}

// scmPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadScmObjects or ebpf.CollectionSpec.LoadAndAssign.
type scmPrograms struct {
	BpfProg *ebpf.Program `ebpf:"bpf_prog"`
}

func (p *scmPrograms) Close() error {
	return _ScmClose(
		p.BpfProg,
	)
}

func _ScmClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed scm_bpfeb.o
var _ScmBytes []byte
