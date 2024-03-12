# Use a imagem oficial do MongoDB
FROM mongo:4.4

# Define a variável de ambiente para o diretório de dados do MongoDB
ENV MONGO_DATA_DIR /data/db

# Define a variável de ambiente para o diretório de logs do MongoDB
ENV MONGO_LOG_DIR /dev/null

# Expõe a porta padrão do MongoDB e a porta para conexões seguras
EXPOSE 27017

# Define o comando padrão para executar ao iniciar o contêiner
CMD ["mongod"]