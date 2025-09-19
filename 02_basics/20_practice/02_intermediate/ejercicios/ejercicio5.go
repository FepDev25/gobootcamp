package ejercicios

import "fmt"

/*
EJERCICIO 5: Sistema de Calificaciones de Estudiantes
====================================================
Gestiona calificaciones de múltiples estudiantes y materias.

Instrucciones:
1. Define estructuras de datos usando maps anidados
2. Almacena: map[estudiante]map[materia][]calificaciones
3. Implementa funciones que retornen múltiples valores:
  - Agregar calificación para estudiante y materia
  - Calcular promedio de un estudiante en una materia
  - Calcular promedio general de un estudiante
  - Obtener el mejor estudiante por materia
  - Generar ranking general de estudiantes

Ejemplo:
Estudiantes: ["Ana", "Luis", "Maria"]
Materias: ["Matemáticas", "Historia", "Ciencias"]
Ana - Matemáticas: [85, 90, 92] → Promedio: 89.0
Luis - Matemáticas: [78, 82, 85] → Promedio: 81.7
Mejor en Matemáticas: Ana (89.0)
Ranking general: [Ana: 87.5, Luis: 83.2, Maria: 91.1]
*/
func Ejercicio5() {
	// TODO: Implementa aquí tu solución
	// Crea la estructura de datos y demuestra todas las operaciones
	sistemaCalificaciones := make(map[string]map[string][]int)

	estudiantes := []string{"Ana", "Luis", "Maria"}
	materias := []string{"Matemáticas", "Historia", "Ciencias"}

	fmt.Println("Estudiantes:", estudiantes)
	fmt.Println("Materias:", materias)

	agregarEstudiante(sistemaCalificaciones, estudiantes...)
	agregarMateriasAEstudiantes(sistemaCalificaciones, materias...)
	simularNotas(sistemaCalificaciones)

	estudiante, materia, notas, promedio := calcularPromedioEstudianteMateria(sistemaCalificaciones, estudiantes[0], materias[0])
	fmt.Printf("%s - %s: %d -> %0.2f\n", estudiante, materia, notas, promedio)

	estudiante2, materia2, notas, promedio := calcularPromedioEstudianteMateria(sistemaCalificaciones, estudiantes[1], materias[1])
	fmt.Printf("%s - %s: %d -> %0.2f\n", estudiante2, materia2, notas, promedio)

	estudiante3, materia3, notas, promedio := calcularPromedioEstudianteMateria(sistemaCalificaciones, estudiantes[2], materias[2])
	fmt.Printf("%s - %s: %d -> %0.2f\n", estudiante3, materia3, notas, promedio)

	rankingEstudiantes := obtenerRankingGeneral(sistemaCalificaciones)
	fmt.Print("Ranking general: [")
	i := 0
	for _, ranking := range rankingEstudiantes {
		for est, puntuacion := range ranking {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s: %.2f", est, puntuacion)
			i++
		}
	}
	fmt.Println("]")

	fmt.Println("Sistema completo:", sistemaCalificaciones)

}

func agregarEstudiante(sistema map[string]map[string][]int, estudiantes ...string) {
	for _, estudiante := range estudiantes {
		sistema[estudiante] = make(map[string][]int)
	}
}

func agregarMateriasAEstudiantes(sistema map[string]map[string][]int, materias ...string) {
	for _, materia := range materias {
		for alumno := range sistema {
			sistema[alumno][materia] = make([]int, 0)
		}
	}
}

func agregarCalificacionEstudianteMateria(sistema map[string]map[string][]int, estudiante string, materia string, nota int) {

	materiaEncontrada := obtenerNotasDeMateriaPorEstudiante(sistema, estudiante, materia)
	if materiaEncontrada == nil {
		return
	}

	materiaEncontrada = append(materiaEncontrada, nota)
	sistema[estudiante][materia] = materiaEncontrada
}

func calcularPromedioEstudianteMateria(sistema map[string]map[string][]int, estudiante string, materia string) (string, string, []int, float64) {
	materiaEncontrada := obtenerNotasDeMateriaPorEstudiante(sistema, estudiante, materia)
	if materiaEncontrada == nil {
		return "", "", nil, 0
	}

	if len(materiaEncontrada) == 0 {
		fmt.Println("Aun no se han registrado notas en la materia")
		return "", "", nil, 0
	}

	suma := 0
	for _, nota := range materiaEncontrada {
		suma += nota
	}

	return estudiante, materia, materiaEncontrada, float64(suma / len(materiaEncontrada))
}

func obtenerRankingGeneral(sistema map[string]map[string][]int) []map[string]float64 {
	ranking := make([]map[string]float64, 0)
	for estudiante, materias := range sistema {

		promedios := make([]float64, 0)

		for materia := range materias {
			_, _, _, promedio := calcularPromedioEstudianteMateria(sistema, estudiante, materia)
			promedios = append(promedios, promedio)
		}

		sumaPromedios := 0.0
		for _, promedio := range promedios {
			sumaPromedios += promedio
		}

		puntuacion := sumaPromedios / float64(len(promedios))

		estudianteRanking := map[string]float64{estudiante: puntuacion}
		ranking = append(ranking, estudianteRanking)
	}

	return ranking
}

func obtenerNotasDeMateriaPorEstudiante(sistema map[string]map[string][]int, estudiante string, materia string) []int {
	materias, okEstudiante := sistema[estudiante]
	if !okEstudiante {
		fmt.Println("Estudiante no encontrado")
		return nil
	}

	materiaEncontrada, okMateria := materias[materia]
	if !okMateria {
		fmt.Println("Materia no encontrada")
		return nil
	}

	return materiaEncontrada
}

func simularNotas(sistemaCalificaciones map[string]map[string][]int) {
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Ana", "Historia", 97)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Ana", "Historia", 100)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Ana", "Historia", 94)

	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Ana", "Matemáticas", 77)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Ana", "Matemáticas", 80)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Ana", "Matemáticas", 74)

	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Ana", "Ciencias", 100)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Ana", "Ciencias", 100)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Ana", "Ciencias", 100)

	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Luis", "Historia", 80)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Luis", "Historia", 84)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Luis", "Historia", 87)

	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Luis", "Matemáticas", 99)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Luis", "Matemáticas", 97)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Luis", "Matemáticas", 96)

	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Luis", "Ciencias", 90)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Luis", "Ciencias", 88)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Luis", "Ciencias", 81)

	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Maria", "Historia", 99)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Maria", "Historia", 99)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Maria", "Historia", 100)

	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Maria", "Matemáticas", 91)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Maria", "Matemáticas", 90)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Maria", "Matemáticas", 92)

	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Maria", "Ciencias", 87)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Maria", "Ciencias", 82)
	agregarCalificacionEstudianteMateria(sistemaCalificaciones, "Maria", "Ciencias", 85)
}
