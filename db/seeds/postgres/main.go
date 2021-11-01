package main

import (
	"context"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/cjmarkham/GoBB/src/domain/user"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun/dbfixture"

	"github.com/cjmarkham/GoBB/config"
	"github.com/cjmarkham/GoBB/src/domain/forum"
	"github.com/cjmarkham/GoBB/src/domain/post"
	"github.com/cjmarkham/GoBB/src/domain/topic"
	"github.com/cjmarkham/GoBB/src/repository"
)

func main() {
	cfg, err := config.ProvidePostgresqlConfig()
	if err != nil {
		log.Error().Err(err).Msg("could not get postgres config")
	}

	db := repository.ProvideClient(cfg)

	db.RegisterModel(
		(*forum.Forum)(nil),
		(*topic.Topic)(nil),
		(*post.Post)(nil),
		(*user.User)(nil),
	)
	fixture := dbfixture.New(db, dbfixture.WithTruncateTables(), dbfixture.WithTemplateFuncs(funcMap()))

	files := []string{"forums.yml", "topics.yml", "posts.yml", "users.yml"}
	if err := fixture.Load(context.TODO(), os.DirFS("db/seeds/postgres"), files...); err != nil {
		panic(err)
	}
}

func funcMap() template.FuncMap {
	return template.FuncMap{
		"now": func() string {
			return time.Now().Format(time.RFC3339)
		},
		"uuid_nil": func() uuid.UUID {
			return uuid.Nil
		},
		"uuid_rand": func() uuid.UUID {
			return uuid.New()
		},
		"x_days_ago": func(days int) (*string, error) {
			parsed := fmt.Sprintf("-%dh", days*24)
			now := time.Now()
			duration, err := time.ParseDuration(parsed)
			if err != nil {
				return nil, err
			}

			then := now.Add(duration)
			formatted := then.Format(time.RFC3339)
			return &formatted, nil
		},
	}
}
