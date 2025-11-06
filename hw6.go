package main

import (
	"fmt"
)

type Product struct {
	id       int
	name     string
	desc     string
	price    float64
	category string
	stock    int
}

type Customer struct {
	id      int
	name    string
	email   string
	phone   string
	address string
}

type CartItem struct {
	product Product
	amount  int
}

type Cart struct {
	customer Customer
	items    []CartItem
	discount float64
}

type Order struct {
	id       int
	customer Customer
	items    []CartItem
	total    float64
	status   string
}

var products []Product
var customers []Customer
var orders []Order
var carts []Cart
var nextProductID, nextCustomerID, nextOrderID int

func main() {
	for {
		fmt.Println("=== Онлайн-магазин 'TechStore' ===")
		fmt.Println("1. Управління товарами")
		fmt.Println("2. Управління клієнтами")
		fmt.Println("3. Кошик покупок")
		fmt.Println("4. Замовлення")
		fmt.Println("5. Вихід")
		fmt.Print("> ")
		var choice int
		fmt.Scan(&choice)
		fmt.Scanln()

		switch choice {
		case 1:
			menuProducts()
		case 2:
			menuCustomers()
		case 3:
			menuCart()
		case 4:
			menuOrders()
		case 5:
			fmt.Println("Вихід з програми...")
			return
		default:
			fmt.Println("Невірний вибір.")
		}
	}
}

func menuProducts() {
	for {
		fmt.Println("--- Меню товарів ---")
		fmt.Println("1. Додати товар")
		fmt.Println("2. Переглянути всі товари")
		fmt.Println("3. Знайти товар за ID")
		fmt.Println("4. Пошук за категорією")
		fmt.Println("5. Оновити товар")
		fmt.Println("6. Повернутися до головного меню")
		fmt.Print("> ")
		var choice int
		fmt.Scan(&choice)
		fmt.Scanln()

		switch choice {
		case 1:
			addProduct()
		case 2:
			listProducts()
		case 3:
			findProductByID()
		case 4:
			searchByCategory()
		case 5:
			updateProduct()
		case 6:
			return
		default:
			fmt.Println("Невірний вибір.")
		}
	}
}

func addProduct() {
	var name, desc, category string
	var price float64
	var stock int

	fmt.Print("Введіть назву товару: ")
	fmt.Scanln(&name)

	fmt.Print("Введіть опис: ")
	fmt.Scanln(&desc)

	fmt.Print("Введіть ціну: ")
	fmt.Scan(&price)
	fmt.Scanln()

	fmt.Print("Введіть категорію: ")
	fmt.Scanln(&category)

	fmt.Print("Введіть кількість на складі: ")
	fmt.Scan(&stock)
	fmt.Scanln()

	nextProductID++
	product := Product{nextProductID, name, desc, price, category, stock}
	products = append(products, product)
	fmt.Println("Товар успішно додано!")
}

func listProducts() {
	if len(products) == 0 {
		fmt.Println("Каталог порожній.")
		return
	}
	for _, p := range products {
		fmt.Printf("ID: %d | %s | Категорія: %s | Ціна: %.2f грн | Наявність: %d шт.\n",
			p.id, p.name, p.category, p.price, p.stock)
	}
}

func findProductByID() {
	var id int
	fmt.Print("Введіть ID товару: ")
	fmt.Scan(&id)
	fmt.Scanln()
	for _, p := range products {
		if p.id == id {
			fmt.Printf("Знайдено: %s | Ціна: %.2f грн | %d шт.\n", p.name, p.price, p.stock)
			return
		}
	}
	fmt.Println("Товар не знайдено.")
}

func searchByCategory() {
	var category string
	fmt.Print("Введіть категорію: ")
	fmt.Scanln(&category)
	found := false
	for _, p := range products {
		if p.category == category {
			fmt.Printf("%s | %.2f грн | %d шт.\n", p.name, p.price, p.stock)
			found = true
		}
	}
	if !found {
		fmt.Println("Товари не знайдено.")
	}
}

func updateProduct() {
	var id int
	fmt.Print("Введіть ID товару для оновлення: ")
	fmt.Scan(&id)
	fmt.Scanln()

	for i := range products {
		if products[i].id == id {
			fmt.Print("Нова ціна: ")
			fmt.Scan(&products[i].price)
			fmt.Scanln()
			fmt.Print("Нова кількість: ")
			fmt.Scan(&products[i].stock)
			fmt.Scanln()
			fmt.Println("Товар оновлено.")
			return
		}
	}
	fmt.Println("Товар не знайдено.")
}

func menuCustomers() {
	for {
		fmt.Println("--- Меню клієнтів ---")
		fmt.Println("1. Реєстрація клієнта")
		fmt.Println("2. Переглянути клієнтів")
		fmt.Println("3. Оновити дані клієнта")
		fmt.Println("4. Повернутися")
		fmt.Print("> ")
		var choice int
		fmt.Scan(&choice)
		fmt.Scanln()

		switch choice {
		case 1:
			addCustomer()
		case 2:
			listCustomers()
		case 3:
			updateCustomer()
		case 4:
			return
		default:
			fmt.Println("Невірний вибір.")
		}
	}
}

