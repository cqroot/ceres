/*
Copyright (C) 2025 Keith Chu <cqroot@outlook.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package tmpl_test

import (
	"os"
	"testing"

	"github.com/cqroot/ceres/pkg/tmpl"
	"github.com/stretchr/testify/require"
)

func PrepareTestData(t *testing.T, content string) error {
	t.Log("Delete the old testdata directory.")
	err := os.RemoveAll("./testdata")
	if err != nil {
		return err
	}

	t.Log("Create the new testdata directory.")
	err = os.MkdirAll("./testdata", 0o777)
	if err != nil {
		return err
	}

	t.Log("Create file ./testdata/test.tmpl.")
	err = os.WriteFile("./testdata/test.tmpl", []byte(content), 0o666)
	if err != nil {
		return err
	}

	return nil
}

func TestExecute(t *testing.T) {
	require.Nil(t, PrepareTestData(t, `ProjectName: {{ .ProjectName }}
Author: {{ .Author }}
StringToUpper: {{ ToUpper .TestString }}
StringToLower: {{ ToLower .TestString }}
`))

	require.Nil(t, tmpl.Execute("./testdata/test.tmpl", "./testdata/test.txt", map[string]any{
		"ProjectName": "ceres",
		"Author":      "cqroot",
		"TestString":  "TestString",
	}))

	content, err := os.ReadFile("./testdata/test.txt")
	require.Nil(t, err)

	require.Equal(t, string(content), `ProjectName: ceres
Author: cqroot
StringToUpper: TESTSTRING
StringToLower: teststring
`)
}
