package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Kullanici struct {
	Isim     string
	Sifre    string
	Rol      string
}

var kullanicilar = []Kullanici{
	{"yonetici", "123", "admin"}, //kullanıcı adı yonetici
	{"misafir", "321", "kullanici"}, //kullanıcı adı misafir
}

func logYaz(logMesaj string) {
	f, hata := os.OpenFile("kayitlar.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if hata != nil {
		fmt.Println("Log dosyasına yazılamadı:", hata)
		return
	}
	defer f.Close()

	t := time.Now().Format("02-01-2006 15:04:05")
	f.WriteString(t + " -> " + logMesaj + "\n")
}

func kullaniciDogrula(ad, sifre string) *Kullanici {
	for i := 0; i < len(kullanicilar); i++ {
		if kullanicilar[i].Isim == ad && kullanicilar[i].Sifre == sifre {
			return &kullanicilar[i]
		}
	}
	return nil
}

func adminPaneli() {
	giris := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Admin Paneli")
		fmt.Println("1 - Yeni Kullanıcı Ekle")
		fmt.Println("2 - Kullanıcı Sil")
		fmt.Println("3 - Logları Gör")
		fmt.Println("4 - Çıkış")
		fmt.Print("Seçiminiz: ")
		islem, _ := giris.ReadString('\n')
		islem = strings.TrimSpace(islem)

		if islem == "1" {
			fmt.Print("Kullanıcı adı: ")
			isim, _ := giris.ReadString('\n')
			isim = strings.TrimSpace(isim)

			fmt.Print("Şifre: ")
			sifre, _ := giris.ReadString('\n')
			sifre = strings.TrimSpace(sifre)

			kullanicilar = append(kullanicilar, Kullanici{Isim: isim, Sifre: sifre, Rol: "kullanici"})
			fmt.Println("Yeni kullanıcı eklendi.")
			logYaz("Yeni kullanıcı eklendi: " + isim)

		} else if islem == "2" {
			fmt.Print("Silinecek kullanıcı adı: ")
			isim, _ := giris.ReadString('\n')
			isim = strings.TrimSpace(isim)

			bulundu := false
			for i := 0; i < len(kullanicilar); i++ {
				if kullanicilar[i].Isim == isim && kullanicilar[i].Rol == "kullanici" {
					kullanicilar = append(kullanicilar[:i], kullanicilar[i+1:]...)
					bulundu = true
					fmt.Println("Kullanıcı silindi.")
					logYaz("Kullanıcı silindi: " + isim)
					break
				}
			}
			if !bulundu {
				fmt.Println("Kullanıcı bulunamadı.")
			}

		} else if islem == "3" {
			veriler, hata := os.ReadFile("kayitlar.txt")
			if hata != nil {
				fmt.Println("Log okunamıyor:", hata)
			} else {
				fmt.Println("Loglar:\n", string(veriler))
			}

		} else if islem == "4" {
			break

		} else {
			fmt.Println("Geçersiz işlem.")
		}
	}
}

func kullaniciPaneli(isim string) {
	giris := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Kullanıcı Menüsü")
		fmt.Println("1 - Profil Gör")
		fmt.Println("2 - Şifre Değiştir")
		fmt.Println("3 - Çıkış")
		fmt.Print("Seçiminiz: ")
		islem, _ := giris.ReadString('\n')
		islem = strings.TrimSpace(islem)

		if islem == "1" {
			fmt.Println("Kullanıcı:", isim)
			logYaz("Kullanıcı profiline bakıldı: " + isim)

		} else if islem == "2" {
			fmt.Print("Yeni şifre: ")
			yeniSifre, _ := giris.ReadString('\n')
			yeniSifre = strings.TrimSpace(yeniSifre)

			for i := 0; i < len(kullanicilar); i++ {
				if kullanicilar[i].Isim == isim {
					kullanicilar[i].Sifre = yeniSifre
					fmt.Println("Şifre değiştirildi.")
					logYaz("Şifre değiştirildi: " + isim)
					break
				}
			}

		} else if islem == "3" {
			break

		} else {
			fmt.Println("Geçersiz işlem.")
		}
	}
}

func main() {
	giris := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1 - Admin Girişi")
		fmt.Println("2 - Kullanıcı Girişi")
		fmt.Print("Seçiminiz: ")
		tip, _ := giris.ReadString('\n')
		tip = strings.TrimSpace(tip)

		fmt.Print("Kullanıcı adı: ")
		isim, _ := giris.ReadString('\n')
		isim = strings.TrimSpace(isim)

		fmt.Print("Şifre: ")
		sifre, _ := giris.ReadString('\n')
		sifre = strings.TrimSpace(sifre)

		k := kullaniciDogrula(isim, sifre)
		if k != nil {
			logYaz("Giriş yapıldı: " + isim)
			if tip == "1" && k.Rol == "admin" {
				adminPaneli()
			} else if tip == "2" && k.Rol == "kullanici" {
				kullaniciPaneli(isim)
			} else {
				fmt.Println("Yanlış giriş tipi.")
			}
		} else {
			fmt.Println("Hatalı giriş.")
			logYaz("Başarısız giriş denemesi: " + isim)
		}
	}
}
