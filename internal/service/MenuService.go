package service

func MenuService() (map[string]string, map[string]string) {
	text := map[string]string{
		"Message": "Вітаю!",
	}

	buttons := map[string]string{
		"Магазин":      "shop",
		"Чат":          "https://t.me/StarMafiaUA",
		"Правила чату": "https://telegra.ph/STAR-MAFIA-07-29",
		"Команди":      "view_commands",
	}

	return text, buttons
}
