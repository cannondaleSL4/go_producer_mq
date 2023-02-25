package main

import (
	"fmt"
	"github.com/google/uuid"
	"go_producer_mq/data"
	"math/rand"
	"time"
)

var street, country, state, city, firstName, secondName []string
var productPrice map[string]float32

func init() {
	street = make([]string, 10)
	country = make([]string, 10)
	state = make([]string, 10)
	city = make([]string, 10)
	firstName = make([]string, 10)
	secondName = make([]string, 10)
	productPrice = make(map[string]float32)

	productPrice["milk"] = 3.25
	productPrice["salt"] = 0.25
	productPrice["bob"] = 2.35
	productPrice["sugar"] = 1.15
	productPrice["bread"] = 1.1
	productPrice["water"] = 0.75
	productPrice["coca-cola"] = 2
	productPrice["TVset"] = 650
	productPrice["sneakers Nike"] = 25
	productPrice["socks"] = 7.25
	productPrice["Iphone"] = 600

	street[0] = "Yanki Kupali"
	street[1] = "Ilyniskaia"
	street[2] = "B.Pokrovskaya"
	street[3] = "Gor'kogo"
	street[4] = "Orechovskaia"
	street[5] = "Monchegorskaia"
	street[6] = "Pr.Gagarina"
	street[7] = "Krasnaia"
	street[8] = "Malaya Iamskaia"
	street[9] = "Medisson St."

	country[0] = "Russia"
	country[1] = "Ukraine"
	country[2] = "Kazahstan"
	country[3] = "Kyrgiztan"
	country[4] = "Belarussia"
	country[5] = "Poland"
	country[6] = "Germanya"
	country[7] = "Gb.Britain"
	country[8] = "Turkey"
	country[9] = "Usa"

	state[0] = "New York"
	state[1] = "Massachusetts"
	state[2] = "Arizona"
	state[3] = "Texas"
	state[4] = "California"
	state[5] = "New Mexico"
	state[6] = "Oregon"
	state[7] = "Florida"
	state[8] = "Washington"
	state[9] = "South Coraline"

	city[0] = "New York"
	city[1] = "Chicago"
	city[2] = "Boston"
	city[3] = "Los Angeles"
	city[4] = "Moscow"
	city[5] = "Nizhniy Novgorod"
	city[6] = "Kstovo"
	city[7] = "Berlin"
	city[8] = "Viena"
	city[9] = "London"

	firstName[0] = "Evgenia"
	firstName[1] = "Kira"
	firstName[2] = "Dmitriy"
	firstName[3] = "Michail"
	firstName[4] = "Jhon"
	firstName[5] = "Mikle"
	firstName[6] = "Igor"
	firstName[7] = "Nikita"
	firstName[8] = "Anon"
	firstName[9] = "Ira"

	secondName[0] = "Smith"
	secondName[1] = "Johnson"
	secondName[2] = "Williams"
	secondName[3] = "Brown"
	secondName[4] = "Jones"
	secondName[5] = "Garcia"
	secondName[6] = "Miller"
	secondName[7] = "Davis"
	secondName[8] = "Rodriguez"
	secondName[9] = "Martinez"
}

func GetOrder() data.UsersOrder {

	rand.Seed(time.Now().UnixNano())

	address := data.UsersOrder_UserAddress{
		Country:         getRandomFromArray(country),
		State:           getRandomFromArray(state),
		City:            getRandomFromArray(city),
		ZipCode:         rand.Int31n(999999),
		Street:          getRandomFromArray(street),
		NumberHouse:     rand.Int31n(100),
		NumberApartment: rand.Int31n(100),
	}

	user := data.UsersOrder_User{
		UserName: fmt.Sprintf("%s %s", getRandomFromArray(firstName), getRandomFromArray(secondName)),
		Age:      rand.Int31n(85),
		Role:     getRole(),
	}

	orders, accountTotal := getProducts()

	myData := data.UsersOrder{
		Uuid: &data.UsersOrder_UUID{
			Value: uuid.NewString(),
		},
		Address:      &address,
		Payed:        intToBool(rand.Intn(2)),
		User:         &user,
		Order:        orders,
		AccountTotal: accountTotal,
		TimeStamp:    timestamp(),
	}

	return myData
}

func intToBool(number int) bool {
	return number == 0
}

func getRole() data.UsersOrder_Role {
	roleInt := rand.Intn(2)
	if roleInt == 0 {
		return data.UsersOrder_USER
	}
	return data.UsersOrder_VIP_USER
}

func getRandomFromArray(array []string) string {
	return array[rand.Intn(9)]
}

func getProducts() ([]*data.UsersOrder_Product, float32) {
	size := rand.Intn(len(productPrice)-4) + 4

	var accountTotal float32
	var products []*data.UsersOrder_Product

	prod := make([]string, 0, len(productPrice))
	for k := range productPrice {
		prod = append(prod, k)
	}

	randomInt := rand.Perm(len(productPrice))

	for i := 1; i < size+1; i++ {

		title := prod[randomInt[i-1]]
		amount := rand.Int31n(10)
		cost := productPrice[title] * float32(amount)
		accountTotal = accountTotal + cost
		products = append(products, &data.UsersOrder_Product{
			Title:  title,
			Price:  productPrice[title],
			Amount: amount,
			Cost:   cost,
		})
	}

	return products, accountTotal
}

func timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
