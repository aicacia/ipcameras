package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

var languageCodeKey = "languageCode"
var countryCodeKey = "countryCode"

var defaultLanguageCode = "en"

func LocaleMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		languageCode, countryCode := GetLanguageAndCountryCodeFromContext(c)
		c.Locals(languageCodeKey, languageCode)
		c.Locals(countryCodeKey, countryCode)
		return c.Next()
	}
}

func GetLanguageAndCountryCodeFromContext(c *fiber.Ctx) (string, string) {
	locale := strings.TrimSpace(c.Get("X-Locale"))
	if len(locale) != 0 {
		localeParts := strings.SplitN(locale, "-", 2)
		switch len(localeParts) {
		case 1:
			return localeParts[0], ""
		case 2:
			return localeParts[0], localeParts[1]
		}
	}
	return defaultLanguageCode, ""
}
