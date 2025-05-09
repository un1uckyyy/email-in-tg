//nolint:lll
package tg

const somethingWentWrong = "Что-то пошло не так."

const helpTemplate = `<b>Вот доступные команды, которые вы можете использовать:</b>

<b>/start 'email' 'password'</b> — Начать работу с ботом
<b>/stop</b> — Прекратить все письма. Для того, чтобы возобновить работу используйте /start
<b>/login 'email' 'password'</b> — Обновить адрес почты и пароль. Для изначальной регистрации используется /start
<b>/subscribe 'sender-email'</b> — Подписаться на письма от конкретного отправителя. Рекомендуется использовать в <b>ветках</b> (темах)
<b>/subscribe_others</b> — Подписаться на <b>все</b> письма. Используйте в отдельной ветке, чтобы получать письма от отправителей не из ваших подписок
<b>/subscriptions</b> — Просмотреть ваши текущие подписки

<i>support:</i> {{.}}.
`

const loginTemplate = `<b>Отправь команду /start в следующем формате:</b>
/start 'email-address' 'password'
Как сгенерировать пароль смотри <a href="{{.}}">здесь</a>,
`

const emailTemplate = `✉️ <b>Новое письмо</b> ✉️
<b>От:</b> {{.MailFrom}}
<b>Кому:</b> {{.MailTo}}
<b>Тема:</b> {{.Subject}}
<b>Дата:</b> {{.Date}}

{{.Text}}`

const subscriptionsMessage = `Активные подписки
Нажмите на подписку, чтобы <b>отписаться</b>`
