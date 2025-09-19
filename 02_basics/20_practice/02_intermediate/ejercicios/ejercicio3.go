package ejercicios

import "strings"

/*
EJERCICIO 3: Análisis de Texto Avanzado
======================================
Procesa texto y genera estadísticas detalladas.

Instrucciones:
1. Recibe un texto (string)
2. Retorna múltiples valores con estadísticas:
  - Mapa de frecuencia de palabras
  - Mapa de frecuencia de caracteres (excluyendo espacios)
  - Palabra más común
  - Carácter más común
  - Número de oraciones (terminadas en ., !, ?)
  - Número de párrafos (separados por \n\n)

Ejemplo para texto="Hola mundo. ¿Cómo estás? ¡Muy bien!":
Frecuencia de palabras: map[Hola:1 mundo:1 Cómo:1 estás:1 Muy:1 bien:1]
Frecuencia de caracteres: map[H:1 o:4 l:2 a:2 ...]
Palabra más común: "mundo" (todas empatan)
Carácter más común: "o"
Oraciones: 3
Párrafos: 1
*/
func Ejercicio3(texto string) (map[string]int, map[rune]int, string, rune, int, int) {
	// TODO: Implementa aquí tu solución
	// Pista: usa strings.Fields(), strings.Count(), range con string

	// Obtener mapa de palabras
	frecuenciaPalabrasc := make(map[string]int)
	var textoSplitPlabras = texto
	textoSplitPlabras = strings.ReplaceAll(textoSplitPlabras, ".", "")
	textoSplitPlabras = strings.ReplaceAll(textoSplitPlabras, "?", "")
	textoSplitPlabras = strings.ReplaceAll(textoSplitPlabras, "¿", "")
	textoSplitPlabras = strings.ReplaceAll(textoSplitPlabras, "!", "")
	textoSplitPlabras = strings.ReplaceAll(textoSplitPlabras, "¡", "")
	textoSplitPlabras = strings.ReplaceAll(textoSplitPlabras, "\n", " ")
	slicePalabras := strings.SplitSeq(textoSplitPlabras, " ")
	for palabra := range slicePalabras {
		v, ok := frecuenciaPalabrasc[palabra]
		if ok {
			frecuenciaPalabrasc[palabra] = v + 1
		} else {
			frecuenciaPalabrasc[palabra] = 1
		}
	}

	// Obtener palabra m[as repetida
	var palabraComun string
	count := -1
	for k, v := range frecuenciaPalabrasc {
		if v > count {
			palabraComun = k
			count = v
		}
	}

	// Obtener mapa de caracteres
	frecuenciaCaracteres := make(map[rune]int)
	var textoSplitCarcacteres = texto
	textoSplitCarcacteres = strings.ReplaceAll(textoSplitCarcacteres, " ", "")
	for _, char := range textoSplitCarcacteres {
		v, ok := frecuenciaCaracteres[char]
		if ok {
			frecuenciaCaracteres[char] = v + 1
		} else {
			frecuenciaCaracteres[char] = 1
		}
	}

	// Obtener caracter mas comun
	var runeComun rune
	count = -1
	for k, v := range frecuenciaCaracteres {
		if v > count {
			runeComun = k
			count = v
		}
	}

	// Obtener numero de oraciones
	var oraciones = strings.Count(texto, ".") + strings.Count(texto, "?") + strings.Count(texto, "!")

	// Obtener numero de parrafos
	var parrafos = strings.Count(texto, "\n")

	return frecuenciaPalabrasc, frecuenciaCaracteres, palabraComun, runeComun, oraciones, parrafos
}
