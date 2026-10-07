// SPDX-License-Identifier: GPL-3.0-or-later

// Package smartphone rewrites (normalizes) and minifies the History Tracers
// smartphone lesson content stored in src/smartphone/<lang>/.
//
// The package is independent from the command line tool: any History Tracers
// project can import "historytracers-publisher/smartphone" and call
// HTMinifyAllFiles, HTValidateSMGameFormats or HTTransformSMGame directly.
package smartphone
