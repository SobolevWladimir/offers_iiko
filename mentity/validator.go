package mentity

import (
	"unicode"

	"github.com/asaskevich/govalidator"
	"gopkg.in/guregu/null.v3"
)

func init() {
	govalidator.CustomTypeTagMap.Set("nullUuidv4", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case null.String:
			if v.Valid {
				return govalidator.IsUUIDv4(v.String)
			}
			return false
		default:
			return false
		}
	}))
	govalidator.CustomTypeTagMap.Set("nullUuid", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case null.String:
			if v.Valid {
				return govalidator.IsUUID(v.String)
			}
			return false
		default:
			return false
		}
	}))
	govalidator.CustomTypeTagMap.Set("nullAlphanum", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case null.String:
			if v.Valid {
				return govalidator.IsAlphanumeric(v.String)
			}
			return false
		default:
			return false
		}
	}))
	govalidator.CustomTypeTagMap.Set("nullEmail", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case null.String:
			if v.Valid {
				return govalidator.IsEmail(v.String)
			}
			return false
		default:
			return false
		}
	}))
	govalidator.CustomTypeTagMap.Set("nullJson", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case null.String:
			if v.Valid {
				return govalidator.IsJSON(v.String)
			}
			return false
		default:
			return false
		}
	}))
	govalidator.CustomTypeTagMap.Set("null_utfletternum", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case null.String:
			if v.Valid {
				return govalidator.IsUTFLetterNumeric(v.String)
			}
			return false
		default:
			return false
		}
	}))
	govalidator.CustomTypeTagMap.Set("null_float", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case null.String:
			if v.Valid {
				return govalidator.IsFloat(v.String)
			}
			return false
		default:
			return false
		}
	}))
	govalidator.CustomTypeTagMap.Set("null_utfletternumspace", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case null.String:
			if v.Valid {
				return IsUTFLetterNumericSpace(v.String)
			}
			return false
		default:
			return false
		}
	}))

	govalidator.CustomTypeTagMap.Set("ingridients", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		switch v := i.(type) {
		case null.String:
			if v.Valid {
				return IsUTFLetterNumericSpace(v.String)
			}
			return false
		default:
			return false
		}
	}))
	govalidator.TagMap["utfletternumspace"] = govalidator.Validator(func(str string) bool {
		return IsUTFLetterNumericSpace(str)

	})
}
func IsUTFLetterNumericSpace(str string) bool {
	if govalidator.IsNull(str) {
		return true
	}
	for _, c := range str {

		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && !unicode.IsSpace(c) { //letters && numbers are ok
			return false
		}
	}
	return true

}
