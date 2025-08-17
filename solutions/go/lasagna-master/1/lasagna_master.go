package lasagna

func PreparationTime(layers []string, averagePreparationTime int) int {
    if averagePreparationTime == 0 {
        averagePreparationTime = 2
    }
    return len(layers) * averagePreparationTime
}

func Quantities(layers []string) (int, float64) {
    const gramsPerNoodleLayer int = 50
    const litersPerSauceLayer float64 = 0.2
    var totalGramsNoodles int = 0
    var totalLitersSauce float64 = 0.0
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            totalGramsNoodles += gramsPerNoodleLayer
        }
        if layers[i] == "sauce" {
            totalLitersSauce += litersPerSauceLayer
        }
    }

    return totalGramsNoodles, totalLitersSauce
}

func AddSecretIngredient(ingredients1 []string, ingredients2 []string) {
    ingredients2[len(ingredients2)-1] = ingredients1[len(ingredients1) - 1] 
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
    totalPortions := make([]float64, len(quantities), cap(quantities))
    for i := 0; i < len(quantities); i++ {
        singlePortion := quantities[i] / 2
        totalPortions[i] = singlePortion * float64(numberOfPortions)
    }
    return totalPortions
}
