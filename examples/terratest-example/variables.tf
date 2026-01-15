####################################################################################################################
# DEFINE VARIABLES HERE, I USUALLY SEPARATE THE VARIABLES INTO SECTIONS TO HELP FIND THEM EASIER/KEEP THEM ORGANIZED
#
#       variable "foo" {
#          type = string
#          description = "value"
#          default = ""
#        }
#
####################################################################################################################

variable "foo" {
  type        = string
  description = "value"
  default     = ""
}

variable "bar" {
  type        = map(any)
  description = "value"
  default     = {}
}

variable "baz" {
  type        = list(string)
  description = "value"
  default     = ["value"]
}