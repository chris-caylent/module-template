################################################################################
# THIS IS WHERE YOU WILL PUT ANY RESOURCES, MODULE CALLS, OR LOCALS
################################################################################

module "foo" {
  source = "../../"

  foo = var.foo
  bar = var.bar
  baz = var.baz

}