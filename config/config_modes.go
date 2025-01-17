// Copyright (c) OpenFaaS Author(s) 2021. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package config

const (
	// ModeStreaming streams the values live to the caller as they are printed by the process.
	ModeStreaming = 1

	// ModeSerializing reads all the response and buffers before returning
	ModeSerializing = 2

	// ModeAfterBurn for performance tuning
	ModeAfterBurn = 3

	// ModeHTTP for routing requests over HTTP
	ModeHTTP = 4

	// ModeStatic for serving static content
	ModeStatic = 5

	// ModeWasm for WebAssembly
	ModeWasm = 6
)

// WatchdogModeConst as a const int
func WatchdogModeConst(mode string) int {
	switch mode {
	case "streaming":
		return ModeStreaming
	case "afterburn":
		return ModeAfterBurn
	case "serializing":
		return ModeSerializing
	case "http":
		return ModeHTTP
	case "static":
		return ModeStatic
	case "wasm":
		return ModeWasm
	default:
		return 0
	}
}

// WatchdogMode as a string
func WatchdogMode(mode int) string {
	switch mode {
	case ModeStreaming:
		return "streaming"
	case ModeAfterBurn:
		return "afterburn"
	case ModeSerializing:
		return "serializing"
	case ModeHTTP:
		return "http"
	case ModeStatic:
		return "static"
	case ModeWasm:
		return "wasm"
	default:
		return "unknown"
	}
}
