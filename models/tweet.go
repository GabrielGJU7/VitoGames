package models

/*Tweet comentario*/
type Tweet struct {
	Mensaje string `bson: "mensaje" json:"mensaje"`
}
