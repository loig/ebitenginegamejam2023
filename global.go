package main

const (
	gUnit = 10

	gScreenWidth  = gUnit * 160
	gScreenHeight = gUnit * 90

	gPlayerWidth         = gUnit * 10
	gPlayerHeight        = gUnit * 15
	gPlayerIncrSpeedY    = float64(gUnit) * 0.2
	gPlayerMaxSpeedY     = float64(gUnit) * 2
	gPlayerSlidingWidth  = gUnit * 15
	gPlayerSlidingHeight = gUnit * 10
	gPlayerSlideDuration = 40
	gPlayesSpeedYMult    = 2

	gBigGhostWidth         = gUnit * 20
	gBigGhostHeight        = gUnit * 30
	gBigGhostIncrSpeedY    = float64(gUnit) * 0.1
	gBigGhostMaxSpeedY     = float64(gUnit) * 1.5
	gBigGhostSlidingWidth  = gUnit * 30
	gBigGhostSlidingHeight = gUnit * 20
	gBigGhostSlideDuration = 40
	gBigGhostSpeedYMult    = 1

	gSmallGhostWidth         = gUnit * 5
	gSmallGhostHeight        = gUnit * 5
	gSmallGhostIncrSpeedY    = float64(gUnit) * 0.4
	gSmallGhostMaxSpeedY     = float64(gUnit) * 3.5
	gSmallGhostSlidingWidth  = gUnit * 5
	gSmallGhostSlidingHeight = gUnit * 5
	gSmallGhostSlideDuration = 20
	gSmallGhostSpeedYMult    = 3

	gTallGhostWidth         = gUnit * 5
	gTallGhostHeight        = gUnit * 25
	gTallGhostIncrSpeedY    = float64(gUnit) * 0.2
	gTallGhostMaxSpeedY     = float64(gUnit) * 2
	gTallGhostSlidingWidth  = gUnit * 25
	gTallGhostSlidingHeight = gUnit * 5
	gTallGhostSlideDuration = 60
	gTallGhostSpeedYMult    = 2

	gGravity = 1.0
)
