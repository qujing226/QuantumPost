package log

type Field struct {
	Key   string
	Value any
}

func Error(err error) Field {
	return Field{
		Key:   "error",
		Value: err,
	}
}

func String(key, value string) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}

func Int32(key string, num int32) Field {
	return Field{
		Key:   key,
		Value: num,
	}
}

func Int64(key string, num int64) Field {
	return Field{
		Key:   key,
		Value: num,
	}
}
