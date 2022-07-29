openssl req -new -sha256 \                                                                      slixperi@slixperi-desktop
    -newkey rsa:2048 \
    -subj "/C=RC/ST=Barcelona/O=pet2cattle/CN=pet2cattle-hook.webhookdemo.svc" \
    -nodes -x509 \
    -days 365 \
    -out server.crt \
    -addext "subjectAltName = DNS:pet2cattle-hook.webhookdemo.svc"