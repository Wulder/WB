package main

import (
	"encoding/json"
	"fmt"
)

type OzonServer struct { //свой тип сервера у Озона (отличается от wildberries)
	name string
}

func (s *OzonServer) Get(request string) string { //Озон так же решил сделать своему серверу метод Get() для получения данных с него, но данный метод не соответствует по сигнатуре методу из интерфейса GetAds с которым работает наш рассыльщие сообщений, поэтому в дальнейшем нам придется писать адаптер к нему
	//будем использовать данный метод для обращения к api озона и получать от него все новости об акциях (опять же в формате, который отличается от того, который используется в wildberries)
	result := fmt.Sprintf(`{"saleName":"Распродажа кроссовков!","server":"%v"}`, s.name) //ответ с сервера

	return result
}

type OzonServerAdapter struct { //наш адаптер, реализует интерфейс который нам необходим
	Server *OzonServer
}

func (a *OzonServerAdapter) Get() Ads { //самая важная часть, имея выходные данные другого типа, мы адаптируем их под подходящий для нас формат и возвращаем с помощью интерфейса

	result := a.Server.Get("Дайте мне текущие акции!!!") //получаем данные сервера через апи озона

	data := make(map[string]string)
	json.Unmarshal([]byte(result), &data)

	ads := Ads{Store: "OZON", Name: data["saleName"], Description: "Какая-то распродажа с озона", Reg: data["server"]}
	return ads
}
