package gen

import (
	"regexp"
	"strings"
)

var (
	importList = new(importPkgS).Add(
		"context",
		"database/sql",
		"strings",
		"",
		"gorm.io/gorm",
		"gorm.io/gorm/schema",
		"gorm.io/gorm/clause",
		"",
		"gen \"github.com/dieagenturverwaltung/gorm-gen\"",
		"github.com/dieagenturverwaltung/gorm-gen/field",
		"github.com/dieagenturverwaltung/gorm-gen/helper",
		"",
		"gorm.io/plugin/dbresolver",
	)
	unitTestImportList = new(importPkgS).Add(
		"context",
		"fmt",
		"strconv",
		"testing",
		"",
		"gorm.io/driver/sqlite",
		"gorm.io/gorm",
	)
)

type importPkgS struct {
	paths []string
}

func (ip importPkgS) Add(paths ...string) *importPkgS {
	purePaths := make([]string, 0, len(paths)+1)
	for _, p := range paths {
		p = strings.TrimSpace(p)
		if p == "" {
			purePaths = append(purePaths, p)
			continue
		}

		if p[len(p)-1] != '"' {
			p = `"` + p + `"`
		}

		reg := regexp.MustCompile(`^.*"(.*)"$`)

		var exists bool
		for _, existsP := range ip.paths {
			submatch := reg.FindStringSubmatch(existsP)
			if len(submatch) < 2 {
				continue
			}

			if submatch[1] == p {
				exists = true
				break
			}

			if p == existsP {
				exists = true
				break
			}
		}
		if !exists {
			purePaths = append(purePaths, p)
		}
	}
	purePaths = append(purePaths, "")

	ip.paths = append(ip.paths, purePaths...)

	return &ip
}

func (ip importPkgS) Paths() []string { return ip.paths }
