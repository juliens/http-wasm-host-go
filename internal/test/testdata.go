package test

import (
	_ "embed"
)

// AuthWasm was generated by the following:
//
//	cd testdata; wat2wasm --debug-names auth.wat
//
//go:embed testdata/auth.wasm
var AuthWasm []byte

// ConfigWasm was generated by the following:
//
//	cd testdata; wat2wasm --debug-names config.wat
//
//go:embed testdata/config.wasm
var ConfigWasm []byte

// LogWasm was generated by the following:
//
//	cd testdata; wat2wasm --debug-names log.wat
//
//go:embed testdata/log.wasm
var LogWasm []byte

// RouterWasm was generated by the following:
//
//	cd testdata; wat2wasm --debug-names router.wat
//
//go:embed testdata/router.wasm
var RouterWasm []byte

// RedactWasm was generated by the following:
//
//	cd testdata; wat2wasm --debug-names redact.wat
//
//go:embed testdata/redact.wasm
var RedactWasm []byte
