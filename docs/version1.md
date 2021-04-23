# Version 1

The Version 1 of TaleSpire slabs follow the byte array scheme below.

## Gzip
The content of the base64 string is a gzip file

## Example

One Rock And Two Grass

```
H4sIAAAAAAAAAzv369xFRgYmBt6LgpbaIsb+81/FWgkcNW9kYmBgeLrn0b/gP6v99uzirSp+4e3JyAADDfYIGoZR5ByQ5TigMtcXb8Ap9yZwBlwMwgapY2AAAFC/RiOgAAAA
```

```json
{
  "magic_hex": ["CE", "FA", "CE", "D1"],
  "version": 1,
  "assets_count": 2,
  "assets": [{
    "id": "0dd11139-2b14-334f-9fea-5d3a10c53781",
    "layouts_count": 2,
    "layouts": [{
      "center": {
        "x": 0,
        "y": 1,
        "z": 0
      },
      "extents": {
        "x": 1,
        "y": 1,
        "z": 1
      }
    }, {
      "center": {
        "x": 0,
        "y": 1,
        "z": 2
      },
      "extents": {
        "x": 1,
        "y": 1,
        "z": 1
      },
      "rotation": 8
    }]
  }, {
    "id": "e5bce2fe-53fc-ab4e-bcba-0d7a73e84b49",
    "layouts_count": 1,
    "layouts": [{
      "center": {
        "x": 0,
        "y": 1.38,
        "z": 2
      },
      "extents": {
        "x": 1,
        "y": 1,
        "z": 1
      },
      "rotation": 8
    }]
  }],
  "bounds": {
    "center": {
      "x": 0,
      "y": 1.19,
      "z": 1
    },
    "extents": {
      "x": 1,
      "y": 1.19,
      "z": 2
    }
  }
}
```