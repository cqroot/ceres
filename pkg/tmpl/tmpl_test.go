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
