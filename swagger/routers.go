package swagger

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"AddPet",
		"POST",
		"/v1/pet",
		AddPet,
	},

	Route{
		"DeletePet",
		"DELETE",
		"/v1/pet/{petId}",
		DeletePet,
	},

	Route{
		"FindPetsByStatus",
		"GET",
		"/v1/pet/findByStatus",
		FindPetsByStatus,
	},

	Route{
		"FindPetsByTags",
		"GET",
		"/v1/pet/findByTags",
		FindPetsByTags,
	},

	Route{
		"GetPetById",
		"GET",
		"/v1/pet/{petId}",
		GetPetById,
	},

	Route{
		"UpdatePet",
		"PUT",
		"/v1/pet",
		UpdatePet,
	},

	Route{
		"UpdatePetWithForm",
		"POST",
		"/v1/pet/{petId}",
		UpdatePetWithForm,
	},

	Route{
		"UploadFile",
		"POST",
		"/v1/pet/{petId}/uploadImage",
		UploadFile,
	},

	Route{
		"DeleteOrder",
		"DELETE",
		"/v1/store/order/{orderId}",
		DeleteOrder,
	},

	Route{
		"GetInventory",
		"GET",
		"/v1/store/inventory",
		GetInventory,
	},

	Route{
		"GetOrderById",
		"GET",
		"/v1/store/order/{orderId}",
		GetOrderById,
	},

	Route{
		"PlaceOrder",
		"POST",
		"/v1/store/order",
		PlaceOrder,
	},

	Route{
		"CreateUser",
		"POST",
		"/v1/user",
		CreateUser,
	},

	Route{
		"CreateUsersWithArrayInput",
		"POST",
		"/v1/user/createWithArray",
		CreateUsersWithArrayInput,
	},

	Route{
		"CreateUsersWithListInput",
		"POST",
		"/v1/user/createWithList",
		CreateUsersWithListInput,
	},

	Route{
		"DeleteUser",
		"DELETE",
		"/v1/user/{username}",
		DeleteUser,
	},

	Route{
		"GetUserByName",
		"GET",
		"/v1/user/{username}",
		GetUserByName,
	},

	Route{
		"LoginUser",
		"GET",
		"/v1/user/login",
		LoginUser,
	},

	Route{
		"LogoutUser",
		"GET",
		"/v1/user/logout",
		LogoutUser,
	},

	Route{
		"UpdateUser",
		"PUT",
		"/v1/user/{username}",
		UpdateUser,
	},

	Route{
		"SwaggerUI",
		"GET",
		"/v1/swagger",
		func(w http.ResponseWriter, r *http.Request) {
			pwd, err := os.Getwd()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			http.ServeFile(w, r, filepath.Join(pwd, "_swagger", "index.html"))
		},
	},

	Route{
		"SwaggerDefinision",
		"GET",
		"/v1/swagger/{filename}",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("GET Swagger")
			filename := mux.Vars(r)["filename"]
			basename := filepath.Base(filename)
			pwd, err := os.Getwd()
			if err != nil {
				log.Printf("failed to get pwd %v", err)
				http.Error(w, err.Error(), 500)
				return
			}
			f, err := os.Open(filepath.Join(pwd, "api", basename))
			if err != nil {
				log.Printf("failed to open %v", err)
				http.Error(w, err.Error(), 500)
				return
			}
			defer f.Close()

			w.Header().Set("Content-Type", "application/json; charset=utf8")
			if _, err := io.Copy(w, f); err != nil {
				log.Printf("failed to copy %v", err)
				http.Error(w, err.Error(), 500)
			}
		},
	},

	Route{
		"StaticFile",
		"GET",
		"/static/{filename}",
		func(w http.ResponseWriter, r *http.Request) {
			filename := mux.Vars(r)["filename"]
			basename := filepath.Base(filename)
			pwd, err := os.Getwd()
			if err != nil {
				log.Printf("failed to get pwd %v", err)
				http.Error(w, err.Error(), 500)
				return
			}
			http.ServeFile(w, r, filepath.Join(pwd, "_swagger", basename))
		},
	},
}
