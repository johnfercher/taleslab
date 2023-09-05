# Slab Versions

TaleSpire have two versions of slab code. The Version 1 is the main supported one, which have some
projects capable to serialize/deserialize it. The Version 2 is the new one, which there is no other
projects (besides this, yet) capable to serialize/deserialize it. TaleSpire is capable to
work with both versions, but when a Version 1 code is pasted into the game, TaleSpire converts the code into a Version 2.


## Version 2

The Version 2 serialization/deserialization was developed based on the Version 1, the first part
of the ByteArray is almost the same, but the last objects are different.

[Version 2: Documentation](/versions/v2/README.md)


# Unuspported versions
These versions are here for documentation purpose only. Current versions DO NOT support them.

## Version 1

The Version 1 serialization/deserialization is based on
[Mercer01/talespireDeserialize](https://github.com/Mercer01/talespireDeserialize),
[brcoding/TaleSpireHtmlSlabGeneration](https://github.com/brcoding/TaleSpireHtmlSlabGeneration)
and [https://github.com/creadth/Creadth.Talespire.DungeonGenerator](https://github.com/creadth/Creadth.Talespire.DungeonGenerator)

[Version 1: Documentation](/versions/v1/README.md)