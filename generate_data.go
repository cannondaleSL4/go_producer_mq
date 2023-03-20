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
	country = make([]string, 20)
	state = make([]string, 10)
	city = make([]string, 100)
	firstName = make([]string, 100)
	secondName = make([]string, 100)
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

	country = []string{"Australia", "Canada", "Saudi Arabia", "United States", "India", "Russia", "South Africa", "Turkey", "Argentina", "Brazil", "Mexico", "France", "Germany", "Italy", "United Kingdom", "China", "Indonesia", "Japan", "South Korea"}

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

	city = []string{
		"London", "Paris", "New York", "Tokyo", "Dubai", "Barcelona", "Rome", "Madrid", "Singapore", "Amsterdam",
		"Prague", "Los Angeles", "Chicago", "San Francisco", "Berlin", "Hong Kong", "Washington", "Beijing", "Dublin",
		"Istanbul", "Las Vegas", "Milan", "Budapest", "Toronto", "Sydney", "Seoul", "Doha", "Abu Dhabi", "Osaka",
		"Bangkok", "Vienna", "San Diego", "São Paulo", "Melbourne", "Zurich", "Boston", "Lisbon", "Warsaw", "Seattle",
		"Orlando", "Munich", "Houston", "Austin", "Buenos Aires", "Naples", "Copenhagen", "Dallas", "Helsinki",
		"Frankfurt", "Atlanta", "Stockholm", "Miami", "Athens", "Rio de Janeiro", "Hamburg", "Denver", "Montreal",
		"Brussels", "Tel Aviv", "Oslo", "Taipei", "Valencia", "Minneapolis", "Philadelphia", "Calgary", "Portland",
		"Nashville", "Auckland", "Vancouver", "Santiago", "Mexico City", "Mumbai", "Shanghai", "San Jose", "Lyon",
		"Bilbao", "Liverpool", "New Orleans", "Brisbane", "Manchester", "Fukuoka", "Seville", "Riyadh", "Jerusalem",
		"Nanjing", "Minsk", "Salt Lake City", "Phoenix", "Jakarta", "Gothenburg", "Perth", "Glasgow", "Nagoya",
		"Baltimore", "Stuttgart", "Ottawa", "Hanoi", "Sendai", "Cologne", "Marseille",
	}

	firstName = []string{
		"Muhammad", "Noah", "Jack", "Theo", "Leo", "Oliver", "George", "Ethan", "Oscar", "Arthur",
		"Charlie", "Freddie", "Harry", "Zayn", "Alfie", "Finley", "Henry", "Luca", "Thomas", "Aiden",
		"Archie", "Teddy", "Lucas", "Ryan", "Kai", "Liam", "Jaxon", "Louie", "William", "Jacob",
		"Ali", "Caleb", "Isaac", "Joshua", "Jude", "James", "Jayden", "Adam", "Arlo", "Daniel",
		"Elijah", "Max", "Tommy", "Ezra", "Mason", "Theodore", "Roman", "Dylan", "Reuben", "Albie",
		"Alexander", "Toby", "Yusuf", "Logan", "Rory", "Alex", "Harrison", "Kayden", "Nathan", "Ollie",
		"Ayaan", "Elliot", "Ahmad", "Kian", "Samuel", "Hudson", "Jason", "Myles", "Rowan", "Benjamin",
		"Finn", "Omar", "Riley", "Zachary", "Brodie", "Michael", "Abdullah", "Matthew", "Sebastian", "Hugo",
		"Jesse", "Junior", "Oakley", "Abdul", "Eli", "Grayson", "Mateo", "Reggie", "Gabriel", "Hunter",
		"Levi", "Ibrahim", "Jasper", "Syed", "Zion", "Luke", "Seth", "Aaron", "Asher", "Blake",
		"Lily", "Sophia", "Olivia", "Amelia", "Ava", "Isla", "Freya", "Aria", "Ivy", "Mia",
		"Elsie", "Emily", "Ella", "Grace", "Isabella", "Evie", "Hannah", "Luna", "Maya", "Daisy",
		"Zoe", "Millie", "Rosie", "Layla", "Isabelle", "Zara", "Fatima", "Harper", "Nur", "Charlotte",
		"Esme", "Florence", "Maryam", "Poppy", "Sienna", "Sophie", "Aisha", "Emilia", "Willow", "Emma",
		"Evelyn", "Eliana", "Maisie", "Alice", "Chloe", "Erin", "Hallie", "Mila", "Phoebe", "Lyla",
		"Ada", "Lottie", "Ellie", "Matilda", "Molly", "Ruby", "Ayla", "Sarah", "Maddison", "Aaliyah",
		"Aurora", "Maeve", "Bella", "Nova", "Robyn", "Arabella", "Eva", "Lucy", "Eden", "Gracie",
		"Jessica", "Amaya", "Anna", "Leah", "Violet", "Eleanor", "Maria", "Olive", "Orla", "Abigail", "Eliza", "Rose", "Talia",
		"Elizabeth", "Gianna", "Holly", "Imogen", "Nancy", "Annabelle", "Hazel", "Margot", "Raya", "Bonnie",
		"Nina", "Nora", "Penelope", "Scarlett", "Anaya", "Delilah", "Iris",
	}

	secondName = []string{
		"Wang", "Li", "Zhang", "Chen", "Liu", "Devi", "Yang", "Huang", "Singh", "Wu",
		"Kumar", "Xu", "Ali", "Zhao", "Zhou", "Nguyen", "Khan", "Ma", "Lu", "Zhu",
		"Maung", "Sun", "Yu", "Lin", "Kim", "He", "Hu", "Jiang", "Guo", "Ahmed",
		"Khatun", "Luo", "Akter", "Gao", "Zheng", "da Silva", "Tang", "Liang", "Das",
		"Wei", "Mohamed", "Islam", "Shi", "Song", "Xie", "Han", "Garcia", "Mohammad",
		"Tan", "Deng", "Bai", "Ahmad", "Yan", "Kaur", "Feng", "Hernandez", "Rodriguez",
		"Cao", "Lopez", "Hassan", "Hussain", "Gonzalez", "Martinez", "Ceng", "Ibrahim",
		"Peng", "Cai", "Xiao", "Tran", "dos Santos", "Cheng", "Yuan", "Rahman", "Yadav",
		"Su", "Perez", "I", "Le", "Fan", "Dong", "Ye", "Ram", "Tian", "Fu", "Hossain",
		"Kumari", "Sanchez", "Du", "Pereira", "Yao", "Zhong", "Jin", "Pak", "Ding",
		"Mohammed", "Lal", "Yin", "Bibi",
	}
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
