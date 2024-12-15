data "talos_image_factory_extensions_versions" "this" {
  talos_version = var.image.version
  filters = {
    names = var.image.extensions
  }
}

resource "talos_image_factory_schematic" "this" {
  schematic = yamlencode(
    {
      customization = {
        systemExtensions = {
          officialExtensions = data.talos_image_factory_extensions_versions.this.extensions_info[*].name
        }
      }
    }
  )
}

data "talos_image_factory_urls" "this" {
  talos_version = var.image.version
  schematic_id  = talos_image_factory_schematic.this.id
  architecture  = var.image.architecture
  platform      = var.image.platform
}