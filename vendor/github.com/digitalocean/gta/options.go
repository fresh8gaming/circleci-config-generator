/*
Copyright 2016 The gta AUTHORS. All rights reserved.

Use of this source code is governed by the Apache 2 license that can be found
in the LICENSE file.
*/
package gta

// Option is an option function used to modify a GTA.
type Option func(*GTA) error

// SetDiffer sets a differ on a GTA.
func SetDiffer(d Differ) Option {
	return func(g *GTA) error {
		g.differ = d
		return nil
	}
}

// SetPackager sets a packager on a GTA.
func SetPackager(p Packager) Option {
	return func(g *GTA) error {
		g.packager = p
		return nil
	}
}

// SetPrefixes sets a list of prefix to be included
func SetPrefixes(prefixes ...string) Option {
	return func(g *GTA) error {
		g.prefixes = prefixes
		return nil
	}
}

// SetTags sets a list of build tags to consider.
func SetTags(tags ...string) Option {
	return func(g *GTA) error {
		g.tags = tags
		return nil
	}
}
