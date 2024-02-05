package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var jwtSecret = []byte("mysecretkey")                           //env
var db, Err = gorm.Open(sqlite.Open("site.db"), &gorm.Config{}) //env
var currentUser LoggedUser

func main() {

	if Err != nil {
		log.Fatal("Unable To Connect Database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Cart{})

	mux := http.NewServeMux()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	handler := cors.Handler(mux)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/login", login)
	mux.Handle("/profile", isAuthenticated(profile))

	mux.Handle("/products", isAuthenticated(allProducts))
	mux.Handle("/latest_products", isAuthenticated(latestProducts))
	mux.Handle("/create_product", isAuthenticated(createProduct))
	mux.Handle("/get_product/", isAuthenticated(getProduct))
	mux.Handle("/update_product", isAuthenticated(updateProduct))
	mux.Handle("/delete_product", isAuthenticated(deleteProduct))

	mux.Handle("/add_cart", isAuthenticated(addCart))
	mux.Handle("/all_cart", isAuthenticated(allCart))
	mux.Handle("/get_cart/", isAuthenticated(getCart))
	mux.Handle("/confirm_cart", isAuthenticated(confirmCart))
	mux.Handle("/delete_cart", isAuthenticated(deleteCart))
	mux.Handle("/all_confirmed_cart", isAuthenticated(allConfirmedCart))

	mux.Handle("/count", isAuthenticated(count))

	fmt.Println("Server Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var product []Product

	db.Find(&product)
	json.NewEncoder(w).Encode(product)
}

func profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	db.First(&user, "id == ?", currentUser.ID)
	json.NewEncoder(w).Encode(&user)
}

// Retrieve user list with eager loading credit card
// func GetAll(db *gorm.DB) ([]User, error) {
// 	var users []User
// 	err := db.Model(&User{}).Preload("CreditCard").Find(&users).Error
// 	return users, err
//   }

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.First(&user, "email = ?", user.Email)
	if result.RowsAffected != 0 {
		json.NewEncoder(w).Encode("User With This Email Already Exists")
		return
	}
	hw_password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return
	}
	u_result := db.Create(&User{
		ID:           uuid.NewString(),
		Name:         user.Name,
		Email:        user.Email,
		Password:     string(hw_password),
		ProfileImage: "default.jpg",
		Role:         user.Role,
	})
	if u_result.Error != nil {
		json.NewEncoder(w).Encode("Failed To Create User")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func login(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Email    string
		Password string
	}
	var body Body
	var user User

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	result := db.Where("email = ?", body.Email).First(&user)
	if result.RowsAffected == 0 {
		http.Error(w, "User Does Not Exist", http.StatusNotFound)
		return
	}
	n_err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if n_err != nil {
		// fmt.Println(n_err.Error())
		// http.Error(w, "Incorrect Password", http.StatusNotFound)
		return
	} else {
		fmt.Println("Hash Success")
	}
	// //Generate Jwt Token7
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		http.Error(w, "Failed to create token", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&tokenString)
}

func isAuthenticated(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Checking Auth Token")

		reqToken := r.Header.Get("Authorization")

		splitToken := strings.Split(reqToken, " ")
		if len(splitToken) != 2 {
			fmt.Printf("Error Reading Token")
			return
		}

		reqToken = strings.TrimSpace(splitToken[1])

		if reqToken != "" {
			// fmt.Println("Header is set! We can serve content!")

			token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(jwtSecret), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if float64(time.Now().Unix()) > claims["exp"].(float64) {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}

				var user User
				db.First(&user, "id = ?", claims["sub"])

				if user.ID == "" {
					http.Error(w, err.Error(), http.StatusUnauthorized)
				}
				currentUser = LoggedUser{ID: user.ID, Name: user.Name, Email: user.Email, Role: user.Role}
			} else {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}
			endpoint(w, r)
		} else {
			fmt.Println("Not Authorized")
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

// func profile(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user User
// 	db.First(&user, "id = ?", currentUser.ID)
// 	json.NewEncoder(w).Encode(user)
// }

// func updateProfile(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	r.ParseMultipartForm(10 * 1024 * 1024)

// 	// var user user
// 	var updateUser User

// 	file, handler, err := r.FormFile("profileImage")

// 	db.First(&updateUser, "id = ?", currentUser.ID)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer file.Close()
// 	fmt.Println("File Name: ", handler.Filename)
// 	fmt.Println("File Size: ", handler.Size)
// 	fmt.Println("File Name: ", handler.Header.Get("Content-Type"))

// 	temp, err2 := ioutil.TempFile("static", "user-*.jpg")
// 	if err2 != nil {
// 		fmt.Println(err2)
// 		return
// 	}
// 	defer temp.Close()

// 	fileBytes, err3 := ioutil.ReadAll(file)
// 	if err3 != nil {
// 		fmt.Println(err3)
// 		return
// 	}
// 	temp.Write(fileBytes)
// 	fmt.Println("Done")
// 	db.Model(&updateUser).Update("profile_image", temp.Name())
// 	json.NewEncoder(w).Encode(updateUser)
// }
