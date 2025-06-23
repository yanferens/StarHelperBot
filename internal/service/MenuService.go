package service

func MenuService() (map[string]string, map[string]string) {
	text := map[string]string{
		"Message": "Вітаю!",
	}

	buttons := map[string]string{
		"Магазин":    "shop",
		"Чат": "link_to_chat",
		"Правила чату":   "link_to_chat_rules",
		"Команди": "view_commands",
	}

	return text, buttons
}
