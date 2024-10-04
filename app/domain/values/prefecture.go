package values

import (
	"fmt"
)

type Prefecture string

const (
	Hokkaido  Prefecture = "Hokkaido"
	Aomori    Prefecture = "Aomori"
	Iwate     Prefecture = "Iwate"
	Miyagi    Prefecture = "Miyagi"
	Akita     Prefecture = "Akita"
	Yamagata  Prefecture = "Yamagata"
	Fukushima Prefecture = "Fukushima"
	Ibaraki   Prefecture = "Ibaraki"
	Tochigi   Prefecture = "Tochigi"
	Gunma     Prefecture = "Gunma"
	Saitama   Prefecture = "Saitama"
	Chiba     Prefecture = "Chiba"
	Tokyo     Prefecture = "Tokyo"
	Kanagawa  Prefecture = "Kanagawa"
	Niigata   Prefecture = "Niigata"
	Toyama    Prefecture = "Toyama"
	Ishikawa  Prefecture = "Ishikawa"
	Fukui     Prefecture = "Fukui"
	Yamanashi Prefecture = "Yamanashi"
	Nagano    Prefecture = "Nagano"
	Gifu      Prefecture = "Gifu"
	Shizuoka  Prefecture = "Shizuoka"
	Aichi     Prefecture = "Aichi"
	Mie       Prefecture = "Mie"
	Shiga     Prefecture = "Shiga"
	Kyoto     Prefecture = "Kyoto"
	Osaka     Prefecture = "Osaka"
	Hyogo     Prefecture = "Hyogo"
	Nara      Prefecture = "Nara"
	Wakayama  Prefecture = "Wakayama"
	Tottori   Prefecture = "Tottori"
	Shimane   Prefecture = "Shimane"
	Okayama   Prefecture = "Okayama"
	Hiroshima Prefecture = "Hiroshima"
	Yamaguchi Prefecture = "Yamaguchi"
	Tokushima Prefecture = "Tokushima"
	Kagawa    Prefecture = "Kagawa"
	Ehime     Prefecture = "Ehime"
	Kochi     Prefecture = "Kochi"
	Fukuoka   Prefecture = "Fukuoka"
	Saga      Prefecture = "Saga"
	Nagasaki  Prefecture = "Nagasaki"
	Kumamoto  Prefecture = "Kumamoto"
	Oita      Prefecture = "Oita"
	Miyazaki  Prefecture = "Miyazaki"
	Kagoshima Prefecture = "Kagoshima"
	Okinawa   Prefecture = "Okinawa"
)

var (
	Prefectures = map[Prefecture]struct{}{
		Hokkaido:  {},
		Aomori:    {},
		Iwate:     {},
		Miyagi:    {},
		Akita:     {},
		Yamagata:  {},
		Fukushima: {},
		Ibaraki:   {},
		Tochigi:   {},
		Gunma:     {},
		Saitama:   {},
		Chiba:     {},
		Tokyo:     {},
		Kanagawa:  {},
		Niigata:   {},
		Toyama:    {},
		Ishikawa:  {},
		Fukui:     {},
		Yamanashi: {},
		Nagano:    {},
		Gifu:      {},
		Shizuoka:  {},
		Aichi:     {},
		Mie:       {},
		Shiga:     {},
		Kyoto:     {},
		Osaka:     {},
		Hyogo:     {},
		Nara:      {},
		Wakayama:  {},
		Tottori:   {},
		Shimane:   {},
		Okayama:   {},
		Hiroshima: {},
		Yamaguchi: {},
		Tokushima: {},
		Kagawa:    {},
		Ehime:     {},
		Kochi:     {},
		Fukuoka:   {},
		Saga:      {},
		Nagasaki:  {},
		Kumamoto:  {},
		Oita:      {},
		Miyazaki:  {},
		Kagoshima: {},
		Okinawa:   {},
	}
)

func (p Prefecture) Validate() error {
	if _, ok := Prefectures[p]; !ok {
		return fmt.Errorf("invalid prefecture")
	}

	return nil
}
