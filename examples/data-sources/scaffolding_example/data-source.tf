terraform {
  required_providers {
    lotr = {
      source  = "exlabs/lotr"
      version = "0.0.1"
    }
  }
  required_version = ">= 1.1.0"
}

data "lotr_character" "example" {
}

output "lotr_name" {
  value = data.lotr_character.example.name
}
