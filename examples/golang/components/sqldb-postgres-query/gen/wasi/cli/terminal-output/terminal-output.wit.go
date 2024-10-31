// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package terminaloutput represents the imported interface "wasi:cli/terminal-output@0.2.0".
package terminaloutput

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// TerminalOutput represents the imported resource "wasi:cli/terminal-output@0.2.0#terminal-output".
//
//	resource terminal-output
type TerminalOutput cm.Resource

// ResourceDrop represents the imported resource-drop for resource "terminal-output".
//
// Drops a resource handle.
//
//go:nosplit
func (self TerminalOutput) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TerminalOutputResourceDrop((uint32)(self0))
	return
}