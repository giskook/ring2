package base

import (
	"fmt"
	"strconv"
)

func Latd2dm(d string) string {
	ld, _ := strconv.ParseFloat(d, 64)
	ld_inter := int(ld)
	ld_decimal := ld - float64(ld_inter)
	ld_decimal_min_inter := int(ld_decimal * 60)
	ld_decimal_min_decimal_inter := int(ld_decimal*60*10000) % 10000
	return fmt.Sprintf("%02d%02d%04d", ld_inter, ld_decimal_min_inter, ld_decimal_min_decimal_inter)
}

func Longd2dm(d string) string {
	ld, _ := strconv.ParseFloat(d, 64)
	ld_inter := int(ld)
	ld_decimal := ld - float64(ld_inter)
	ld_decimal_min_inter := int(ld_decimal * 60)
	ld_decimal_min_decimal_inter := int(ld_decimal*60*10000) % 10000
	return fmt.Sprintf("%03d%02d%04d", ld_inter, ld_decimal_min_inter, ld_decimal_min_decimal_inter)

}
