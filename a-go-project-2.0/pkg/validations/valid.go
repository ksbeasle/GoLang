package validations

import (
	"errors"
	"strings"
	"sort"
	"fmt"
	"strconv"
	"time"
)

//TODO: create custom error for invalid month

var validMonths = []string{"April", "August", "December", "February", "January", "July", "June", "March", "May", "November", "October", "September"}

//Validate game title
func ValidTitle(s string) (bool, error) {
	//Check if the title is empty
	if s == "" {
		return false, errors.New("Title cannot be empty")
	}
	return true, nil
}

//Validate game genre
func ValidGenre(s string) (bool, error) {
	//Check if the genre is empty
	if s == "" {
		return false, errors.New("Genre cannot be empty")
	}
	return true, nil
}

//Validate game rating
func ValidRating(i int) (bool, error) {
//Check if the rating is less than 1 or greater than 10
//This way we can make sure that the rating was sent in properly and isn't the 0 value of int
	if i < 1 || i > 10 {
		return false, errors.New("Rating Must be between 1 and 10")
	}
	return true, nil
}

//Validate game platform
func ValidPlatform(s string) (bool, error){
	//Check if the platform is empty
	if s == "" {
		return false, errors.New("Platform cannot be empty")
	}
	return true, nil
}

//Validate game release date
func ValidReleaseDate(s string) (bool, error){
	//Check if the release date is empty
	if s == "" {
		return false, errors.New("Release Date cannot be empty")
	}
	
	//Split the string to check the date
	splitDate := strings.Split(s, " ")

	year, err:= strconv.Atoi(splitDate[2])
	if err != nil {
		return false, errors.New("Invalid year")
	}

	month := splitDate[0]
	fmt.Println()
	day, err := strconv.Atoi(strings.TrimSuffix(splitDate[1], ","))
	if err != nil {
		return false, errors.New("Invalid month")
	}
	var date = ""
	
	if year <= time.Now().Year() && year > 1958 {
		date = date + splitDate[2] + "-"
	} else {
		return false, errors.New("Invalid month")
	}


	//check if month is actually valid
	check := sort.SearchStrings(validMonths, splitDate[0])
	if check < len(validMonths) && validMonths[check] == splitDate[0] {
		switch month {
		case "January":
			date = date + "01" +  "-"
		case "February":
			date = date + "02" +  "-"
		case "March":
			date = date + "03" +  "-"
		case "April":
			date = date + "04" +  "-"
		case "May":
			date = date + "05" +  "-"
		case "June":
			date = date + "06" +  "-"
		case "July":
			date = date + "07" +  "-"
		case "August":
			date = date + "08" +  "-"
		case "September":
			date = date + "09" +  "-"
		case "October":
			date = date + "10" +  "-"
		case "November":
			date = date + "11" +  "-"
		case "December":
			date = date + "12" +  "-"
		default:
			return false, errors.New("Invalid Month")
		}
	}else {
		return false, errors.New("Invalid Month")
	}

	if day > 0 && day <= 31 {
		date = date + strings.TrimSuffix(splitDate[1], ",")
	} 
	fmt.Println(date)
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false, errors.New("Invalid date")
	}
	fmt.Println(t)
	return true, nil
}
