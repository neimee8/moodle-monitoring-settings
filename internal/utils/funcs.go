package utils

import "os"

func AtomicWrite(
	path, tmpPath string,
	data []byte,
	perm os.FileMode,
) error {
	err := os.WriteFile(tmpPath, data, perm)

	if err != nil {
		return err
	}

	return os.Rename(tmpPath, path)
}
