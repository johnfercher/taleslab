# TaleSlab
TaleSpire Slab Creator

## Examples

**Code (Random Cube)**
```golang
package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/assetloaderv2"
	slab2 "github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slab/slabv2"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
	"math/rand"
)

func main() {
	loader := assetloaderv2.NewAssetLoaderV2()

	constructors, err := loader.GetConstructors()
	if err != nil {
		log.Fatalln(err)
	}

	builder := slabdecoder.NewSlabEncoderBuilder()
	encoder := builder.Build()

	slab := &slabv2.Slab{
		MagicBytes:  slab2.MagicBytes,
		Version:     2,
		AssetsCount: 1,
		Assets: []*slabv2.Asset{
			{
				Id: constructors[0].Id,
			},
		},
	}

	xSize := 10
	ySize := 10
	zSize := 10

	for i := xSize; i > 0; i-- {
		for j := ySize; j > 0; j-- {
			for k := zSize; k > 0; k-- {
				if rand.Int()%2 == 0 {
					layout := &slabv2.Bounds{
						Coordinates: &slabv2.Vector3d{
							X: int16(slabv2.GainX * i),
							Y: int16(slabv2.GainY * j),
							Z: int16(slabv2.GainZ * k),
						},
						Rotation: 0,
					}

					slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layout)
					slab.Assets[0].LayoutsCount++
				}
			}
		}
	}

	base64, err := encoder.Encode(&slab2.Aggregator{
		SlabV2: slab,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}
```

**Base64 (Random Cube)**
```bash
H4sIAAAAAAAE/zSXvc3rRheEl79ioOBLlTHRV4AKEBgYcKACxALeAq4zJhd4C3DAwIkjA66APThh4gs5dQWswU6MZx7daDDc2bO7c47Orr798+2vulSllP/9/kf17y+//fDzj7/e/77/9P8/q1Jep9epfAWPRhzfONfytRJfp88v6KZeXDpxbMS1EodT+UC3teLYiHMtrpX4On0+0a2VuBdxOJUH36de3FpxrsW1Epfu845urcS9iK9TufF96sWlE7dW3Is49Z9XdEsn7kV8ncqF71MvHo24dJ9nvm+teDTi+Ma5lq8VOLDe11IGzh/Ux4Fzf+G7/jEu17eBdT8Y1z/G5fo3ECfjexGXDt8G9vNk3tiIcy2+4udA/Ifjon4mXr7rZ+Lc0c21qL/sG98S58b40YhrJepn5l8ZHxtxrUT9TZwL40N8jk/h+pzzhs+1+Drp5/DGqZebD+bLzcfAvs/E/56HvcCnfjiRh6nfWtF6n3rreuqt66k/ko+pt66n3rqe+r2QJ/T4PvVLJx7JD3q59Y0O/1lXnHrRfGWdZzFu0DyhJx+J/2B8S/2jl1v/7Ec+9eQJvTg2onlDJx+Sv6mfa/LFOuQl866uI5o39o//6MS1Eq1/9qOv+rx0U4+vS7d0ov1j6fYiH9I/0OEjOlG/l26u5fq9dK/4jB5fl06fl85+gt7v9pGF+M9i3ODWikd+D+jl9pnEf6DX38QP3+Mr4/jG+qK+ch65/YVz4Gfm34g35nfBenJ/H8yTv+J79ntF7++GOPIteYDjd857Qaf/2c8Zbr9Z8CvcOt9a63xrre+ttZ9vrf5m/KM4HtRPxuX6iQ6/tta6JZ7c/rK19u2ttW7R4x96celE/WW+fHyjdby1+s18fEUn6jf7klvP6OX6v7HOjfPY35kvNw/Mlw/xnf3h89bqM+Ny+1P8uhBvL6L9J/s7893+w7py+3/8yLh5OBrv1aPx93A09p2jMS9H4+8CHXV/NP4+GJfbbxgnL4yL5uto7DOJ81ESn/xk/hNunWf+A+69ejTeq9lPvh/Jx0HccO+FxA9/JS+sj+/MF4/kCZ3cezZxbqxn3Wd/4fYVOH4TD3+JJ26tODai9Y4OX4kr2m/Gt78j5/paysi5gvo8EifcPjS+/UaPv+hF72H0cn8n6PB15FwfxPc+HvEr3DyMxA/3d8M4/oP4R1xRH7PunXjet8yXW9+M41fWvaGzrhMvXF85Pz6yT9F3DDr5nLrm3HLv28S/EHdJfxmJH+79OxI/3PuWfeB74p6ZZ71n3+dS5vqVe3WurfO51v+5ts7n2v4z1+YBPT7PtfU+1967zPe79wH67xx/E+eD9XwHzbX1P9e+gxLvybj9Hr3c/sQ4ecj+HujsQ6wjtw+hIx+Zf0fnO4j15N6j7Jt8JN4Nne+hnO8Kt68zLvddSXz8zfwLOvNAfLl9J3Hir/0m/oR7/+Kv3HzgA3yt7P9r5X28VuZlrXx/rpV9Z61896DD57XyvbNW1j86v/vOWSvvX76Th7XyfoXjM3HwMes9SsZF+wvj+Me4aP9gX3Lf7+jwk32I3pd8x8eseyW+/YRz4Bvn9vz6lXXOpezFetyLdbgX63Av1t9erL+92Af2Yv2h4/x78T2yF/sx8zg/OtE6jP7Jet6L0Yfbd7OPB+P6kf2EW3fZx51x+wLjcutvL/6vQYcvWf+G3nuO9eX+v0GPT+xXtD8k7pV59oecL9z7j3n4mXUu6kTrknF9tT+gk1uXiR/ffZ/sZS+f51JK+S8AAP//pjtcFogOAAA=
```

**Printscreen (Random Cube)**
![version_size](./docs/images/randomcube.png)

More examples [here]().

## Slab Versions

TaleSpire have two versions of slab code. The Version 1 is the main supported one, which have some 
projects capable to serialize/deserialize it. The Version 2 is the new one, which there is no other
projects (besides this, yet) capable to serialize/deserialize it. TaleSpire is capable to
work with both versions, but when a Version 1 code is pasted into the game, TaleSpire converts the code into a Version 2.

### Version 1

The Version 1 serialization/deserialization is based on 
[Mercer01/talespireDeserialize](https://github.com/Mercer01/talespireDeserialize), 
[brcoding/TaleSpireHtmlSlabGeneration](https://github.com/brcoding/TaleSpireHtmlSlabGeneration) 
and [https://github.com/creadth/Creadth.Talespire.DungeonGenerator](https://github.com/creadth/Creadth.Talespire.DungeonGenerator)

[Version 1: Documentation](docs/versions/version1.md)

### Version 2

The Version 2 serialization/deserialization was developed based on the Version 1, the first part
of the ByteArray is almost the same, but the last objects are different.

[Version 2: Documentation](docs//versions/version2.md)