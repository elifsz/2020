(define (doitHelper n doitPartial)
	(if ( = n 0)
		doitPartial
		(doitHelper(- n 1) (+ n doitPartial)))
)
(define (doit n)
	(doitHelper n 0))