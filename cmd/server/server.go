package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"GoNews/pkg/storage/mongo"
	"GoNews/pkg/storage/postgres"
	"log"
	"net/http"
	"os"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	// Создаём объекты баз данных.
	//
	// БД в памяти.
	db := memdb.New()

	// Реляционная БД PostgreSQL.
	pwd := os.Getenv("pgpass")
	connstr := "postgres://postgres:" + pwd + "@0.0.0.0/catalog3"
	db2, err := postgres.New(connstr)
	if err != nil {
		log.Fatal(err)
	}

	// Документная БД MongoDB.
	db3, err := mongo.New("data", "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	_ = db
	_ = db2
	_ = db3

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	http.ListenAndServe(":8080", srv.api.Router()) //<-- для теста закомментить это

	//для теста раскомментить это
	// test area ------------------------------------------------------------------
	/*
		_ = srv
		s := db3

			//Тест на пустой БД
			fmt.Println(s.Posts())
			fmt.Println("Тест на пустой БД выполнен")

			//Тест создания одной записи и вывода её
			err = s.AddPost(storage.Post{ID: 1, Title: "first", Content: "first task"})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Создание записи выполнено, id=", 1)
			fmt.Println(s.Posts())
			fmt.Println("Вывод содержимого БД выполнен")

			//Тест удаления записи
			err = s.DeletePost(storage.Post{ID: 1})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(s.Posts())
			fmt.Println("Удаление записи выполнено")

			//Тест создания массива записей и вывода их двумя методами
			var posts []storage.Post
			for i := 2; i <= 10; i++ {
				t := storage.Post{ID: i, Title: strconv.Itoa(i), Content: strconv.Itoa(i)}
				posts = append(posts, t)
			}
			fmt.Println("Создание массива записей выполнено, posts=", posts)

			var postsIds []int
			for _, post := range posts {
				err = s.AddPost(post)
				if err != nil {
					log.Fatal(err)
				}
				postsIds = append(postsIds, post.ID)
			}

			fmt.Println("Создание записей выполнено, postsIds=", postsIds)
			fmt.Println(s.Posts())
			fmt.Println("Вывод содержимого БД выполнен")

			//Тест изменения записи
			err = s.UpdatePost(storage.Post{
				ID: 2, Title: "Second", Content: "Second task",
				PublishedAt: time.Now().UnixNano()})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(s.Posts())
			fmt.Println("Изменение записи выполнено")

			//Удаление записей
			for _, id := range postsIds {
				err = s.DeletePost(storage.Post{ID: id})
				if err != nil {
					log.Fatal(err)
				}
			}

			fmt.Println(s.Posts())
			fmt.Println("Удаление записи выполнено")
	*/
}
