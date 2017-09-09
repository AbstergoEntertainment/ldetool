/* This file was autogenerated via
 ----------------------------------------
 ldetool generate --package main rule.lde
 ----------------------------------------
do not touch it with bare hands!
*/

package main

import (
	"bytes"
)

// Line autogenerated parser
type Line struct {
	rest  []byte
	Name  []byte
	Count []byte
}

// Extract autogenerated method of Line
func (p *Line) Extract(line []byte) (bool, error) {
	var pos int
	var tmp []byte
	p.rest = line

	// Take until '|' as Name(string)
	pos = bytes.IndexByte(p.rest, '|')

	if pos >= 0 {
		tmp = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}
	p.Name = tmp

	// Looking for '|' and then pass it
	pos = bytes.IndexByte(p.rest, '|')
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Looking for '|' and then pass it
	pos = bytes.IndexByte(p.rest, '|')
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Looking for '|' and then pass it
	pos = bytes.IndexByte(p.rest, '|')
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Take until '|' as Count(string)
	pos = bytes.IndexByte(p.rest, '|')

	if pos >= 0 {
		tmp = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}
	p.Count = tmp

	return true, nil
}
