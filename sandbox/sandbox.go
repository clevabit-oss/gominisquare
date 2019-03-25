package sandbox

type ObjectBuilder interface {

}

type Sandbox interface {
	NewObject(func(builder ObjectBuilder))
}
