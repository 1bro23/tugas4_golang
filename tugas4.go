package main

import "fmt"

var mahasiswa =  map[string]int{"Aldo":182,"Yosep":178}

func main(){
  tampil_mahasiswa()
}

func tampil_mahasiswa(){
  fmt.Println("Aldo :",mahasiswa["Aldo"],"cm")
  fmt.Println("Yosep :",mahasiswa["Yosep"],"cm")
}
