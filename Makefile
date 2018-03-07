NAME = github.com/lnxbil/dmidecode

build:
	@docker build --quiet --pull --tag $(NAME) .

test: build
	@docker run --privileged --tty --volume ${PWD}:/go $(NAME) go test -cover
