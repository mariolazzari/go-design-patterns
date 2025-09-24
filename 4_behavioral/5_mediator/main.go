package main

import "fmt"

// Component Interface

type Package interface {
	ship()
	deliver()
	allowShipping()
}

// Mediator interface
type Mediator interface {
	canShip(Package) bool
	notifyAboutDelivery()
}

// Concrete Component A
type UPSPackage struct {
	mediator Mediator
}

func (u *UPSPackage) ship() {
	if !u.mediator.canShip(u) {
		fmt.Println("UPS Package Shipping blocked...waiting")
		return
	}
	fmt.Println("UPS Package Shipped")
}

func (u *UPSPackage) deliver() {
	fmt.Println("Delivering UPS package")
	u.mediator.notifyAboutDelivery()
}

func (u *UPSPackage) allowShipping() {
	fmt.Println("UPS Package: Ready to ship")
	u.ship()
}

// Concrete Component B
type FedExPackage struct {
	mediator Mediator
}

func (f *FedExPackage) ship() {
	if !f.mediator.canShip(f) {
		fmt.Println("FedEx Shipping blocked...waiting..")
		return
	}
	fmt.Println("FedEx Package Shipped")
}

func (f *FedExPackage) deliver() {
	fmt.Println("Delivering FedEx package")
	f.mediator.notifyAboutDelivery()
}

func (f *FedExPackage) allowShipping() {
	fmt.Println("FedEx Package: Ready to ship")
	f.ship()
}

// Concrete Mediator
type PackageManager struct {
	isPackagePacked bool
	packages        []Package
}

func newPackageManger() *PackageManager {
	return &PackageManager{
		isPackagePacked: true,
	}
}

func (pm *PackageManager) canShip(p Package) bool {
	if pm.isPackagePacked {
		pm.isPackagePacked = false
		return true
	}
	pm.packages = append(pm.packages, p)
	return false
}

func (pm *PackageManager) notifyAboutDelivery() {
	if !pm.isPackagePacked {
		pm.isPackagePacked = true
	}

	if len(pm.packages) > 0 {
		firstPackage := pm.packages[0]
		pm.packages = pm.packages[1:]
		firstPackage.allowShipping()
	}

}

func main() {
	packageManager := newPackageManger()

	upsPackage := &UPSPackage{
		mediator: packageManager,
	}
	fedExPackage := &FedExPackage{
		mediator: packageManager,
	}

	upsPackage.ship()
	fedExPackage.ship()
	upsPackage.deliver()
}
