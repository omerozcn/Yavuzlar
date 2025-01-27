package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func VeriyiBul(html, baslangic, bitis string) []string {
	var sonuclar []string
	for {
		basIndex := strings.Index(html, baslangic)
		if basIndex == -1 {
			break
		}
		html = html[basIndex+len(baslangic):]
		bitIndex := strings.Index(html, bitis)
		if bitIndex == -1 {
			break
		}
		sonuclar = append(sonuclar, html[:bitIndex])
		html = html[bitIndex+len(bitis):]
	}
	return sonuclar
}

func TarihiTemizle(tarihHtml string) string {
	temizlenmisTarih := strings.ReplaceAll(tarihHtml, "<i class='icon-font icon-calendar'>&#59394;</i>", "")
	return strings.TrimSpace(temizlenmisTarih)
}

func ArstechnicaVerisiCekVeKaydet() error {
	url := "https://arstechnica.com/"
	fmt.Printf("Veri çekiliyor: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("URL'ye erişim sağlanamadı: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP hatası: %d", resp.StatusCode)
	}

	var htmlBuilder strings.Builder
	buf := bufio.NewScanner(resp.Body)
	for buf.Scan() {
		htmlBuilder.WriteString(buf.Text())
	}
	html := htmlBuilder.String()

	basliklar := VeriyiBul(html, `<h2 class="mb-2 mt-3 font-serif text-3xl font-bold leading-none md:text-4xl">`, `</h2>`)
	aciklamalar := VeriyiBul(html, `<p class="text-gray-550 dark:text-gray-250 dusk:text-gray-250 mb-2 text-lg leading-tight">`, `</p>`)
	tarihlerRaw := VeriyiBul(html, `<time`, `</time>`)
	links := VeriyiBul(html, `<a href="`, `">`)

	if len(basliklar) == 0 || len(aciklamalar) == 0 || len(tarihlerRaw) == 0 || len(links) == 0 {
		return fmt.Errorf("Veri çekilemedi, sonuçlar boş")
	}

	var temizBasliklar []string
	for _, baslik := range basliklar {
		baslik = strings.Split(baslik, `href="`)[1]
		baslik = strings.Split(baslik, `">`)[0]
		temizBasliklar = append(temizBasliklar, strings.TrimSpace(baslik))
	}

	var tarihler []string
	for _, tarihHtml := range tarihlerRaw {
		tarih := strings.Split(tarihHtml, `datetime="`)[1]
		tarih = strings.Split(tarih, `" `)[0]
		tarihler = append(tarihler, tarih)
	}

	dosyaAdi := fmt.Sprintf("arstechnica_veriler_%d.txt", time.Now().Unix())
	dosya, err := os.Create(dosyaAdi)
	if err != nil {
		return fmt.Errorf("Dosya oluşturulamadı: %v", err)
	}
	defer dosya.Close()

	writer := bufio.NewWriter(dosya)
	for i := 0; i < len(temizBasliklar) && i < len(aciklamalar) && i < len(tarihler) && i < len(links); i++ {
		link := strings.TrimSpace(strings.ReplaceAll(links[i], "\"", ""))

		cikti := fmt.Sprintf("Başlık: %s\nAçıklama: %s\nTarih: %s\nLink: %s\n\n", temizBasliklar[i], aciklamalar[i], tarihler[i], link)
		writer.WriteString(cikti)
	}
	writer.Flush()

	fmt.Printf("Veri başarıyla kaydedildi: %s\n", dosyaAdi)
	return nil
}

func VeriyiCekVeKaydet(url string) error {
	fmt.Printf("Veri çekiliyor: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("URL'ye erişim sağlanamadı: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP hatası: %d", resp.StatusCode)
	}

	var htmlBuilder strings.Builder
	buf := bufio.NewScanner(resp.Body)
	for buf.Scan() {
		htmlBuilder.WriteString(buf.Text())
	}
	html := htmlBuilder.String()

	basliklar := VeriyiBul(html, `<h2 class='home-title'>`, `</h2>`)
	aciklamalar := VeriyiBul(html, `<div class='home-desc'>`, `</div>`)
	tarihlerRaw := VeriyiBul(html, `<span class='h-datetime'>`, `</span>`)

	if len(basliklar) == 0 || len(aciklamalar) == 0 || len(tarihlerRaw) == 0 {
		return fmt.Errorf("Veri çekilemedi, sonuçlar boş")
	}

	var tarihler []string
	for _, tarihHtml := range tarihlerRaw {
		tarihler = append(tarihler, TarihiTemizle(tarihHtml))
	}

	dosyaAdi := fmt.Sprintf("thehackernews_veriler_%d.txt", time.Now().Unix())
	dosya, err := os.Create(dosyaAdi)
	if err != nil {
		return fmt.Errorf("Dosya oluşturulamadı: %v", err)
	}
	defer dosya.Close()

	writer := bufio.NewWriter(dosya)
	for i := 0; i < len(basliklar) && i < len(aciklamalar) && i < len(tarihler); i++ {
		cikti := fmt.Sprintf("Başlık: %s\nAçıklama: %s\nTarih: %s\n\n", basliklar[i], aciklamalar[i], tarihler[i])
		writer.WriteString(cikti)
	}
	writer.Flush()

	fmt.Printf("Veri başarıyla kaydedildi: %s\n", dosyaAdi)
	return nil
}

func AnaMenu() {
	fmt.Println("\n--- Web Veri Çekme Aracı ---")
	fmt.Println("1 - 'https://thehackernews.com/' sitesinden veri çek")
	fmt.Println("2 - 'https://arstechnica.com/' sitesinden veri çek")
	fmt.Println("3 - Çıkış yap")
	fmt.Print("Seçiminizi yapın: ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		AnaMenu()

		scanner.Scan()
		secim := strings.TrimSpace(scanner.Text())

		switch secim {
		case "1":
			err := VeriyiCekVeKaydet("https://thehackernews.com/")
			if err != nil {
				fmt.Printf("Hata: %v\n", err)
			}
		case "2":
			err := ArstechnicaVerisiCekVeKaydet()
			if err != nil {
				fmt.Printf("Hata: %v\n", err)
			}
		case "3":
			fmt.Println("Çıkılıyor... Görüşmek üzere!")
			return
		default:
			fmt.Println("Geçersiz seçim, lütfen tekrar deneyin.")
		}
	}
}
