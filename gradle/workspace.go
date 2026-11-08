package gradle

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/JXSnack/flatmod/follow"
	"github.com/JXSnack/flatmod/modfile"
)

type Workspace struct {
	Follow follow.Info
	File   *modfile.File
}

func NewWorkspace(file *modfile.File) (*Workspace, error) {
	cfg, err := follow.GetConfig(file.Follow, file.FollowBranch)
	if err != nil {
		return nil, err
	}

	return &Workspace{
		File: file,
		Follow: follow.Info{
			Config: cfg,
			URL:    file.Follow,
			Branch: file.FollowBranch,
		},
	}, nil
}

func (w *Workspace) Dir() string {
	return filepath.Join(".flatmod", "generated", w.Follow.Checksum())
}

func (w *Workspace) Generate() error {
	dir := w.Dir()

	// make directory
	err := os.MkdirAll(w.Dir(), 0700)
	if err != nil {
		return fmt.Errorf("workspace: %w", err)
	}

	// get zip
	data, err := follow.Get(w.Follow.URL, w.Follow.Branch)
	if err != nil {
		return err
	}

	zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return fmt.Errorf("workspace: %w", err)
	}

	// unpack zip
	for _, zf := range zr.File {
		r, err := zf.Open()
		if err != nil {
			return fmt.Errorf("workspace: %w", err)
		}

		path := filepath.Join(dir, zf.Name)

		if zf.FileInfo().IsDir() {
			r.Close()
			if err := os.MkdirAll(path, zf.Mode()); err != nil {
				return fmt.Errorf("workspace: %w", err)
			}
		} else {
			if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				r.Close()
				return fmt.Errorf("workspace: %w", err)
			}

			data, err := io.ReadAll(r)
			r.Close()
			if err != nil {
				return fmt.Errorf("workspace: %w", err)
			}

			content := modfile.FillProperties(w.File, string(data))

			if err := os.WriteFile(path, []byte(content), zf.Mode()); err != nil {
				return fmt.Errorf("workspace: %w", err)
			}
		}
	}

	// move files
	err = moveFiles(w, dir)
	if err != nil {
		return err
	}

	return nil
}
