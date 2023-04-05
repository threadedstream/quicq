package queue

type Queue interface {
	Put(val any)
	Get() any
	Purge()
	Remove()
	Len() int
}

type LLQ struct {
	val  any
	len  int
	prev *LLQ
}

func New(val any) *LLQ {
	q := new(LLQ)
	q.val = val
	q.len += 1
	return q
}

func Put(q **LLQ, val any) {
	cur := new(LLQ)
	cur.val = val
	oldCur := *q
	(*q) = cur
	(*q).prev = oldCur
	(*q).len += oldCur.len + 1
}

func (q *LLQ) Get() any {
	return q.val
}

func Remove(q **LLQ) {
	newCur := (*q).prev
	(*q) = newCur
}

func Purge(q **LLQ) {
	for (*q).prev != nil {
		Remove(q)
	}
	(*q).val = nil
	(*q).len = 0
}

func (q *LLQ) Purge() {
	for q.prev != nil {
		Remove(&q)
	}
}

func (q *LLQ) Len() int {
	return q.len
}
