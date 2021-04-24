# Pyramid

**[Code](../../cmd/examples/pyramid/main.go)**
```golang
package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/assetloaderv2"
	slab2 "github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slab/slabv2"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
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

	xSize := 20
	ySize := 20
	zSize := 20

	for k := zSize; k > 0; k-- {
		for i := xSize-k; i > k; i-- {
			for j := ySize-k; j > k; j-- {
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

	base64, err := encoder.Encode(&slab2.Aggregator{
		SlabV2: slab,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}
```

**Base64**
```bash
H4sIAAAAAAAE/zTWv2ojVxvH8SPrzxxJtnG7nZoNqnMBZopADLmAuFSRC0gghRrDtIYULoKZVAHdgFUKXKZRtUxgi5AunZqA6qRZPuc5qj5orfN8f17eN8mn/z79eZVGKaW73R+j/3/9/Ztfvv3t/u/7n776eZJSn3PTPYfpKaWhic/0eTNvZ+k1FX2vfH6qn3/0Ln7e57b8nH5OPx+a+Dm9p5/Tz3P9Of2cfk4/Xy+2026XinaUz8/181N4+V76wa74/ma+Ld+n79P36fv0/T7H9+k+fZ++T9+n7w9NfJ++T9+n79P36fu5fp++T9+n79P36fvtLO7T9+n79H36Pn3/vNhP0lsq+nsqn1/r5+fQO3/uHS/vuu/9/cX79WJf3tN7ek/v6T29p/ebebynPr2n9/Se3tN7et/neE/v6T29p/f0nt7T+6GJ9/Se3tN7ek/v6T29z/U9vaf39J7e03t6T+/bWfTpPb2n9/Se3tN7er+dxnt6T+/pPb2n9/Se3h+Wp3H3nor+d1A+7+rn19Adf+4O3aE7vNxJ3/nfS9w7L07lHt2je3SP7tE9ukf36N663qN9dI/u0T26R/foHt2je5t57KN7dI/u0T26R/foHt2je32Oe3SP7tE9ukf36B7do3t0b2jiHt2je3SP7tE9ukf36B7dy/Ue3aN7dI/u0T26R/foHt1rZ7GP7tE9ukf36B7do3t0j+5tp3GP7tE9ukf36B7do3t0j+7tJ3GP7tE9ukf36B7do3t0j+5116txOqai/5+Uz2/18y5015+7S3fpLt2lu7zc7e79/yvuH5arcp/u0326T/fpPt2n+3Sf7tP98yLu0366T/fpPt2n+3Sf7tN9uk/31/U+3af7dJ/u0326T/fpPt2n+3R/M4/9dJ/u0326T/fpPt2n+3Sf7tP9Psd9uk/36T7dp/t0n+7TfbpP9+n+0MR9uk/36T7dp/t0n+7TfbpP9+l+rvfpPt2n+3Sf7tN9uk/36T7dp/vtLPbTfbpP9+k+3af7dJ/u0326T/e307hP9+k+3af7dJ/u0326T/fpPt3fT+I+3af7dJ/u0326T/fpPt2n+3T/NI77dJ/u0326T/fpPt2n+3Sf7tP9h5vHq+5zKvrnUPn8Xj+/hTr+XIc61KEOdahDHV466Wv/PIted/1YetSjHvWoRz3qUY961KMe9ah3WEaPfj/qUY961KMe9ahHPepRj3rUOy+iRz3qUY961KMe9ahHPepRj3rUW9ce9ahHPepRj3rUox71qEc96lFvM4/fj3rUox71qEc96lGPetSjHvWo1+foUY961KMe9ahHPepRj3rUox71hiZ61KMe9ahHPepRj3rUox71qEe9XHvUox71qEc96lGPetSjHvWoR712Fr8f9ahHPepRj3rUox71qEc96lFvO40e9ahHPepRj3rUox71qEc96lFvP4ke9ahHPepRj3rUox71qEc96lHvNI4e9ahHPepRj3rUox71qEc96lFvVXvUox71qEc96lGPetSjHvWoR72725dR+icV/XuifD7Wz++hrj/XpS51qUtd6lKXutTlpdt99O+f6D/cvJQ+9alPfepTn/rUpz71qU996lOf+t119On3pz71qU996lOf+tSnPvWpT33qU/+wjD71qU996lOf+tSnPvWpT33qU5/61D8vok996lOf+tSnPvWpT33qU5/61Kc+9de1T33qU5/61Kc+9alPfepTn/rUpz71N/P4/alPfepTn/rUpz71qU996lOf+tSnfp+jT33qU5/61Kc+9alPfepTn/rUpz71hyb61Kc+9alPfepTn/rUpz71qU996lM/1z71qU996lOf+tSnPvWpT33qU5/61G9n8ftTn/rUpz71qU996lOf+tSnPvWpT/3tNPrUpz71qU996lOf+tSnPvWpT33qU38/iT71qU996lOf+tSnPvWpT33qU5/61D+No0996lOf+tSnPvWpT33qU5/61Kc+9Ve1T33qU5/61Kc+9alPfepTn/rUpz71H6/i96c+9alPfepTn/rUpz71qU996lOf+n/dHlP3byr674Dy+XP9fAzt8Od20A7aQTtoB+2gHbSDdtAO2sHLjvTBf2/EnrvbY9lDe2gP7aE9tIf20B7aQ3toD+2hPbSH9tAe2vNwE3vo74f20B7aQ3toD+2hPbSH9tAe2kN7aA/toT20p7uOPbSH9tAe2kN7aA/toT20h/bQHtpDe2gP7aE9tOewjD20h/bQHtpDe2gP7aE9tIf20B7aQ3toD+2hPbTnvIg9tIf20B7aQ3toD+2hPbSH9tAe2kN7aA/toT20Z1330B7aQ3toD+2hPbSH9tAe2kN7aA/toT20h/bQns08/n5oD+2hPbSH9tAe2kN7aA/toT20h/bQHtpDe2hPn2MP7aE9tIf20B7aQ3toD+2hPbSH9tAe2kN7aA/tGZrYQ3toD+2hPbSH9tAe2kN7aA/toT20h/bQHtpDe3LdQ3toD+2hPbSH9tAe2kN7aA/toT20h/bQHtpDe9pZ/P3QHtpDe2gP7aE9tIf20B7aQ3toD+2hPbSH9tCe7TT20B7aQ3toD+2hPbSH9tAe2kN7aA/toT20h/bQnv0k9tAe2kN7aA/toT20h/bQHtpDe2gP7aE9tIf20J7TOPbQHtpDe2gP7aE9tIf20B7aQ3toD+2hPbSH9tCeVd1De2gP7aE9tIf20B7aQ3toD+2hPbSH9tAe2kN7Hq/i74f20B7aQ3toD+2hPbSH9tAe2kN7aA/toT20h/a8jGIP7aE9tIf20B7aQ3toD+2hPbSH9tAe2kN72H1M6WV0TOlDSil9CQAA///gy5r9wCMAAA==
```

**Printscreen**
![version_size](../../docs/images/pyramid.png)