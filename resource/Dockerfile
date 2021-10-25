FROM python:2.7.18-slim-buster

WORKDIR /

COPY . .

RUN pip install --no-cache-dir -r requirements.txt

CMD [ "python", "./server.py" ]