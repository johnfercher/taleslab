# TaleSlab
TaleSpire Slab Creator

This project serialize/deserialize Slabs from a Base64 string into a Go Structure, i.e, make this:
```
H4sIAAAAAAAACzv369xFJgYmBgaG7pL+4F2SeZ7t9wUSuyZLlDACxXYovn/klGjqOCNyUuOKTWF/QOr6gRIBgg4sQCZDAKsDkzQziOUAlAIAoQYiAEwAAAA=
```

Into this:
```json
{
  "magic_hex": "zvrO0Q==",
  "version": 2,
  "assets_count": 2,
  "assets": [{
    "id": "AACLdI9TuhluSYffEGGKkxh0",
    "layouts_count": 1,
    "layouts": [{
      "coordinates": {
        "x": 399,
        "y": 4432,
        "z": 0
      },
      "rotation": 1088
    }]
  }, {
    "id": "AAC4Ie/iQmE1QZhZkoGoslb8",
    "layouts_count": 2,
    "layouts": [{
      "coordinates": {
        "x": 0,
        "y": 1360,
        "z": 0
      },
      "rotation": 576
    }, {
      "coordinates": {
        "x": 795,
        "y": 0,
        "z": 0
      },
      "rotation": 576
    }]
  }]
}
```

And enable you to generate/modify this:
![version_size](./docs/images/version2photo.png)

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

[Version 1: Documentation](docs/version1.md)

### Version 2

The Version 2 serialization/deserialization was developed based on the Version 1, the first part
of the ByteArray is almost the same, but the last objects are different.

[Version 2: Documentation](docs/version2.md)