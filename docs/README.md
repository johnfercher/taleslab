# TaleSlab

**WARNING: this project is in "alpha", so it's in heavy development and there is no concerns with breaking changes at this
development phase.**

A map generator for the [TaleSpire](https://store.steampowered.com/app/720620/TaleSpire/) (rpg tabletop). This project
support the creation of maps in different sizes, biomes and characteristics. The maps can be built using [procedural algorithms](http://johnfercher.com/taleslab/#/codes/procedurals/beach)
or even with [georeferencing data](http://johnfercher.com/taleslab/#/codes/elevations/petropolis).

This project uses this other projects to do all the magic:
* [talescoder](https://github.com/johnfercher/talescoder) to serde TaleSpire code into Go objects.
* [tessadem-sdk](https://github.com/johnfercher/tessadem-sdk) to retrieve georeferencing data from [tessadem-api](https://tessadem.com/elevation-api/).
* [go-rrt](https://github.com/johnfercher/go-rrt) to generate rivers procedurally in a relief environment.

## Map Example
![version_size](https://raw.githubusercontent.com/johnfercher/taleslab/main/cmd/elevations/danielabeach/image.png)

## TaleSpire Code Example
[filename](https://raw.githubusercontent.com/johnfercher/taleslab/main/cmd/procedurals/beach/data.txt ':include :type=code')

## Go Code Example
[filename](https://raw.githubusercontent.com/johnfercher/taleslab/main/cmd/procedurals/beach/main.go ':include :type=code')

## Video Example
<iframe width="560" height="560" src="https://www.youtube.com/embed/oCb4TEgpAt0?si=2rvhn-63IDkJeV4M" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

## Doc Structure
[filename](_sidebar.md ':include :type=markdown')