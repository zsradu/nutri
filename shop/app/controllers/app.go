package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"shop/app"
	"shop/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) AddItemsFromApiToDatabase(queryParams string) {
	url := "https://edamam-food-and-grocery-database.p.rapidapi.com/parser?ingr=" + queryParams

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "'your token'")
	req.Header.Add("X-RapidAPI-Host", "edamam-food-and-grocery-database.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var APIresult map[string]interface{}
	err := json.Unmarshal(body, &APIresult)
	if err != nil {
		fmt.Println(err)
	}

	var foodItems []interface{}
	foodItems = APIresult["parsed"].([]interface{})

	for _, foodItem := range foodItems {
		var product models.Products
		product.Name = foodItem.(map[string]interface{})["food"].(map[string]interface{})["label"].(string)
		product.ImageName = foodItem.(map[string]interface{})["food"].(map[string]interface{})["image"].(string)
		product.Calories = foodItem.(map[string]interface{})["food"].(map[string]interface{})["nutrients"].(map[string]interface{})["ENERC_KCAL"].(float64)
		product.Protein = foodItem.(map[string]interface{})["food"].(map[string]interface{})["nutrients"].(map[string]interface{})["PROCNT"].(float64)
		product.Carbs = foodItem.(map[string]interface{})["food"].(map[string]interface{})["nutrients"].(map[string]interface{})["CHOCDF"].(float64)
		product.Fat = foodItem.(map[string]interface{})["food"].(map[string]interface{})["nutrients"].(map[string]interface{})["FAT"].(float64)
		app.DB.Create(&product)
	}

}

func (c App) InitDatabase() {
	c.AddItemsFromApiToDatabase("apple")
	c.AddItemsFromApiToDatabase("banana")
	c.AddItemsFromApiToDatabase("orange")
	c.AddItemsFromApiToDatabase("milk")
	c.AddItemsFromApiToDatabase("bread")
	c.AddItemsFromApiToDatabase("cheese")
	c.AddItemsFromApiToDatabase("chicken")
	c.AddItemsFromApiToDatabase("beef")
	c.AddItemsFromApiToDatabase("pork")
	c.AddItemsFromApiToDatabase("fish")
	c.AddItemsFromApiToDatabase("rice")
	c.AddItemsFromApiToDatabase("potato")
	c.AddItemsFromApiToDatabase("tomato")
	c.AddItemsFromApiToDatabase("onion")
	c.AddItemsFromApiToDatabase("garlic")
	c.AddItemsFromApiToDatabase("egg")
	c.AddItemsFromApiToDatabase("butter")
	c.AddItemsFromApiToDatabase("sugar")
	c.AddItemsFromApiToDatabase("salt")
	c.AddItemsFromApiToDatabase("pepper")
	c.AddItemsFromApiToDatabase("flour")
	c.AddItemsFromApiToDatabase("oil")
	c.AddItemsFromApiToDatabase("water")
	c.AddItemsFromApiToDatabase("coffee")
	c.AddItemsFromApiToDatabase("tea")
	c.AddItemsFromApiToDatabase("beer")
	c.AddItemsFromApiToDatabase("wine")
	c.AddItemsFromApiToDatabase("chocolate")
	c.AddItemsFromApiToDatabase("ice cream")
	c.AddItemsFromApiToDatabase("yogurt")
	c.AddItemsFromApiToDatabase("cucumber")
	c.AddItemsFromApiToDatabase("lettuce")
	c.AddItemsFromApiToDatabase("carrot")
	c.AddItemsFromApiToDatabase("broccoli")
	c.AddItemsFromApiToDatabase("peas")
	c.AddItemsFromApiToDatabase("beans")
	c.AddItemsFromApiToDatabase("corn")
	c.AddItemsFromApiToDatabase("pepper")
	c.AddItemsFromApiToDatabase("mushroom")
	c.AddItemsFromApiToDatabase("avocado")
	c.AddItemsFromApiToDatabase("olive")
	c.AddItemsFromApiToDatabase("strawberry")
	c.AddItemsFromApiToDatabase("blueberry")
	c.AddItemsFromApiToDatabase("raspberry")
	c.AddItemsFromApiToDatabase("lemon")
	c.AddItemsFromApiToDatabase("lime")
	c.AddItemsFromApiToDatabase("coconut")
	c.AddItemsFromApiToDatabase("cherry")
	c.AddItemsFromApiToDatabase("kiwi")
	c.AddItemsFromApiToDatabase("pineapple")
	c.AddItemsFromApiToDatabase("melon")
	c.AddItemsFromApiToDatabase("papaya")
	c.AddItemsFromApiToDatabase("peach")
	c.AddItemsFromApiToDatabase("apricot")
	c.AddItemsFromApiToDatabase("plum")
	c.AddItemsFromApiToDatabase("grape")
	c.AddItemsFromApiToDatabase("grapefruit")
	c.AddItemsFromApiToDatabase("pomegranate")
	c.AddItemsFromApiToDatabase("watermelon")
	c.AddItemsFromApiToDatabase("cabbage")
	c.AddItemsFromApiToDatabase("cauliflower")
	c.AddItemsFromApiToDatabase("spinach")
	c.AddItemsFromApiToDatabase("asparagus")
	c.AddItemsFromApiToDatabase("zucchini")
	c.AddItemsFromApiToDatabase("peanut butter")
	c.AddItemsFromApiToDatabase("honey")
	var products []models.Products
	app.DB.Find(&products)
	for _, product := range products {
		randomPrice := float32(rand.Float64() * 10)
		randomPrice = float32(math.Floor(float64(randomPrice*100)) / 100)
		product.Price = randomPrice
		app.DB.Save(&product)
	}
}

