package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//мы захотели сделать рассылку об акция на Wildberries подписчикам
func main() {

	var servers []GetAds //все опрашиваемые сервера. слайс серверов, реализующих метод GetAds который возвапащает нам ответ с сервера в виде объявления об акциях

	//создадим сервера wildberries
	s1 := AdsServer{name: "Wildberries 1", region: "Moscow"}
	s2 := AdsServer{name: "Wildberries 2", region: "Kazan"}
	s3 := AdsServer{name: "Wildberries 3", region: "Rostov"}
	servers = append(servers, &s1, &s2, &s3)

	for i := 0; i < 5; i++ {
		//каждую 1 секунду делаем запрос на сервера и отправляем ответы (информацию об акциях) подписчикам

		broadcast(servers...)
		time.Sleep(time.Second * 1)
	}

	fmt.Println("Выкупили озон и подключили их сервера к нашему приложению...")
	time.Sleep(time.Second * 1)
	//неожиданно wildberries выкупили Ozon и решили добавить и акции с озона в этот рассыльщик акций, после покупки Озона, нам предоставили их апи с некими похожим наработакми (файл Ozon.go)

	ozonS1 := OzonServer{name: "MoscowServer"}
	ozonS2 := OzonServer{name: "VolgogradServer"}
	//ozonS2 := OzonServer{name: "MoscowServer"}

	//servers = append(servers, &ozonS1, &ozonS2) просто добавить сервера озона в список всех серверов мы не можем, так как он не реализует интерфейс который нам нужен, нужно писать адаптер тк как по каким-то причинам мы не можем реализовать этот интерфейс в OzonServer
	//напишем этот адаптер в ozonApi.go

	ozonS1Adapter := OzonServerAdapter{Server: &ozonS1} //создаем адаптер, теперь мы можем внети этот адаптер в список всех серверов и получать с него данные
	ozonS2Adapter := OzonServerAdapter{Server: &ozonS2}
	servers = append(servers, &ozonS1Adapter, &ozonS2Adapter)

	//проверяем
	for i := 0; i < 5; i++ {

		broadcast(servers...)
		time.Sleep(time.Second * 1)
	}
}

func broadcast(servers ...GetAds) { //сообщение подписчикам об акции (наши подписчики это терминал в данном случае)
	for i := 0; i < len(servers); i++ {
		data := servers[i].Get()
		fmt.Println(data.Store, "\n", "Акция:", data.Name, "\n", data.Description, "\nМесто проведения:", data.Reg)
		fmt.Println()
	}

}

type AdsServer struct { //сервера вайлдберриз которые будут отправлять нам информацию о текущих акциях
	name, region string
}

type GetAds interface { //интерфейс, позволяющий нам получить объявления об акциях
	Get() Ads
}

func (s *AdsServer) Get() Ads { //реализуем интерфейс GetAds серверу AdsServer

	answerFromServer := fmt.Sprintf(`{"server":"%v","name": "Зимняя распродажа","desc": "Скидки на все обогреватели - 50%% !!!","reg":"%v"}`, s.name, s.region) //данные которые приходят нам с сервера

	a := Ads{}
	a.Fill(answerFromServer)
	return a

}

type Ads struct { //тип описывающие сообщение об акции для пользователей
	Store       string `json:"server"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Reg         string `json: "reg"`
}

func (a *Ads) Fill(j string) { //парсинг из json
	json.Unmarshal([]byte(j), a)
}
