
echo "Esperando o banco de dados iniciar..."
sleep 5  


echo "Rodando migrations..."
tern migrate --config ./migrations/tern.conf --migrations ./migrations 
