// El paquete main es el paquete predeterminado para ejecutar el programa.
package main

// Importamos los paquetes necesarios.
import (
	"context"
	"encoding/json" // Para decodificar la respuesta JSON de la API.
	"fmt"           // Para imprimir en la consola.
	"log"
	"net/http" // Para hacer la solicitud HTTP a la API.
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func ExampleClient_GetCompletions() {
	azureOpenAIKey := ""
	modelDeployment := "gpt-35-turbo-instruct" // ex: gpt-4-vision-preview"

	// Ex: "https://<your-azure-openai-host>.openai.azure.com"
	azureOpenAIEndpoint := "https://poligpt.upv.es/gpt/ufasu-ia"

	if azureOpenAIKey == "" || modelDeployment == "" || azureOpenAIEndpoint == "" {
		fmt.Fprintf(os.Stderr, "Skipping example, environment variables missing\n")
		return
	}

	keyCredential := azcore.NewKeyCredential(azureOpenAIKey)

	// In Azure OpenAI you must deploy a model before you can use it in your client. For more information
	// see here: https://learn.microsoft.com/azure/cognitive-services/openai/how-to/create-resource
	client, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, nil)

	if err != nil {
		//  TODO: Update the following line with your application specific error handling logic
		log.Fatalf("ERROR: %s", err)
	}

	resp, err := client.GetCompletions(context.TODO(), azopenai.CompletionsOptions{
		Prompt:         []string{"Cuentame sobre la historia de la IA"},
		MaxTokens:      to.Ptr(int32(500)),
		Temperature:    to.Ptr(float32(0.0)),
		DeploymentName: &modelDeployment,
	}, nil)

	if err != nil {
		//  TODO: Update the following line with your application specific error handling logic
		log.Fatalf("ERROR: %s", err)
	}

	for _, choice := range resp.Choices {
		fmt.Fprintf(os.Stderr, "Resultado: %s\n", *choice.Text)
	}

	// Output:
}

func main() {
	ExampleClient_GetCompletions()
}

// WeatherResponse es la estructura que coincide con la respuesta JSON de la API.
type WeatherResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"` // Solo nos interesa la temperatura en Celsius.
	} `json:"current"`
}

// La función main es el punto de entrada del programa.
/*func main() {
	// Llamamos a la función obtenerClima para diferentes ciudades.
	obtenerClima("Jaen", "Spain")
	obtenerClima("Madrid", "Spain")
	obtenerClima("Valencia", "Spain")

}*/

// obtenerClima hace una solicitud a la API del clima y muestra la temperatura en Celsius.
func obtenerClima(ciudad, codigoPais string) {
	// Formateamos la URL de la API con la ciudad y el código del país.
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=018eb411b2bc419bb18111159241604&q=%s,%s", ciudad, codigoPais)

	// Hacemos la solicitud HTTP.
	respuesta, err := http.Get(url)
	if err != nil {
		// Si hay un error, lo imprimimos y salimos de la función.
		fmt.Println("Error al obtener el clima:", err)
		return
	}
	// Nos aseguramos de que el cuerpo de la respuesta se cierre al final de la función.
	defer respuesta.Body.Close()

	// Decodificamos la respuesta en la estructura WeatherResponse.
	var weatherResponse WeatherResponse
	err = json.NewDecoder(respuesta.Body).Decode(&weatherResponse)
	if err != nil {
		// Si hay un error, lo imprimimos y salimos de la función.
		fmt.Println("Error al decodificar la respuesta:", err)
		return
	}

	// Imprimimos la temperatura en Celsius.
	fmt.Printf("La temperatura actual en %s, %s es %.2f grados Celsius.\n", ciudad, codigoPais, weatherResponse.Current.TempC)
}
