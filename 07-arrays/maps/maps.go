package maps

import "fmt"

// what is map? key-value pairs.


func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	// unlike arrays, map is much flexible with adding more key-value pairs
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}