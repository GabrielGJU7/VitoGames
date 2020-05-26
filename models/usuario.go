package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es la tabla de registro de usuarios*/

type Usuario struct {
	ID              primitive.ObjectID `bson: "_id, omitempty" json: "id"`
	Nombre          string             `bson: "nombre" json: "nombre, omitempty"`
	Apellidos       string             `bson: "apellidos" json: "apellidos, omitempty"`
	FechaNacimiento time.Time          `bson: "fechaNacimiento" json: "fechaNacimiento, omitempty"`
	Email           string             `bson: "email" json: "email"`
	Password        string             `bson: "password" json: "password, omitempty"`
	Avatar          string             `bson: "avatar" json: "avatar, omitempty"`
	Banner          string             `bson: "banner" json: "banner, omitempty"`
	biografia       string             `bson: "biografia" json: "biografia, omitempty"`
	ubicacion       string             `bson: "ubicacion" json: "ubicacion, omitempty"`
	sitioWeb        string             `bson: "sitioWeb" json: "sitioWeb, omitempty"`
}
