connect-db:
	docker compose exec db psql localhost -U app -d subscriptions
