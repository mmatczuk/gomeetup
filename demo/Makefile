SUBJECT = /C=PL/ST=Mazowieckie/L=Warsaw/O=CodiLime/OU=RnD

.PHONY: clean
clean:
	rm *.crt
	rm *.key

.PHONY: cert
cert:
	openssl req -x509 -nodes -newkey rsa:2048 -sha256 -subj "$(SUBJECT)" -keyout client.key -out client.crt
	openssl req -x509 -nodes -newkey rsa:2048 -sha256 -subj "$(SUBJECT)" -keyout server.key -out server.crt

.PHONY: deploy
deploy:
	@scp -i ~/.ssh/mmatczuk.pem server.* ec2-user@mmatczuk.hopto.org:
	@scp -i ~/.ssh/mmatczuk.pem tunneld ec2-user@mmatczuk.hopto.org:

.PHONY: ssh
ssh:
	ssh -i ~/.ssh/mmatczuk.pem ec2-user@mmatczuk.hopto.org
