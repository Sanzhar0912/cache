Методы
NewCache() *Cache
Функция NewCache() создает и возвращает новый экземпляр кэша.

SetCache(key string, value interface{})
Метод SetCache() устанавливает значение value для указанного key в кэше.

GetCache(key string) interface{}
Метод GetCache() возвращает значение, связанное с указанным key из кэша. Если такого ключа нет, возвращается nil.

DeleteCache(key string)
Метод DeleteCache() удаляет значение, связанное с указанным key из кэша.