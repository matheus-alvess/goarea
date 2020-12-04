package goarea // por padrão o nome do pacote é o mesmo nome da pasta em que esta
// contido
import "math"

// PI é uma proporcao numerica
const PI = math.Pi

// Circ calcula a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

// Rect calcula a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}

