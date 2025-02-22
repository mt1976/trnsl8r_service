package jobs

import (
	"context"
	"fmt"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dao/actions"
	"github.com/mt1976/frantic-core/dao/database"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/timing"
	"github.com/mt1976/trnsl8r_service/app/business/domains"
	"github.com/mt1976/trnsl8r_service/app/business/translation"
	"github.com/mt1976/trnsl8r_service/app/dao/textstore"
)

type localeUpdate struct {
}

func (job *localeUpdate) Run() error {

	j := timing.Start(domains.JOBS.String(), actions.VALIDATE.GetCode(), "Locales")
	textstore.Initialise(context.TODO())
	set := commonConfig.Get()

	locales := set.GetLocales()
	noLocales := len(locales)

	logHandler.ServiceLogger.Printf("[%v] Number of Locales=[%v]", domain.String(), noLocales)

	// Check that the locales follow the structure zz_ZZ, otherwise error out
	for _, locale := range locales {
		logHandler.ServiceLogger.Printf("[%v] Checking locale formatting for [%v] %v", domain.String(), locale.Key, locale.Name)
		if len(locale.Key) != 5 {
			msg := fmt.Sprintf("locale [%v] is not in the correct format", locale.Key)
			logHandler.ErrorLogger.Println(msg)
			j.Stop(0)
			return fmt.Errorf("locale is not in the correct format, length is not 5")
		}
		if locale.Key[2] != '_' {
			msg := fmt.Sprintf("locale [%v] is not in the correct format", locale.Key)
			logHandler.ErrorLogger.Println(msg)
			j.Stop(0)
			return fmt.Errorf("locale is not in the correct format, 3rd character is not an underscore")
		}
		if locale.Key[0] < 'a' || locale.Key[0] > 'z' {
			msg := fmt.Sprintf("locale [%v] is not in the correct format", locale.Key)
			logHandler.ErrorLogger.Println(msg)
			j.Stop(0)
			return fmt.Errorf("locale is not in the correct format, 1st character is not a lowercase letter")
		}
		if locale.Key[1] < 'a' || locale.Key[1] > 'z' {
			msg := fmt.Sprintf("locale [%v] is not in the correct format", locale.Key)
			logHandler.ErrorLogger.Println(msg)
			j.Stop(0)
			return fmt.Errorf("locale is not in the correct format, 2nd character is not a lowercase letter")
		}
		if locale.Key[3] < 'A' || locale.Key[3] > 'Z' {
			msg := fmt.Sprintf("locale [%v] is not in the correct format", locale.Key)
			logHandler.ErrorLogger.Println(msg)
			j.Stop(0)
			return fmt.Errorf("locale is not in the correct format, 4th character is not an uppercase letter")
		}
		if locale.Key[4] < 'A' || locale.Key[4] > 'Z' {
			msg := fmt.Sprintf("locale [%v] is not in the correct format", locale.Key)
			logHandler.ErrorLogger.Println(msg)
			j.Stop(0)
			return fmt.Errorf("locale is not in the correct format, 5th character is not an uppercase letter")
		}
	}
	logHandler.ServiceLogger.Printf("[%v] Locales are in the correct format", domain.String())
	// We can assume that all the locales are in the correct format
	textData, err := textstore.GetAll()
	if err != nil {
		logHandler.ErrorLogger.Println(err)
		j.Stop(0)
		return err
	}
	noText := len(textData)
	logHandler.ServiceLogger.Printf("[%v] Number of Translations to Verify=[%v]", domain.String(), noText)
	upgradeText := false
	for thisPos, text := range textData {
		if len(text.Localised) != noLocales {
			msg := fmt.Sprintf("text [%v] does not have the correct number of locales, has [%v], want [%v]", text.Signature, len(text.Localised), noLocales)
			logHandler.ServiceLogger.Println(msg)
			upgradeText = true
		}
		// Check that the locales in localised map match the locales in the settings
		for key := range text.Localised {
			found := false
			for _, locale := range locales {
				if key == locale.Key {
					found = true
					break
				}
			}
			if !found {
				msg := fmt.Sprintf("text [%v] has locale [%v] that is not in the settings", text.Signature, key)
				logHandler.ServiceLogger.Println(msg)
				upgradeText = true
			}
		}

		if upgradeText {
			msg := fmt.Sprintf("Upgrading text (%v/%v) [%v]", thisPos+1, noText, text.Signature)
			logHandler.ServiceLogger.Println(msg)
		}

		if upgradeText {
			// Add the missing locales to the text.Localised map
			newTextLocalised := make(map[string]string)

			for _, locale := range locales {

				// if the locale map is empty create it
				if text.Localised == nil {
					text.Localised = make(map[string]string)
				}

				newTextLocalised[locale.Key] = text.Localised[locale.Key]

			}
			text.Localised = newTextLocalised
			// Update the text in the database
			msg := fmt.Sprintf("Localisation Upgrade, update to [%v]", locales)
			err = text.Update(context.TODO(), msg)
			if err != nil {
				logHandler.ErrorLogger.Println(err)
				j.Stop(thisPos)
				return err
			}
		}
	}

	logHandler.ServiceLogger.Printf("[%v] Locales Updated", domain.String())

	j.Stop(noText)
	return nil
}

func (job *localeUpdate) Service() func() {
	return func() {
		job.Run()
	}
}

func (job *localeUpdate) Schedule() string {
	return "10 7 * * *"
}

func (job *localeUpdate) Name() string {
	return translation.Get("Update Locales", "")
}

func (job *localeUpdate) AddFunction(fn func() (*database.DB, error)) {
	// do nothing
}

func (t *localeUpdate) Description() string {
	return "Update Locales"
}
