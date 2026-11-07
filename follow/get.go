package follow

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/helper"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

func makeGitURL(rawURL string) string {
	rawURL = strings.TrimPrefix(rawURL, "https://")
	rawURL = strings.TrimPrefix(rawURL, "http://")
	rawURL = strings.TrimSuffix(rawURL, ".git")

	if helper.ForceGitHTTPS() {
		return "https://" + rawURL
	}

	host, path, _ := strings.Cut(rawURL, "/")
	return "git@" + host + ":" + path + ".git"
}

func Get(url string, branch string) (Config, error) {
	// connect
	nameSum := md5.Sum([]byte(url + branch))
	name := hex.EncodeToString(nameSum[:])
	url = makeGitURL(url)

	remote := git.NewRemote(nil, &config.RemoteConfig{
		Name: name,
		URLs: []string{url},
	})

	refs, err := remote.List(&git.ListOptions{})
	if err != nil {
		return Config{}, fmt.Errorf("get remote: %w", err)
	}

	// get hash
	hash := ""
	for _, ref := range refs {
		if ref.Name().String() == plumbing.NewBranchReferenceName(branch).String() {
			hash = ref.Hash().String()
			break
		}
	}

	if hash == "" {
		return Config{}, fmt.Errorf("branch %s does not exist", branch)
	}

	// read repo
	data, err := cache.Read(filepath.Join("follow", name), md5.Sum([]byte(hash)), func() ([]byte, error) {
		dir, _ := os.MkdirTemp("", "repo-*")
		defer os.RemoveAll(dir)

		// clone
		helper.Vprintf("follow: cloning %s:%s\n", url, branch)
		_, err = git.PlainClone(dir, false, &git.CloneOptions{URL: url})
		if err != nil {
			return []byte{}, fmt.Errorf("failed to clone: %w", err)
		}

		// zip
		helper.Vprintln("follow: zipping repository")
		data, err := zipPath(dir)
		if err != nil {
			return []byte{}, err
		}

		return data, nil
	})
	if err != nil {
		return Config{}, err
	}

	// read fmod.json
	r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return Config{}, err
	}

	for _, file := range r.File {
		fmt.Println(fmt.Sprintf("file: %s", file.Name))
	}

	return Config{}, nil
}

func zipPath(dir string) ([]byte, error) {
	zipPath := filepath.Join(dir, "repository.zip")

	f, err := os.Create(zipPath)
	if err != nil {
		return nil, err
	}

	zw := zip.NewWriter(f)

	err = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		// Don't include the ZIP itself.
		if path == zipPath {
			return nil
		}

		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		w, err := zw.Create(rel)
		if err != nil {
			return err
		}

		src, err := os.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()

		_, err = io.Copy(w, src)
		return err
	})

	if err != nil {
		zw.Close()
		f.Close()
		return nil, fmt.Errorf("while zipping: %w", err)
	}

	if err := zw.Close(); err != nil {
		f.Close()
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	// read zip as cache result
	data, err := os.ReadFile(zipPath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
