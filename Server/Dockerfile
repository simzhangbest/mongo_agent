from ubuntu
RUN mkdir /home/zhangming/ \
	&& cd /home/zhangming/
COPY config.json .
COPY server .
EXPOSE 8088
ENTRYPOINT ["./server"]

