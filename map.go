package vbcore

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

	//Map generation seeds

	// BlockWaterSeed is the seed for water
	BlockWaterSeed = 0.05
	// BlockLightDirtSeed was used to be Beach
	BlockLightDirtSeed = 0.15
	// BlockGrasslandSeed is the seed for grass
	BlockGrasslandSeed = 0.2
	// BlockTreeSeed was used to be forest
	BlockTreeSeed = 0.6
	// BlockLightMountainSeed is the seed for mountain
	BlockLightMountainSeed = 0.7
	// BlockMountainSeed was used to be snow
	BlockMountainSeed = 0.95
)
