
all: ff-wallet

depend:
	git submodule update --init --recursive
	make -C extern/filecoin-ffi all

clean:
	rm ff-wallet
	make -C extern/filecoin-ffi clean


ff-wallet: depend
	go build ./
