(import "wordpress")
(import "memcached")
(import "mysql")
(import "haproxy")

(define (New zone nMemcache nSql nWordpress nHaproxy)
  (let ((memcd (memcached.New (+ zone "-memcd") nMemcache))
        (db (mysql.New (+ zone "-mysql") nSql))
        (wp (wordpress.New (+ zone "-wp") nWordpress db memcd))
        (hap (haproxy.New (+ zone "-hap") nHaproxy wp)))
    hap))
