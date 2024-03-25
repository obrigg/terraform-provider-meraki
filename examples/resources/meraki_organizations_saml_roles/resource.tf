
resource "meraki_organizations_saml_roles" "example" {

  networks = [{

    access = "full"
    id     = "N_24329156"
  }]
  org_access      = "none"
  organization_id = "string"
  role            = "myrole"
  tags = [{

    access = "read-only"
    tag    = "west"
  }]
}

output "meraki_organizations_saml_roles_example" {
  value = meraki_organizations_saml_roles.example
}