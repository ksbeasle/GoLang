package validations

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"
)

// **IGNORING LEAP YEARS**
//Errors
var (
	errEmptyTitle         = errors.New("title cannot be empty")
	errEmptyGenre         = errors.New("genre cannot be empty")
	errInvalidRating      = errors.New("rating must be between 1 and 10")
	errEmptyPlatform      = errors.New("platform cannot be empty")
	errEmptyReleaseDate   = errors.New("release date cannot be empty")
	errInvalidYear        = errors.New("year must be between 1958 and the current year")
	errInvalidMonth       = errors.New("month must be (January, February, etc...)")
	errInvalidDay         = errors.New("invalid day less than 0 or greater than 31")
	errInvalidDay31       = errors.New("for this particular month the day must be between 1 and 31")
	errInvalidDay30       = errors.New("for this particular month the day must be between 1 and 30")
	errInvalidDayFeb      = errors.New("for this particular month the day must be 28")
	errInvalidReleaseDate = errors.New("date is not valid")
)

//Sorted because sort package must work with sorted data in order to use some methods
var validMonths = []string{"April", "August", "December", "February", "January", "July", "June", "March", "May", "November", "October", "September"}

//ValidTitle - Validate game title
func ValidTitle(s string) error {
	//Check if the title is empty
	if s == "" {
		return errEmptyTitle
	}
	return nil
}

//Validate game genre
func ValidGenre(s string) error {
	//Check if the genre is empty
	if s == "" {
		return errEmptyGenre
	}
	return nil
}

//Validate game rating
func ValidRating(i int) error {
	//Check if the rating is less than 1 or greater than 10
	//This way we can make sure that the rating was sent in properly and isn't the 0 value of int
	if i < 1 || i > 10 {
		return errInvalidRating
	}
	return nil
}

//Validate game platform
func ValidPlatform(s string) error {
	//Check if the platform is empty
	if s == "" {
		return errEmptyPlatform
	}
	return nil
}

//Validate game release date
func ValidReleaseDate(s string) error {
	//Check if the release date is empty
	if s == "" {
		return errEmptyReleaseDate
	}

	//Split the string to check the date
	splitDate := strings.Split(s, " ")

	year, err := strconv.Atoi(splitDate[2])
	if err != nil {
		return errInvalidYear
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
		return errInvalidYear
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
			return errInvalidMonth
		}
	} else {
		return errInvalidMonth
	}
	//check valid day
	if day < 1 || day > 31 {
		return errInvalidDay
	}
	//Check months with 31 days
	if day == 31 {
		if month == "January" || month == "March" || month == "May" || month == "July" || month == "August" || month == "October" || month == "December" {
			//Do nothing on purpose
		} else {
			return errInvalidDay31
		}
	}
	//check months for 30 days
	if day == 30 {
		if month == "April" || month == "June" || month == "September" || month == "November" {
			//Do nothing on purpose
		} else {
			return errInvalidDay30
		}
	}

	if month == "February" && day > 28 {
		return errInvalidDayFeb
	}

	date = date + strings.TrimSuffix(splitDate[1], ",")

	_, err = time.Parse("2006-01-02", date)
	if err != nil {
		return errInvalidReleaseDate
	}

	return nil
}
