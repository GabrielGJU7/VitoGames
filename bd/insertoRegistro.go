package bd

import (
	"context"
	"time"

	"github.com/GabrielGJU7/VitoGames/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InseroRegistro para insertar usuarios*/

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("vitogames")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
