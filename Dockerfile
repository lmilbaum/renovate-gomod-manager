FROM python

COPY renovate-gomod-manager/main.py ./

CMD [ "./main.py" ]