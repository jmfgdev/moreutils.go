/*
Copyright (C) <2017>  <jimmybot@teknik.io>
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.
You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import "fmt"
import "strings"

func main() {
	var input string
	fmt.Println("---DOWNCASETHIS---")
	fmt.Printf("Input a string you would like to downcase -->  ")
	fmt.Scanf("%s", &input)
	// shortened from 2 -> 1 lines, thanks to @Strum355
	fmt.Println("Your downcased string is --> " + strings.ToLower(input))
}