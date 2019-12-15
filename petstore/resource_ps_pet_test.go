package petstore
import (
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"os"
	"fmt"
)

func TestAccPSPet_basic(t *testing.T) {
	resourceName := "petstores_pet.pet"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPSPetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPSPetConfig_basic("Winston",2,"Siamese"),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Winston"),
					resource.TestCheckResourceAttr(resourceName, "species","cat"),
					resource.TestCheckResourceAttr(resourceName, "age", 2),
				),
			},
		},
	})
}

func testAccCheckPSPetDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "petstore_pet" {
			continue
		}
		fileName := rs.Primary.Attributes["petstore_pet.test.environment.filename"]
		if _, err := os.Stat(fileName); os.IsExist(err) {
			return fmt.Errorf("Shell Script file failed to cleanup")
		}
	}
	return nil
}

func testAccPSPetConfig_basic(name string, species string, age int) string {
	return fmt.Sprintf(`
	resource "petstore_pet" "pet" {
		name = "%s"
		species = "%s"
		age = %d
	  }
`, name, species, age)
}