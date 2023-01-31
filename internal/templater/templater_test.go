package templater_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cqroot/ceres/internal/templater"
)

func TestExecuteTemplate(t *testing.T) {
	input := "./testdata/test.txt.tmpl"
	output := "./testdata/test.txt"

	err := templater.ExecuteTemplate(
		input,
		output,
		map[string]string{"message": "test message"},
	)
	require.Nil(t, err)

	require.FileExists(t, "./testdata/test.txt", nil)

	file, err := os.Open(output)
	require.Nil(t, err, nil)
	defer file.Close()

	data, err := io.ReadAll(file)
	require.Nil(t, err, nil)

	require.Equal(t, "text: test message\n", string(data))

	testFile(t, "text: test message\n", "./testdata/test.txt")
}

func TestTemplaterExecute(t *testing.T) {
	tmpl := templater.New(
		"./testdata/template",
		"./testdata/output",
		map[string]string{
			"message": "test message",
		},
		nil, nil,
	)
	tmpl.SetVerbose(true)
	err := tmpl.Execute()
	require.Nil(t, err, nil)

	require.FileExists(t, "./testdata/output/test_1.txt", nil)
	require.FileExists(t, "./testdata/output/test_2.txt", nil)
	require.FileExists(t, "./testdata/output/subdir/test_1.txt", nil)
	require.FileExists(t, "./testdata/output/subdir/test_2.txt", nil)
	require.FileExists(t, "./testdata/output/subdir/subsubdir/test_1.txt", nil)
	require.FileExists(t, "./testdata/output/subdir/subsubdir/test_2.txt", nil)

	testFile(t, "text: test message\n", "./testdata/output/test_1.txt")
	testFile(t, "text: test message\n", "./testdata/output/test_2.txt")
	testFile(t, "text: test message\n", "./testdata/output/subdir/test_1.txt")
	testFile(t, "text: test message\n", "./testdata/output/subdir/test_2.txt")
	testFile(t, "text: test message\n", "./testdata/output/subdir/subsubdir/test_1.txt")
	testFile(t, "text: test message\n", "./testdata/output/subdir/subsubdir/test_2.txt")
}

func testFile(t *testing.T, content string, path string) {
	file, err := os.Open(path)
	require.Nil(t, err, nil)
	defer file.Close()

	data, err := io.ReadAll(file)
	require.Nil(t, err, nil)

	require.Equal(t, content, string(data))
}
