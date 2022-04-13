package gym 

type BaseExercise interface {
	Create()
	Hydrate() (string, error)
}