func (c App) Home() revel.Result {
	age := "25"
	gender := "male"
	weight := "80"
	height := "180"
	activity := "level_1"
	// TODO: Get this data from form
	url := "https://fitness-calculator.p.rapidapi.com/dailycalorie?" +
		"age=" + age +
		"&gender=" + gender +
		"&height=" + height +
		"&weight=" + weight +
		"&activitylevel=" + activity // 1 through 6
	/*
		‘level _ 1’ : “Sedentary: little or no exercise”,
		‘level _ 2’ : “Exercise 1-3 times/week”,
		‘level _ 3’ : “Exercise 4-5 times/week”,
		‘level _ 4’ : “Daily exercise or intense exercise 3-4 times/week”,
		‘level _ 5’ : “Intense exercise 6-7 times/week”,
		‘level _ 6’ : “Very intense exercise daily, or physical job”
	*/

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "'your token'")
	req.Header.Add("X-RapidAPI-Host", "fitness-calculator.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var APIresult map[string]interface{}
	err := json.Unmarshal(body, &APIresult)
	if err != nil {
		fmt.Println(err)
	}
	maintainWeight := APIresult["data"].(map[string]interface{})["goals"].(map[string]interface{})["maintain weight"].(float64)

	return c.Render(maintainWeight)
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Shop() revel.Result {
	//c.InitDatabase()
	var products []models.Products
	app.DB.Limit(30).Find(&products)
	return c.Render(products)
}

func (c App) AddToCart(id int) revel.Result {
	var product models.Products
	app.DB.First(&product, id)
	app.DB.Create(&models.Cart{ProductID: product.ID})
	return c.Redirect(App.Shop)
}

func (c App) Cart() revel.Result {
	var cart []models.Cart
	app.DB.Find(&cart)
	var products []models.Products
	var cartTotalData models.CartTotalData
	for _, item := range cart {
		var product models.Products
		app.DB.First(&product, item.ProductID)
		products = append(products, product)
		cartTotalData.Price += product.Price
		cartTotalData.Calories += product.Calories
		cartTotalData.Protein += product.Protein
		cartTotalData.Fat += product.Fat
		cartTotalData.Carbs += product.Carbs
	}

	var recipeAPIQuery string
	for _, product := range products {
		recipeAPIQuery += product.Name + " "
	}
	recipeObjects := EdamamAPIRecipes(recipeAPIQuery)

	return c.Render(products, recipeObjects, cartTotalData)
}

func EdamamAPIRecipes(queryParams string) []models.Recipes {

	url := "https://edamam-recipe-search.p.rapidapi.com/search?q=" + queryParams

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "'your token'")
	req.Header.Add("X-RapidAPI-Host", "edamam-recipe-search.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var APIresult map[string]interface{}
	err := json.Unmarshal(body, &APIresult)
	if err != nil {
		fmt.Println(err)
	}
	var recipes []interface{}
	recipes = APIresult["hits"].([]interface{})

	var recipeObjects []models.Recipes
	for _, recipe := range recipes {
		var recipeObject models.Recipes
		recipeObject.Name = recipe.(map[string]interface{})["recipe"].(map[string]interface{})["label"].(string)
		recipeObject.Image = recipe.(map[string]interface{})["recipe"].(map[string]interface{})["image"].(string)
		recipeObject.RecipeURL = recipe.(map[string]interface{})["recipe"].(map[string]interface{})["url"].(string)
		recipeObjects = append(recipeObjects, recipeObject)
	}
	return recipeObjects
}

func (c App) RemoveFromCart(key string) revel.Result {
	var product models.Products
	app.DB.Where("name = ?", key).First(&product)
	app.DB.Where("product_id = ?", product.ID).Delete(&models.Cart{})
	return c.Redirect(App.Cart)
}

func (c App) RecognizeFood() revel.Result {
	recognizedFoods := c.RecognizeImageApi()
	var recognizedItems []models.Recognize
	for _, food := range recognizedFoods {
		var recognizedItem models.Recognize
		recognizedItem.Name = food
		if food != "" {
			if !c.CheckIfFoodExists(food) {
				c.AddItemsFromApiToDatabase(food)
			}
		}
		id := c.GetFoodId(food)
		recognizedItem.Id = id
		recognizedItems = append(recognizedItems, recognizedItem)
	}
	return c.Render(recognizedItems)
}

func (c App) AddToCartFromRecognizeFood(id uint) revel.Result {
	var product models.Products
	app.DB.First(&product, id)
	app.DB.Create(&models.Cart{ProductID: product.ID})
	return c.Redirect(App.RecognizeFood)
}

func (c App) CheckIfFoodExists(food string) bool {
	var product models.Products
	app.DB.Where("name = ?", food).First(&product)
	if product.Name == food {
		return true
	}
	return false
}

func (c App) GetFoodId(food string) uint {
	var product models.Products
	app.DB.Where("name = ?", food).First(&product)
	return product.ID
}

func (c App) RecognizeImageApi() []string {
	url := "https://api.logmeal.es/v2/image/segmentation/complete/v1.0?language=eng"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("public/img/tomato-cucumber.jpg")
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("image", filepath.Base("public/img/tomato-cucumber.jpg"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return nil
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	bearerToken := "Bearer 'your token'"
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearerToken)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = os.WriteFile("app/scriptResponses/ingredients.json", body, 0644)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// [segmentation_results][0][recognition_results]
	// [0][name], [1][name], [2][name]
	var APIresult map[string]interface{}
	err = json.Unmarshal(body, &APIresult)
	if err != nil {
		fmt.Println(err)
	}
	var ingredients []interface{}
	ingredients = APIresult["segmentation_results"].([]interface{})[0].(map[string]interface{})["recognition_results"].([]interface{})
	var recognizedFoods []string
	recognizedFoods = append(recognizedFoods, ingredients[0].(map[string]interface{})["name"].(string))
	recognizedFoods = append(recognizedFoods, ingredients[1].(map[string]interface{})["name"].(string))
	recognizedFoods = append(recognizedFoods, ingredients[2].(map[string]interface{})["name"].(string))
	return recognizedFoods
}

func (c App) Form() revel.Result {
	return c.Render()
}

func (c App) Contact() revel.Result {
	return c.Render()
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) Checkout() revel.Result {
	return c.Render()
}

func (c App) Signup() revel.Result {
	return c.Render()
}

func (c App) Promocode() revel.Result {
	return c.Render()
}

func (c App) Thankyou() revel.Result {
	return c.Render()
}

func (c App) Deficit() revel.Result {
	var products []models.Products
	app.DB.Order("Calories asc").Limit(30).Find(&products)
	return c.Render(products)
}

func (c App) Surplus() revel.Result {
	var products []models.Products
	app.DB.Order("Calories desc").Limit(30).Find(&products)
	return c.Render(products)
}

func (c App) Healthier() revel.Result {
	var products []models.Products
	app.DB.Order("Protein desc").Limit(30).Find(&products)
	return c.Render(products)
}

func (c App) Formular() revel.Result {
	age := "25"
	gender := "male"
	weight := "80"
	height := "180"
	activity := "level_1"
	// TODO: Get this data from form
	url := "https://fitness-calculator.p.rapidapi.com/dailycalorie?" +
		"age=" + age +
		"&gender=" + gender +
		"&height=" + height +
		"&weight=" + weight +
		"&activitylevel=" + activity // 1 through 6
	/*
		‘level _ 1’ : “Sedentary: little or no exercise”,
		‘level _ 2’ : “Exercise 1-3 times/week”,
		‘level _ 3’ : “Exercise 4-5 times/week”,
		‘level _ 4’ : “Daily exercise or intense exercise 3-4 times/week”,
		‘level _ 5’ : “Intense exercise 6-7 times/week”,
		‘level _ 6’ : “Very intense exercise daily, or physical job”
	*/

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "'your token'")
	req.Header.Add("X-RapidAPI-Host", "fitness-calculator.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var APIresult map[string]interface{}
	err := json.Unmarshal(body, &APIresult)
	if err != nil {
		fmt.Println(err)
	}
	maintainWeight := APIresult["data"].(map[string]interface{})["goals"].(map[string]interface{})["maintain weight"].(float64)

	return c.Render(maintainWeight)
}
