package main

import (
	"fmt"
	"strconv"
	"strings"
)

func predictGOL(input []string) []string {

	// parse to output type
	var ds goldata
	ds.input = input
	ds.parse()
	ds.countNeighbours()

	output := ds.parseOutput()

	return output //[]string{""}
}

type goldata struct {
	input                           []string
	xd, yd                          int
	input2d, neighbours2d, output2d [][]int
}

func (v *goldata) parse() {
	var inputData [][]int
	for i0, line := range v.input {
		if i0 == 0 {
			fields := strings.Split(line, " ")
			for i1, dim := range fields {
				if i1 == 0 {
					n, _ := strconv.ParseUint(dim, 10, 32)
					v.yd = int(n) - 1
				}
				if i1 == 1 {
					n, _ := strconv.ParseUint(dim, 10, 32)
					v.xd = int(n) - 1
				}
			}
		}
		if i0 >= 1 {
			tmpSlice := make([]int, v.xd+1)
			for i2, char := range line {
				// fmt.Println(char, i2)
				switch char {
				case rune(46): // ascii char 42 = .
					tmpSlice[i2] = 0 // append(tmpSlice, tmpVal)
				case rune(42): // ascii char 42 = *
					tmpSlice[i2] = 1 // append(tmpSlice, []int{1})
				}
			}
			inputData = append(inputData, tmpSlice)
		}
		v.input2d = inputData
	}
}

func (v *goldata) countNeighbours() {
	//  Y-X-  Y-X  Y-X+
	//   YX-   YX   YX+
	//  Y+X-  Y+X  Y+X+
	// take cell position Y? X?

	// calc neighbours and build 2d array for that
	// apply rules
	// build output 2d array
	var neighbourCount [][]int
	var outputArray [][]int
	for i0 := 0; i0 <= v.yd; i0++ {
		tmpSlice := make([]int, v.xd+1)
		tmpSliceOut := make([]int, v.xd+1)
		for i1 := 0; i1 < v.xd; i1++ {
			ncount := 0
			ypos := i0
			xpos := i1
			yposMin := ypos - 1
			yposMax := ypos + 1
			xposMin := xpos - 1
			xposMax := xpos + 1

			if xposMin >= 0 {
				if yposMin >= 0 {
					if v.input2d[yposMin][xposMin] == 1 {
						ncount++
					}
				}
				if v.input2d[ypos][xposMin] == 1 {
					ncount++
				}
			}
			if yposMin >= 0 {
				if v.input2d[yposMin][xpos] == 1 {
					ncount++
				}
				if xposMax <= v.xd {
					if v.input2d[yposMin][xposMax] == 1 {
						ncount++
					}
				}
			}
			if xposMax <= v.xd {
				if v.input2d[ypos][xposMax] == 1 {
					ncount++
				}
				if yposMax <= v.yd {
					if v.input2d[yposMax][xposMax] == 1 {
						ncount++
					}
				}
			}
			if yposMax <= v.yd {
				if v.input2d[yposMax][xpos] == 1 {
					ncount++
				}
				if xposMin >= 0 {
					if v.input2d[yposMax][xposMin] == 1 {
						ncount++
					}
				}
			}
			tmpSlice[i1] = ncount
			// check if we are at edge and apply our rules

			deadCell := true
			if ypos > 0 {
				if ypos < v.yd {
					if xpos > 0 {
						if xpos < v.xd {
							if ncount == 3 {
								if v.input2d[ypos][xpos] == 0 {
									deadCell = false
								}
							}
							if v.input2d[ypos][xpos] == 1 {
								if ncount >= 2 {
									if ncount <= 3 {
										deadCell = false
									}
								}
							}
						}
					}
				}
			}
			if deadCell {
				tmpSliceOut[i1] = 0
			} else {
				tmpSliceOut[i1] = 1
			}
		}
		neighbourCount = append(neighbourCount, tmpSlice)
		outputArray = append(outputArray, tmpSliceOut)
	}
	v.neighbours2d = neighbourCount
	v.output2d = outputArray
}

func (v *goldata) parseOutput() []string {
	var tmpstr []string
	tmpstr = append(tmpstr, fmt.Sprint(v.yd+1)+" "+fmt.Sprint(v.xd+1))
	for _, line := range v.output2d {
		strbld := ""
		for _, char := range line {
			if char == 0 {
				strbld = strbld + "."
			}
			if char == 1 {
				strbld = strbld + "*"
			}
		}
		tmpstr = append(tmpstr, strbld)
	}
	return tmpstr
}
