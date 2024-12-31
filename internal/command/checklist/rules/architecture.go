package rules

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	c "github.com/switchover/eGovFrameChecker/internal/constant"
	"github.com/switchover/eGovFrameChecker/internal/result"
)

func InputArchitectureRule() {
	result.CheckResult.PresentationLayerRule = true
	result.CheckResult.ServiceLayerRule = true
	result.CheckResult.DataAccessLayerRule = true

	inputPresentationLayer()
	if result.CheckResult.PresentationLayerRule {
		fmt.Printf("%s MVC Architecture Rule: %sCompatible%s\n", c.IconOkay, c.LightBlue, c.Reset)
	} else {
		fmt.Printf("%s MVC Architecture Rule: %sNot compatible%s\n", c.IconNotOkay, c.LightBlue, c.Reset)
	}
	fmt.Println("--------------------------------------------------------------------------------")

	inputServiceLayer()
	if result.CheckResult.ServiceLayerRule {
		fmt.Printf("%s Service Architecture Rule: %sCompatible%s\n", c.IconOkay, c.LightBlue, c.Reset)
	} else {
		fmt.Printf("%s Service Architecture Rule: %sNot compatible%s\n", c.IconNotOkay, c.LightBlue, c.Reset)
	}
	fmt.Println("--------------------------------------------------------------------------------")

	inputDataAccessLayer()
	if result.CheckResult.DataAccessLayerRule {
		fmt.Printf("%s Data Access Architecture Rule: %sCompatible%s\n", c.IconOkay, c.LightBlue, c.Reset)
	} else {
		fmt.Printf("%s Data Access Architecture Rule: %sNot compatible%s\n", c.IconNotOkay, c.LightBlue, c.Reset)
	}
	fmt.Println("--------------------------------------------------------------------------------")
}

func inputPresentationLayer() {
	checkPrompt := promptui.Select{
		Label: fmt.Sprintf("Does it consist of Controller classes annotated with "+
			"%s@Controller%s or %s@RestController%s? [Yes/No]", c.LightBlue, c.Reset, c.LightBlue, c.Reset),
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err := checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.PresentationLayerRule = false
	}

	checkPrompt = promptui.Select{
		Label: fmt.Sprintf("Is handler mapping based on Spring MVC annotations "+
			"(%s@RequestMapping%s, %s@GetMapping%s, etc.) applied? [Yes/No]", c.LightBlue, c.Reset, c.LightBlue, c.Reset),
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err = checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.PresentationLayerRule = false
	}

	checkPrompt = promptui.Select{
		Label: "Are the DAO methods not being called directly from Controller? [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err = checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.PresentationLayerRule = false
	}

	checkPrompt = promptui.Select{
		Label: "Is the Controller not calling any data store services(NoSQL, Message Queue, Cache, etc.)? [Yes/No]?",
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err = checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.PresentationLayerRule = false
	}

	checkPrompt = promptui.Select{
		Label: "Is the business logic being performed through the Service injected into the interface? [Yes/No]?",
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err = checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.PresentationLayerRule = false
	}
}

func inputServiceLayer() {
	checkPrompt := promptui.Select{
		Label: fmt.Sprintf("Does it consist of Service classes annotated with "+
			"%s@Service%s or %s@Component%s? [Yes/No]", c.LightBlue, c.Reset, c.LightBlue, c.Reset),
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err := checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.ServiceLayerRule = false
	}

	checkPrompt = promptui.Select{
		Label: fmt.Sprintf("Does the Service class extend %sEgovAbstractServiceImpl%s? [Yes/No]", c.LightBlue, c.Reset),
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err = checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.ServiceLayerRule = false
	}

	checkPrompt = promptui.Select{
		Label: "Does the Service class implement a separate interface? [Yes/No]?",
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err = checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.PresentationLayerRule = false
	}
}

func inputDataAccessLayer() {
	checkPrompt := promptui.Select{
		Label: fmt.Sprintf("Does it consist of Dao/Mapper classes annotated with "+
			"%s@Repository%s? [Yes/No]", c.LightBlue, c.Reset),
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err := checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.DataAccessLayerRule = false
	}

	dataAccessPrompt := promptui.Select{
		Label:     "Select the data access method",
		CursorPos: 1,
		Items:     []string{"iBatis", "MyBatis", "JPA/Hibernate", "Others"},
	}

	_, dataAccess, err := dataAccessPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if dataAccess == "Others" {
		result.CheckResult.DataAccessLayerRule = false
		return
	}

	if dataAccess == "iBatis" {
		checkPrompt = promptui.Select{
			Label: fmt.Sprintf("Does the Dao class extend %sEgovAbstractDao%s? [Yes/No]", c.LightBlue, c.Reset),
			Items: []string{"Yes", "No"},
		}
		_, yesOrNo, err = checkPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
		if yesOrNo == "No" {
			result.CheckResult.DataAccessLayerRule = false
		}
	} else if dataAccess == "MyBatis" {
		checkPrompt = promptui.Select{
			Label: fmt.Sprintf("Does the Mapper class extend %sEgovAbstractMapper%s? [Yes/No]", c.LightBlue, c.Reset),
			Items: []string{"Yes", "No"},
		}
		_, yesOrNo, err = checkPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
		if yesOrNo == "Yes" {
			return
		}

		checkPrompt = promptui.Select{
			Label: fmt.Sprintf("When applying the Mapper interface, "+
				"are the %sMapperConfigurer%s and %s@Mapper%s provided by eGovFrame being used? [Yes/No]",
				c.LightBlue, c.Reset, c.LightBlue, c.Reset),
			Items: []string{"Yes", "No"},
		}
		_, yesOrNo, err = checkPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
		if yesOrNo == "No" {
			result.CheckResult.DataAccessLayerRule = false
		}
	} else {
		checkPrompt = promptui.Select{
			Label: fmt.Sprintf("Does the Repository class extend %sJpaRepository%s "+
				"(or its superclasses %sCrudRepository%s or %sPagingAndSortingRepository%s)? [Yes/No]",
				c.LightBlue, c.Reset, c.LightBlue, c.Reset, c.LightBlue, c.Reset),
			Items: []string{"Yes", "No"},
		}
		_, yesOrNo, err = checkPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
		if yesOrNo == "Yes" {
			return
		}

		checkPrompt = promptui.Select{
			Label: fmt.Sprintf("Or is %sHibernateTemplate%s or %sEntityManager%s injected and used?? [Yes/No]",
				c.LightBlue, c.Reset, c.LightBlue, c.Reset),
			Items: []string{"Yes", "No"},
		}
		_, yesOrNo, err = checkPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
		if yesOrNo == "No" {
			result.CheckResult.DataAccessLayerRule = false
		}
	}
}
