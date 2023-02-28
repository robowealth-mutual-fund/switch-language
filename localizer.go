package acceptlang

import (
	"context"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func GetLocalizer(ctx context.Context) (*i18n.Localizer, *i18n.Bundle) {
	acceptLangs := FromContext(ctx)
	bundle := i18n.NewBundle(language.English)
	lang := "th"

	if len(acceptLangs.Languages()) > 0 {
		lang = acceptLangs.Languages()[0]
	}

	return i18n.NewLocalizer(bundle, lang), bundle
}

func GetLocalizedField(fieldID string, fields map[language.Tag]string, localizer *i18n.Localizer, bundle *i18n.Bundle) (string, error) {
	if len(fields) > 0 {
		for tag, field := range fields {
			message := i18n.Message{
				ID:    fieldID,
				Other: field,
			}

			err := bundle.AddMessages(tag, &message)
			if err != nil {
				return "", err
			}
		}
	}

	localizeConfig := i18n.LocalizeConfig{
		MessageID: fieldID,
	}

	message, err := localizer.Localize(&localizeConfig)
	if err != nil {
		return "", err
	}

	return message, nil
}
