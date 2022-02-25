package main

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

func main() {
	pci, err := ghw.PCI()
	if err != nil {
		fmt.Printf("Error getting PCI info: %v", err)
	}
	fmt.Printf("host PCI devices:\n")
	fmt.Println("====================================================")

	for _, device := range pci.Devices {
		vendor := device.Vendor
		vendorName := vendor.Name
		if len(vendor.Name) > 20 {
			vendorName = string([]byte(vendorName)[0:17]) + "..."
		}
		product := device.Product
		productName := product.Name
		if len(product.Name) > 40 {
			productName = string([]byte(productName)[0:37]) + "..."
		}
		fmt.Printf("%-12s\t%-20s\t%-40s\n", device.Address, vendorName, productName)
	}
}
