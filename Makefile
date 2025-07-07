run:
	docker build --cache-from=ex -f Dockerfile -t ex . && docker run -e PORT=5555 -p 5555:5555 ex
