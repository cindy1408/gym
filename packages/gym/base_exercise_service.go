package gym 

type BaseExercise interface{
	Create()
	Hydrate() error
}
