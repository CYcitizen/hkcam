package hkcam

import (
	"github.com/brutella/hap/characteristic"
)

const TypeGetAsset = "6A6C39F5-67F0-4BE1-BA9D-E56BD27C9606"

// GetAsset is used to get the raw data of an asset.
// After writing a valid JSON to this characteristic,
// the characteristic value will be the raw data of the requested asset.
// A valid JSON looks like this. `{"id":"1.jpg","width":320,"height":240}`
type GetAsset struct {
	*characteristic.Bytes
}

func NewGetAsset() *GetAsset {
	b := characteristic.NewBytes(TypeGetAsset)
	b.Permissions = []string{characteristic.PermissionRead, characteristic.PermissionWrite}
	b.SetValue([]byte{})

	return &GetAsset{b}
}

type GetAssetRequest struct {
	ID     string `json:"id"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}
