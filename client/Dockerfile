FROM python:3.10-slim-buster

ADD requirements.txt /
RUN pip install -r requirements.txt

ADD . .

CMD ["python", "./client.py"] 
