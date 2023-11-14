package entity

import (
	"math/rand"
	"time"

	"github.com/srik007/sensor-api/domain/valueObjects"
)

var oceanFishNames = []string{
	"AtlanticBluefinTuna", "AtlanticCod", "AtlanticGoliathGrouper", "AtlanticSalmon", "AtlanticTrumpetfish",
	"AtlanticWolffish", "BandedButterflyfish", "BelugaSturgeon", "BlueMarlin", "BlueTang", "BluebandedGoby",
	"BlueheadWrasse", "CaliforniaGrunion", "ChileanCommonHake", "ChileanJackMackerel", "ChinookSalmon",
	"ClownTriggerfish", "Coelacanth", "CommonClownfish", "CommonDolphinfish", "CommonFangtooth", "DeepSeaAnglerfish",
	"FlashlightFish", "FrenchAngelfish", "GreatBarracuda", "GreenMorayEel", "GuineafowlPuffer", "JohnDory",
	"LeafySeadragon", "LongsnoutSeahorse", "MexicanLookdown", "NassauGrouper", "NorthernRedSnapper", "Oarfish",
	"OceanSunfish", "OrangeRoughy", "PacificBlackdragon", "PacificHalibut", "PacificHerring", "PacificSardine",
	"PatagonianToothfish", "PeruvianAnchoveta", "PinkSalmon", "PygmySeahorse", "QueenAngelfish", "QueenParrotfish",
	"RedLionfish", "Sailfish", "SarcasticFringehead", "ScarletFrogfish", "Scorpionfish", "SkipjackTuna",
	"SlenderSnipeEel", "SmalltoothSawfish", "SockeyeSalmon", "SpottedMoray", "SpottedPorcupinefish", "SpottedRatfish",
	"Stonefish", "StoplightLoosejaw", "SummerFlounder", "Swordfish", "TanBristlemouth", "ThreespotDamselfish",
	"TropicalTwoWingFlyingfish", "Wahoo", "WhiptailGulper", "WhiteRingGardenEel", "YellowfinTuna",
}

func GenerateFakeSensorData(sensor Sensor) SensorData {
	temparature := getTemparature(sensor.Coordinate.Z)
	transparency := getTransparency(sensor.Coordinate.Z)
	fishSpecies := pickRandomOceanSpecies()
	return SensorData{
		Transparency: uint(transparency),
		Temparature:  temparature,
		Species:      fishSpecies,
		SensorId:     sensor.ID,
	}

}

func getTemparature(depth float64) valueObjects.Temparature {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return valueObjects.Temparature{Value: 25.0 - 0.1*depth + rand.Float64()*5.0, Scale: "Celisus"}
}

func getTransparency(depth float64) int {
	maxDepth := 100.0
	return int(100.0 - (depth/maxDepth)*100.0)
}

func pickRandomOceanSpecies() []Specie {
	var oceanSpecies []Specie
	numberONewSpecies := rand.Intn(5)
	for i := 0; i < numberONewSpecies; i++ {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		randomIndex := rand.Intn(len(oceanFishNames))
		oceanSpecies = append(oceanSpecies, Specie{Name: oceanFishNames[randomIndex], Count: 1 + rand.Intn(100)})
	}
	return oceanSpecies
}
