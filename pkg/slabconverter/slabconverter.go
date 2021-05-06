package slabconverter

import (
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/mappers"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
)

type slabConverter struct {
	decoder talespirecoder.Decoder
	encoder talespirecoder.Encoder
}

func New(decoder talespirecoder.Decoder, encoder talespirecoder.Encoder) *slabConverter {
	return &slabConverter{
		decoder: decoder,
		encoder: encoder,
	}
}

func (self *slabConverter) ConvertToSlab(base64Slab, assetName, assetType string) (*assetloader.AssetInfo, error) {

	decodedSlab, err := self.decoder.Decode(base64Slab)
	if err != nil {
		return nil, err
	}

	convertedSlab := mappers.AssetInfoFromTalespireContract(assetName, assetType, decodedSlab)

	return convertedSlab, nil
}

func (self *slabConverter) ConvertFromSlab(info assetloader.AssetInfo) (string, error) {

	//encodedSlab, err := self.encoder.Encode()
	return "", nil

}
