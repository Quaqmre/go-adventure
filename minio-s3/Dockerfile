from minio/minio

ENV MINIO_ACCESS_KEY=testkey
ENV MINIO_SECRET_KEY=secretkey
VOLUME  "C:\\Users\\mehmetakif.tatar\\Desktop\\go-adventure\\minio-s3\\data" "/data"
EXPOSE 9000:9000
ENTRYPOINT [ "minio","server","/data" ]