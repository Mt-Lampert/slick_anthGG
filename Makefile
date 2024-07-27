
gen_templ:
	templ generate

dev: gen_templ
	go run app/*

# vim: ts=4 sw=4 fdm=indent
