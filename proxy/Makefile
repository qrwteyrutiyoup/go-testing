include make.conf

all:
	cd src; make
	$(GOCC) main.go
	$(GOLD) main.$(MAGIC)
clean:
	cd src; make clean
	rm -f main.$(MAGIC)
	rm -f $(MAGIC).out
