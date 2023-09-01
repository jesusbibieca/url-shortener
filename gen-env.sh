# generate env.yaml file in the current directory
# Usage: ./gen-env.sh

echo |
awk 'BEGIN {
    FS=": "
    print "REDIS_PORT:"
    print "REDIS_DB:"
    print "REDIS_ADDRESS:"
    print "APP_PORT:"
    print "APP_ADDRESS:"
    print "DB_DRIVER:"
    print "DB_SOURCE:"
}' > .env.yaml