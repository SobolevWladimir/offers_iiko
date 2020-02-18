package coupon

//@todo   при генерации купонов на данный момен если купон уже есть в для этого подразделения то он просто не всталяется.
// хорошо бы   не пропускать уже имеющийся купон в бд а сгенерировать еще раз
import (
	"altegra_offers/service/coupon"
	"altegra_offers/service/coupon_category"
	"crypto/rand"
	"errors"
	"math/big"

	"gopkg.in/guregu/null.v3"
)

type GData struct {
	Category     int         `json:"category"`
	Prefix       string      `json:"prefix"`
	Count        int         `json:"count"`
	RandomLeght  int         `json:"random_leght"`
	Status       bool        `db:"status" json:"status" valid:"-"`
	Comment      null.String `db:"comment" json:"comment" valid:"-"`
	Type         int         `json:"type"`
	CharasterSet int         `json:"charaster_set"` // 0 -only digital 1-only latin 2-all
}

func (entity GData) GetValue(field string) (interface{}, error) {
	switch field {
	case "city":
		cat, _ := coupon_category.FindOneById(entity.Category)
		return cat.City, nil
	case "category":
		return entity.Category, nil
	default:
		return nil, errors.New("cant not find field:" + field)
	}
}

// генерируем купоны
func (gdata *GData) Generate() (coupon.Coupons, error) {
	result := coupon.Coupons{}
	if gdata.RandomLeght <= gdata.GetCountLenght() {
		return result, errors.New("small length of random part")
	}
	names := generateNames(gdata.RandomLeght, gdata.Count, gdata.GetChars())
	for _, name := range names {
		result = append(result, gdata.NewCoupon(name))
	}
	return result, nil
}

// кол-во разрядов в длине
func (gdata *GData) GetCountLenght() int {
	result := 0
	for i := gdata.Count; i > 0; i = i / 10 {
		result++
	}
	return result
}
func (gdata *GData) NewCoupon(name string) coupon.Coupon {
	result := coupon.Coupon{}
	result.Type = gdata.Type
	result.Category = gdata.Category
	result.Status = gdata.Status
	result.Comment = gdata.Comment
	result.Name = gdata.Prefix + name
	return result
}

func (gdata *GData) GetChars() []byte {
	result := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	switch gdata.CharasterSet {
	case 0:
		result = []byte("0123456789")
	case 1:
		result = []byte("abcdefghijklmnopqrstuvwxyz")
	case 2:
		result = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	}
	return result
}
func generateNames(length int, count int, chars []byte) []string {
	result := []string{}
	len_arr := len(chars)
	max_count := degreeInt(len_arr, length)
	min_count := degreeInt(len_arr, length-1)
	step := (max_count - min_count) / count
	for i := min_count; i < max_count; i += step {
		num, _ := IntRange(i, i+step-1)
		result = append(result, translateByChars(num, chars))
	}

	return result
}
func addNumInMask(mask []int, max int, pos int) []int {
	result := mask
	num := mask[pos] + 1
	result[pos] = num
	if num >= max && pos+1 < len(mask) {
		result[pos] = 0
		result = addNumInMask(result, max, pos+1)
	}
	return result
}

// переводит десятичное число в  строку
func translateByChars(num int, chars []byte) string {
	result := []byte{}
	arr_leght := len(chars)
	for i := num; i > 0; i = i / arr_leght {
		result = append(result, chars[i%arr_leght])
	}
	return string(result)
}
func getNameByMask(mask []int, chars []byte) string {
	result := []byte{}
	for _, num := range mask {
		result = append(result, chars[num])
	}
	return string(result)
}
func IntRange(min, max int) (int, error) {
	var result int
	switch {
	case min > max:
		// Fail with error
		return result, errors.New("Min cannot be greater than max.")
	case max == min:
		result = max
	case max > min:
		maxRand := max - min
		b, err := rand.Int(rand.Reader, big.NewInt(int64(maxRand)))
		if err != nil {
			return result, err
		}
		result = min + int(b.Int64())
	}
	return result, nil
}

func degreeInt(num, deg int) int {
	result := num
	for i := 1; i < deg; i++ {
		result = result * num
	}
	return result
}
