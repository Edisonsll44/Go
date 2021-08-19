package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

//Funcion de conexion
func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := ""
	Nombre := "sistema_go"
	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

//Funcion principal
func main() {

	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)
	http.HandleFunc("/eliminar", Eliminar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/actualizar", Actualizar)
	log.Println("Servidor corriendo......")
	http.ListenAndServe(":8080", nil)
}

//Creacion de estructura
type Empleado struct {
	Id       int
	Nombre   string
	Apellido string
	Correo   string
	Foto     string
}

//Funcion de lectura
func Inicio(w http.ResponseWriter, r *http.Request) {
	// var estructuraEmpleado modelos.Empleado
	conexionEstablecida := conexionBD()
	registros, err := conexionEstablecida.Query("SELECT * FROM empleado")
	if err != nil {
		panic(err.Error())
	}
	empleado := Empleado{}
	arregloEmpleados := []Empleado{}

	for registros.Next() {
		var id int
		var nombre, apellido, correo, foto string
		err := registros.Scan(&id, &nombre, &apellido, &correo, &foto)

		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Apellido = apellido
		empleado.Correo = correo
		empleado.Foto = foto
		arregloEmpleados = append(arregloEmpleados, empleado)
	}
	plantillas.ExecuteTemplate(w, "inicio", arregloEmpleados)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	conexionEstablecida := conexionBD()
	registro, err := conexionEstablecida.Query("SELECT * FROM empleado WHERE id = ?", idEmpleado)
	if err != nil {
		panic(err.Error())
	}
	empleado := Empleado{}
	for registro.Next() {
		var id int
		var nombre, apellido, correo, foto string
		err := registro.Scan(&id, &nombre, &apellido, &correo, &foto)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Apellido = apellido
		empleado.Correo = correo
		empleado.Foto = foto
	}
	plantillas.ExecuteTemplate(w, "editar", empleado)
}

func Eliminar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	conexionEstablecida := conexionBD()
	borrarRegistro, err := conexionEstablecida.Prepare("DELETE FROM empleado WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	borrarRegistro.Exec(idEmpleado)
	http.Redirect(w, r, "/", 301)
}

func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		apellido := r.FormValue("apellido")
		correo := r.FormValue("correo")
		conexionEstablecida := conexionBD()
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleado(nombre, apellido,correo) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, apellido, correo)
		http.Redirect(w, r, "/", 301)
	}
}

func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		apellido := r.FormValue("apellido")
		correo := r.FormValue("correo")
		conexionEstablecida := conexionBD()
		actualizarRegistros, err := conexionEstablecida.Prepare("UPDATE empleado SET nombre=?, apellido=?,correo=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		actualizarRegistros.Exec(nombre, apellido, correo,id)
		http.Redirect(w, r, "/", 301)
	}
}