func addCustomer() {
	var name, email, phone, address string
	fmt.Print("Ім'я клієнта: ")
	fmt.Scanln(&name)
	fmt.Print("Email: ")
	fmt.Scanln(&email)
	fmt.Print("Телефон: ")
	fmt.Scanln(&phone)
	fmt.Print("Адреса: ")
	fmt.Scanln(&address)

	nextCustomerID++
	c := Customer{nextCustomerID, name, email, phone, address}
	customers = append(customers, c)
	fmt.Println("Клієнт зареєстрований.")
}

func listCustomers() {
	if len(customers) == 0 {
		fmt.Println("Список клієнтів порожній.")
		return
	}
	for _, c := range customers {
		fmt.Printf("ID: %d | %s | %s | %s\n", c.id, c.name, c.email, c.phone)
	}
}

func updateCustomer() {
	var id int
	fmt.Print("Введіть ID клієнта: ")
	fmt.Scan(&id)
	fmt.Scanln()

	for i := range customers {
		if customers[i].id == id {
			fmt.Print("Новий телефон: ")
			fmt.Scanln(&customers[i].phone)
			fmt.Print("Нова адреса: ")
			fmt.Scanln(&customers[i].address)
			fmt.Println("Дані оновлено.")
			return
		}
	}
	fmt.Println("Клієнта не знайдено.")
}

func menuCart() {
	fmt.Println("--- Меню кошика ---")
	fmt.Println("1. Додати товар до кошика")
	fmt.Println("2. Переглянути кошик")
	fmt.Println("3. Оформити замовлення")
	fmt.Println("4. Назад")
	fmt.Print("> ")
	var choice int
	fmt.Scan(&choice)
	fmt.Scanln()

	switch choice {
	case 1:
		addToCart()
	case 2:
		viewCart()
	case 3:
		checkout()
	case 4:
		return
	default:
		fmt.Println("Невірний вибір.")
	}
}

func addToCart() {
	var custID, prodID, qty int
	fmt.Print("Введіть ID клієнта: ")
	fmt.Scan(&custID)
	fmt.Scanln()
	fmt.Print("Введіть ID товару: ")
	fmt.Scan(&prodID)
	fmt.Scanln()
	fmt.Print("Кількість: ")
	fmt.Scan(&qty)
	fmt.Scanln()

	var customer Customer
	for _, c := range customers {
		if c.id == custID {
			customer = c
		}
	}

	var product Product
	for _, p := range products {
		if p.id == prodID {
			product = p
		}
	}

	item := CartItem{product, qty}
	carts = append(carts, Cart{customer, []CartItem{item}, 0})
	fmt.Println("Товар додано до кошика.")
}

func viewCart() {
	if len(carts) == 0 {
		fmt.Println("Кошик порожній.")
		return
	}
	for _, cart := range carts {
		fmt.Printf("Клієнт: %s\n", cart.customer.name)
		var sum float64
		for _, i := range cart.items {
			fmt.Printf("%s x%d - %.2f грн\n", i.product.name, i.amount, i.product.price*float64(i.amount))
			sum += i.product.price * float64(i.amount)
		}
		fmt.Printf("Загальна сума: %.2f грн\n", sum)
	}
}

func checkout() {
	var custID int
	fmt.Print("Введіть ID клієнта: ")
	fmt.Scan(&custID)
	fmt.Scanln()

	for _, cart := range carts {
		if cart.customer.id == custID {
			var sum float64
			for _, i := range cart.items {
				sum += i.product.price * float64(i.amount)
			}
			delivery := 150.0
			total := sum + delivery
			nextOrderID++
			orders = append(orders, Order{nextOrderID, cart.customer, cart.items, total, "pending"})
			fmt.Printf("Замовлення #%d створено. Сума: %.2f грн\n", nextOrderID, total)
			return
		}
	}
	fmt.Println("Кошик не знайдено.")
}

func menuOrders() {
	for {
		fmt.Println("--- Меню замовлень ---")
		fmt.Println("1. Переглянути всі замовлення")
		fmt.Println("2. Змінити статус")
		fmt.Println("3. Назад")
		fmt.Print("> ")
		var choice int
		fmt.Scan(&choice)
		fmt.Scanln()

		switch choice {
		case 1:
			listOrders()
		case 2:
			updateOrderStatus()
		case 3:
			return
		default:
			fmt.Println("Невірний вибір.")
		}
	}
}

func listOrders() {
	if len(orders) == 0 {
		fmt.Println("Немає замовлень.")
		return
	}
	for _, o := range orders {
		fmt.Printf("Замовлення #%d | Клієнт: %s | Сума: %.2f грн | Статус: %s\n",
			o.id, o.customer.name, o.total, o.status)
	}
}

func updateOrderStatus() {
	var id int
	var status string
	fmt.Print("Введіть ID замовлення: ")
	fmt.Scan(&id)
	fmt.Scanln()
	fmt.Print("Новий статус: ")
	fmt.Scanln(&status)

	for i := range orders {
		if orders[i].id == id {
			orders[i].status = status
			fmt.Println("Статус оновлено.")
			return
		}
	}
	fmt.Println("Замовлення не знайдено.")
}
