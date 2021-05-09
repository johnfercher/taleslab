# Fire Spiral

**Printscreen**
![version_size](../../docs/images/firespiral.png)

**Base64**
```bash
H4sIAAAAAAAE/6TWoY7bThDH8XHizd1Jf/RXSEF1JkUNqE4FRa1BS3ohRTWpdCCvUMkwILSSpSN5DD9AQFhTck15HsCgMKio+u6udK6a2jMO+mi985PjeHfWD78efowkEZGX2zff77/+fH//bv3iy7ent59HIvlERG66zV+Fea3L16Fe6/ZtqMeDE7VyK77eav4h5LB0Ynb5UXyu7dSJHw9x+0mkTiW6MzmP9RplIb4Om/FukKuY05rFeq2bUfhdWotYb/WYhPu05T0w/tMqCeNgNdBZzCHrxupOKp87JeuY6xoXUvm6tuwfxlbZr+S6FKn8fFv6AmOt9B3qg8WoT/qZSGF2+1x8TuvyWajXml+HelzwHNdswaJ3LE9CndXt/yGHOynMLv8Tn2s7Swo/fjQbMz9TyFlTJdkgjzGnkX+ZOixG2SA3Mac1G4f7aF3FeqtNzPXb+Oeep3/Le+B6l3Xa+HmrUxdyyLqwWrrG507JOua6xYNrfD2yf4bIPiXXZT5p/Hxb+gLjYJ32Sd/JJ7VZ+h05rfRP6rXSp6m3yrlwcLVazh/qh8r5V7paaelC/aOcy6UrvdMz5HuhTsuznMd8l3zHMY/NuBzkKua0ZrFe62YUfle/+YTn2EQLo8ck5NvyHhh3WSW5n7c6izlk3VjdSe5zp2Rdct3iQnJfj+yfIbJfyQUvL04p8Xpb+gJjrfQd6q3S38hppZ9Sr5X+TL1VzgVyuJDLXjmHqLO5j/X7C86/3RlyLpPHWbIfLN8PVbI/y2PMawzfc/tBbmJOazYO9/m360t+TxZdDbSJOa3zNNy3Le+BcZd1uvbzVqcu5JD1YrV0a587JeuY6xYPbu3rkf3z6N1VGPfLPj24u07zSZhvS19grJW+Q71V+hs5uRHJJ3dXciMi8jsAAP//xWz89MgTAAA=
```

**[Code](../../cmd/examples/firespiral/main.go)**
```golang
package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/internal/talespireadapter/talespirecoder"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"log"
	"math"
)

func main() {
	propRepository, err := taleslabrepositories.NewPropRepository()
	if err != nil {
		log.Fatal(err.Error())
	}

	compressor := bytecompressor.New()
	encoder := talespirecoder.NewEncoder(compressor)

	assets := taleslabentities.Assets{}

	asset := propRepository.GetProp("fire")

	radius := 8

	for i := 0.0; i < 4.0*3.14; i += 0.02 {
		cos := math.Cos(i)
		sin := math.Sin(i)

		xRounded := fix(float64(radius)*cos, 1)
		yRounded := fix(float64(radius)*sin, 1)

		xPositiveTranslated := radius + xRounded
		yPositiveTranslated := radius + yRounded

		asset := &taleslabentities.Asset{
			Id: asset.Parts[0].Id,
			Coordinates: &taleslabentities.Vector3d{
				X: xPositiveTranslated,
				Y: yPositiveTranslated,
				Z: int(i),
			},
			Rotation: 0,
		}

		assets = append(assets, asset)
	}

	taleSpireSlab := taleslabmappers.TaleSpireSlabFromAssets(assets)

	base64, err := encoder.Encode(taleSpireSlab)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}

func fix(value float64, fixValue int) int {
	division := value / float64(fixValue)

	divisionRounded := math.Round(division)
	top := int(divisionRounded * float64(fixValue))

	return top
}
```