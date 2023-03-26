package sample

import (
	"laptop-grpc/pb"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}

}
func randomId() string {
	return uuid.New().String()
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomLaptopBrand() string {
	return randomStingFromSet("Apple", "Dell", "Lenov")
}
func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStingFromSet("Mackbook Air", "Macbook pro")
	case "Dell":
		return randomStingFromSet("Latitude", "Vostro", "XPS", "Alience")
	default:
		return randomStingFromSet("Thinkpad p1", "Thinkpad x1", "Thinkpad p53")

	}

}
func randomCPUBrand() string {
	return randomStingFromSet("Intel", "Amd")
}

func randomGPUBrand() string {
	return randomStingFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStingFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1070",
		)
	}
	return randomStingFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
	)
}
func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStingFromSet(
			"Xenon E-22687",
			"Core I9-22343",
			"Core I7-22443",
		)
	}
	return randomStingFromSet(
		"RYZON 7 PRO 8978",
		"RYZON 8 PRO 8678",
		"RYZON 9 PRO 8578",
	)
}
func randomStingFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
func randomBool() bool {
	return rand.Intn(2) == 1

}
