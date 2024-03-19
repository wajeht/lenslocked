push:
	# make test
	# make format
	# make lint
	git add -A
	./commit.sh
	git push --no-verify
