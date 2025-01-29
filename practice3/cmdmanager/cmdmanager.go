package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmdm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please Enter Your Prices. Confirm every price with Enter. Input 0 to finish.")

	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmdm CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
