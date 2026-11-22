package updater

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/JXSnack/flatmod/cache"
	"github.com/JXSnack/flatmod/helper"
)

// GetUpdate sends an hourly cached request to the specified url and unmarshals
// the return value into an Update struct.
func GetUpdate(url string) (Update, error) {
	err := cache.Init()
	if err != nil {
		return Update{}, err
	}

	data, err := cache.Read("update.json", helper.HourlySum(), func() ([]byte, error) {
		resp, err := http.Get(url)
		if err != nil {
			return []byte{}, err
		}
		defer resp.Body.Close()
		return io.ReadAll(resp.Body)
	})
	if err != nil {
		return Update{}, err
	}

	var update Update
	err = json.Unmarshal(data, &update)
	if err != nil {
		return Update{}, err
	}

	return update, nil
}

// CheckUpdates compares the result from GetUpdate with the local Current version.
// Returns the update message when there is a new version, otherwise the string is
// empty.
func CheckUpdates() (string, error) {
	update, err := GetUpdate(URL)
	if err != nil {
		return "", err
	}

	if update.Latest > Current {
		return update.Message, nil
	}
	return "", nil
}
