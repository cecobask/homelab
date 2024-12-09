data "talos_image_factory_extensions_versions" "this" {
  talos_version = var.release
  filters = {
    names = var.extensions
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
  talos_version = var.release
  schematic_id  = talos_image_factory_schematic.this.id
  architecture  = var.architecture
  platform      = var.platform
}
