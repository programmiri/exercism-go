//Package raindrops adds a method to translate numbers in raindrop-speak
package raindrops

import "strconv"

// Convert converts a number to a string in raindrop-speak
func Convert(num int) string {
        var result string
        factorable := 0
        if num%3 == 0 {
                result = result + "Pling"
                factorable++
        }
        if num%5 == 0 {
                result = result + "Plang"
                factorable++
        }
        if num%7 == 0 {
                result = result + "Plong"
                factorable++
        }
        if factorable == 0 {
                return strconv.Itoa(num)
        } else {
                return result
        }

}
