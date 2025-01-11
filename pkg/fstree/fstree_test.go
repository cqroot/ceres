package fstree_test

import (
	"os"
	"testing"

	"github.com/cqroot/ceres/pkg/fstree"
	"github.com/stretchr/testify/require"
)

func PrepareTestData(t *testing.T) error {
	t.Log("Delete the old testdata directory.")
	err := os.RemoveAll("./testdata")
	if err != nil {
		t.Logf("Failed to delete the old testdata directory: %s\n", err)
		return err
	}

	t.Log("Create the new testdata directory.")
	err = os.MkdirAll("./testdata", 0o777)
	if err != nil {
		t.Logf("Failed to create the new testdata directory: %s\n", err)
		return err
	}

	err = os.MkdirAll("./testdata/dir", 0o777)
	if err != nil {
		t.Logf("Failed to create the dir directory: %s\n", err)
		return err
	}

	err = os.WriteFile("./testdata/dir/test0", []byte("test"), 0o666)
	if err != nil {
		t.Logf("Failed to create the test0 file: %s\n", err)
		return err
	}

	err = os.WriteFile("./testdata/test1", []byte("test"), 0o666)
	if err != nil {
		t.Logf("Failed to create the test1 file: %s\n", err)
		return err
	}

	return nil
}

func TestFileInfos(t *testing.T) {
	require.Nil(t, PrepareTestData(t))

	infos, err := fstree.FileInfos("./testdata")
	require.Nil(t, err)
	require.Equal(t, 4, len(infos))
	t.Logf("infos: %+v\n", infos)

	require.Equal(t, "testdata", infos[0].Name)
	require.Equal(t, ".", infos[0].RelPath)

	require.Equal(t, "dir", infos[1].Name)
	require.Equal(t, "dir", infos[1].RelPath)

	require.Equal(t, "test0", infos[2].Name)
	require.Equal(t, "dir/test0", infos[2].RelPath)

	require.Equal(t, "test1", infos[3].Name)
	require.Equal(t, "test1", infos[3].RelPath)
}
