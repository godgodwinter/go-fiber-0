package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Lakukan verifikasi atau pengecekan otentikasi di sini
	// Contoh: Periksa header atau token otentikasi

	// Jika otentikasi gagal, kembalikan response error
	if !isAuthenticated() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Jika otentikasi berhasil, lanjutkan ke handler berikutnya
	fmt.Println("Auth middleware executed")
	return c.Next()
}

// Fungsi sederhana untuk mendemonstrasikan verifikasi otentikasi
func isAuthenticated() bool {
	// Lakukan verifikasi otentikasi sesuai kebutuhan
	// Contoh: Periksa token atau database user

	// Return true jika otentikasi berhasil
	return true
}
