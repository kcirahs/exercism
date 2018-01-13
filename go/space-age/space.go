//Package space contains the information needed to calculate how old
//a person would be on different planets, given an age in earth seconds.
package space

//secinyear represents the amount of Earth seconds in one Earth year.
const secinyear float64 = 31557600

//Planet was defined in the test case data as its ow
type Planet string

//Age takes a given age in earth seconds and a planet, then
//calculates the equivalent age on that planet.
func Age(s float64, p Planet) float64 {
	spaceAgeConv := map[Planet]float64{
		"Earth":   1 * secinyear,
		"Mercury": 0.2408467 * secinyear,
		"Venus":   0.61519726 * secinyear,
		"Mars":    1.8808158 * secinyear,
		"Jupiter": 11.862615 * secinyear,
		"Saturn":  29.447498 * secinyear,
		"Uranus":  84.016846 * secinyear,
		"Neptune": 164.79132 * secinyear,
	}
	return s / spaceAgeConv[p]
}
