.PHONY: tests
tests:
	@echo "--- Unitary tests ---"
	go test ./api -run=^Test_U_

	@echo "--- Functional tests ---"
	go test ./api -run=^Test_F_ -coverprofile=functional.out

.PHONY: clean
clean:
	rm functional.out
