// Haversine formula implementation from: http://rosettacode.org/wiki/Haversine_formula#Go
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"math"
)


func main() {

	file, err := os.Open("/Users/ianlewis/Dropbox/customer-venue-latlng.csv")
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		fmt.Println("Error:", err)
		return
	}

	// automatically call Close() at the end of current method
	defer file.Close()
	//
	reader := csv.NewReader(file)

	fout, err1 := os.Create("/Users/ianlewis/Dropbox/customer-venue-latlng-dist.csv")
	if err1 != nil {
		// err is printable
		// elements passed are separated by space automatically
		fmt.Println("Error 1:", err1)
		return
	}
	defer fout.Close()

	writer := csv.NewWriter(fout)


	lineCount := 0
	for {

		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()

		if lineCount == 0 {
			lineCount += 1
			writer.Write(append(record, "distance_km"))
			continue
		}

		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// record is an array of string so is directly printable
		//fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
		// and we can iterate on top of that
		//for i := 0; i < len(record); i++ {
		//	fmt.Println(" ", record[i])
		//}

		lat1, _ := strconv.ParseFloat(record[2], 64)
		lng1, _ := strconv.ParseFloat(record[3], 64)
		lat2, _ := strconv.ParseFloat(record[4], 64)
		lng2, _ := strconv.ParseFloat(record[5], 64)


		d := hsDist(degPos(lat1, lng1), degPos(lat2, lng2))

		fmt.Println(lineCount, record[0], lat1, lng1, lat2, lng2, "--> ", d)
		writer.Write(append(record, strconv.FormatFloat(d, (byte)('f'), 6, 64)))
		writer.Flush()
		fmt.Println()

		//if 10 == lineCount {
		//	break
		//}
		lineCount += 1
	}

}

func haversine(θ float64) float64 {
	return .5 * (1 - math.Cos(θ))
}

type pos struct {
	φ float64 // latitude, radians
	ψ float64 // longitude, radians
}

func degPos(lat, lon float64) pos {
	return pos{lat * math.Pi / 180, lon * math.Pi / 180}
}

const rEarth = 6372.8 // km

func hsDist(p1, p2 pos) float64 {
	return 2 * rEarth * math.Asin(math.Sqrt(haversine(p2.φ-p1.φ)+
			math.Cos(p1.φ)*math.Cos(p2.φ)*haversine(p2.ψ-p1.ψ)))
}



