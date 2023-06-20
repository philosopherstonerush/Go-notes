package pack2

// package is whatever the folder name is

func getName() string { // private function when used outside the pack2 package environment
	return "pack2 file"
}

func GetName() string { // public can be accessed by anyone
	return "pack2 file public"
}

// Lowercase letter for the name ---> private
// Uppercase letter for the name ---> public
