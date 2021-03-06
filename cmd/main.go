package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	activityhandlers "github.com/telegram-go-bot/go_bot/app/activity_handlers"
	onpickfirstorsecond "github.com/telegram-go-bot/go_bot/app/activity_handlers/on_pick_first_or_second"
	onzagadka "github.com/telegram-go-bot/go_bot/app/activity_handlers/on_zagadka"
	collywrapper "github.com/telegram-go-bot/go_bot/app/common/web_scrapper/colly_wrapper"
	"github.com/telegram-go-bot/go_bot/app/domain"
	in "github.com/telegram-go-bot/go_bot/app/input/activities/telegram"
	presenters "github.com/telegram-go-bot/go_bot/app/output/presenters"
	"github.com/telegram-go-bot/go_bot/app/output/views/telegram"
)

var (
	botToken   = os.Getenv("HEROKU_BOT_ID")
	baseURL    = os.Getenv("HEROKU_BASE_URL")
	dbURL      = os.Getenv("DATABASE_URL")
	botUuids   = strings.Split(os.Getenv("BOT_UIDS"), ",")
	vkLogin    = os.Getenv("VK_LOGIN")
	vkPwd      = os.Getenv("VK_PASSWORD")
	botAdminID = os.Getenv("OWNER_ID")
)

var (
	tgView      = telegram.NewTelegramAPIView(botToken)
	tgPresenter = presenters.NewActivityPresenter(tgView)
	tgReader    = in.NewMessageReader(botToken)
	scrapper    = collywrapper.Scrapper{}

	commandHandlers = []activityhandlers.ICommandHandler{
		onpickfirstorsecond.New(tgPresenter),
		onzagadka.New(tgPresenter, scrapper)}
)

func main() {
	tgBot := domain.NewActivityHandlerFacade(commandHandlers)
	tgBot.ProcessActivities(tgReader)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<3")
}
