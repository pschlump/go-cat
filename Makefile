all:
	go build

test:
	go build
	./go-cat -o testdata/out1.out testdata/in1.txt testdata/in1.txt
	diff testdata/out1.out testdata/out1.ref

