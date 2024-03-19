push:
	make test
	make format
	make lint
	git add -A
	./commit.sh
	git push --no-verify

test:
	echo "testing.."

format:
	echo "formatting.."

lint:
	echo "linting.."

kill:
	lsof -ti :8080 | xargs -r kill -9
