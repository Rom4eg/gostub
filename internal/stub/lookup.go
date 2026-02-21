package stub

import (
	"os"
	"path/filepath"
	"strings"
)

func (s *Service) Lookup(path string) (files []string, err error) {
	fp := filepath.Join(s.root, path)
	stats, err := os.Stat(fp)
	if err != nil {
		return []string{}, err
	}

	if stats.IsDir() {
		idx := filepath.Join(path, "index")
		return s.Lookup(idx)
	}

	files = append(files, fp)
	baseDir := filepath.Dir(fp)
	res, err := s._lookupResources(baseDir)
	if err != nil {
		return []string{}, err
	}

	files = append(files, res...)
	return files, nil
}

func (s *Service) _lookupResources(dir string) (files []string, err error) {
	stat, err := os.Stat(dir)
	if err != nil {
		return []string{}, err
	}

	if !stat.IsDir() {
		return []string{}, ErrDirectoryExpected
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}

	for _, e := range entries {
		if e.IsDir() || !strings.HasPrefix(e.Name(), "_") {
			continue
		}

		n := filepath.Join(dir, e.Name())
		files = append(files, n)
	}

	return files, nil
}
