package validations

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	emptyTitleErr       = errors.New("Title cannot be empty.")
	emptyGenreErr       = errors.New("Genre cannot be empty.")
	invalidRatingErr    = errors.New("Rating must be between 1 and 10.")
	emptyPlatformErr    = errors.New("Platform cannot be empty.")
	emptyReleaseDateErr = errors.New("Release date cannot be empty.")
	invalidYearErr      = errors.New("Year must be between 1958 and the current year")
	invalidMonthErr     = errors.New("Month must be (January, February, etc...)")
	invalidDayErr       = errors.New("Day must be between 1 and 31")
	invalidReleaseDate  = errors.New("Date is not valid.")
)

//Sorted because sort package must work with sorted data in order to use some methods
var validMonths = []string{"April", "August", "December", "February", "January", "July", "June", "March", "May", "November", "October", "September"}

//Validate game title
func ValidTitle(s string) error {
	//Check if the title is empty
	if s == "" {
		return emptyTitleErr
	}
	return nil
}

//Validate game genre
func ValidGenre(s string) error {
	//Check if the genre is empty
	if s == "" {
		return emptyGenreErr
	}
	return nil
}

//Validate game rating
func ValidRating(i int) error {
	//Check if the rating is less than 1 or greater than 10
	//This way we can make sure that the rating was sent in properly and isn't the 0 value of int
	if i < 1 || i > 10 {
		return invalidRatingErr
	}
	return nil
}

//Validate game platform
func ValidPlatform(s string) error {
	//Check if the platform is empty
	if s == "" {
		return emptyPlatformErr
	}
	return nil
}

//Validate game release date
func ValidReleaseDate(s string) error {
	//Check if the release date is empty
	if s == "" {
		return emptyReleaseDateErr
	}

	//Split the string to check the date
	splitDate := strings.Split(s, " ")

	year, err := strconv.Atoi(splitDate[2])
	if err != nil {
		return invalidYearErr
	}

	month := splitDate[0]

	day, err := strconv.Atoi(strings.TrimSuffix(splitDate[1], ","))
	if err != nil {
		return err
	}
	var date = ""

	if year <= time.Now().Year() && year > 1958 {
		date = date + splitDate[2] + "-"
	} else {
		return invalidYearErr
	}

	//check if month is actually valid
	check := sort.SearchStrings(validMonths, splitDate[0])
	if check < len(validMonths) && validMonths[check] == splitDate[0] {
		switch month {
		case "January":
			date = date + "01" + "-"
		case "February":
			date = date + "02" + "-"
		case "March":
			date = date + "03" + "-"
		case "April":
			date = date + "04" + "-"
		case "May":
			date = date + "05" + "-"
		case "June":
			date = date + "06" + "-"
		case "July":
			date = date + "07" + "-"
		case "August":
			date = date + "08" + "-"
		case "September":
			date = date + "09" + "-"
		case "October":
			date = date + "10" + "-"
		case "November":
			date = date + "11" + "-"
		case "December":
			date = date + "12" + "-"
		default:
			return invalidMonthErr
		}
	} else {
		return invalidMonthErr
	}

	if day > 0 && day <= 31 {
		date = date + strings.TrimSuffix(splitDate[1], ",")
	}

	_, err = time.Parse("2006-01-02", date)
	if err != nil {
		return invalidReleaseDate
	}

	return nil
}
