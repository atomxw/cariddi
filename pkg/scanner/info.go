/*
==========
Cariddi
==========
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:  https://github.com/edoardottt/cariddi

	@Author:      edoardottt, https://www.edoardottt.com

	@License: https://github.com/edoardottt/cariddi/blob/main/LICENSE

*/

package scanner

import (
	"regexp"
	"sync"
)

// Info struct.
// Name = the name that identifies the information.
// Regex = The regular expression to be matched.
type Info struct {
	Name  string
	Regex regexp.Regexp
}

// InfoMatched struct.
// Info = Info struct.
// Url = url in which the information is found.
// Match = the string matching the regex.
type InfoMatched struct {
	Info  Info
	URL   string
	Match string
}

var (
	infos     []Info    //nolint: gochecknoglobals
	onceInfos sync.Once //nolint: gochecknoglobals
)

// GetInfoRegexes returns all the info structs.
func GetInfoRegexes() []Info {
	onceInfos.Do(func() {
		infos = []Info{
			{
				"Email address",
				*regexp.MustCompile(`(?i)([a-zA-Z0-9_.+-]+@[a-zA-Z0-9]+[a-zA-Z0-9-]*\.[a-zA-Z0-9-.]*[a-zA-Z0-9]{2,})`),
			},
			{
				"HTML comment",
				*regexp.MustCompile(`(?i)(\<![\s]*--[\-!@#$%^&*:;ºª.,"'(){}\w\s\/\\[\]]*--[\s]*\>)`),
			},
			{
				"Internal IP address",
				*regexp.MustCompile(`((172\.\d{1,3}\.\d{1,3}\.\d{1,3})|(192\.168\.\d{1,3}\.\d{1,3})|` +
					`(10\.\d{1,3}\.\d{1,3}\.\d{1,3})|([fF][eE][89aAbBcCdDeEfF]::))`),
			},
			{
				"IPv4 address",
				*regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`),
			},
			/*
				TOO MANY FALSE POSITIVES
				{
					"BTC address",
					`([13]|bc1)[A-HJ-NP-Za-km-z1-9]{27,34}`,
				},
			*/
			/*
				HOW TO AVOID VERY VERY LONG BASE64 IMAGES ???
				{
					"Base64-encoded JSON",
					`ey(A|B)[A-Za-z0-9+\/]{20,}(={0,2})`,
				},
			*/
		}
	})

	return infos
}

// RemoveDuplicateInfos removes duplicates from Infos found.
func RemoveDuplicateInfos(input []InfoMatched) []InfoMatched {
	keys := make(map[string]bool)
	list := []InfoMatched{}

	for _, entry := range input {
		if _, value := keys[entry.Match]; !value {
			keys[entry.Match] = true
			list = append(list, entry)
		}
	}

	return list
}
