package vbcore

import "image/color"

const (
	// BlockWater represents the block water and is accessable by a player but
	// the player gets damage
	BlockWater = "water"
	// BlockLightDirt represents the block light dirt
	BlockLightDirt = "dirt_light"
	// BlockGrass represents the block grass
	BlockGrass = "grass"
	// BlockTree represents the block tree
	BlockTree = "tree"
	// BlockLightMountain represents the block mountain
	BlockLightMountain = "mountain_light"
	// BlockMountain represents the block mountain and isn't accessable by a
	// palyer
	BlockMountain = "mountain"

	// BlockEndOfMap represents the block endofmap which appears for blocks
	// that are not on the map
	BlockEndOfMap = "endofmap"

	// BlockWaterIsAccessable describes if a block can be accessed by a player
	BlockWaterIsAccessable = true
	// BlockLightDirtIsAccessable describes if a block can be accessed by a
	// player
	BlockLightDirtIsAccessable = true
	// BlockGrassIsAccessable describes if a block can be accessed by a player
	BlockGrassIsAccessable = true
	// BlockTreeIsAccessable describes if a block can be accessed by a player
	BlockTreeIsAccessable = true
	// BlockLightMountainIsAccessable describes if a block can be accessed by
	// a player
	BlockLightMountainIsAccessable = true
	// BlockMountainIsAccessable describes if a block can be accessed by a player
	BlockMountainIsAccessable = false

	//Map generation seeds

	// BlockWaterSeed is the seed for water
	BlockWaterSeed = 0.05
	// BlockLightDirtSeed was used to be Beach
	BlockLightDirtSeed = 0.15
	// BlockGrassSeed is the seed for grass
	BlockGrassSeed = 0.2
	// BlockTreeSeed was used to be forest
	BlockTreeSeed = 0.6
	// BlockLightMountainSeed is the seed for mountain
	BlockLightMountainSeed = 0.7
	// BlockMountainSeed was used to be snow
	BlockMountainSeed = 0.95
)

var (
	// BlockWaterClr is the color in rgb for water
	BlockWaterClr = color.RGBA{62, 96, 193, 0}
	// BlockLightDirtClr is the color in rgb for lightdirt
	BlockLightDirtClr = color.RGBA{93, 128, 253, 0}
	// BlockGrassClr is the color in rgb for grass
	BlockGrassClr = color.RGBA{116, 169, 99, 0}
	// BlockTreeClr is the color in rgb for a tree
	BlockTreeClr = color.RGBA{62, 126, 98, 0}
	// BlockLightMountainClr is the color in rgb for a mountain
	BlockLightMountainClr = color.RGBA{130, 132, 128, 0}
	// BlockMountainClr is the color in rgb for a mountain (snow)
	BlockMountainClr = color.RGBA{210, 210, 215, 0}
)
